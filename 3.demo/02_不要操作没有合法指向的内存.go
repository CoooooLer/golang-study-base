package main

import "fmt"

func main() {
	var p *int
	fmt.Printf("p = %v\n", p)

	// *p = 666  // error 不是合法的指向
	// fmt.Printf("p = %v\n", p)

	a := 111
	p = &a // p 指向 a
	*p = 666
	fmt.Printf("a = %v\n", a)
}
