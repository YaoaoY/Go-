package main

import "fmt"

//4）定义全局变量
var a = 100
var b = 300
var s = "heheheh"

//5)一次性声明多个全局变量
var (
	n1 = 12
	n2 = 'c'
	n3 = "jack"
)

func main(){
	//1),指定变量类型，不初始化，系统初始化为默认值
	var i int
	fmt.Println("i = ",i)

	//2)声明变量时不指定变量类型，进行初始化，让编译器自动推导变量类型
	var num = 10.123
	fmt.Printf("num type is %T\n",num)//float64

	//3)使用`:=`省略var和类型进行声明+初始化
	name := "hhh"
	fmt.Printf("name Type is %T\n",name)//string
	fmt.Println(name)

	//6)一次性声明多个同类型变量
	var a1,a2,a3,a4 int
	fmt.Println(a1,a2,a3,a4)

	//7),一次性声明多个变量并且初始化
	var b1,b2,b3 = 1,2,3
	fmt.Println(b1,b2,b3)

	//8)使用`:=`一次性声明并初始化多个变量
	c1,c2,c3 := 1,2,3
	fmt.Println(c1,c2,c3)


}