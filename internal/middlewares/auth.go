package middlewares

import "github.com/labstack/echo/v4"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		if token == nil {
			return c.NoContent(echo.ErrForbidden.Code)
		}

		return next(c)
	}
}
