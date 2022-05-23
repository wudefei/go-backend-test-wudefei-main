package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"user-api/internal/config"
	"user-api/internal/db"
	"user-api/internal/db/redis"
	"user-api/internal/router"
	"user-api/internal/util"
)

func main() {
	// 释放资源
	debug.SetGCPercent(80)

	logID := util.Uniqid()
	addr := config.Config.GinHost

	engine := router.InitRouter()

	if config.Config.Env != "MOCK" {
		//db初始化
		db.Init()

		//redis初始化
		redis.InitRedis()
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	SignalHandler(logID, server)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(logID, "ListenAndServe", err)
	}
}

// SignalHandler 优雅退出
func SignalHandler(logID string, server *http.Server) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP, syscall.SIGABRT,
		syscall.SIGBUS, syscall.SIGFPE, syscall.SIGKILL, syscall.SIGSEGV, syscall.SIGPIPE, syscall.SIGALRM, syscall.SIGTERM)
	go func() {
		s := <-c //阻塞等待
		fmt.Println(logID, "收到信号:"+s.String())
		if s == syscall.SIGTERM || s == syscall.SIGINT {
			maxSecond := time.Duration(60)
			ctx, cancel := context.WithTimeout(context.Background(), maxSecond*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				fmt.Println(logID, maxSecond.String()+"超时，非平滑关闭")
				os.Exit(0)
			}
			fmt.Println(logID, "服务已优雅退出")
			os.Exit(0)
		}
	}()
}
