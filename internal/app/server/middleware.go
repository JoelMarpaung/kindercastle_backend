package server

import (
	"net/http"
	"strings"

	"kindercastle_backend/internal/app/service/firebase"

	"github.com/labstack/echo/v4"
)

type midware struct {
	firebaseSvc firebase.Contract
}

func NewMidleware(firebaseSvc firebase.Contract) midware {
	return midware{
		firebaseSvc: firebaseSvc,
	}
}

type ResponseMiddleware struct {
	Message string `json:"message"`
}

func (j *midware) isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			c.Response().Header().Set("WWW-Authenticate", "Bearer realm=\"Authorization Required\"")
			return c.JSON(http.StatusForbidden, ResponseMiddleware{
				Message: http.StatusText(http.StatusForbidden),
			})
		}
		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		if idToken == "" {
			return c.JSON(http.StatusUnauthorized, ResponseMiddleware{
				Message: "Please provide a valid token",
			})
		}

		ctx := c.Request().Context()

		token, err := j.firebaseSvc.VerifyToken(ctx, idToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired ID token")
		}

		c.Set("uid", token.UID)

		return next(c)
	}
}
