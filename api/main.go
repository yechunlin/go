package main

import (
	"api/middleware/logger"
	"api/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//初始化路由，启动服务
	logger.WriteLogStr("初始化路由,启动服务")
	//server := endless.NewServer(":8081", route.RouteInit())
	//server.BeforeBegin = func(add string) {
	//	util.FilePutContents("pid.txt", strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	//}
	//server.ListenAndServe()
	r := route.RouteInit()
	ser := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	go func() {
		err := ser.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {

		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Panicln("server shutDown...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ser.Shutdown(ctx)
	if err != nil {

	}
	log.Panicln("server exit!")
}
