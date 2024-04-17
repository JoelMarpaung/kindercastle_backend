package firebase

import (
	"context"
	"kindercastle_backend/internal/app/repository"
	"kindercastle_backend/internal/pkg"

	"firebase.google.com/go/auth"
)

type Contract interface {
	VerifyToken(ctx context.Context, tokenId string) (*auth.Token, error)
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
