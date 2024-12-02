package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	}

}
