package main

import "fmt"

/**
slice本身是一个对数组array的封装

	slice 是对底层数组的抽象和控制。它包含 Go 需要对底层数组管理的三种元数据，分别是：

	指向底层数组的指针
	slice 中元素的长度
	slice 的容量(可供增长的最大值)

make 内建函数：
make([]T, len)
make([]T, len, capability) // same as make([]T capability)[:len]  //slice指定初始化容量

append函数: 给指定的slice添加 一定的元素，原来的slice不变，返回新的slice.
copy(dst, src):
    copy()函数是把一个slice的“所有”元素拷贝到另外一个slice中。——是拷贝，而不是追加(append)。
    需要注意的是：如果目标slice的长度小于源slice的长度，那么就不会拷贝源slice的所有原始。
 */
func main() {
	fmt.Println("dsfsdfs");
	//内建的函数 make
	//slice := make([]string, 5);
	//slice := make([]int, 3, 5)

	//不允许创建长度大于容量的 slice：
	//slice := make([]int, 5, 3)//error
	//=============== start test ==============
	//slice2();
	Slice_copy();
}

func Slice_copy()  {
	slice1 := []int {1,2,3}
	slice2 := make([]int, 2) //2 为长度
	slice3 := make([]int, 4)

	slice3[0] = 10
	debug_slice(slice1, "slice1:")
	debug_slice(slice2, "slice2:")
	debug_slice(slice3, "slice3:")

	//将slice1硬拷贝到slice2
	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("------------------------ after copy -----------------------")
	debug_slice(slice1, "slice1:")
	debug_slice(slice2, "slice2:")
	debug_slice(slice3, "slice3:")
}

func Slice1() {
	//go slice 相当于java的 ArrayList
	var values []int // (1)

	values = make([]int, 5)  // (2)
	//同数组
	for i := range values {
		values[i] = i + 1
	}
	for i, item := range values {
		fmt.Printf("values[%d]=%d\n", i, item)
	}
}
/**
append函数: 给指定的slice添加 一定的元素，原来的slice不变，返回新的slice.
 */
func slice2()  {
	var the_slice []int
	the_slice = make([]int, 5, 10)
	debug_slice(the_slice, "length = 5, capacity = 10");

	the_slice1 := []int {1,2,3}
	the_slice2 := append(the_slice1, 4, 5)

	//append后 原来的slice并无变化
	debug_slice(the_slice1, "slice1:")
	debug_slice(the_slice2, "slice2:")

	the_slice1 = append(the_slice1, 4, 5, 6)
	debug_slice(the_slice1, "slice1:")
	debug_slice(the_slice2, "slice2:")

}
func debug_slice(the_slice []int, msg string) {
	fmt.Println(msg)
	for i, item := range the_slice {
		fmt.Printf("the_slice[%d]=%d\n", i, item)
	}
}