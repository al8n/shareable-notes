package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
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


