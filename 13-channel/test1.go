package main

import "fmt"

// 无缓冲channel
func test1() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine running...")
		c <- 666
	}()

	a := <-c
	fmt.Println(a)
	fmt.Println("main goroutine end")

}

func main() {
	//test1()
	Test2()
}
