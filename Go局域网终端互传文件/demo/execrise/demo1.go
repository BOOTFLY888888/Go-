package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zserge/lorca"
)

//这段代码的好在哪？
// 1。关闭UI后 主线程会自动退出
// 2. 中断主线程后 UI会自动退出
func main() {
	//声明ui变量是lorca.UI接口类型，那么ui就能够调用该接口的方法
	var ui lorca.UI
	//使用new开启一个chrome窗口，取消了同步跟翻译功能，返回的ui代表着对这个窗口的所有操作
	ui, _ = lorca.New("https://baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")
	// 创建一个接收系统信号的channel，1代表缓存必须写上
	chSignal := make(chan os.Signal, 1)
	// 使用notify，接收os.Signal类型的channel  后面的参数是os.signal类型  （中断信号和中止信号）
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	//select监听选择UI的信号或者选择系统的信号执行     select随机轮询，会【等待】第一个可以读或写的ch进行【操作】 你没操作，我就干等
	//只要select语句中有一个出现了，就把ui关闭。代码全部执行完默认return结束main主进程
	// select{}含义：挂起forever 睡觉，不消耗资源 for{} 死循环，但会消耗资源
	select {
	case <-ui.Done(): //ui是个接口，通过Done（）方法，返回一个channel
	case <-chSignal:
	}

	ui.Close()
}
