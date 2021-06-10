package main

import "fmt"

func main() {
	//创建一个切片
	//相当于声明一个未指定长度的数组
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 3, 5)
	fmt.Println(s1, s2) //[1 2 3 4 5] [0 0 0]

	//切片初始化
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:]
	fmt.Println(arr, slice)

	//len()和cap()
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) //[0 0 0] 3 5

	//切片截取
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := slice2[1:4]
	fmt.Println(newSlice)   //[2 3 4]
	fmt.Println(slice2[:])  //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(slice2[0:]) //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(slice2[:4]) //[1 2 3 4]
	fmt.Println("===============================")

	//append()和copy()
	//往切片中添加元素
	slice3 := []int{1, 2, 3, 4, 5}
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 7)
	fmt.Println(slice3) //[1 2 3 4 5 6 7]
	slice4 := make([]int, len(slice3), cap(slice3)*2)
	fmt.Println(slice4)
	//拷贝切片slice3到slice4中
	copy(slice4, slice3)
	fmt.Println(slice4)
	slice4 = append(slice4, 1, 1, 1, 1)
	fmt.Println(slice4, len(slice4), cap(slice4))

}
