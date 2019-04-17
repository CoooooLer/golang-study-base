package main

import "fmt"

func main() {
	var arr [50]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}
