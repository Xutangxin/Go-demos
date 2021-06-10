package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

//将channel通道传入函数
func sum(arr []int, ch chan int) {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	//将sum赋值给通道
	ch <- sum

}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	//关闭通道
	close(ch)

}

func main() {

	// arr := []int{7, 2, 8, -9, 4, 0}
	// ch := make(chan int)
	// go sum(arr[:len(arr)/2], ch)
	// go sum(arr[len(arr)/2:], ch)
	// x, y := <-ch, <-ch
	// fmt.Println(x, y, x+y)

	//通道缓冲区
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//遍历通道与关闭通道
	ch1 := make(chan int, 10)
	go fibonacci(cap(ch1), ch1)
	for i := range ch1 {
		fmt.Println(i)

	}
}
