package main

import "fmt"
import "errors"

func main() {
	res, err := MyDiv(1, 0)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Printf("res: %d\n", res)
	}
}

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("0 不能做除数！！！")
	} else {
		result = a / b
	}
	return // 有返回值，必须写 return
}
