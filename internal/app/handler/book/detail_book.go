package book

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"kindercastle_backend/internal/model/payload"
)

// DetailBook func
//
//	@Summary		Get detail book.
//	@Description	Get detail book.
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			book_id	path		string	true	"book ID"
//	@Success		200		{object}	payload.ResponseDataDetailBook
//	@Router			/v1/books/{book_id} [get]
func (h Handler) DetailBook(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("book_id")
	)

	res, err := h.Book.Detail(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, payload.ResponseData[any]{Data: res})
}
