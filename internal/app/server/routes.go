package server

import (
	"kindercastle_backend/internal/app/handler/book"
	"kindercastle_backend/internal/app/handler/utils"
	"kindercastle_backend/internal/model/payload"
	"net/http"

	_ "kindercastle_backend/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title						KinderCastle API
//	@schemes					https http
//	@BasePath					/
//	@securityDefinitions.apikey	JWTToken
//	@in							header
//	@name						Authorization
//	@securityDefinitions.basic	FirebaseAuth
//
// Router list
func (srv *Server) initRoutes() {
	var (
		bookHandler    = book.New(srv.services)
		utilsHandler   = utils.New(srv.services)
		authMiddleware = NewMidleware(srv.services.FirebaseSvc)
	)

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

	v1 := srv.E.Group("/v1")

	v1utils := v1.Group("", authMiddleware.isAuthenticated)
	v1utils.POST("/image", utilsHandler.UploadImage)

	v1books := v1.Group("/books")
	v1books.Use(authMiddleware.isAuthenticated)
	v1books.POST("", bookHandler.CreateBook)
	v1books.GET("", bookHandler.ListBook)
	v1books.PUT("/:book_id", bookHandler.UpdateBook)
	v1books.GET("/:book_id", bookHandler.DetailBook)
	v1books.DELETE("/:book_id", bookHandler.DeleteBook)

}
