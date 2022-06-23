package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("sdsd"))
		})
		router.Run(":8080")
	}()
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080") //exec开得是一个进程 命令行任务
	cmd.Start()
	<-chSignal //没有值就会阻塞  通过channel监听系统信号 防止僵尸进程：main程结束了，chrome经常还在
	cmd.Process.Kill()
}
