package main

import (
	"fmt"
	"test"
)

func main() {
	test.Test()

	var s test.Stu
	s.Id = 111
	s.Name = "goooo"
	fmt.Println("s = ", s)
}
