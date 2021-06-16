package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := `
	hello golang
	hello world
	how are you?
	`
	fmt.Println(str)

	// number := "12345"
	myStr := "hello world"

	// for _, val := range number {
	// 	// fmt.Println(reflect.typeOf(val))
	// 	fmt.Println(reflect.TypeOf(val))
	// }

	for _, val := range myStr {
		fmt.Println(reflect.TypeOf(val))
	}

	str2 := "你好世界"

	for _, val := range str2 {
		fmt.Println(reflect.TypeOf(val))
	}

	//go里没有隐式类型转换，只有强制类型转换
	intNum := 22
	fmt.Println(reflect.TypeOf(intNum))
	fmt.Println(float64(intNum))
	newIntNUm := float64(intNum)
	fmt.Println(reflect.TypeOf(newIntNUm))
}
