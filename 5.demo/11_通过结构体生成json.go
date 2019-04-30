package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量名首字母必须大写
type IT struct {
	// Company string
	// Subject []string
	// Isok    bool
	// Price   float64
	Company string   `json: "-"`       // 此字段不会输出到屏幕
	Subject []string `json:"subject"`  // 二次编码
	Isok    bool     `json: ",string"` // 类型改为 string 类型
	Price   float64  `json: ",string"` // 类型改为 string 类型
}

func main() {
	// 定义一个结构体变量，同时初始化
	s := IT{"company", []string{"go", "c++", "python"}, true, 11.23}
	// 编码，根据内容生成json字符串
	// buf, err := json.Marshal(s)

	// 格式化编码
	buf, err := json.MarshalIndent(s, " ", "  ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf)) // string(buf)  必须转字符串
}
