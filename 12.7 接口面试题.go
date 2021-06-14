package main

import "fmt"

type People interface {
	speak(string) string
}

type Student struct {
}

func (s *Student) speak(think string) (talk string) {
	if think == "sb" {
		talk = "you are a sb"
	} else {
		talk = "hello, how are you"
	}
	return
}

func main() {
	// var peo People = Student{} //报错，这里应该接收指针类型。Student没有实现People接口，speak方法应该接受一个指针
	//正确的写法
	var peo People = &Student{}
	think := "hello world"
	fmt.Println(peo.speak(think))

}
