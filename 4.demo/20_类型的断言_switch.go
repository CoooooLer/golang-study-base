package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	var i = make([]interface{}, 3)
	i[0] = 111
	i[1] = "hello go"
	i[2] = Student{1, "mike"}

	// 类型查询 ， 又叫类型断言
	for index, data := range i {
		switch value := data.(type) { // data.(type) switch特有的
		case int:
			fmt.Printf("x[%d]的类型是int,值为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]的类型是string，值为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]的类型是Student{}, 值id = %d, name = %s\n", index, value.id, value.name)
		}
	}
}
