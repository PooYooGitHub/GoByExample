package main

import "fmt"

func main() {
	sum, a := plus(1, 2)
	fmt.Println(sum, a)
	fmt.Println(plusplus(1, 2, 3, 4))
	nums := []int{1, 2, 3, 4}
	fmt.Println("使用slice展开:", plusplus(nums...))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}

// 2. 多返回值
func plus(a int, b int) (int, int) {
	return a + b, a
}

// 1.可变参数，nums就是一个数组。
func plusplus(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 3. 闭包
// 函数名:intSeq,无参数。返回值(返回的这个函数是一个闭包)：func(),无参数，返回int
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
