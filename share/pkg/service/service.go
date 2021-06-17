package service

import (
	"context"
	"github.com/al8n/shareable-notes/share/internal/repositories"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"
)

const (
	ShareNoteServiceName = "ShareNote"
	PrivateNoteServiceName = "PrivateNote"
	GetNoteServiceName = "GetNote"
)

type Service interface {
	ShareNote(ctx context.Context, name, content string) (url, sharedID string, err error)
	PrivateNote(ctx context.Context, id string) (err error)
	GetNote(ctx context.Context, id string) (name, content string, err error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger, counters map[string]metrics.Counter, tracer stdopentracing.Tracer) (svc Service,  err error) {

	svc, err = NewBasicService()
	if err != nil {
		return nil, err
	}

	svc = LoggingMiddleware(logger)(svc)
	svc = InstrumentingMiddleware(counters)(svc)
	svc = TracingMiddleware(tracer)(svc)

	return svc, nil
}

type basicService struct {
	repo *repositories.Repo
}

func (svc basicService) ShareNote(ctx context.Context, name, content string) (url, sharedID string, err error) {
	return svc.repo.ShareNote(ctx, name, content)
}

func (svc basicService) PrivateNote(ctx context.Context, id string) (err error) {
	return svc.repo.PrivateNote(ctx, id)
}

func (svc basicService) GetNote(ctx context.Context, id string) (name,content string, err error) {
	return svc.repo.GetNote(ctx, id)
}

func NewBasicService() (svc Service, err error ) {
	var repo *repositories.Repo

	repo, err = repositories.NewRepo()
	if err != nil {
		return
	}

	return &basicService{
		repo: repo,
	}, nil
}