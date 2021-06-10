package main

import "fmt"

func main() {
	a := 1
	b := 2
	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is bigger than b")
	}

	n := 0
	for n < 5 {
		switch n {
		case 0, 2, 4:
			fmt.Println("even")
		case 1, 3:
			fmt.Println("odd")
		default:
			fmt.Println("none")
		}
		n++
	}

	sum := 0
	for num := 0; num < 5; num++ {
		sum += num
	}
	fmt.Println(sum)

	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

}
