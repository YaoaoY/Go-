	package main

import (
	"fmt"
)

//演示Golang中string类型的使用

func main(){

	//1)基本使用
	var address string = "大连大学"
	fmt.Println(address)

	//2)字符串一旦赋值就不能修改;
	  //字符串内部使用"UTF-8"编码
	//var str = "hhh"
	//str[0] = 'b'//这种用法是错误的，不能去修改字符串中的内容

	//3）如何修改字符串
	//要修改字符串，需要将其转化为[]rune/[]byte，完成后再转换为string.
	//注意，无论哪种转换，都会重新分配内存并且复制字节数组

	s1 := "yaoyao"
	//强制类型转化为byte数组
	bytes1 := []byte(s1)
	bytes1[0] = 'b'
	fmt.Println(string(bytes1))

	//强制转换为rune数组
	s2 := "瑶瑶子"
	runes2 := []rune(s2)
	runes2[2] = '籽'
	fmt.Println(string(runes2))


	//4)使用反引号，以字符串原生形式输出，包含换行和特殊字符，可以实现防止攻击
	s3 :=`
	func main(){
		scoreMap := make(map[string]int)
	
		scoreMap["11"] = 90
		scoreMap["22"] = 100
	
		v,c := scoreMap["12"]
	
		if c{
			fmt.Print(v)
		}else{
			fmt.Print("查无此人")
		}
		
	}
	`
	fmt.Println(s3)

	//5)使用'+'可以直接拼接字符串
	s4 := "hello"+"yyz"
	s4 += "!!!"
	fmt.Println(s4)
	//当需要拼接的字符串比较长时，可以分行写，但是注意把'+'留在上一行
	s5 := "hello"+"yyz"+"hello"+"yyz"+"hello"+"yyz"+"hello"+"yyz"+
	"hello"+"yyz"+"!!!"
	fmt.Println(s5)

	
}