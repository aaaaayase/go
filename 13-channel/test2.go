package main

import (
	"fmt"
	"time"
)

// 有缓冲channel
func Test2() {

	c := make(chan int, 3)
	go func() {
		defer fmt.Println("go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go程正在运行，cap=%d,len=%d", cap(c), len(c), i)
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num)
	}

	fmt.Println("main go程结束")
	for {

	}
}
