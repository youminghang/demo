package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/youminghang/go-gin-example/initialize"
	"github.com/youminghang/go-gin-example/pkg/setting"
	"github.com/youminghang/go-gin-example/pkg/util/register/consul"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	registry := consul.NewRegistry(setting.ServerConfig.ConsulInfo.Host, setting.ServerConfig.ConsulInfo.Port)
	serviceId := uuid.NewV4().String()
	registry.Register(setting.ServerConfig.Host, setting.ServerConfig.Port, setting.ServerConfig.Name, setting.ServerConfig.Tags, serviceId)

	zap.S().Debug("启动服务器，端口: ", setting.ServerConfig.Port)
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", setting.ServerConfig.Port)); err != nil {
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
