package main

import "fmt"

// 定义一个新的类型
type myint int

// 定义一个结构体类型
type Book struct {
	title  string
	author string
}

func changeBook1(book Book) {
	book.title = "xiyou"
}

func changeBook2(book *Book) {
	(*book).title = "xiyou"
}

func test1() {
	var a myint = 10
	fmt.Println(a)
	var book1 Book
	book1.title = "shuihu"
	book1.author = "yun"
	fmt.Printf("%v\n", book1)
	changeBook1(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}

func main() {
	//test1()
	Test3()

}
