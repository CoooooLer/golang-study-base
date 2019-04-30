package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Contains  "hellogo"字符串是否包含"hello",包含则返回true，不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "golang"))

	// strings.Joins 组合字符串
	s := []string{"abc", "gooo", "mike"}
	buf := strings.Join(s, "@")
	fmt.Println("buf = ", buf)

	// strings.Index 查找子串的位置
	fmt.Println(strings.Index("abchellogo", "hello"))
	fmt.Println(strings.Index("abchellogo", "aaa")) // 不包含，返回-1

	// strings.Split 以指定字符分隔字符串
	buf = "heloo@abc@go"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	// strings.Trim 去掉字符串两头的字符
	buf = strings.Trim("111   aru u ok ?    11", "1")
	// fmt.Println("Trim = ", buf)
	fmt.Printf("trim = ##%s##\n", buf)

	//
	s3 := strings.Fields("   are u ok ?      ")
	// fmt.Println("s3 = ", s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}
}
