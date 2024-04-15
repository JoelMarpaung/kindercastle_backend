package server

import (
	"kindercastle_backend/internal/model/payload"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title						KinderCastle API
//	@schemes					https http
//	@BasePath					/
//	@securityDefinitions.basic	BasicAuth
//
// Router list
func (srv *Server) initRoutes() {

	srv.E.GET("/", func(c echo.Context) error {
		resp := payload.ResponseData[map[string]string]{
			Data: map[string]string{
				"message": "OK",
			},
		}
		return c.JSON(http.StatusOK, resp)
	})

	if srv.conf.EnableDocs {
		srv.E.GET("/docs/*", echoSwagger.WrapHandler)
	}

}
