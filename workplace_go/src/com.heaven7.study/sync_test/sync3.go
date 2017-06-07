package main

import "fmt"

func main() {
	chan1 := make(chan int, 1);
	chan2 := make(chan int, 2);

	select {
	case <-chan1:    // 如果chan1成功读到数据，则进行该case处理语句
	       fmt.Println( " chan1: ", chan1)
	case chan2 <- 1: // 如果成功向chan2写入数据，则进行该case处理语句
		fmt.Println( " chan2: " , chan2)
	default:         // 如果上面都没有成功，则进入default处理流程
	}
}


/**
可以看出，select不像switch，后面并不带判断条件，而是直接去查看case语句。每个
case语句都必须是一个面向channel的操作。比如上面的例子中，第一个case试图从chan1读取
一个数据并直接忽略读到的数据，而第二个case则是试图向chan2中写入一个整型数1，如果这
两者都没有成功，则到达default语句。
 */