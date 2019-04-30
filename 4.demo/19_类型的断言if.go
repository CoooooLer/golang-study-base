package main

import "fmt"

type Student struct {
	id   int
	name string
}

// 类型查询，也叫 类型断言
func main() {
	i := make([]interface{}, 3) // 新建一个长度、容量为3的接口类型的切片
	i[0] = "hello"
	i[1] = 123
	i[2] = Student{66666, "mike"}

	// 第一个返回的是元素的下标，第二个返回的是元素的值
	for index, data := range i { // 递归
		// 第一个返回的是值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int,内容为 %d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为 string, 内容为 %s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为 Student{}, 内容为id = %d, name = %s\n", index, value.id, value.name)
		}
	}

}
