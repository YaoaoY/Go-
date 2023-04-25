package main

import "fmt"

func test01() {
	//声明没有指定数组元素的值，默认为0
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

}

func test02() {
	//在声明数组的时候进行初始化
	var arr = [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr)

	//短声明
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//数组部分初始化
	arr2 := [5]int{1, 2}
	fmt.Println(arr2)

	//通过指定索引，对数组某几个元素进行复制
	arr3 := [5]int{1: 2, 3: 4}
	fmt.Println(arr3)

	//使用...让编译器计算数组长度
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

}

func test3() {
	//对于数组来说，数组长度是类型的一部分
	var arr1 [3]int
	var arr2 [5]int
	fmt.Printf("arr1 type:%T,arr2 type:%T\n", arr1, arr2)
}

func test04() {
	//定义多维数组
	arr := [3][2]string{
		{"11", "22"},
		{"33", "44"},
		{"55", "66"}}
	fmt.Println(arr)
}

//使用内置的len函数返回数组中元素的个数

func arrLen() {
	arr := [...]string{"yaoyao", "hhh", "111"}
	fmt.Println("数组长度：", len(arr))

}

// 数组遍历：for range循环获取数组每个索引以及索引上的值
func show() {
	arr := [...]string{"北京", "上海", "武汉"}
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	//用'_'占位
	for _, value := range arr {
		fmt.Printf("value = %s\n", value)
	}
}

func main() {
	test01()
	test02()
	test3()
	test04()
	arrLen()
	show()
	//注意：数组是值类型，而不是引用类型。当数组赋值给一个新的数组变量时，该变量得到的是一个原始数组的一个副本。
	//对这个新的变量进行修改，不会影响原始数值
	arr := [...]string{"北京", "上海", "武汉"}
	arr1 := arr
	arr1[0] = "天津"
	fmt.Println(arr, arr1)
}
