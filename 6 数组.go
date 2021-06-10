package main

import "fmt"

func main() {
	//数组声明 数组名+数组长度+元素类型
	var nums [10]int
	fmt.Println(nums)

	//数组初始化
	//方式1
	var testArr = [5]int{1, 2, 3, 4, 5}
	//方式2
	testArr2 := [5]int{2, 3, 4, 5, 6}
	fmt.Println(testArr)
	fmt.Println(testArr2)

	//如果数组长度不确定，可用...代替数组长度
	testArr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(testArr3)
	fmt.Println(len(testArr3))

	//还可以指定下标来初始化数组
	testArr4 := [5]int{2: 10, 3: 10}
	fmt.Println(testArr4) //[0 0 10 10 0]

	//忽略设置数组大小，数组会自动根据元素个数来设置数组大小
	testArr5 := []int{2, 3, 5}
	fmt.Println(testArr5)
	fmt.Println(len(testArr5)) //3

	//访问设置数组元素
	testArr[0] = 100
	fmt.Println(testArr[0], testArr[1])

	//设置多维数组
	testArr6 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(testArr6) //[[1 2 3] [4 5 6]]

}
