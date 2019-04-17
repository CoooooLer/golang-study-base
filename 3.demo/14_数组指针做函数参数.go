package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	Modify(&a)
	fmt.Println("a = ", a)
}

// p指向数组a,它是指向数组，它是数组指针
// *p 代表指针所指向的内存地址，就是实参a
func Modify(p *[5]int) {

	(*p)[0] = 666
}
