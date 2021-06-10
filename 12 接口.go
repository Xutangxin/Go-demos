package main

import "fmt"

//定义接口
type Phone interface {
	call()
	getPrice()
}

//定义结构体
type HuaWei struct {
}

type IPhone struct {
}

//实现方法
func (huawei HuaWei) call() {
	fmt.Println("I am HuaWei, I am call you")
}
func (huawei HuaWei) getPrice() {
	fmt.Println("My price is 5000")
}

func (iphone IPhone) call() {
	fmt.Println("I am IPhone, I am call you")
}
func (iphone IPhone) getPrice() {
	fmt.Println("My price is 8000")
}

type People interface {
	eat()
	sleep()
	work()
}

type Man struct {
	name string
	age  int
	sex  string
}

func (man Man) eat() {
	fmt.Println("I am eating")
}
func (man Man) sleep() {
	fmt.Println("I am sleeping")
}
func (man Man) work() {
	fmt.Println("I am working")
}

func main() {
	var phone Phone
	phone = new(HuaWei)
	phone.call()
	phone.getPrice()

	phone = new(IPhone)
	phone.call()
	phone.getPrice()

	tom := Man{name: "tom", age: 22, sex: "male"}
	fmt.Println(tom)
	tom.eat()
	tom.sleep()
	tom.work()

}
