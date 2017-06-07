package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	// 使用post_form形式,注意必须要设置Post的type,同时此方法中忽略URL中带的参数,所有的参数需要从Body中获得
	resp,_ := http.Post("http://0.0.0.0:8888/test8", "application/x-www-form-urlencoded",strings.NewReader("message=8888888&extra=999999"))
	helpRead3(resp)
}

// 用于读取resp的body
func helpRead3(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}