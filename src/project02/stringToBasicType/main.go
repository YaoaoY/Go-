package main

import (
	"fmt"
	"strconv"
)

//golang中string转基本数据类型

func main(){
	//1)字符串转布尔类型
	str1 := "true"
	var b bool

	b , _ = strconv.ParseBool(str1)//把字符串解析成布尔值
	// func strconv.ParseBool(str string) (bool, error)
	// ParseBool returns the boolean value represented by the string. It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.

	fmt.Printf("b type %T, b = %v\n",b,b)

	//2）字符串转整形
	str2 := "123456"

	var n1 int64
	n1,_ = strconv.ParseInt(str2,10,64)
	// ParseInt(s string, base int, bitSize int) (i int64, err error)
	// ParseInt interprets a string s in the given base (0, 2 to 36) and
	// bit size (0 to 64) and returns the corresponding value i.
	fmt.Printf("n1 type is %T, n1 = %v\n",n1,n1)

	//3）字符串转浮点型
	var f1 float64
	str3 := "3.13"
	f1,_ = strconv.ParseFloat(str3,64)
	//ParseFloat(s string, bitSize int) (float64, error)
	fmt.Printf("f1 type is %T\n",f1)



}