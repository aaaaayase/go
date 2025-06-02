package main

import (
	"fmt"
	"time"
)

func Test2() {

	// 匿名方法
	//go func() {
	//	defer fmt.Println("A.defer")
	//
	//	return
	//	func() {
	//		defer fmt.Println("B.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()

	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(20, 10)

	// 死循环
	for {
		time.Sleep(time.Second * 1)
	}
}
