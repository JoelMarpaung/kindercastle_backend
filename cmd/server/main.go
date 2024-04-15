package server

import (
	"context"

	"kindercastle_backend/internal/app/config"
	"kindercastle_backend/internal/app/server"
)

func Start(ctx context.Context) {
	conf := config.GetConfig()

	s := server.NewHTTPServer(&conf)
	defer s.Stop()

	s.Serve()
}
