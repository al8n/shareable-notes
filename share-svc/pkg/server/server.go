package server

import (
	"fmt"
	"github.com/al8n/shareable-notes/share-svc/config"
	sharepb "github.com/al8n/shareable-notes/share-svc/pb"
	shareendpoint "github.com/al8n/shareable-notes/share-svc/pkg/endpoint"
	shareservice "github.com/al8n/shareable-notes/share-svc/pkg/service"
	sharetransport "github.com/al8n/shareable-notes/share-svc/pkg/transport"
	bootapi "github.com/al8n/micro-boot/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	consulsd "github.com/go-kit/kit/sd/consul"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
)

var (
	once sync.Once
	server *Server
)

func GetServer() *Server {
	once.Do(func() {
		server = &Server{}
	})
	return server
}

type Server struct {
	httpListener net.Listener
	httpConsulRegister *consulsd.Registrar

	httpsListener net.Listener
	httpsConsulRegister *consulsd.Registrar
	router *mux.Router

	shareServer  sharepb.ShareServer
	grpcServer  *grpc.Server
	grpcListener net.Listener
	grpcConsulRegister *consulsd.Registrar

	tracerCloser io.Closer

	logger log.Logger
	wg sync.WaitGroup
}

func (s *Server) Serve() (err error) {
	var cfg = config.GetConfig()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		s.logger = logger
	}


	var (
		tracer stdopentracing.Tracer
		jaegerCfg *jaegerconfig.Configuration
	)
	{
		jaegerCfg, err = jaegerconfig.FromEnv()
		if err != nil {
			return err
		}

		jaegerCfg.Reporter.LogSpans = true

		jaegerCfg.Sampler = &jaegerconfig.SamplerConfig{
			Type: "const",
			Param: 1,
		}

		jaegerCfg.ServiceName = cfg.Service.Name


		tracer, s.tracerCloser, err = jaegerCfg.NewTracer(jaegerconfig.Logger(jaeger.StdLogger))
		if err != nil {
			return err
		}
		stdopentracing.SetGlobalTracer(tracer)
	}

	// Create the (sparse) metrics we'll use in the service. They, too, are
	// dependencies that we pass to components that use them.
	var ctrs  = make(map[string]metrics.Counter)
	{
		// Business-level metrics.
		opts := cfg.Prom.CounterOptions
		for key, opt := range opts {
			ctrs[key] = prometheus.NewCounterFrom(
				stdprometheus.CounterOpts{
					Name: opt.Name,
					Namespace: opt.Namespace,
					Subsystem: opt.Subsystem,
					Help: opt.Help,
				},
				opt.LabelNames)
		}
	}

	var duration = make( map[string]metrics.Histogram )
	{
		// Endpoint-level metrics.
		opts := cfg.Prom.SummaryOptions
		for key, opt := range opts {
			duration[key] = prometheus.NewSummaryFrom(
				opt.Standardize(),
				opt.LabelNames)
		}
	}

	var (
		service shareservice.Service
		endpoints *shareendpoint.Set
		httpAddr  = ":" + cfg.HTTP.Port
		httpsAddr  = ":" + cfg.HTTPS.Port
		grpcAddr = ":" + cfg.GRPC.Port
	)
	{
		service, err = shareservice.New(logger, ctrs, tracer)
		if err != nil {
			logger.Log("err", err)
			return err
		}

		endpoints, err = shareendpoint.New(service, logger, duration, tracer)
		if err != nil {
			logger.Log("err", err)
			return err
		}

		if cfg.HTTP.Runnable || cfg.HTTPS.Runnable {
			s.router = sharetransport.NewHTTPHandler(*endpoints, tracer, logger, cfg.Service.APIs)
			s.router.Handle(cfg.Prom.Path, promhttp.Handler())
		} else {
			r := mux.NewRouter()
			r.Handle(cfg.Prom.Path, promhttp.Handler())
			http.ListenAndServe(":9090", r)
		}

		if cfg.HTTP.Runnable {
			err = s.serveHTTP(httpAddr, logger)
			if err != nil {
				return err
			}
		}

		if cfg.HTTPS.Runnable {
			err = s.serveHTTPS(httpsAddr, logger, cfg.HTTPS.Cert, cfg.HTTPS.Key)
			if err != nil {
				return err
			}
		}

		if cfg.GRPC.Runnable {
			err = s.serveRPC(grpcAddr, *endpoints, logger, tracer, cfg.Service.APIs)
			if err != nil {
				return err
			}
		}
	}

	s.wg.Wait()
	return
}

func (s *Server) Close() (err error) {
	var cfg = config.GetConfig()

	if cfg.HTTP.Runnable {
		s.httpConsulRegister.Deregister()
		s.logger.Log("transport", "HTTP", "op", "Close", "error", s.httpListener.Close())
	}

	if cfg.HTTPS.Runnable {
		s.httpsConsulRegister.Deregister()
		s.logger.Log("transport", "gRPC", "op", "Close", "error", s.httpsListener.Close())
	}

	if cfg.GRPC.Runnable {
		s.grpcConsulRegister.Deregister()
		s.logger.Log("transport", "gRPC", "op", "Close", "error", s.grpcListener.Close())
	}

	s.tracerCloser.Close()

	return nil
}

func (s *Server) serveRPC(address string,  endpoints shareendpoint.Set, logger log.Logger, tracer stdopentracing.Tracer, apis bootapi.APIs) (err error) {
	s.grpcListener, err = net.Listen("tcp", address)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		return err
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	s.shareServer = sharetransport.NewGRPCServer(endpoints, tracer, logger, apis)

	sharepb.RegisterShareServer(s.grpcServer, s.shareServer)

	s.grpcConsulRegister, err = NewConsulGRPCRegister(logger)
	if err != nil {
		return err
	}

	go func() {
		s.wg.Add(1)
		logger.Log("transport", "gRPC", "addr", address)
		s.grpcConsulRegister.Register()
		s.grpcServer.Serve(s.grpcListener)
		s.wg.Done()
	}()

	return nil
}

func (s *Server) serveHTTP(address string, logger log.Logger) (err error)  {
	s.httpListener, err = net.Listen("tcp", address)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		return err
	}

	s.httpConsulRegister, err = NewConsulHTTPRegister(logger)
	if err != nil {
		return err
	}

	go func() {
		s.wg.Add(1)
		logger.Log("transport", "HTTP", "addr", address)
		s.httpConsulRegister.Register()
		http.Serve(s.httpListener, s.router)
		s.wg.Done()
	}()

	return nil
}

func (s *Server) serveHTTPS(address string, logger log.Logger, cert, key string) (err error)  {
	s.httpsListener, err = net.Listen("tcp", address)
	if err != nil {
		logger.Log("transport", "HTTPS", "during", "Listen", "err", err)
		return err
	}

	s.httpsConsulRegister, err = NewConsulHTTPSRegister(logger)
	if err != nil {
		return err
	}

	go func() {
		s.wg.Add(1)
		logger.Log("transport", "HTTPS", "addr", address)
		s.httpsConsulRegister.Deregister()
		http.ServeTLS(s.httpsListener, s.router, cert, key)
		s.wg.Done()
	}()

	return nil
}

func NewConsulGRPCRegister(logger log.Logger) (register *consulsd.Registrar, err error)  {
	var (
		cfg = config.GetConfig()
		consulClient *api.Client
		client consulsd.Client
		apiCfg *api.Config
	)

	apiCfg = cfg.Consul.Client.Standardize()

	consulClient, err = api.NewClient(apiCfg)

	if err != nil {
		return nil, err
	}

	client = consulsd.NewClient(consulClient)


	reg := cfg.Consul.Agent.Standardize()
	reg.ID = fmt.Sprintf("%v-%v-%v", cfg.Service.Name, cfg.Host, cfg.GRPC.Port)
	reg.Name = cfg.GRPC.Name
	reg.Port = cfg.GRPC.GetIntPort()

	return consulsd.NewRegistrar(client, reg, logger), nil
}

func NewConsulHTTPRegister(logger log.Logger) (register *consulsd.Registrar, err error)  {
	var (
		cfg = config.GetConfig()
		consulClient *api.Client
		client consulsd.Client
		apiCfg *api.Config
	)


	apiCfg = cfg.Consul.Client.Standardize()
	consulClient, err = api.NewClient(apiCfg)
	if err != nil {
		return nil, err
	}

	client = consulsd.NewClient(consulClient)

	reg := cfg.Consul.Agent.Standardize()

	reg.ID = fmt.Sprintf("%v-%v-%v", cfg.Service.Name, cfg.Host, cfg.HTTP.Port)
	reg.Name = cfg.HTTP.Name
	reg.Port = cfg.HTTP.GetIntPort()

	return consulsd.NewRegistrar(client, reg, logger), nil
}

func NewConsulHTTPSRegister(logger log.Logger) (register *consulsd.Registrar, err error)  {
	var (
		cfg = config.GetConfig()
		consulClient *api.Client
		client consulsd.Client
		apiCfg *api.Config
	)

	apiCfg = cfg.Consul.Client.Standardize()

	consulClient, err = api.NewClient(apiCfg)
	if err != nil {
		return nil, err
	}

	client = consulsd.NewClient(consulClient)

	reg := cfg.Consul.Agent.Standardize()
	reg.ID = fmt.Sprintf("%v-%v-%v", cfg.Service.Name, cfg.Host, cfg.HTTPS.Port)
	reg.Name = cfg.HTTPS.Name
	reg.Port = cfg.HTTPS.GetIntPort()

	return consulsd.NewRegistrar(client, reg, logger), nil
}