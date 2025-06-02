package main

import "fmt"

func Test3() {

	c := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {

			c <- i

		}

		// 关闭channel
		close(c)

	}()

	for {
		// ok为true表示channel未关闭 反之已经关闭
		//if data, ok := <-c; ok {
		//	fmt.Println(data)
		//} else {
		//	break
		//}
		for data := range c {
			fmt.Println(data)
		}
	}

	fmt.Println("finished....")

}
