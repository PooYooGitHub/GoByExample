package main

import "fmt"

func main() {
	nums := make([]int, 2)
	nums[0] = 1
	nums[1] = 2

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, c := range "hello" {
		fmt.Println(i, c)
	}

}
