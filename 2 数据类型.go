package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := "hello world"
	var b int = 11
	fmt.Println(b, reflect.TypeOf(a))
	fmt.Println(a, reflect.TypeOf(b))

}
