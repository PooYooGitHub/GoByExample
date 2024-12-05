package main

import (
	"fmt"
	"time"
)

func besides(messages chan string) {
	for {
		select {
		case message, ok := <-messages:
			if ok {
				fmt.Println(message)
			} else {
				return
			}
		default:
			// 没有消息时执行的操作
			fmt.Println("No message received")
			time.Sleep(time.Second) // 防止忙等待
		}
	}
}

func main() {
	messages := make(chan string)
	go besides(messages)

	// 模拟发送消息
	time.Sleep(2 * time.Second)
	messages <- "hello goroutine"
	close(messages)

	// 等待 goroutine 处理完所有消息
	time.Sleep(3 * time.Second)
}
