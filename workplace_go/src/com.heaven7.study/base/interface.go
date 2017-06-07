package main

import (
	"fmt"
)

func main() {

	//定义一个 People interface 类型的变量p1
	var p1 People
	p1 = Person{"Rain", 23}
	p1.printMsg()           //I am Rain, and my age is 23.
	p1.drink("orange juice")//print result: Rain is drinking orange juice

	var p2 PeopleEat
	p2 = Person{"Sun", 24}
	p2.eat("chaffy dish") //print result: Sun is eating chaffy dish ...

	//不同类也可以实现同一个 interface
	var p3 PeopleEat
	p3 = Foodie{"James"}
	p3.eat("noodle")//print result: I am foodie, James. My favorite food is the noodle


	//interface 赋值
	p3 = p1  //p3 中的方法会被 p1 中的覆盖 (PeopleEat 被 People 覆盖)
	p3.eat("noodle")

	//interface 查询
	//将PeopleEat 转为 People 类型. p4为强制转化后的对象，ok为bool,  相当于转化后是否为null
	if p4, ok := p2.(People); ok {
		p4.drink("water") //调用 People interface 中有而 PeopleEat 中没有的方法
		fmt.Println(p4)
	}

}
//接口继承和实现

type Person struct {
	name string
	age int
}

func (p Person) printMsg() {
	fmt.Printf("Person: I am %s, and my age is %d.\n", p.name, p.age)
}

func (p Person) eat(s string) {
	fmt.Printf("Person: %s is eating %s ...\n", p.name, s)
}

func (p Person) drink(s string) {
	fmt.Printf("Person:  %s is drinking %s ...\n", p.name, s)
}

//people 继承了PeopleEat, PeopleEatDrink
type People interface {
	printMsg()
	PeopleEat    //组合
	PeopleDrink
	//eat() //不能出现重复的方法
}
/*
//与上面等价
type People interface {
    printMsg()
    eat()
    drink()
}
*/
type PeopleDrink interface {
	drink(s string)
}

type PeopleEat interface {
	eat(s string)
}

type PeopleEatDrink interface {
	eat(s string)
	drink(s string)
}
//===============================================
type Foodie struct {
	name string
}
func (p Foodie) eat(s string) {
	fmt.Printf("Foodie:  %s is eating %s ...\n", p.name, s)
}