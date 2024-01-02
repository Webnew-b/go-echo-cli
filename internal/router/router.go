package router

import (
	"github.com/labstack/echo/v4"
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/internal/controller/attendeesController"
	"wscmakebygo.com/internal/controller/eventsController"
	"wscmakebygo.com/internal/middleware"
)

func hookEventsRoute(api *echo.Group) {
	api.GET("/events", eventsController.GetEvents)
}

func hookEventDetailRoute(api *echo.Group) {
	api.GET("/organizers/:organizerSlug/events/:eventSlug", eventsController.GetEventDetail)
}

func hookFetchEventReg(api *echo.Group) {
	api.GET("/registrations",
		middleware.WithMiddleware(eventsController.FetchEventReg, middleware.UserAuthMiddleware))
}

func hookEventReg(api *echo.Group) {
	api.POST("/organizers/:organizerSlug/events/:eventSlug/registration",
		middleware.WithMiddleware(eventsController.EventReg, middleware.UserAuthMiddleware))
}

func hookLoginRoute(api *echo.Group) {
	api.POST("/login", attendeesController.AttendeesLogin)
}

func hookLogoutRoute(api *echo.Group) {
	api.POST("/logout", attendeesController.AttendeesLogout)
}

func HookRoute() {
	var api = route.GetRoute().Group("/api/v1")
	hookEventsRoute(api)
	hookEventDetailRoute(api)
	hookLoginRoute(api)
	hookLogoutRoute(api)
	hookEventReg(api)
	hookFetchEventReg(api)
}
