package jsontest

import "fmt"

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func newStudent() *Student {
	return &Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
}

func (s *Student)ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	/*for _, a := range s.Classes {
		24
		fmt.Printf("%s ", a)
		25
	}*/
	fmt.Println("")
}
