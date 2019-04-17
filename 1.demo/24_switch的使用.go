package main

import "fmt"

func main() {
	var i int
	fmt.Print("请按下楼层：")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("按下的是1楼")
		// break  // go语言默认添加 break
		fallthrough // fallthrough 不跳出 switch 语句，后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的是X楼")
	}
}
