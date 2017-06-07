package main

import "fmt"
/**
下划线的意义
http://www.jianshu.com/p/309f55a152db
 */
type Element interface {}

func main() {
	var e Element = new(S);

	//type关键字不能和fallthrough 一起使用.
	//只能判断接口类型， 不会走结构体
	switch value := e.(type) {
	case S:
		fmt.Println("SSS", value)
		value.log("SSS")
	case T:
		fmt.Println("TTT", value)
	case I:
		fmt.Println("III", value)
		value.log("III")
	default:
		fmt.Println("unknown", value)
	}
}

//结构体内部不能有方法. 接口才有
type T struct{}
type S struct{
    age int
}

func ( s1 S )log(msg string){
     fmt.Println("msg = " + msg)
}

type I interface {
	log(msg string);
}
