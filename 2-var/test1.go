package main

import "fmt"

// 声明全局变量 第四种方法不可以
var aG = 100

func main() {
	var a int
	fmt.Println(a)

	var b int = 100
	fmt.Println(b)

	var c = 100
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	var d string = "111"
	fmt.Printf("%T\n", d)

	e := 100
	fmt.Println(e)
	fmt.Printf("%T\n", e)

	fmt.Println(aG)

	// 多个变量
	var aa, bb = 100, "123"

	fmt.Println(aa, bb)
}
