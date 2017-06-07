package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

//goroutine之间通过 channel进行通信,channel是和类型相关的 可以理解为  是一种类型安全的管道。
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
		fmt.Println("Count", i)
	}
	for i, ch := range chs {
		<-ch
		fmt.Println("Counting", i)
	}
}