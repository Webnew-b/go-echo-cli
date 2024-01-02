package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/global/redisConn"
	"wscmakebygo.com/global/route"
	"wscmakebygo.com/start"
	"wscmakebygo.com/tools/logUtil"
)

func main() {
	start.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stopServe(ctx)
}

func stopServe(ctx context.Context) {
	route.StopEcho(ctx)
	database.CloseDatabase()
	redisConn.StopRedis()
	logUtil.Close()
}
