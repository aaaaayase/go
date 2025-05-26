package main

import "fmt"

func test3() {

	// 1
	slice1 := []int{1, 2, 3}

	// 2
	var slice2 []int = make([]int, 3)

	// 3
	slice3 := make([]int, 3)

	support(slice1)
	support(slice2)
	support(slice3)

}

func support(slice []int) {
	for i := 0; i < len(slice); i++ {

		fmt.Println(slice[i])
	}
}
