package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4), sum(1, 2, 3, 4, 5))
	defer clear() //defer 延迟执行函数
	fmt.Println("hello world")
}

//可变参数（不定数量参数）
func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func clear() {
	fmt.Println("clear resources")
}
