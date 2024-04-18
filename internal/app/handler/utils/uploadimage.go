package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Upload Image
//
//	@Description	upload image
//	@Tags			utils
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			payload	body		payload.UploadImage	true	"Upload Image File"
//	@Success		200		{object}	payload.UploadImageResponse
//	@Router			/v1/image [post]
func (h Handler) UploadImage(c echo.Context) error {
	ctx := c.Request().Context()

	image, err := c.FormFile("image")
	if err != nil {
		return err
	}

	result, err := h.Utils.UploadImage(ctx, image)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
