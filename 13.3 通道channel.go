package main

import "fmt"

func main() {
	arr := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go getSum(arr[:len(arr)/2], ch)
	go getSum(arr[len(arr)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

}

func getSum(arr []int, ch chan int) {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	ch <- sum
}
