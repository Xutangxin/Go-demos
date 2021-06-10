package main

import "fmt"

func main() {

	a := 2333
	//声明指针
	var pointer *int
	//为指针赋值
	pointer = &a
	fmt.Println("a 的地址是", &a)
	//使用指针访问变量地址
	fmt.Println("指针的值是", pointer)
	fmt.Println("使用指针访问变量值", *pointer)

	// //空指针
	// var nullPointer *int
	// fmt.Printf("nullPointer的值为：%x\n", nullPointer)//nullPointer的值为：0

	// //判断空指针
	// fmt.Println(nullPointer==nil)//true

	//指针数组

}
