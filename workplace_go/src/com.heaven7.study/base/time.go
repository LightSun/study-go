package main

import (
	"time"
	"fmt"
)

func main() {
	 //2016-02-25T12:00:00Z
	formate:="2006-01-02T15:04:05Z"
	now := time.Now()
	local1, err1 := time.LoadLocation("UTC") //输入参数"UTC"，等同于""
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") //本地local
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now.In(local1).Format(formate))
	fmt.Println(now.In(local2).Format(formate))
	fmt.Println(now.In(local3).Format(formate))
	//output:
	/*
	2017-06-12T07:12:40Z
	2017-06-12T15:12:40Z
	2017-06-12T00:12:40Z
	*/
}
