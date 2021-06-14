package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

//接口嵌套
type Animal interface {
	Sayer
	Mover
}

type Cat struct {
}

func (cat Cat) say() {
	fmt.Println("miao~")
}

func (cat Cat) move() {
	fmt.Println("the cat is moving")
}

func show(a string) {
	fmt.Println(a)
}

//空接口作为函数参数
//空接口可以接受任何类型的参数
func showData(a interface{}) {
	fmt.Println(a)
}

func main() {
	// cat := Cat{}
	// cat.say()
	// cat.move()

	//空接口
	//空接口可以存储任何类型的变量
	var x interface{}
	x = "hello"
	x = false
	fmt.Println(x)

	test := "fsociety"
	test2 := []int{1, 2, 3, 4, 5}
	show(test)
	showData(test)
	showData(test2)

}
