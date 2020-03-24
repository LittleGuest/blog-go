package main

import (
	"blog/consts"
	"blog/repo/db"
	"blog/routers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db.InitDB()

	// gin.DisableConsoleColor()
	// logFile, err := os.OpenFile(ServerLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// if err != nil {
	// 	log.Fatalf("创建日志文件失败： %s", err)
	// }
	// defer logFile.Close()
	// gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	// log.SetOutput(logFile)

	router := routers.InitRouter()

	server := &http.Server{
		Addr:         ":65530",
		Handler:      router,
		ReadTimeout:  consts.ReadTimeout,
		WriteTimeout: consts.WriteTimeout,
		IdleTimeout:  consts.IdleTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err == http.ErrServerClosed {
			log.Fatalf("服务启动失败: %s", err)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	log.Println("关闭服务。。。")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务关闭异常: %s", err)
	}
	log.Println("服务安全退出。。。")
}
