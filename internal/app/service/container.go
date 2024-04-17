package service

import (
	"kindercastle_backend/internal/app/service/book"
	"kindercastle_backend/internal/app/service/firebase"
)

type Container struct {
	Book        book.IService
	FirebaseSvc firebase.Contract
}
