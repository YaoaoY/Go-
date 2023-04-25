package main
import(
	"fmt"
)

//演示Golang中指针类型
func main(){
	var i int = 123
	//打印i的地址
	fmt.Println("i address is ",&i)

	var ptr *int = &i//*代表ptr是指针变量,int表示该变量指向的是int类型变量
	fmt.Printf("ptr = %v\n",ptr)
	fmt.Printf("ptr 地址 = %v\n",&ptr)
	//和c一样，对指针解引用，可以得到指向变量的值
	fmt.Printf("ptr 指向的值 = %v",*ptr)

	// ptr = 0xc0000a6058
	// ptr 地址 = 0xc0000ca020
	// ptr 指向的值 = 123
}