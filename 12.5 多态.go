package main

import "fmt"

type Code string

//定义接口
type Programmer interface {
	WriteHelloWorld() Code
}

//定义结构体
type GoProgrammer struct {
}

//结构体实现接口方法
func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(hello world)"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.print.out(hello world)"
}

//接口方法
func writeFirstProgrammer(p Programmer) {
	// fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
	fmt.Println(p, p.WriteHelloWorld())
}

func main() {
	goer := new(GoProgrammer)
	javaer := new(JavaProgrammer)
	fmt.Println(goer.WriteHelloWorld())
	fmt.Println(javaer.WriteHelloWorld())
	writeFirstProgrammer(goer)
	writeFirstProgrammer(javaer)
}
