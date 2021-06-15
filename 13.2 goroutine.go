package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go hello()
	// fmt.Println("goroutine done")
	// time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个goroutine就加1
		go hello(i)
	}
	wg.Wait()
}

func hello(i int) {
	defer wg.Done() //结束一个goroutine就登记-1
	fmt.Println("hello goroutine", i)

}
