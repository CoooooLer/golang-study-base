package main

import "fmt"

func main() {
	var i interface{}
	i = 1
	fmt.Println("i = ", i)

	i = "abv"
	fmt.Println("i = ", i)
}

// func xxx(args ...interface()){

// }
