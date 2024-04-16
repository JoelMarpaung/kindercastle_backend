package custom

import (
	"net/http"
)

//go:generate easytags $GOFILE json

type FieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (fe FieldError) Error() string {
	return fe.Reason
}

type ErrorCustom struct {
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	Errors     []FieldError `json:"errors,omitempty"`
}

func (ce ErrorCustom) Error() string {
	return ce.Message
}

var (
	ErrForbiddenAccess       = &ErrorCustom{StatusCode: http.StatusForbidden, Message: "forbidden access"}
	ErrUnauthorized          = &ErrorCustom{StatusCode: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)}
	ErrInternalServer        = &ErrorCustom{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
	ErrInvalidToken          = &ErrorCustom{StatusCode: http.StatusBadRequest, Message: "failed to parse token"}
	ErrInvalidCredential     = &ErrorCustom{StatusCode: http.StatusUnprocessableEntity, Message: "mismatch credential"}
	ErrNoFileTypeAvailable   = &ErrorCustom{StatusCode: http.StatusBadRequest, Message: "invalid file type"}
	ErrNoFileSelected        = &ErrorCustom{StatusCode: http.StatusBadRequest, Message: "no file selected"}
	ErrMaxFileSize           = &ErrorCustom{StatusCode: http.StatusBadRequest, Message: "file uploaded exceed maximum size"}
	ErrInvalidImageExtension = &ErrorCustom{StatusCode: http.StatusBadRequest, Message: "file is not an image / image not supported"}

	ErrInvalidBook = &ErrorCustom{StatusCode: http.StatusNotFound, Message: "invalid book"}
)
