//go实现“继承”
//go是使用组合来实现的继承

package main

import "fmt"

//模拟动物行为的接口
type IAnimal interface {
	Eat()
	Bark()
}

//动物类（结构体）
type Animal struct {
	name string
	bark string
}

//动物类实现动物行为接口
func (a *Animal)Eat(){
	fmt.Printf("%s is eating food\n", a.name)
}

func (a *Animal)Bark(){
	fmt.Println(a.bark)
}

//动物类的构造函数
func newAnimal(name string, bark string)*Animal{
	return &Animal{
		name:name,
		bark:bark,
	}
}

//猫类（结构体）
type Cat struct {
	*Animal
}

//重写Eat()方法
func (cat *Cat)Eat(){
	fmt.Println("eating eating")
}

//猫的构造函数
//使用组合继承Animal类的属性和方法
func newCat(name string, bark string)*Cat {
	return &Cat{
		Animal:newAnimal(name, bark),
	}
}


func main() {
	animal :=newAnimal("animal", "wang~")
	fmt.Println(animal)
	animal.Eat()
	animal.Bark()

	cat :=newCat("mimi", "miao~")
	fmt.Println(cat)
	cat.Eat()
	cat.Bark()

}
