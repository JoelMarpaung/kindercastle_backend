package service

import (
	"kindercastle_backend/internal/app/service/book"
)

type Container struct {
	Book book.IService
}
