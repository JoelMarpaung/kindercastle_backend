package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteBook func
//
//	@Summary		Delete book.
//	@Description	Delete book.
//	@Tags			book
//	@Accept			json
//	@Param			book_id	path	string	true	"book ID"
//	@Produce		json
//	@Success		204
//	@Router			/v1/books/{book_id} [delete]
func (h Handler) DeleteBook(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("book_id")
	)

	if err := h.Book.Delete(ctx, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
