package main

import "fmt"

func main() {
	//map的值可以是函数
	m := map[int]func(val int) int{}
	m[1] = func(val int) int {
		return val
	}
	m[2] = func(val int) int {
		return val * val
	}

	fmt.Println(m[1](2), m[2](2))

	//使用map实现set
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	delete(mySet, 1)
	isExist(1, mySet)
	isExist(2, mySet)
	isExist(3, mySet)

}

//判断set中是否存在某个值
func isExist(val int, m map[int]bool) {
	if m[val] {
		fmt.Println(val, " is existing")
	} else {
		fmt.Println(val, "is not existing")
	}
}
