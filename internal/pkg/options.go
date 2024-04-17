package pkg

import (
	"github.com/jmoiron/sqlx"

	"kindercastle_backend/internal/app/config"
	"kindercastle_backend/internal/pkg/firebaseclient"
)

type Options struct {
	Config         *config.Config
	DB             *sqlx.DB
	FirebaseClient firebaseclient.FirebaseClient
}
