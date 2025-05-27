package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func Test5() {

	m := movie{Title: "喜剧之王", Year: 2000}

	jsonstr, err := json.Marshal(m)

	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	fmt.Println(string(jsonstr))

	m1 := movie{}

	err = json.Unmarshal(jsonstr, &m1)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(m1)

}
