package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "company"
	m["subject"] = []string{"go", "c++", "python"}
	m["isok"] = false
	m["price"] = 11.21

	res, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
