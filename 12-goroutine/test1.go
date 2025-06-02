package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func test1() {
	// 创建一个go程去执行newTask流程
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("test1 Goroutine : i = %d\n", i)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	//test1()
	Test2()
}
