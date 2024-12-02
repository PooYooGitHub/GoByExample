package main

import "fmt"

func main() {
	s := make([]int, 8)
	le := []string{"1", "2"}
	//re := [...]string{"1", "2"}
	le = append(le, "2", "3")
	fmt.Println(le, "   ", cap(s), " ", len(s))
	fmt.Println(s)
	s = append(s, 2, 3)
	fmt.Println(s)
}
