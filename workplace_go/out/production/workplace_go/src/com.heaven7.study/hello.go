package main

import (
	"fmt"
)
/**
go数据类型：http://studygolang.com/articles/9929
 */
func main() {
	fmt.Println("Hello World")
	var x,y int = swap(1,2);
	fmt.Println(x,y);
}
//函数返回多个
func swap(a, b int) (int, int) {
	return b, a
}
/**
命名返回值
命名返回值有点像某些语言的out参数。使用命名返回值的话，return语句就不需要添加返回值了，因为返回值已经在函数签名上指定了。
 */
func divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
//递归函数
func facterial(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return facterial(n-1) * n
	}
}