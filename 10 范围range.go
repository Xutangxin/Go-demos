package main

import "fmt"

func main() {

	//使用range遍历数组和切片的形式
	//for index, item of range:arr
	//index是元素索引，item是数组中元素
	nums := []int{1, 2, 3, 4, 5}
	for index, item := range nums {
		fmt.Println(index, item)
	}
	//不写索引
	sum := 0
	for _, item := range nums {
		sum += item
	}
	fmt.Println(sum)

	//range也可以遍历字符串
	//index 是元素索引，item是字符元素的Unicode
	// for index, item :=range "hello, world"{
	//     fmt.Println(index, item)
	// }

	//使用range遍历map

}
