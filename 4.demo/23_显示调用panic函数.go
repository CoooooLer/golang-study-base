package main

import "fmt"

func main() {
	TestA()
	TestB()
	TestC()
}

func TestA() {
	fmt.Println("aaaaaaaaaaaa")
}

func TestB() {
	// 显式调用 panic() 函数，导致程序中断
	panic("this is panic test")
	fmt.Println("bbbbbbbbbbbbbb")
}

func TestC() {
	fmt.Println("CCCCCCCCCCCCC")
}
