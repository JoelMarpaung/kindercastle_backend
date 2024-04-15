package custom

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ResponseError[T any] struct {
	Error T `json:"error"`
}

func EchoCustomErrorHandler(err error, c echo.Context) {
	var (
		code            = http.StatusInternalServerError
		message         string
		errs            []FieldError
		tpErrCustom     *ErrorCustom
		tpErrHttp       *echo.HTTPError
		tpErrValidation validator.ValidationErrors
	)

	switch {
	case errors.As(err, &tpErrCustom):
		message = tpErrCustom.Message
		code = tpErrCustom.StatusCode
	case errors.As(err, &tpErrValidation):
		code = http.StatusBadRequest
		message = "Validation Error"
		for _, valErr := range tpErrValidation {
			errs = append(errs, FieldError{Field: valErr.Field(), Reason: messageForTag(valErr)})
		}
	case errors.As(err, &tpErrHttp):
		code = tpErrHttp.Code
		message = fmt.Sprint(tpErrHttp.Message)
	default:
		message = err.Error()
	}

	errResp := ResponseError[ErrorCustom]{
		Error: ErrorCustom{
			StatusCode: code,
			Message:    message,
		},
	}

	if len(errs) > 0 {
		errResp.Error.Errors = errs
	}

	_ = c.JSON(code, errResp)
}

func messageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "base64", "datauri":
		return "Invalid base64 format"
	case "required":
		return "This field is required"
	case "required_with":
		return fmt.Sprintf("This field is required with %s", fe.StructField())
	case "max":
		return fmt.Sprintf("Reach maximum %s", fe.Param())
	case "min":
		return fmt.Sprintf("Does not meet minimum %s", fe.Param())
	case "invalid_image":
		return fmt.Sprintf("invalid image in %s", fe.Field())
	case "unsupported_image_type":
		return fmt.Sprintf("file is not an image / unsupported type : %s", fe.Param())
	}
	return fe.Error()
}
