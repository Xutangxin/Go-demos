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

//类型断言 判断空接口类型函数
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("the type of x is string, the value is ", v)
	case int:
		fmt.Println("the type of x is int, the value is ", v)
	case bool:
		fmt.Println("the type of x is bool, the value is ", v)
	default:
		fmt.Println("unsupport type", v)
	}

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
	//类型断言，判断空接口存储类型的值
	v, ok := x.(bool)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	justifyType(x)

	test := "fsociety"
	test2 := []int{1, 2, 3, 4, 5}
	show(test)
	showData(test)
	showData(test2)

}
