package main

import "fmt"

type People struct {
	name string
	sex  string
	age  int
}

func main() {

	people := People{"jake", "man", 22}
	//声明并初始化一个结构体指针
	var pointer *People
	pointer = &people
	showPeopleInfo(pointer)

}

func showPeopleInfo(people *People) {
	//结构体指针可使用.访问结构体元素
	fmt.Printf("people name:%s\n", people.name)
	fmt.Printf("people sex:%s\n", people.sex)
	fmt.Printf("people age:%d\n", people.age)
}
