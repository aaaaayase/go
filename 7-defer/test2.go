package main

import "fmt"

func deferTest() int {
	fmt.Println("deferTest")
	return 0
}

func returnTest() int {
	fmt.Println("returnTest")
	return 0
}

func DeferAndReturn() int {
	defer deferTest()
	return returnTest()
}
