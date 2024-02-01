package util

import (
	"fmt"
	"os"
	"syscall"

	"go.uber.org/zap"
)

func GenSigHupSignal() {
	pid := os.Getpid()
	// 打开当前进程
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Error finding process: %v\n", err)
		return
	}
	// 向自身进程发送 SIGHUP 信号
	err = process.Signal(syscall.SIGHUP)
	if err != nil {
		fmt.Printf("Error sending SIGHUP signal: %v\n", err)
		return
	}
	zap.S().Info("SigHup信号发送成功")
}
