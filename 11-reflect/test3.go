package main

import (
	"fmt"
	"reflect"
)

type book struct {
	Title string
	Auth  string
}

func (this book) getTitle() {

	fmt.Println("sdsd ")
}

func reflectFunc(arg interface{}) {
	inputType := reflect.TypeOf(arg)
	inputValue := reflect.ValueOf(arg)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func Test3() {
	b := book{Title: "g", Auth: "a"}

	reflectFunc(b)
}
