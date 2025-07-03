package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func CustomLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		stop := time.Now()

		log.Printf("[CustomLog] %s %s %s (%s)",
			c.RealIP(),
			c.Request().Method,
			c.Request().URL.Path,
			stop.Sub(start),
		)

		return err
	}
}

