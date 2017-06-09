package main

import (
	"fmt"
	"sort"
	"github.com/golang/protobuf/jsonpb/jsonpb_test_proto"
)
func main() {
	test := map[string]string{
		"c":"d123" ,
		"a": "b123",
		"b": "c123",
	};
	fmt.Println("before sort: " + test);
	sort.Sort(test);
	fmt.Println("after sort: " + test);
	jsonpb.Maps{}
}
