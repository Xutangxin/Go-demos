package main

import "fmt"

func main() {
	a := 11
	fmt.Println(&a)
	first := "hello"
	second := "world"
	fmt.Println(first == second)
	first = "world"
	fmt.Println(first)
	fmt.Println(first == second)

	age := 22
	age2 := 22
	if age == age2 {
		fmt.Println("age equals age2")
	}

	num := 1
	num += 10
	fmt.Println(num)

}
