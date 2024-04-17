package book

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"kindercastle_backend/internal/model/payload"
)

// CreateBook func
//
//	@Summary		Create a new book.
//	@Description	Create a new book.
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	payload.CreateBookPayload	true	"Create a new book"
//	@Success		201
//	@Router			/v1/books [post]
func (h Handler) CreateBook(c echo.Context) error {
	var (
		body   payload.CreateBookPayload
		ctx    = c.Request().Context()
		userID = c.Get("uid").(string)
	)

	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(&body); err != nil {
		return err
	}

	if err := h.Book.Create(ctx, body, userID); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
