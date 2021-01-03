package transport

import (
	"github.com/ALiuGuanyan/margin/share/internal/codec/httpcodec"
	"github.com/ALiuGuanyan/margin/share/internal/codec/httpcodec/httpdecode"
	"github.com/ALiuGuanyan/margin/share/internal/codec/httpcodec/httpencode"
	serviceendpoint "github.com/ALiuGuanyan/margin/share/pkg/endpoint"
	shareservice "github.com/ALiuGuanyan/margin/share/pkg/service"
	bootapi "github.com/ALiuGuanyan/micro-boot/api"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"net/url"
	"strings"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints serviceendpoint.Set, otTracer stdopentracing.Tracer, logger log.Logger, apis bootapi.APIs) *mux.Router {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(httpcodec.ErrorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	//if zipkinTracer != nil {
	//	// Zipkin HTTP Server Trace can either be instantiated per endpoint with a
	//	// provided operation name or a global tracing service can be instantiated
	//	// without an operation name and fed to each Go kit endpoint as ServerOption.
	//	// In the latter case, the operation name will be the endpoint's http method.
	//	// We demonstrate a global tracing service here.
	//	options = append(options, zipkin.HTTPServerTrace(zipkinTracer))
	//}


	var (
		r *mux.Router
		sn bootapi.API
		pn bootapi.API
		gn bootapi.API
	)
	{
		r = mux.NewRouter()
		sn = apis[shareservice.ShareNoteServiceName]


		r.Methods(sn.Method).Path(sn.Path).Handler(	httptransport.NewServer(
			endpoints.ShareNoteEndpoint,
			httpdecode.ShareNoteRequest,
			httpencode.ShareNoteResponse,
			append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "ShareNote", logger)))...,
		))

		pn = apis[shareservice.PrivateNoteServiceName]
		r.Methods(pn.Method).Path(pn.Path).Handler(httptransport.NewServer(
			endpoints.PrivateNoteEndpoint,
			httpdecode.PrivateNoteRequest,
			httpencode.PrivateNoteResponse,
			append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "PrivateNote", logger)))...,
		))

		gn = apis[shareservice.GetNoteServiceName]
		r.Methods(gn.Method).Path(gn.Path).Handler(httptransport.NewServer(
			endpoints.GetNoteEndpoint,
			httpdecode.GetNoteRequest,
			httpencode.GetNoteResponse,
			append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "GetNote", logger)))...,
		))

	}

	return r
}

func NewHTTPClient(instance string, otTracer stdopentracing.Tracer, logger log.Logger, apis bootapi.APIs) (shareservice.Service, error) {
	// Quickly sanitize the instance string.
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}


	// Each individual endpoint is an http/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// made your own client library, you'd do this work there, so your server
	// could rely on a consistent set of client behavior.
	var shareNoteEndpoint endpoint.Endpoint
	{
		var (
			name = shareservice.ShareNoteServiceName
			sn = apis[name]
		)

		shareNoteEndpoint = httptransport.NewClient(
			sn.Method,
			copyURL(u, sn.Path),
			httpencode.GenericRequest,
			httpdecode.ShareNoteResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(otTracer, logger)),
		).Endpoint()

		shareNoteEndpoint = opentracing.TraceClient(otTracer, name)(shareNoteEndpoint)

		// We construct a single ratelimiter middleware, to limit the total outgoing
		// QPS from this client to all methods on the remote instance. We also
		// construct per-endpoint circuitbreaker middlewares to demonstrate how
		// that's done, although they could easily be combined into a single breaker
		// for the entire remote instance, too.
		shareNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(sn.RateLimit.Duration),
				sn.RateLimit.Delta),
			)(shareNoteEndpoint)

		shareNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				sn.Breaker.Standardize()),
			)(shareNoteEndpoint)
	}


	var privateNoteEndpoint endpoint.Endpoint
	{
		var (
			name = shareservice.PrivateNoteServiceName
			pn = apis[name]
		)

		privateNoteEndpoint = httptransport.NewClient(
			pn.Method,
			copyURL(u, pn.Path),
			httpencode.GenericRequest,
			httpdecode.PrivateNoteResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(otTracer, logger)),
		).Endpoint()
		privateNoteEndpoint = opentracing.TraceClient(otTracer, name)(privateNoteEndpoint)


		privateNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(pn.RateLimit.Duration),
				pn.RateLimit.Delta))(privateNoteEndpoint)

		privateNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				pn.Breaker.Standardize()),
			)(privateNoteEndpoint)
	}

	var getNoteEndpoint endpoint.Endpoint
	{
		var (
			name = shareservice.GetNoteServiceName
			gn = apis[name]
		)

		getNoteEndpoint = httptransport.NewClient(
			gn.Method,
			copyURL(u, gn.Path),
			httpencode.GenericRequest,
			httpdecode.PrivateNoteResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(otTracer, logger)),
		).Endpoint()
		getNoteEndpoint = opentracing.TraceClient(otTracer, name)(getNoteEndpoint)

		getNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(gn.RateLimit.Duration),
				gn.RateLimit.Delta))(getNoteEndpoint)

		getNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				gn.Breaker.Standardize()),
		)(getNoteEndpoint)
	}

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return serviceendpoint.Set{
		ShareNoteEndpoint:    shareNoteEndpoint,
		PrivateNoteEndpoint: privateNoteEndpoint,
		GetNoteEndpoint: getNoteEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}