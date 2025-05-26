package main

func test4() {
	// 长度 容量
	arr := make([]int, 3, 5)

	arr = append(arr, 1) // 追加元素
	// 超出容量 会动态多开辟一份capacity这样的空间
}
