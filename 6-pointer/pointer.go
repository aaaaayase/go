package main

import "fmt"

func main() {
	var a int = 20
	var b int = 30

	swap(&a, &b)
	fmt.Println(a, b)

	// 二级指针
	var p *int = &a
	var q **int = &p
	fmt.Println(**q)
}

func swap(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}
