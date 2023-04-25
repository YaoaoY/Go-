package main	//说明hello.go这个文件在main这个包中

//导入内置包，可以使用其中函数等
import 	(
	"fmt"
)

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