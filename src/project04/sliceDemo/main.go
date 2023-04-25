package main

import "fmt"

func appendSliceData() {
	s := []string{"yyz"}
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, "yy")
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, "12", "45")
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, []string{"微服务", "分布式锁"}...)
	fmt.Println(s)
	fmt.Println(cap(s))
}
func mSlice() {
	numList := [][]string{
		{"1", "G"},
		{"2", "Go"},
		{"3", "goo"},
	}
	fmt.Println(numList)
}
