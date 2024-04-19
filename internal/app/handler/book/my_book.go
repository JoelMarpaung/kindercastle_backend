package book

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"

	"kindercastle_backend/internal/model/payload"
)

// MyBook func
//
//	@Summary		Get list paginated book.
//	@Description	Get list paginated book.
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			search	query		string	false	"search query"
//	@Param			limit	query		int		false	"limit"
//	@Param			page	query		int		false	"page"
//	@Success		200		{object}	payload.ResponseListBook
//	@Router			/v1/books/me [get]
func (h Handler) MyBook(c echo.Context) error {
	var (
		ctx    = c.Request().Context()
		param  payload.PagingAndFilterPayload
		userID = c.Get("uid").(string)
	)

	if err := c.Bind(&param); err != nil {
		return err
	}

	param.Normalize()

	items, count, err := h.Book.GetMyBook(ctx, param, userID)
	if err != nil {
		return err
	}

	if items == nil {
		items = []payload.Book{}
	}

	return c.JSON(http.StatusOK, payload.PaginatedResponse[payload.Book]{
		Page:      param.Page,
		Limit:     param.Limit,
		Search:    param.Search,
		TotalRow:  int(count),
		TotalPage: int(math.Ceil(float64(count) / float64(param.Limit))),
		Items:     items,
	})
}
