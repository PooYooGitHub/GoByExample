package main

import "fmt"

func main() {
	//defer一定会执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("程序被恢复了:", r)
		}
	}()
	panic("this is a panic")

	fmt.Println("这行代码不会执行，因为已经触发了 panic")
}
