package pkg

import (
	"github.com/jmoiron/sqlx"

	"kindercastle_backend/internal/app/config"
)

type Options struct {
	Config *config.Config
	DB     *sqlx.DB
}
