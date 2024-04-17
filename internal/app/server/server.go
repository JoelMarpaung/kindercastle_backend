package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"kindercastle_backend/internal/app/repository"
	"kindercastle_backend/internal/app/repository/book"
	bookSvc "kindercastle_backend/internal/app/service/book"
	"kindercastle_backend/internal/app/service/firebase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"kindercastle_backend/internal/app/appcontext"
	"kindercastle_backend/internal/app/config"
	"kindercastle_backend/internal/app/service"
	"kindercastle_backend/internal/pkg"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
)

type Server struct {
	E        *echo.Echo
	services *service.Container
	conf     *config.Config
	opt      *pkg.Options
}

func NewHTTPServer(conf *config.Config) Server {
	srv := Server{
		E:    echo.New(),
		conf: conf,
	}

	// don't change the init order
	srv.initContainers()
	srv.initMiddleware()
	srv.initValidator()
	srv.initRoutes()

	srv.E.HTTPErrorHandler = custom.EchoCustomErrorHandler
	return srv
}

func (srv *Server) Serve() {
	go func() {
		if err := srv.E.Start(fmt.Sprintf(":%s", srv.conf.AppPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			srv.E.Logger.Fatal("shutting down the server")
		}
	}()
	defer srv.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func (srv *Server) Stop() {
	//nolint:gomnd // expected
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = srv.opt.DB.Close()

	if err := srv.E.Shutdown(ctx); err != nil {
		srv.E.Logger.Fatal(err)
	}
}

func (srv *Server) initValidator() {
	srv.E.Validator = custom.NewValidator()
}

func (srv *Server) initMiddleware() {
	srv.E.Use(middleware.RequestID())
	srv.E.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogMethod:    true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogUserAgent: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// ignore docs page
			if strings.HasPrefix(v.URI, "/docs/") {
				return nil
			}

			logging.
				Logger().
				Info().
				Str("time", v.StartTime.String()).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("URI", v.URI).
				Str("method", v.Method).
				Str("user_agent", v.UserAgent).
				Int("status", v.Status).
				Str("latency", v.Latency.String()).
				Msg("")
			return nil
		},
	}))
	srv.E.Use(middleware.Recover())
	srv.E.Use(middleware.CORS())
}

func (srv *Server) initContainers() {
	appCtx := appcontext.NewAppContext(srv.conf)

	opts := &pkg.Options{
		Config:         srv.conf,
		DB:             appCtx.GetDBConnection(),
		FirebaseClient: appCtx.GetFirebaseCLient(),
	}

	repositories := &repository.Container{
		Book: book.New(opts),
	}

	services := &service.Container{
		Book:        bookSvc.New(opts, repositories),
		FirebaseSvc: firebase.New(opts, repositories),
	}

	srv.opt = opts
	srv.services = services
}
