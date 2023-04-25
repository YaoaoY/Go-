package main

import "fmt"

type student struct {
	name string
	age  int
	sex  byte
}

func main() {
	// 使用字段名创建结构体
	student1 := student{
		name: "yyz",
		age:  18,
		sex:  'F',
	}
	// 不使用字段名创建结构体
	student2 := student{"yyz", 18, 'F'}

	fmt.Println("lesson1 ", student1)
	fmt.Println("lesson2 ", student2)
}
