package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)
//字符串拼接
func main() {
	//第一种连接方法（最快）
	var buffer bytes.Buffer
	s := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("test is here\n")
	}
	buffer.String() // 拼接结果
	e := time.Now()
	fmt.Println("1 time is ", e.Sub(s).Seconds())

	//第二种方法
	s = time.Now()
	var sl []string
	for i := 0; i < 100000; i++ {
		sl = append(sl, "test is here\n")
	}
	strings.Join(sl, "")
	e = time.Now()
	fmt.Println("2 time is", e.Sub(s).Seconds())

	//第三种方法
	s = time.Now()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "test is here\n"
	}
	e = time.Now()
	fmt.Println("3 time is ", e.Sub(s).Seconds())

	//第四种方法
	s = time.Now()
	str4 := ""
	for i := 0; i < 100000; i++ {
		str4 = str4 + "test is here"
	}
	e = time.Now()
	fmt.Println("4 time is ", e.Sub(s).Seconds())
}
/**
运行结果如下

1 time is 0.00402775
2 time is 0.019025534
3 time is 7.162544528
4 time is 6.538371249



// bytes.Buffer的方法, 将给定字符串追加(append)到Buffer
func  (b *Buffer) WriteString(sstring) (nint, err error)

// 字符串拼接, 把slice通过给定的sep连接成一个字符串
func  Join(a []string, sepstring)  string

 */