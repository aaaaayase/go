package main

import "fmt"

func func1() {
	fmt.Println("test1")
}

func func2() {
	fmt.Println("test2")
}

func func3() {
	fmt.Println("test3")
}

func main() {

	defer func1()
	defer func2()
	defer func3()
	fmt.Println("the end")

	DeferAndReturn()
}
