package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/global/route"
)

func hookExampleRoute(api *echo.Group) {
	api.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
}

func HookRoute() {
	var api = route.GetRoute().Group("/api")
	hookExampleRoute(api)
}
