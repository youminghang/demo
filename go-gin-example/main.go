package main

import (
	"fmt"
	"net/http"
	"time"

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

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(setting.ServerConfig.ReadTimeOut),
		WriteTimeout:   time.Duration(setting.ServerConfig.WriteTimeOut),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
