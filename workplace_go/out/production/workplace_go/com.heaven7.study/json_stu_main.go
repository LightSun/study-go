package main

import (//alia
	json_heaven7 "com.heaven7.study/jsontest"
	"fmt"
	"encoding/json"
)
/**
go 同一个目录下面不能．package多个.
 */
func main() {
	st := &json_heaven7.Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		35,
	}
	fmt.Println("before JSON encoding :")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println(string(b))
	}
	//concurrent
	//chan indicate a channel of concurrent
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		//<-”被称为接收操作符
		c <- str //send to c.
	}(ch, string(b))

	//fmt.Println("-------------- ch = " + ch + "------------------") //error
	strData := <-ch
	fmt.Println("--------------------------------")

	stb := &json_heaven7.Student{}
	stb.ShowStu()

	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
}

