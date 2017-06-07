package main

import "fmt"

func main() {
	fmt.Println("gfgd");

	c := new(Child)
	c.MingZi = "张小明"
	c.Name = "Tom Zhang"
	c.cName = "fdgfdgdgd";
	//c.Say(); //多继承冲突
	fmt.Println(c.Father.Say())
	fmt.Println(c.Mother.Say())
	fmt.Println(c.Cname())
}

//继承
type Father struct {
	MingZi string
}

func (this *Father) Say() string {
	return "大家好，我叫 " + this.MingZi
}

//有父亲当然有母亲，母亲是个外国人：
type Mother struct {
	Name string
}
func (this *Mother) Say() string {
	return "Hello, my name is " + this.Name
}

//父亲和母亲结合有了孩子类，孩子类继承了父亲和母亲
type Child struct {
	Father
	Mother

	cName string;
}
func (this *Child) Cname() string {
	return "Hello, my Child name is " + this.cName
}

type Formatter interface {
	GoString() string
}