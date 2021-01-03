package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Middleware func(Service) Service


type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{logger: logger, next: next}
	}
}

func (mw loggingMiddleware) PrivateNote(ctx context.Context, id string) (err error) {
	defer func() {
		mw.logger.Log("method", "PrivateNote", "id", id, "err", err)
	}()
	return mw.next.PrivateNote(ctx, id)
}

func (mw loggingMiddleware) ShareNote(ctx context.Context, name, content string) (url, shareID string, err error)  {
	defer func() {
		mw.logger.Log("method", "ShareNote", "name", name, "err", err)
	}()
	return mw.next.ShareNote(ctx, name, content)
}

func (mw loggingMiddleware) GetNote(ctx context.Context, id string) (name, content string, err error) {
	defer func() {
		mw.logger.Log("method", "GetNote", "id", id, "err", err)
	}()
	return mw.next.GetNote(ctx, id)
}


type instrumentingMiddleware struct {
	ctrs map[string]metrics.Counter
	next  Service
}

func (mw instrumentingMiddleware) PrivateNote(ctx context.Context, id string) (err error) {
	err = mw.next.PrivateNote(ctx, id)
	mw.ctrs[PrivateNoteServiceName].Add(1)
	return
}

func (mw instrumentingMiddleware) ShareNote(ctx context.Context, name, content string) (url, shareID string, err error) {
	url, shareID, err = mw.next.ShareNote(ctx, name, content)
	mw.ctrs[ShareNoteServiceName].Add(1)
	return
}

func (mw instrumentingMiddleware) GetNote(ctx context.Context, id string) (name, content string, err error)  {
	name, content, err = mw.next.GetNote(ctx, id)
	mw.ctrs[GetNoteServiceName].Add(1)
	return
}

func InstrumentingMiddleware(ctrs map[string]metrics.Counter) Middleware  {
	return func(next Service) Service {
		return instrumentingMiddleware{
			ctrs: ctrs,
			next:  next,
		}
	}
}

type tracerMiddleware struct {
	tracer stdopentracing.Tracer
	next Service
}

func TracingMiddleware(tracer stdopentracing.Tracer) Middleware {
	return func(next Service) Service {
		return tracerMiddleware{
			tracer: tracer,
			next:  next,
		}
	}
}

func (mw tracerMiddleware) PrivateNote(ctx context.Context, id string) (err error) {
	var (
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, "Private Note Service")
	defer span.Finish()

	span.SetTag(string(ext.Component), "ServerMiddleware")
	span.SetTag("id", id)

	err = mw.next.PrivateNote(spanCtx, id)
	span.LogKV("error", err)
	return
}

func (mw tracerMiddleware) ShareNote(ctx context.Context, name, content string) (url, shareID string, err error) {
	var (
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, "Share Note Service")
	defer span.Finish()

	url, shareID, err = mw.next.ShareNote(spanCtx, name, content)
	span.SetTag("url", url)
	span.LogKV("error", err)
	return
}

func (mw tracerMiddleware) GetNote(ctx context.Context, id string) (name, content string, err error)  {
	var (
		span stdopentracing.Span
		spanCtx context.Context
	)

	span, spanCtx = stdopentracing.StartSpanFromContext(ctx, "Get Note Service")
	defer span.Finish()

	name, content, err = mw.next.GetNote(spanCtx, id)
	span.SetTag("name", name)
	span.LogKV("error", err)
	return
}
