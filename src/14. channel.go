package main

import (
	"fmt"
)

func main() {
	//通道 默认阻塞式
	//我需要main等待携程执行完才能退出
	channel := make(chan string)
	waiting := make(chan bool)
	go func() {
		fmt.Println("开始receive")
		fmt.Println(<-channel)
		waiting <- true

	}()
	channel <- "消息"
	<-waiting
}
