package main

import (
	"fmt"
)

type Element interface {}

//type关键字
func main() {
	var e Element = 100
	switch value := e.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
	}
}
