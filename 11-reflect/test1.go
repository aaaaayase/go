package main

import "fmt"

func test1() {
	a := "111"
	// Pair<type:string,value:"111">
	// 这个会一直传递下去
	var b interface{}
	// 变量内置的pair会传递
	b = a
	value, ok := b.(string)
	fmt.Println(value, ok)
}

func main() {
	//test1()
	//Test2()
	//Test3()
	//Test4()
	Test5()
}
