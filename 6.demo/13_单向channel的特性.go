package main

func main() {
	// 创建一个channel，双向的
	ch := make(chan int)
	// 双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch // 只能写，不能读
	var readCh <-chan int = ch  // 只能读，不能写

	writeCh <- 666 // 写
	//<-writeCh // err

	<-readCh // 读
	//readCh <- 666 // err

	// 单向channel无法转换为双向channel
	var ch2 chan int = writeCh

}
