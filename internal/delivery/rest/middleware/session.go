package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CheckSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/")
			}

			// Check if user session exist
			_, ok := sess.Values["email"]
			if !ok {
				return c.Redirect(http.StatusTemporaryRedirect, "/")
			}

			// Check if session is correct
			requestURI := c.Request().RequestURI
			roleSession, ok := sess.Values["role"]
			if !ok {
				return c.Redirect(http.StatusTemporaryRedirect, "/")
			}
			role := fmt.Sprintf("%v", roleSession)
			if !strings.Contains(requestURI, role) {
				return c.Redirect(http.StatusTemporaryRedirect, "/")
			}
			
			return next(c)
		}
	}
}
