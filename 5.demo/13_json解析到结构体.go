package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string   `json:company` // 因为原json字段都是小写，所以结构体字段必须重命名为小写
	Subject []string `json:subject`
	Isok    bool     `json:isok`
	Price   float32  `json:price`
}

func main() {
	// 原json数据
	jsonBuf := `{
        "company": "company",
        "isok": false,
        "price": 11.21,
        "subject": [
                "go",
                "c++",
                "python"
        ]
	}`

	var tmp IT                                   // 定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) // 第二个参数是地址传递 必须是 byte 类型
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp = %+v\n", tmp) // %+v 打印详细信息

	// 取部分数据
	type IT2 struct {
		Subject []string `json:subject`
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}
