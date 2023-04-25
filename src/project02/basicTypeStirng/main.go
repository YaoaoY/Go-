package main
import(
	"fmt"
	_ "unsafe"
	"strconv"
)
//Go语言中基本数据类型转string类型
func main() {
	var num1 int  = 99
	var num2 float64 = 3.14
	var b bool = true
	var c byte = 'h'
	var str string	//空串

	//第一种方式

	str = fmt.Sprintf("%d",num1)//相当于把变量num1按照十进制数字输出到字符串中，即把这个数字转换成了字符串
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串

	str = fmt.Sprintf("%t",b) //%t是打印布尔类型变量
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串

	str = fmt.Sprintf("%c",c)
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串

	//第二种方式：strconv函数;需要导入"strconv"包
	var num3 int  = 99
	var num4 float64 = 3.14
	var b2 bool = true
	

	str = strconv.FormatInt(int64(num3),10) //func strconv.FormatInt(i int64, base int) string,第一个参数是int64类型，所以需要强制类型转换，第二个参数表示进制
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串
	
	//FormatFloat(f float64, fmt byte, prec int, bitSize int) string;第二个参数为'f'格式的小数（小数有多种表示形式），第3个参数表示精度（多少个小数），最后一个表示这个小数占多少位
	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串


	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串

	// Itoa(i int) string
	// Itoa is equivalent to FormatInt(int64(i), 10).
	var num5 int64 = 1234
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str = %q\n",str,str)//%T是打印类型，%q是带双引号打印字符串
	
	
}