package main

import (
	"fmt"
	"goImooc/go-gin-test/models"
	"goImooc/go-gin-test/pkg/logging"
	"goImooc/go-gin-test/pkg/setting"
	"goImooc/go-gin-test/routers"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Hello, api 正在启动中...")
	setting.SetUp() //初始化配置文件
	logging.SetUp() //设置日志文件
	models.SetUp()  //设置数据库

	router := routers.InitRouter() //初始化路由

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort), //设置端口号
		Handler:        router,                                             //http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
		ReadTimeout:    setting.ServerSetting.ReadTimeout,                  //允许读取的最大时间
		WriteTimeout:   setting.ServerSetting.WriteTimeout,                 //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                                            //请求头的最大字节数
	}

	/*
	   使用 http.Server - Shutdown() 优雅的关闭http服务
	*/
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("程序服务关闭退出")
}
