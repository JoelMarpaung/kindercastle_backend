package book

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"kindercastle_backend/internal/model/payload"
)

// UpdateBook func
//
//	@Summary		Update the book.
//	@Description	Update the book.
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	payload.EditBookPayload	true	"Edit book"
//	@Param			book_id	path	string					true	"book ID"
//	@Success		200
//	@Router			/v1/books/{book_id} [put]
func (h Handler) UpdateBook(c echo.Context) error {
	var (
		body   payload.EditBookPayload
		ctx    = c.Request().Context()
		id     = c.Param("book_id")
		userID = c.Get("uid").(string)
	)

	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	body.ID = id
	if err := h.Book.Edit(ctx, body, userID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
