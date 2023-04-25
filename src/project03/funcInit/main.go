package main

import (
	"fmt"
	"funcInit/utils"
)

//执行流程：全局变量->init->main函数

var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}


func main(){

	fmt.Println("main()...age = ",age)
	fmt.Println("Age = ",utils.Age,"Name =",utils.Name)
}