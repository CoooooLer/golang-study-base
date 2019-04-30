package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "  abc  go golang cc c++"
	// reg1 := regexp.MustCompile(`a.c`)
	// fmt.Println("reg1 = ", reg1)
	// reg2 := regexp.MustCompile(`a[0-9]c`)
	// fmt.Println("reg2 = ", reg2)
	// 1) 解析规则，它会解析正则表达式，如果成功返回解释器
	reg3 := regexp.MustCompile(`a[a-z]c`)
	fmt.Println(reg3)
	if reg3 == nil { // 解释失败，返回nil
		fmt.Println("rexexp error")
		return
	}

	//2) 根据规则提取关键信息
	result1 := reg3.FindAllStringSubmatch(buf, -1) // -1 返回所有匹配的结果
	fmt.Println("result1 = ", result1)

}
