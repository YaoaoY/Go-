package main

import "fmt"

//切片是对数组的一个连续片段的引用
//1)是引用类型，本身不拥有任何数据。只是对现有数组的引用
//2）每个切片值都会将数组作为底层数据结构

//一个slice由3部分组成：
//1）指针：指向slice指向的底层数组地址（不一定是底层数组的第一个元素
//2）长度：slice中元素个数（切片中元素个数
//3）容量：slice的开始位置到底层数据的结尾位置（从创建切片索引开始的底层数组中的元素个数

func appendSliceData() {
	s := []string{"yaoyao"}
	fmt.Println(s)
	fmt.Println(cap(s))

	//接收参数可变的叫作可变参数，append是一个可变参数函数
	s = append(s, "hehe", "zzz") //func append(slice []Type, elems ...Type) []Type
	fmt.Println(s)
	fmt.Println(cap(s))

}

func main() {
	//使用[]Type声明，一个Type类型的切片
	var numList []int //不初始化，也是一个空切片

	//声明一个空切片
	var numEmptyList = []int{}
	fmt.Println(numList, numEmptyList)

	//使用make函数构造切片
	numList2 := make([]int, 3, 5) //([]Type,size,cap)
	fmt.Println(numList2)

	//对数组进行片段截取创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s1 = arr[1:4] //左开右闭区间
	fmt.Println(arr, s1)

	//内置的len和cap用于返回slice的长度和容量
	s := make([]string, 3, 5) //长度为3，容量为5
	fmt.Println(len(s), cap(s))

	//对切片操作超出上限将导致‘panic'异常
	//fmt.Println(s[10])panic: runtime error: index out of range [10] with length 3

	//由于切片是引用类型，不赋值默认值为Nil
	//这里使用前面但未使用的空切片
	fmt.Println(numList == nil) //true

	//判断切片是否为空，就是元素个数是否为0，需要用len(s)==0来判断

	fmt.Println(len(s1) == 0)

	//切片元素的修改
	//切片元素不拥有任何数据类型，它只是底层数组的一种表示，对切片所做的任何修改都会反映在底层数组中
	var arr2 = [5]int{1, 2, 3, 4, 5}
	s2 := arr2[1:4]
	s2[0] = 1
	fmt.Println(arr2, s2)

	//追加切片元素
	//使用append可以将新元素追加到切片上。
	appendSliceData()

	//多维切片，类似于数组，切片也有多维
	numList3 := [][]string{
		{"1", "yyz"},
		{"2", "yzz"},
		{"3", "zzy"}}
	fmt.Println(numList3)
}
