// 关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
// 不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Cat struct {
}

type Dog struct {
}

func (c Cat) say() { //值接受实现接口
	fmt.Println("miao")
}

func (d Dog) say() { //值接受实现接口
	fmt.Println("wang")
}

//一个类型可以同时实现多个接口
func (d *Dog) move() { //指针接受实现接口
	fmt.Println("dog is moving")
}

//多个类型也可以实现同一个接口
//我们定义一个Car类型来实现move接口，上面的Dog也实现了move接口
type Car struct {
	brand string
}

func (car Car) move() {
	fmt.Printf("%s is moving", car.brand)
}

func main() {
	cat := Cat{}
	dog := Dog{}

	cat.say()
	dog.say()

	var sayer Sayer
	sayer = cat
	sayer = dog
	sayer.say()

	cat1 := Cat{}
	cat2 := &Cat{}
	cat1.say()
	cat2.say()

	var mover Mover
	wangcai := &Dog{}
	mover = wangcai
	fugui := &Dog{}
	mover = fugui
	mover.move()

	car := Car{"奥迪"}
	fmt.Println(car)
	car.move()

}
