package middleware

import (
	"github.com/labstack/echo/v4"
	"time"
	"wscmakebygo.com/tools/logUtil"
)

func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		err := next(c)

		request := c.Request()
		response := c.Response()
		logUtil.RouteLog.Printf("%s %s %d %s",
			request.Method,
			request.URL,
			response.Status,
			time.Since(startTime).String(),
		)

		return err
	}
}
