package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	// 调用最基本的GET,并获得返回值
	resp, _ := http.Get("http://0.0.0.0:8888/test1")
	helpRead(resp)

	// 调用最基本的POST,并获得返回值
	resp, _ = http.Post("http://0.0.0.0:8888/test2", "", strings.NewReader(""))
	helpRead(resp)
}

// 用于读取resp的body
func helpRead(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}