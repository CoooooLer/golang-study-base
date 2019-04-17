package main

import "fmt"

func main() {
	a, b := 1, 9
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1 // 操作指针，记得写 *
}
