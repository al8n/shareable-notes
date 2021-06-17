package transport

import (
	"context"
	"github.com/al8n/shareable-notes/share-svc/internal/codec/grpccodec/grpcdecode"
	"github.com/al8n/shareable-notes/share-svc/internal/codec/grpccodec/grpcencode"
	"github.com/al8n/shareable-notes/share-svc/pb"
	serviceendpoint "github.com/al8n/shareable-notes/share-svc/pkg/endpoint"
	shareservice "github.com/al8n/shareable-notes/share-svc/pkg/service"
	bootapi "github.com/al8n/micro-boot/api"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	shareNote grpctransport.Handler
	privateNote grpctransport.Handler
	getNote grpctransport.Handler
}

func (g GRPCServer) ShareNote(ctx context.Context, request *pb.ShareNoteRequest) (*pb.ShareNoteResponse, error) {
	_, resp, err := g.shareNote.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ShareNoteResponse), nil
}

func (g GRPCServer) PrivateNote(ctx context.Context, request *pb.PrivateNoteRequest) (*pb.PrivateNoteResponse, error) {
	_, resp, err := g.privateNote.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PrivateNoteResponse), nil
}

func (g GRPCServer) GetNote(ctx context.Context, request *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	_, resp, err := g.getNote.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetNoteResponse), nil
}

func NewGRPCServer(set serviceendpoint.Set, otTracer stdopentracing.Tracer, logger log.Logger, apis bootapi.APIs) pb.ShareServer  {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	//if zipkinTracer != nil {
	//	// Zipkin GRPC Server Trace can either be instantiated per gRPC method with a
	//	// provided operation name or a global tracing service can be instantiated
	//	// without an operation name and fed to each Go kit gRPC server as a
	//	// ServerOption.
	//	// In the latter case, the operation name will be the endpoint's grpc method
	//	// path if used in combination with the Go kit gRPC Interceptor.
	//	//
	//	// In this example, we demonstrate a global Zipkin tracing service with
	//	// Go kit gRPC Interceptor.
	//	options = append(options, zipkin.GRPCServerTrace(zipkinTracer))
	//}

	return &GRPCServer{
		shareNote:   grpctransport.NewServer(
			set.ShareNoteEndpoint,
			grpcdecode.ShareNoteRequest,
			grpcencode.ShareNoteResponse,
			append(
				options,
				grpctransport.ServerBefore(
					opentracing.GRPCToContext(
						otTracer,
						"ShareNote",
						logger)),
				)...,
			),
		privateNote: grpctransport.NewServer(
			set.PrivateNoteEndpoint,
			grpcdecode.PrivateNoteRequest,
			grpcencode.PrivateNoteResponse,
			append(
				options,
				grpctransport.ServerBefore(
					opentracing.GRPCToContext(
						otTracer,
						"PrivateNote",
						logger)),
			)...,
		),
		getNote:    grpctransport.NewServer(
			set.GetNoteEndpoint,
			grpcdecode.GetNoteRequest,
			grpcencode.GetNoteResponse,
			append(
				options,
				grpctransport.ServerBefore(
					opentracing.GRPCToContext(
						otTracer,
						"GetNote",
						logger)),
			)...,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, logger log.Logger, apis bootapi.APIs) shareservice.Service {

	// global client middlewares
	var (
		options []grpctransport.ClientOption
		serviceName = "pb.Share"
	)

	// Each individual endpoint is an grpc/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// made your own client library, you'd do this work there, so your server
	// could rely on a consistent set of client behavior.
	var shareNoteEndpoint endpoint.Endpoint
	{
		var (

			name = shareservice.ShareNoteServiceName
			rl = apis[name].RateLimit
			bkr = apis[name].Breaker
		)

		shareNoteEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			name,
			grpcencode.ShareNoteRequest,
			grpcdecode.ShareNoteResponse,
			pb.ShareNoteResponse{},
			append(
				options,
				grpctransport.ClientBefore(
					opentracing.ContextToGRPC(
							otTracer,
							logger,
						),
					),
				)...,
		).Endpoint()

		shareNoteEndpoint = opentracing.TraceClient(otTracer, name)(shareNoteEndpoint)

		// We construct a single ratelimiter middleware, to limit the total outgoing
		// QPS from this client to all methods on the remote instance. We also
		// construct per-endpoint circuitbreaker middlewares to demonstrate how
		// that's done, although they could easily be combined into a single breaker
		// for the entire remote instance, too.
		shareNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(rl.Duration),
				rl.Delta),
			)(shareNoteEndpoint)

		shareNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				bkr.Standardize()),
			)(shareNoteEndpoint)
	}

	var privateNoteEndpoint endpoint.Endpoint
	{
		var (
			name = shareservice.PrivateNoteServiceName
			rl = apis[name].RateLimit
			bkr = apis[name].Breaker
		)

		privateNoteEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			name,
			grpcencode.PrivateNoteRequest,
			grpcdecode.PrivateNoteResponse,
			pb.PrivateNoteResponse{},
			append(
				options,
				grpctransport.ClientBefore(
					opentracing.ContextToGRPC(otTracer, logger)),
				)...,
		).Endpoint()

		privateNoteEndpoint = opentracing.TraceClient(otTracer, name)(privateNoteEndpoint)

		privateNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(
					rl.Duration),
					rl.Delta),
			)(privateNoteEndpoint)

		privateNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				bkr.Standardize()),
			)(privateNoteEndpoint)
	}

	var getNoteEndpoint endpoint.Endpoint
	{
		var (
			name = shareservice.GetNoteServiceName
			rl = apis[name].RateLimit
			bkr = apis[name].Breaker
		)

		getNoteEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			name,
			grpcencode.GetNoteRequest,
			grpcdecode.GetNoteResponse,
			pb.GetNoteResponse{},
			append(
				options,
				grpctransport.ClientBefore(
					opentracing.ContextToGRPC(otTracer, logger)),
			)...,
		).Endpoint()

		getNoteEndpoint = opentracing.TraceClient(otTracer, name)(getNoteEndpoint)

		getNoteEndpoint = ratelimit.NewErroringLimiter(
			rate.NewLimiter(
				rate.Every(
					rl.Duration),
					rl.Delta),
		)(getNoteEndpoint)

		getNoteEndpoint = circuitbreaker.Gobreaker(
			gobreaker.NewCircuitBreaker(
				bkr.Standardize()),
		)(getNoteEndpoint)
	}

	// Returning the endpoint.Endpoints as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return serviceendpoint.Set{
		ShareNoteEndpoint: shareNoteEndpoint,
		PrivateNoteEndpoint: privateNoteEndpoint,
		GetNoteEndpoint: getNoteEndpoint,
	}
}

