package main

import (
	"fmt"
)

func printMap(mapT map[string]string) {
	for key, value := range mapT {
		fmt.Println(key, value)
	}
}

func test2() {

	map1 := make(map[string]string)
	map1["name"] = "dcy"
	printMap(map1)

	map1["name"] = "ggg"
	printMap(map1)

	delete(map1, "name")
	printMap(map1)

}
