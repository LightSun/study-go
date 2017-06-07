package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

/**
defer : means called before exit. same as java: try--finally
 */
func main() {
	fmt.Println(defer1())
	fmt.Println(defer2())
	fmt.Println(defer3())
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

//1
func defer1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//5
func defer2() (r int) {
	t := 5
	defer func() { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		t = t + 5
	}()
	return t
}

//1
func defer3() (r int) {
	defer func(r int) { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		r = r + 5
	}(r)
	return 1
}