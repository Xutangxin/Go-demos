package main

import "fmt"

//定义结构体
//结构体允许我们定义一系列不同类型的变量在同一集合中
type People struct {
	name string
	sex  string
	age  int
}

func main() {

	//初始化一个结构体
	//方式1
	a := People{"xutangxin", "man", 22}
	//方式2
	b := People{name: "xutangxin", sex: "man", age: 22}
	fmt.Println(a)
	fmt.Println(b)

	//访问结构体元素
	fmt.Println(a.name, a.sex, a.age)

	people := People{"jake", "man", 33}
	showInfo(people)

}

//结构体作为函数参数
func showInfo(a People) {
	fmt.Println(a)
}
