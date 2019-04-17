package main

import "fmt"

func main() {
	str := "abc"

	// 通过 for 循环打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	// range 迭代打印每个元素，默认返回两个值：一个是元素的下标，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for i := range str { // 第2个返回值，默认丢弃，返回元素的下标
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str { // 第2个返回值，默认丢弃，返回元素的下标
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

}
