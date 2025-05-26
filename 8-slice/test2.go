package main

import "fmt"

func Test2() {
	arr := []int{1, 1, 1, 1}
	test(arr)
	fmt.Println("-------------------------")
	fmt.Printf("%T\n", arr)
	fmt.Println(arr[0])
}

func test(arr []int) {

	arr[0] = 1000
}
