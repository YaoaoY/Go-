package main

import (
	"fmt"
)

//声明结构体

// 定义一个名字为student的结构体
// 这样的结构体称为：命名结构体(named Structure).相当于是创建了名字为student的新类型，可以用这个类型来创建变量
type Student struct {
	name string
	age  int
	sex  byte
}

// 给结构体定义方法：
// 定义一个与结构体类型为Student绑定的方法（括号内为方法的接收者
func (s Student) ShowInfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}

// 创建匿名字段的结构体
type Cat struct {
	string
	int
}

// 如果绑定的结构体的方法要改变实例的属性时，必须使用指针作为方法的接收者
func (s *Student) changeAge() {
	s.age = 20
}

// 匿名结构体：不用声明一个新的类型
var Lesson struct {
	name, target string
	spend        int
}

// 创建一个嵌套结构体
type Work struct {
	time    string
	num     int
	student Student
}

func main() {

	//创建结构体变量
	//1)使用字段名创建结构体变量(注意语法，每条赋值后面有一个英文逗号
	student01 := Student{
		name: "yyz",
		age:  18,
		sex:  'F',
	}
	fmt.Println(student01)
	//2）不适用字段名赋值，按照属性定义顺序依次传参形式赋值
	student02 := Student{"yyz", 18, 'F'}
	fmt.Println(student02)

	//3)创建匿名结构体变量
	student03 := struct {
		name string
		age  int
		sex  byte
	}{"yy", 12, 'M'}
	fmt.Println(student03)

	//4)当结构体变量没有被初始时，字段被赋予相应类型的零值
	student04 := Student{}
	fmt.Println(student04)

	//5)访问结构体字段、赋值
	fmt.Println(student01.name, student01.age, student01.sex)
	student01.age = 19
	fmt.Println(student01)

	//指向结构体的指针
	student05 := &Student{"yyz", 18, 'F'}
	fmt.Println(student05.name)
	fmt.Println((*student05).name)
	//使用指针也是用点操作符，是把C和JavA整合了一下

	//匿名字段，在创建结构体时，字段可以只有类型没有名字，这种字段称为：”匿名字段“：Anonymous Field
	cat := Cat{"mm", 12}
	fmt.Println(cat.string, cat.int) //匿名字段的名称默认是它的类型

	//嵌套结构体
	//结构体字段可能是另外一个结构体，这样的结构体称为：嵌套结构体（Nested，Structs)
	work := Work{
		time: "23/4/24",
		num:  12,
	}
	work.student = Student{"yyz", 2, 'f'}

	fmt.Println(work)

	//结构体的比较，前提：结构体的全部成员都可以互相比较，那么结构体就是可比较的
	//通过'=='或者'!='或者DeeplEqual()函数比较两个结构相同的类型并包含相同的字符段
	fmt.Println(student01 == student03)

	//给结构体定义方法
	//和C一样，Go也是无法直接在结构体内部定义方法
	fmt.Println(student01)
	student01.changeAge()
	fmt.Println(student01)

}
