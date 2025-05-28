package main

import "fmt"

// 对类名包括属性名首字母大写起到public的作用 反之就会一定程度上限制访问 即是否能够被其它包访问
type hero struct {
	name  string
	ad    int
	level int
}

//func (this hero) show() {
//	fmt.Println(this)
//}
//
//func (this hero) getName() string {
//	return this.name
//}
//
//func (this hero) setName(name string) {
//	this.name = name
//}

func (this *hero) show() {
	fmt.Println(this)
}

func (this *hero) getName() string {
	return this.name
}

func (this *hero) setName(name string) {
	this.name = name
}

func main() {
	// 创建一个对象
	hero1 := hero{name: "zhang3", ad: 3, level: 1}
	hero1.show()
	hero1.setName("111")
	hero1.show()
}
