package start

import (
	"log"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/global/redisConn"
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/internal/middleware"
	"wscmakebygo.com/internal/router"
	"wscmakebygo.com/tools/logUtil"
)

func Init() {
	log.Println("Configuration is initializing")
	envConfig.InitVal()
	logUtil.CreateLogger()
	logUtil.Log.Println("Server is Starting")
	database.InitVal()
	redisConn.InitVal()
	go routeInit()
	logUtil.Log.Println("Server is Started")
	log.Println("Server is Started")
}

func routeInit() {
	route.InitVal()
	router.HookRoute()
	middleware.HookMiddleware()
	route.StartRoute()
}

func StartDbConnect() {
	envConfig.InitVal()
	database.InitVal()
}

func StopServe() {
	logUtil.Log.Println("stopping Server")
	// todo 停止日志读写以及清退存在的协程
}
