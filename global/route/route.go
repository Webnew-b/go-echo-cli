package route

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"sync"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/tools/logUtil"
)

var (
	router *echo.Echo
	once   sync.Once
)

func GetRoute() *echo.Echo {
	if router == nil {
		panic("database not initialized")
	}
	return router
}

func StopEcho(ctx context.Context) {
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}
}

func StartRoute() {
	router.Logger.Fatal(router.Start(createServerAddr()))
}

func InitVal() {
	once.Do(func() {
		log.Println("starting http Server")
		router = echo.New()
		router.Logger.SetOutput(logUtil.GetEchoLogFile())
	})
}

func createServerAddr() string {
	serveAddr := fmt.Sprintf("%s:%d",
		envConfig.GetConfig().App.Host,
		envConfig.GetConfig().App.Port)
	return serveAddr
}
