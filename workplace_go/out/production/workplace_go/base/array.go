package main

import "fmt"

func main() {
	a := [] int {-1,1,2,3,4,5,6,7,8,9}
	s := a[0:5]  //从a数组中拷贝从Index= 0开始的5个元素

	fmt.Print(a, "a")
	fmt.Print(s, "s")

	s1 := append(s, 0)  //在s数组的末尾增加一个元素0
	fmt.Printf("=================\n")

	fmt.Print(a, "a")
	fmt.Print(s, "s")
	fmt.Println(s1, "s1")

	//定义了一个 int类型数组。元素个数20个。将index =19 的元素附值为1
	a2 := [20]int{19:1};
	fmt.Println(a2, "a2")

	a3 := [20]int{15:1};
	fmt.Println(a3, "a3")

	a4 := [3]int{1,2,3}
	var p * [3]int = &a4 //这种是指针数组 我们看到可以直接输出指向数组的指针

	x , y := 1 ,3
	a5 := [...]*int{&x ,&y}
	fmt.Println(a5) //输出这样[0xc080000068 0xc080000070]的地址 这就是数组指针
}
