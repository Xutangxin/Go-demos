package main

import "fmt"

func main() {
	//变量定义
	//定义变量方式1
	var a string = "hello world"
	//定义变量方式2
	b := 1
	//可以组合定义变量
	var (
		c string = "nice to meet you"
		d bool   = false
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("------------------------------------------------------")

	//常量定义
	const name string = "jake"
	const age int = 22
	const height float32 = 1.7
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	//快速设置连续值
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday)

	//实现斐波拉契数列
	fmt.Println("斐波拉契数列")
	first := 1
	second := 1
	fmt.Print(first)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", second)
		temp := first
		first = second
		second = temp + first
	}
	fmt.Println()

}
