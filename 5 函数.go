package main

import "fmt"

func main() {
	a := 100
	b := 200
	var result int
	result = getMax(a, b)
	fmt.Printf("最大值是 %d\n", result)

	str1, str2 := swapVal("hello", "world")
	fmt.Println(str1, str2)

	number := plusNumber()
	fmt.Println(number())
	fmt.Println(number())

}

//函数定义（返回两个数最大值）
//其余函数需在main函数外定义
func getMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

//函数可以返回多个值
//以下函数实现了交换两个字符串值的功能
func swapVal(val1, val2 string) (string, string) {
	return val2, val1
}

//函数闭包
func plusNumber() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}
