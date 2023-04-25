package main

import (
	"fmt"
)

//在Go中，mao是哈希表的作用，是一个拥有键值对元素的无序集合
//1）键key,唯一

func main() {
	//创建map:使用make函数，传入键和值的类型，可以创建mao

	scores := make(map[string]int)
	fmt.Println(scores)

	//也可以使用map字面值来创建mao，同时可以指定初始化一些key和vaule
	//语法注意：每一对键值对后面都有一个英文逗号
	var steps map[string]string = map[string]string{
		"step1": "eat",
		"step2": "sleep",
	}
	fmt.Println(steps)

	//使用短声明初始化
	steps2 := map[string]string{
		"step1": "eat",
		"step2": "sleep",
	}
	fmt.Println(steps2)

	//添加元素：直接使用map[key] = value进行元素添加
	steps["step3"] = "getting up"
	fmt.Println(steps)

	//更新元素,使用map[key] = ..更改对应key的值
	steps["step2"] = "crying"
	fmt.Println(steps)

	//获取元素，使用map[ke]即可像数组一样获取到对应key的value值;如果key不存在，则返回valie类型的零值
	fmt.Println(steps["step1"])

	//删除元素，使用delete(map,key)函数，可以删除对应哈希表中其键值对;若key不存在，则静默处理
	delete(steps, "step2")
	fmt.Println(steps)

	//判断key是否存在
	v1, ok := steps["step2"]
	fmt.Println(v1, ok)
	v2, ok1 := steps["step1"]
	fmt.Println(v2, ok1)

	//使用for range遍历map
	for key, value := range steps {
		fmt.Println("key:", key, "value:", value)
	}

	//可以使用len函数获取数组长度
	fmt.Println("steps长度为：", len(steps))

	//注意:map是引用类型，当被赋值为，当把map类型赋值给一个新的遍历时，它们指向的是同一个内部数据结构
	//因此改变其中一个变量，也会影响另外一个变量
	steps3 := steps
	fmt.Println("step3:", steps3, "steps:", steps)
	steps3["step4"] = "outing"
	fmt.Println("step3:", steps3, "steps:", steps)

}
