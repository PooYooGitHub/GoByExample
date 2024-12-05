package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	waiting := make(chan bool)
	channel <- "h"
	channel <- "e"
	channel <- "l"
	go func() {
		for i := range channel {
			fmt.Print(i)
		}
		fmt.Println()
		fmt.Print("遍历结束")
		<-waiting
	}()
	close(channel)
	waiting <- true

}
