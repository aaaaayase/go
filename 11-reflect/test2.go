package main

import (
	"fmt"
	"reflect"
)

func func1(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func Test2() {
	a := 1.222
	func1(a)
}
