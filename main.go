package main

import (
	"context"
	"gin-demo/config"
	models "gin-demo/model"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	log.Println(config.DatabaseConfig.Port)

	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库基础数据初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	r := router.InitRouter()
	srv := &http.Server{
		Addr: config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil  && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server Run ", config.ApplicationConfig.Host+":"+config.ApplicationConfig.Port)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}