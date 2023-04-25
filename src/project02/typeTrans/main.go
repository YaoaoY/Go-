package main
import (
	"fmt"
)
//演示Golang中基本数据类型之间的转换
func main(){
	//注意：Golang是强类型语言，不存在自动类型转换，类型转换需要通过手动强制来完成，注意高精度强转低精度会发生截断

	var i int32  = 128
	//i -> float
	var n1 float32 = float32(i)
	//高精度转低精度
	var n2 int8 = int8(i)

	//低精度转高精度
	var n3 int64 = int64(i)

	fmt.Printf("i = %v n1 = %v n2 = %v n3 = %v\n",i,n1,n2,n3)

	//注意，强转的是该变量的存储的值，而不是把变量本身的类型转换了
	fmt.Printf("i type is %T\n",i)

}