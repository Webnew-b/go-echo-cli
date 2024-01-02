package middleware

import (
	"github.com/labstack/echo/v4"
	"wscmakebygo.com/global/route"
)

func WithMiddleware(h echo.HandlerFunc, middleware echo.MiddlewareFunc) echo.HandlerFunc {
	return middleware(h)
}

func HookMiddleware() {
	var hook = route.GetRoute()
	hook.Use(logMiddleware)
}
