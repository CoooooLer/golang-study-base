package main

import (
	"encoding/json"
	"fmt"
)

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

	// 创建一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m = %v\n", m)

	// var str string
	// str = string(m["company"]) // error 必须使用断言
	// fmt.Println("str = ", str)

	// 类型断言
	var str string
	for key, value := range m {
		// fmt.Printf("%v======>%v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的类型是string, value = %v\n", key, str)
		case bool:
			fmt.Printf("map[%s]的类型是bool, value = %v\n", key, data)
		case float32:
			fmt.Printf("map[%s]的类型是float, value = %v\n", key, data)
		case []string: //
			fmt.Printf("map[%s]的类型是[]string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的类型是[]interface{}, value = %v\n", key, data)
		}
	}

}
