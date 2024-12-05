package main

import "fmt"

func main() {
	//字符串类型值不能改变，但是引用可以改变
	str1 := "haha"
	str1 = "hello world"
	str2 := "ni好"
	//字符串的本质是UTF-8编码的字节序列，而‘h’占一个字节的大小
	//'好'占三个字节
	fmt.Println(len(str1), len(str2))

	//遍历
	for _, str := range str1 {
		fmt.Print(str)
		fmt.Print(" ")
	}
	//子串，左闭右开
	fmt.Println(str1[0:1]) // h
	//长度，字符串所占的字节数。使用UTF-8编码
	fmt.Println(len("你好")) // 6
	//遍历
	for _, c := range str1 {
		fmt.Printf("%c", c)
	} // hello world
}
