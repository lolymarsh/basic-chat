package middleware

import (
	"github.com/centrifugal/centrifuge"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		ctx := req.Context()

		cred := &centrifuge.Credentials{
			UserID: "",
		}

		newCtx := centrifuge.SetCredentials(ctx, cred)
		req = req.WithContext(newCtx)
		c.SetRequest(req)

		return next(c)
	}
}
