package main

import "fmt"

func mapByReference() {
	steps4 := map[string]string{
		"第一步": "eat",
		"第二步": "sleep",
		"第三步": "cry",
	}
	fmt.Println("steps4: ", steps4)
	// steps4:  map[第一步:Go语言极简一本通 第三步:从0到Go语言微服务架构师 第二步:Go语言微服务架构师核心22讲]
	newSteps4 := steps4
	newSteps4["第一步"] = "1"
	newSteps4["第二步"] = "2"
	newSteps4["第三步"] = "3"
	fmt.Println("steps4: ", steps4)

	fmt.Println("newSteps4: ", newSteps4)

}
