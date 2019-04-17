package main

import "fmt"

func main() {
	oldSlice := []int{1, 2, 3}
	newSlice := []int{0, 0, 0, 0, 0}
	copy(oldSlice, newSlice)
	fmt.Println("oldSlice = ", oldSlice)
}
