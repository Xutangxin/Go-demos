package main

import "fmt"

//接口定义
type Programmer interface {
	WriteHelloWorld() string
}

//定义结构体（接口实现）
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func main() {
	g := GoProgrammer{}
	fmt.Println(g.WriteHelloWorld())
}
