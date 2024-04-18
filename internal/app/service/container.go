package service

import (
	"kindercastle_backend/internal/app/service/book"
	"kindercastle_backend/internal/app/service/firebase"
	"kindercastle_backend/internal/app/service/utils"
)

type Container struct {
	Book        book.IService
	FirebaseSvc firebase.Contract
	Utils       utils.IService
}
