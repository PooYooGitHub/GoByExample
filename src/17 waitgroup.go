package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	//主线程等待多个线程同步；如果是等待单线程直接通道就可以以了
	wg.Wait()
	fmt.Println("All workers completed.")
}
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// 模拟工作
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)

}
