package main

import "fmt"

type People struct {
	name   string
	age    int
	gender string
}

func (p *People) Speak() {
	fmt.Println("hello my name is " + p.name)
}

func (p *People) Eat() {
	fmt.Println("I am eating now")
}

func (p *People) showMyself() {
	fmt.Println(*p)
}

func main() {
	jake := People{"jake", 22, "male"}
	jake.Speak()
	jake.Eat()
	jake.showMyself()
}
