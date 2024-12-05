package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
}

// 结构体中的方法
func (p person) print() {
	p.name = "jhah"
	fmt.Println("名字为", p.name, "年龄为", p.age)
}
func main() {
	//新建对象
	p1 := new(person)
	//p3 := person{name: "ysy", age: 20}
	fmt.Println(*p1)
	p2 := p1
	p1.print()
	p2.age = 40
	p1.print()

	i1 := 1
	ptr := &i1
	fmt.Println(*ptr) // 1

}
