package main

func test5() {

	slice1 := []int{1, 2, 3, 4, 5}
	// 浅拷贝 指向的还是一个地址
	slice2 := slice1[0:2]

	// 指向的是不同的地址
	slice3 := make([]int, 2)
	copy(slice3, slice2)

}
