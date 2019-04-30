package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "11.2 .2 3. 333.3 56.4 4.4 553.4"

	// 解释正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`) // esc 下面的 `
	if reg == nil {
		fmt.Println("compile err")
		return
	}

	// 提取关键信息
	// res := reg.FindAllString(buf, -1) // -1 返回所有满足的结果
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(res)
}
