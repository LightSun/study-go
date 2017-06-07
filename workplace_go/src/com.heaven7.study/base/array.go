package main

import "fmt"
/**
len（arr） 代表数组的长度
 */
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

	//a4 := [3]int{1,2,3}
	//var p * [3]int = &a4 //这种是指针数组 我们看到可以直接输出指向数组的指针

	x , y := 1 ,3
	a5 := [...]*int{&x ,&y}
	fmt.Println(a5) //输出这样[0xc080000068 0xc080000070]的地址 这就是数组指针

	VisitArray()
}

func VisitArray(){
	//最好不要用 arr := [...]int {9: 1}初始化，否则, 若某个函数接收可变个数的数组时，会报错.
	//arr := [...]int {9: 1} //未指定个数，则，这里9代表前面9个元素,1代表最后一个元素附值为1. 元素个数为9+1
	arr := []int {9: 1}
	fmt.Println(arr)
	fmt.Println(len(arr))

	//遍历数组
	//相比java: 省略括号
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	// range具有两个返回值，第一个返回值i是元素的数组下标，第二个返回值v是元素的值。
	for i, v := range arr {
		fmt.Println(i, v)
	}
	modifyArray(arr...);
	//modifyArray(1,2,3);
	modifyArray2(arr);
}
//cannot use arr (type [10]int) as type int in argument to modifyArray
//error: xxx [...]int
func modifyArray(xxx ...int ){
	xxx[0] = 10
	fmt.Println("In modify(), arr values:", xxx)
}
func modifyArray2(xxx []int ){
	xxx[0] = 10
	fmt.Println("In modify(), arr values:", xxx)
}
/*
//错误。不能这么写
func modifyArray3(xxx [...]int ){
	xxx[0] = 10
	fmt.Println("In modify(), arr values:", xxx)
}
*/

/**
// 常规声明方法
var a [5]byte //长度为5的数组，每个元素为一个字节
var b [2*N] struct { x, y int5 } //复杂类型数组
var c [5]*int // 指针数组
var d [2][3]int //二维数组
var e [2][3][4]int //等同于[2]([3]([4]int))

//直接声明并初始化
a := [3]byte{'1', '2', '3'} //声明并初始化一个长度为3的byte数组
a := [...]byte{'1', '2', '3'} //可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
d := [2][3]int{[3]int{1,2,3},[3]int{4,5,6}}
d := [2][3]int{{1,2,3},{4,5,6}} //如果内部的元素和外部的一样，那么上面的声明可以简化，直接忽略内部的
类型
 */