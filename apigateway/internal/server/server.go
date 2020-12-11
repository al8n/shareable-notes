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
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	var zipkinTracer *stdzipkin.Tracer
	{
		if cfg.Zipkin.Reporter.URL != "" {
			var (
				err error
				lep *model.Endpoint
			)

			lep, err = cfg.Zipkin.Tracer.LocalEndpoint.Standardize()
			if err != nil {
				logger.Log("err", err)
				panic(err)
			}

			zipkinTracer, err = stdzipkin.NewTracer(cfg.Zipkin.StandardReporter(), stdzipkin.WithLocalEndpoint(lep))
			if err != nil {
				logger.Log("err", err)
				panic(err)
			}

			if !cfg.Zipkin.Bridge {
				logger.Log("tracer", "Zipkin", "type", "Native", "URL", cfg.Zipkin.Reporter.URL)
			}
		}
	}

	// Transport domain.
	var tracer stdopentracing.Tracer // no-op
	{
		if cfg.Zipkin.Bridge && zipkinTracer != nil {
			logger.Log("tracer", "Zipkin", "type", "OpenTracing", "URL", cfg.Zipkin.Reporter.URL)
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		}
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
			factory := sharesvcFactory(shareendpoint.MakeShareNoteEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(cfg.RetryMax, cfg.RetryTimeout, balancer)
			endpoints.ShareNoteEndpoint = retry
		}
		{
			factory := sharesvcFactory(shareendpoint.MakePrivateNoteEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(cfg.RetryMax, cfg.RetryTimeout, balancer)
			endpoints.PrivateNoteEndpoint = retry
		}
		{
			factory := sharesvcFactory(shareendpoint.MakeGetNoteEndpoint, tracer, zipkinTracer, logger)
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
						zipkinTracer,
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

func sharesvcFactory(makeEndpoint func(service shareservice.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		svc := sharetransport.NewGRPCClient(conn, tracer, zipkinTracer, logger, config.GetConfig().ShareSVC.APIs)
		ep := makeEndpoint(svc)

		return ep, conn, nil
	}
}

func sharesvcHTTPFactory(makeEndpoint func(service shareservice.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {

		svc, err := sharetransport.NewHTTPClient(instance, tracer, zipkinTracer, logger, config.GetConfig().ShareSVC.APIs)
		if err != nil {
			return nil, nil, err
		}


		ep := makeEndpoint(svc)

		return ep, nil, nil
	}
}