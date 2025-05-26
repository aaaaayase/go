package main

import "fmt"

// 返回单个值
func func1(a int, b int) int {
	fmt.Println(a, b)

	return a + b
}

// 返回多个值
func func2(a string, b int) (string, int) {

	return a, b
}

// 返回有名称的
func func3(a string, b int) (c1 string, c2 int) {
	c1 = "sda "
	c2 = 100
	return c1, c2
}

func main() {
	fmt.Println(func1(10, 20))
	fmt.Println(func2("<UNK>", 100))
	fmt.Println(func3("<UNK>", 100))
}
