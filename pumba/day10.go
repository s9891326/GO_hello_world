package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main10() {
	p1 := Person{"EDDY", 18}
	fmt.Println(p1.name, p1.age)

	p2 := Person{age: 20, name: "aaaaa"}
	fmt.Println(p2)
	p2.name = "edddy222"
	fmt.Println(p2)
}
