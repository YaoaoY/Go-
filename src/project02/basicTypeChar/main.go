package main

import(
	"fmt"
)

func main(){

	//Golang中没有char类型，用两个字节的byte存储字符型
	//byte = uint8

	var c1 byte = 'a'
	var c2 byte = 'c'

	//当我们直接用Println输出字符时，输出的是该字符的码值
	fmt.Println("c1 = ",c1)
	fmt.Println("c2 = ",c2)

	//如果想以字符格式输出，那就使用格式输出的方法
	fmt.Printf("c1 = %c,c2 = %c\n",c1,c2)

	//由于汉字是3-4个字节，用byte存储会overflow,
	//可以使用rune = int32(4个字节)，代表UTF-8字符

	var c3 rune = '呵'
	fmt.Println("c3 = ",c3)//这样还是会输出码值
	fmt.Printf("c3 = %c\n",c3)//呵


}