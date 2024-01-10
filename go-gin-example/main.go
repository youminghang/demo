package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	registry := consul.NewRegistry(setting.ServerConfig.ConsulInfo.Host, setting.ServerConfig.ConsulInfo.Port)
	serviceId := uuid.NewV4().String()
	//registry.Register(setting.ServerConfig.Host, setting.ServerConfig.Port, setting.ServerConfig.Name, setting.ServerConfig.Tags, serviceId)
	registry.RegisterByUrl(setting.ServerConfig.Url, setting.ServerConfig.Name, setting.ServerConfig.Tags, serviceId)
	zap.S().Debug("启动服务器，端口: ", setting.ServerConfig.Port)
	go func() {
		if err := r.Run(fmt.Sprintf(":%d", setting.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := registry.Deregister(serviceId); err != nil {
		zap.S().Info("注销失败 err：", err.Error())
	}
	zap.S().Info("注销成功")

}
