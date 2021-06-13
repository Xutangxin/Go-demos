package main

import "fmt"

// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//可以自定义类型
type mystr string

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person //匿名字段
	id     int
	addr   string
	hobby  mystr //mystr自定义类型
}

type Woman struct {
	Person
	name string
}

type Man struct {
	*Person
	name string
}

func main() {
	p := Person{"jake", "male", 22}
	fmt.Println(p)

	//同时给匿名字段和自身字段赋值
	s := Student{Person{"jake", "male", 22}, 1, "America", "reading books"}
	fmt.Println(s)

	w := Woman{}
	//给自己字段赋值
	w.name = "lucy"
	//给父类同名字段赋值
	w.Person.name = "LUCY"

	w.Person.sex = "female"
	w.Person.age = 22
	fmt.Println(w)

	m := Man{&Person{"tom", "male", 22}, "tomson"}
	fmt.Println(m)
	fmt.Println(m.name)
	fmt.Println(m.Person.name)

}
