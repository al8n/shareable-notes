package endpoint

import (
	"context"
	"github.com/ALiuGuanyan/margin/share/config"
	"github.com/ALiuGuanyan/margin/share/internal/utils"
	"github.com/ALiuGuanyan/margin/share/model/requests"
	"github.com/ALiuGuanyan/margin/share/model/responses"
	shareservice "github.com/ALiuGuanyan/margin/share/pkg/service"
	bootapi "github.com/ALiuGuanyan/micro-boot/api"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/ratelimit"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

type MakeEndpointFunc = func(shareservice.Service) endpoint.Endpoint

type Set struct {
	ShareNoteEndpoint endpoint.Endpoint
	PrivateNoteEndpoint endpoint.Endpoint
	GetNoteEndpoint endpoint.Endpoint
}

func (s Set) GetNote(ctx context.Context, id string) (name, content string, err error)  {
	var (
		resp interface{}
		response *responses.GetNoteResponse
	)

	resp, err = s.GetNoteEndpoint(ctx, requests.GetNoteRequest{
		NoteID: id,
	})

	if err != nil {
		return "", "", err
	}

	response = resp.(*responses.GetNoteResponse)
	return response.Name, response.Content, utils.Str2Err(response.Error)
}

func (s Set) ShareNote(ctx context.Context, name, content string) (url, noteID string, err error)  {
	var (
		resp interface{}
		response *responses.ShareNoteResponse
	)

	resp, err = s.ShareNoteEndpoint(ctx, requests.ShareNoteRequest{
		Name: name,
		Content: content,
	})

	if err != nil {
		return "", "", err
	}

	response = resp.(*responses.ShareNoteResponse)
	return response.URL, response.NoteID, utils.Str2Err(response.Error)
}

func (s Set) PrivateNote(ctx context.Context, id string) (err error)  {
	var (
		resp interface{}
		response *responses.PrivateNoteResponse
	)

	resp, err = s.PrivateNoteEndpoint(ctx, requests.PrivateNoteRequest{
		NoteID: id,
	})
	if err != nil {
		return  err
	}
	response = resp.(*responses.PrivateNoteResponse)
	return utils.Str2Err(response.Error)
}

func New(svc shareservice.Service, logger log.Logger, duration map[string]metrics.Histogram, tracer stdopentracing.Tracer) (set *Set, err error) {
	apis := config.GetConfig().Service.APIs

	set = &Set{
		ShareNoteEndpoint: MakeEndpoint(
			svc,
			apis[shareservice.ShareNoteServiceName],
			logger,
			duration[shareservice.ShareNoteServiceName],
			tracer,
			MakeShareNoteEndpoint),

		PrivateNoteEndpoint:    MakeEndpoint(
			svc,
			apis[shareservice.PrivateNoteServiceName],
			logger,
			duration[shareservice.PrivateNoteServiceName],
			tracer,
			MakePrivateNoteEndpoint),

		GetNoteEndpoint:    MakeEndpoint(
			svc,
			apis[shareservice.GetNoteServiceName],
			logger,
			duration[shareservice.GetNoteServiceName],
			tracer,
			MakeGetNoteEndpoint),
	}

	return
}

func MakeEndpoint(svc shareservice.Service, api bootapi.API, logger log.Logger, duration metrics.Histogram, tracer stdopentracing.Tracer, makeFN MakeEndpointFunc) endpoint.Endpoint {
	var ep endpoint.Endpoint
	{
		ep = makeFN(svc)
		// RegisterByEmail is limited to 1000 requests per second with burst of 1 request.
		// Note, rate is defined as a time interval between requests.
		ep = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(api.RateLimit.Duration), api.RateLimit.Delta))(ep)

		ep = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(api.Breaker.Standardize()))(ep)

		ep = LoggingMiddleware(log.With(logger, api.GetGoKitLoggerKVs()))(ep)
		ep = InstrumentingMiddleware(duration.With(api.Instrument...))(ep)

	}
	return ep
}

func MakeShareNoteEndpoint(svc shareservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var (
			req requests.ShareNoteRequest
			url, noteid string
			span stdopentracing.Span
		)

		span = stdopentracing.SpanFromContext(ctx)
		defer span.Finish()

		req = request.(requests.ShareNoteRequest)
		url, noteid, err = svc.ShareNote(ctx, req.Name, req.Content)
		if err != nil {

			return responses.ShareNoteResponse{
				Error: err.Error(),
			}, nil
		}

		return responses.ShareNoteResponse{
			URL:    url,
			NoteID: noteid,
			Error:    "",
		}, nil
	}
}

func MakeGetNoteEndpoint(svc shareservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var (
			req requests.GetNoteRequest
			name, content string
			span stdopentracing.Span
		)

		span = stdopentracing.SpanFromContext(ctx)
		defer span.Finish()

		req = request.(requests.GetNoteRequest)
		name, content, err = svc.GetNote(ctx, req.NoteID)
		if err != nil {
			return responses.GetNoteResponse{
				Error: err.Error(),
			}, nil
		}

		return responses.GetNoteResponse{
			Content:    content,
			Name: name,
			Error:    "",
		}, nil
	}
}

func MakePrivateNoteEndpoint(svc shareservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var (
			req requests.PrivateNoteRequest
		)

		req = request.(requests.PrivateNoteRequest)
		err = svc.PrivateNote(ctx, req.NoteID)
		if err != nil {
			return responses.PrivateNoteResponse{
				Error: err.Error(),
			}, nil
		}

		return responses.PrivateNoteResponse{
			Error:    "",
		}, nil
	}
}