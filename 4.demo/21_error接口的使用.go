package main

import "fmt"
import "errors"

func main() {
	var err01 error = fmt.Errorf("this is a normal err01")
	fmt.Printf("err01: %v\n", err01)

	err02 := fmt.Errorf("this is normal err02")
	fmt.Printf("err02: %s\n", err02)

	err03 := errors.New("this is new normal error03")
	fmt.Printf("err03: %s\n", err03)
}
