package main

import "fmt"

//定义接口
type Behavior interface {
	speak()
}

//定义结构体
type Person struct {
	name string
	age  int
	sex  string
}

//实现方法
func (person Person) speak() {
	fmt.Println("Hello, my name is", person.name)
}

func main() {
	tom := Person{"tom", 22, "male"}
	tom.speak() //Hello, my name is tom
}
