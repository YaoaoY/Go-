package main

import "fmt"

// 定义多个常量
const (
	BEIJING = 10 * iota
	SHANGHAI
	DALIAN
)

func main() {
	fmt.Println("hi")

	//常量
	const length = 10
	fmt.Println(length)

}
