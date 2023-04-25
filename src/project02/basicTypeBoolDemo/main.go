package main
import(
	"fmt"
	"unsafe"//unsafe包提供一些跳过go语言类型安全限制操作
)

//golang中bool类型的使用

func main(){
	var b = false
	fmt.Println("b = ",b)

	//1)bool类型占用存储空间为1字节！
	fmt.Println("b所占内存空间为：",unsafe.Sizeof(b))//Sizeof返回类型本身数据所占字节数，返回值是”顶层“数据占用字节//
	//例如v是切片，它会返回切片本身描述符大小，，而非切片底层引用的那个内存的大小

	//2)bool类型只有两个取值：true或者false

	
}