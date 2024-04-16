package book

import (
	"kindercastle_backend/internal/app/service"
)

type Handler struct {
	*service.Container
}

func New(services *service.Container) Handler {
	return Handler{Container: services}
}
