package main

import "fmt"

// const定义枚举
const (
	BEIJING = 1
)

// const中使用iota默认为0 剩余会递增加一 将iota乘以10 后续递增也会按照+10这样的规律
const (
	A = iota
	B
	C
)

func main() {

	const length = 100
	fmt.Println(length)
}
