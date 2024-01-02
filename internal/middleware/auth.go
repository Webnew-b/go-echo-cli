package middleware

import (
	"github.com/labstack/echo/v4"
)

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//todo 用户校验逻辑
		return next(c)
	}
}
