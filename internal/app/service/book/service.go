package book

import (
	"context"

	"kindercastle_backend/internal/app/repository"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg"
)

type IService interface {
	Create(ctx context.Context, data payload.CreateBookPayload) error
	Edit(ctx context.Context, data payload.EditBookPayload) error
	Delete(ctx context.Context, bookID string) error
	Detail(ctx context.Context, bookID string) (payload.Book, error)
	GetAll(ctx context.Context, param payload.PagingAndFilterPayload) ([]payload.Book, int64, error)
}

type service struct {
	*pkg.Options
	*repository.Container
}

func New(opts *pkg.Options, repos *repository.Container) IService {
	return service{
		Options:   opts,
		Container: repos,
	}
}
