package main

import "fmt"

/**
ok关键字查询
map:
     key 必须是支持==或 !=比较运算的类型, 不可以是函数. map或者slice.
     map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍.
     map使用 make（）创建， 支持:= 简写.
 */
func main() {

	var m map [int] string

	m = map[int] string{}
	fmt.Println(m)

	var n map[int]string
	n = make(map[int]string)
	fmt.Println(n)

	var b = make(map[int]string)
	fmt.Println(b)

	a := make(map[int]string)
	fmt.Println(a)
	a[1] = "ok"
	fmt.Println(a) //输出map
	aa := a[1]
	fmt.Println(aa) //输出map的键对应的值

	//复杂map  value 也是map, 使用map需要初始化
	var x map[int]  map[int]string
	x = make(map[int] map[int]string )
	x[1] = make(map[int]string)

	fmt.Println("----------------- 键为1的map ---------------")
	//键为1的map的键为1的值
	x[1][1] = "ok"
	xx := x[1][1]
	yy := x[1][2]
	fmt.Println(xx)
	fmt.Println(yy) //空字符串

	//p返回value,  ok返回bool类型, 查询键值对是否存在
	p, ok := x[1][2]
	if !ok {
		x[1][2] = "aaa"
		p, ok = x[1][2]
	}
	fmt.Println(p, ok)

	fmt.Println("--------------迭代-----------------")
	//以map为类型的slice,长度为5
	sm := make([]map[int]string, 5, 10)
	//迭代操作---索引、值
	//得到的v是拷贝，不会对sm造成影响
	for i, v := range sm { //, i-v,是key-value
		v = make(map[int]string, 2)
		v[1] = "ok1"
		v[2] = "ok2"
		fmt.Println(i, v)
	}
	fmt.Println(sm)

	fmt.Println("---------遍历切片slice类型的map集合,相当于foreach---------")
	//slice 。元素类型为map
	smm := make([]map[int]string, 5)
	for iii := range smm {
		//smm[iii] = make(map[int]string) 这样也可以
		smm[iii] = make(map[int]string, 1)
		smm[iii][1] = "aaaaaa"
		smm[iii][2] = "aaaddaaa"
		fmt.Println(iii, smm[iii])
	}
	fmt.Println(smm)

	fmt.Println("---------遍历map的键---- value 同理 -----")
	r := map[int]string {1: "a", 2: "b", 4: "c"} //初始化一个map
	sortKey := make([]int, len(r))               //初始化一个存map键的切片
	//遍历操作map-我需要的k，value不需要
	i := 0
	for k := range r {
		sortKey[i] = k
		i++
	}
	fmt.Println(sortKey)
}
