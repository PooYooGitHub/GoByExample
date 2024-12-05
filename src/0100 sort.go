package main

import (
	"fmt"
	"sort"
)

type desc []int

func (receiver desc) Len() int {
	return len(receiver)
}
func (receiver desc) Less(i, j int) bool {
	return receiver[i] > receiver[j]
}
func (receiver desc) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func main() {
	//sort默认升序排序
	strs := []int{1, 2, 3, 4}
	sort.Ints(strs)
	fmt.Println(strs)
	//自定义排序
	sort.Sort(desc(strs))
	fmt.Println(strs)
	//临时自定义排序
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	fmt.Println(strs)
}
