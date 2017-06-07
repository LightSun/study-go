package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	// GET传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp,_ = http.Get("http://0.0.0.0:8888/test3/name=TAO/passwd=123")
	helpRead2(resp)

	// POST传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp,_ = http.Post("http://0.0.0.0:8888/test4/name=PT/passwd=456", "",strings.NewReader(""))
	helpRead2(resp)

	// 注意Param中':'和'*'的区别
	resp,_ = http.Get("http://0.0.0.0:8888/test5/name=TAO/passwd=789")
	helpRead2(resp)
	resp,_ = http.Get("http://0.0.0.0:8888/test5/name=TAO/")
	helpRead2(resp)
}

// 用于读取resp的body
func helpRead2(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}