package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	s := fmt.Sprintf("名字叫 %s ; 年龄为 %d", p.name, p.age)
	return s
}

type family struct {
	mom person
	dad person
}

func (f family) String() string {
	return fmt.Sprint("爸爸是" + f.dad.String() + " 妈妈是" + f.mom.String())
}

func main() {
	dad := person{age: 20, name: "xmh"}
	mom := person{age: 19, name: "lhk"}
	family := family{dad: dad, mom: mom}
	fmt.Println(family)
	mom.age = 18
	fmt.Println(family)

	mom2 := &person{age: 19, name: "lhk"}
	fmt.Println(mom2)

}
