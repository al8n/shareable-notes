package server

import (
	"github.com/ALiuGuanyan/margin/apigateway/config"
	shareendpoint "github.com/ALiuGuanyan/margin/share/pkg/endpoint"
	shareservice "github.com/ALiuGuanyan/margin/share/pkg/service"
	sharetransport "github.com/ALiuGuanyan/margin/share/pkg/transport"
	"github.com/go-kit/kit/log"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"sync"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"os"
)

type Server struct {
	tracerCloser io.Closer
	handler http.Handler
	wg sync.WaitGroup
}

func (s *Server) Serve() (err error) {
	var cfg = config.GetConfig()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(cfg.ConsulAddr) > 0 {
			consulConfig.Address = cfg.ConsulAddr
		}
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	// Transport domain.
	var tracer stdopentracing.Tracer // no-op
	{
		jaegerCfg := &jaegerconfig.Configuration{
			ServiceName: "API Gateway",
			Sampler: &jaegerconfig.SamplerConfig{
				Type: "const",
				Param: 1,
			},
			Reporter: &jaegerconfig.ReporterConfig{
				LogSpans: true,
			},
		}

		tracer, s.tracerCloser, err = jaegerCfg.NewTracer(jaegerconfig.Logger(jaeger.StdLogger))
		if err != nil {
			return err
		}
		stdopentracing.SetGlobalTracer(tracer)
	}

	var r = mux.NewRouter()
	// share routes
	{
		var (
			tags        = []string{}
			passingOnly = true
			endpoints   = shareendpoint.Set{}
			instancer   = consulsd.NewInstancer(client, logger, cfg.ShareSVC.Name, tags, passingOnly)
		)
		{
			factory := sharesvcFactory(shareendpoint.MakeShareNoteEndpoint, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(cfg.RetryMax, cfg.RetryTimeout, balancer)
			endpoints.ShareNoteEndpoint = retry
		}
		{
			factory := sharesvcFactory(shareendpoint.MakePrivateNoteEndpoint, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(cfg.RetryMax, cfg.RetryTimeout, balancer)
			endpoints.PrivateNoteEndpoint = retry
		}
		{
			factory := sharesvcFactory(shareendpoint.MakeGetNoteEndpoint, tracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(cfg.RetryMax, cfg.RetryTimeout, balancer)
			endpoints.GetNoteEndpoint = retry
		}

		r.PathPrefix("/share").Handler(
				http.StripPrefix(
					"/share",
					sharetransport.NewHTTPHandler(
						endpoints,
						tracer,
						logger,
						cfg.ShareSVC.APIs),
				),
			)

		r.Handle("/metrics", promhttp.Handler())
	}


	s.handler = r
	go func() {
		s.wg.Add(1)
		logger.Log("transport", "HTTP", "addr", cfg.HTTP.Port)
		http.ListenAndServe(":" + cfg.HTTP.Port, r)
	}()
	s.wg.Wait()
	return nil
}

func (s *Server) Close() (err error) {
	s.wg.Done()
	return nil
}

func sharesvcFactory(makeEndpoint func(service shareservice.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		svc := sharetransport.NewGRPCClient(conn, tracer, logger, config.GetConfig().ShareSVC.APIs)
		ep := makeEndpoint(svc)

		return ep, conn, nil
	}
}
