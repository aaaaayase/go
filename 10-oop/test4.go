package main

import "fmt"

// 接口是个指针 类只要实现了接口中的全部方法就是继承了接口
type animalI interface {
	sleep()
	getColor() string
	getType() string
}

type cat struct {
	color string
}

type dog struct {
	color string
}

func (this *cat) sleep() {

	fmt.Println("cat sleeping")
}

func (this *cat) getColor() string {
	return this.color
}

func (this *cat) getType() string {
	return "cat"
}

func (this *dog) sleep() {
	fmt.Println("dog sleeping")
}

func (this *dog) getColor() string {
	return this.color
}

func (this *dog) getType() string {
	return "dog"
}

func Test4() {
	var c animalI
	c = &cat{color: "red"}
	c.sleep()
	fmt.Println(c.getColor())
}
