package main

import "fmt"

type book struct {
	auth string
}

// 空接口可以接受任何类型的参数
func func1(arg interface{}) {
	fmt.Println("ok--------------")

	// 可以使用断言来判断具体是哪一种类型
	value, ok := arg.(string)

	if !ok {
		fmt.Println("不是string类型")
	} else {
		fmt.Println(value)
	}

}

func Test5() {

	b := book{auth: "1111"}
	func1(b)

	c := 10
	func1(c)

	d := "12121"
	func1(d)

}
