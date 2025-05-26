package main

import "fmt"

// 静态数组给方法传参还是拷贝的副本 动态数组不是
func Test(arr [4]int) {

	arr[0] = 1000
}

func main() {
	var arr1 [10]int
	arr2 := [10]int{1, 2, 3, 4}
	arr3 := [4]int{1, 1, 1, 1}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr2 {

		fmt.Println(index, value)
	}

	fmt.Printf("%T\n", arr3)
	Test(arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	Test2()
}
