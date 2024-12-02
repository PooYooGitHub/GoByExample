package main

import "fmt"

func main() {
	A := make(map[int]string)
	A[1] = "s"
	A[2] = "y"
	delete(A, 1)
	fmt.Println(A[1], A[2])
	fmt.Println(A)
}
