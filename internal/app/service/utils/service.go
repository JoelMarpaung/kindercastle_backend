package utils

import (
	"context"
	"mime/multipart"

	"kindercastle_backend/internal/app/repository"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg"
)

type IService interface {
	UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (payload.UploadImageResponse, error)
}

type Service struct {
	*pkg.Options
	*repository.Container
}

func New(opts *pkg.Options, repos *repository.Container) Service {
	return Service{
		Options:   opts,
		Container: repos,
	}
}
