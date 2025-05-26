package main

import "fmt"

func main() {
	//test1()
	test2()
}

func test1() {
	// 1
	var m1 map[string]string
	m1 = make(map[string]string, 10)
	fmt.Println(m1)

	// 2
	m2 := make(map[string]string, 10)
	fmt.Println(m2)

	// 3
	m3 := map[string]string{"g": "G", "k": "K"}
	fmt.Println(m3)
}
