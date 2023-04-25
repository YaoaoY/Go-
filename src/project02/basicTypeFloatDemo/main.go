package main
import(
	"fmt"
	//"unsafe"
)
func main(){
	//Golang中浮点型默认声明为float64
	var n1 = 3.14
	fmt.Printf("n1 type is %T\n",n1)

	//小数存储不准确，可能造成精度丢失
	var n2 float32 = -123.0000901
	var n3 float64 = -123.0000901
	fmt.Println("n2 = ",n2,"n3 = ",n3)

	//十进制表示的小数（必须有小数点）
	n4 := 5.12
	n5 := .123//0.123
	fmt.Println("n4 = ",n4,"n5 = ",n5)

	//科学技术法表示小数
	n6 := 3.14e2 //3.14*(10^2)
	n7 := 3.14E2
	n8 := 3.14e-2//3.14*(10^-2) 

	fmt.Println("n6 = ",n6,"n7 = ",n7,"n8 = ",n8)



}