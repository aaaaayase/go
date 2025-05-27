package main

import "fmt"

type human struct {
	name string
	age  int
}

type superman struct {
	human
	sex string
}

func (this *human) eat() {
	fmt.Println("human eat---------")
}

func (this *human) walk() {
	fmt.Println("human walk---------")
}

func (this *superman) fly() {
	fmt.Println("superman fly---------")
}

func Test3() {

	h := human{name: "li4", age: 20}

	h.walk()
	h.eat()

	fmt.Println(h)

	var s superman
	s.name = "zhang3"
	s.age = 30
	s.sex = "nv"

	fmt.Println(s)
	s.walk()
	s.eat()
	s.fly()

}
