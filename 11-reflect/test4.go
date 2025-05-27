package main

import (
	"fmt"
	"reflect"
)

type act struct {
	Title string `info:"title" doc:"标题"`
	Sex   string `info:"sex" doc:"性别"`
}

func support(arg interface{}) {
	intputType := reflect.TypeOf(arg)
	for i := 0; i < intputType.NumField(); i++ {
		taginfo := intputType.Field(i).Tag.Get("info")
		tagdoc := intputType.Field(i).Tag.Get("doc")

		fmt.Println(taginfo, tagdoc)
	}

}

func Test4() {
	a := act{Title: "sdsa", Sex: "sd"}

	support(a)

}
