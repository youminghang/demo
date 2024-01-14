package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"os/exec"

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
		case <-reload:
			zap.S().Info("正在重启服务")
			if err := forkProcess(); err != nil {
				zap.S().Info("重启失败 err：", err.Error())
			}
		}
	}

}

func forkProcess() error {
	// 获取当前可执行文件的路径
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	// 准备新进程的环境变量和参数
	env := os.Environ()
	args := os.Args
	cmd := exec.Command(exe, args...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 为新进程设置额外的文件描述符
	// 这可以用于继承如网络监听套接字等资源
	cmd.ExtraFiles = []*os.File{}

	// 启动新进程
	if err := cmd.Start(); err != nil {
		return err
	}

	// 在此处等待，确保新进程有足够的时间启动
	time.Sleep(time.Second * 5)

	// 安全退出当前进程
	// 注意: 这里不使用syscall.Kill，因为我们不想立即杀死当前进程
	// 而是让它有机会完成正在处理的请求
	if ppid := os.Getppid(); ppid != 1 {
		syscall.Kill(ppid, syscall.SIGTERM)
	}

	return nil
}
