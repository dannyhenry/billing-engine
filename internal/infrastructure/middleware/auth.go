package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func TokenAuth(expectedToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "missing authorization header",
				})
			}

			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid authorization header format",
				})
			}

			token := parts[1]

			if token != expectedToken {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid token",
				})
			}

			return next(c)
		}
	}
}
