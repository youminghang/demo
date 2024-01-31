package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/youminghang/go-gin-example/initialize"
	"github.com/youminghang/go-gin-example/pkg/setting"
	"github.com/youminghang/go-gin-example/pkg/util/register/consul"
	"github.com/youminghang/go-gin-example/routers"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	r := routers.InitRouter()
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(setting.ServerConfig.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(setting.ServerConfig.WriteTimeOut) * time.Second,
	}
	registry := consul.NewRegistry(setting.ServerConfig.ConsulInfo.Host, setting.ServerConfig.ConsulInfo.Port)
	serviceId := uuid.NewV4().String()
	//registry.Register(setting.ServerConfig.Host, setting.ServerConfig.Port, setting.ServerConfig.Name, setting.ServerConfig.Tags, serviceId)
	registry.RegisterByUrl(setting.ServerConfig.Url, setting.ServerConfig.Name, setting.ServerConfig.Tags, serviceId)
	zap.S().Debug("启动服务器，端口: ", setting.ServerConfig.Port)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				zap.S().Info("HTTP 服务器已关闭")
			} else {
				zap.S().Error("HTTP 服务器错误: ", err.Error())
			}
		}
	}()
	quit := make(chan os.Signal)
	reload := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(reload, syscall.SIGHUP)

	for {
		select {
		case <-quit:
			zap.S().Info("正在退出服务")
			if err := registry.Deregister(serviceId); err != nil {
				zap.S().Info("consul注销失败 err：", err.Error())
			} else {
				zap.S().Info("consul注销成功")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				zap.S().Panic("server关闭失败:", err)
			}
			zap.S().Info("sever退出成功")
			return
		}
	}
}
