package main

import "fmt"

type person struct {
	name   string
	height int
}

type group struct {
	name string
	person
}

func (p person) Greeting1() {
	fmt.Println("Hi~")
}

func (g group) Greeting2() {
	fmt.Println("Hi group")
}

func greeting(p person) {
	fmt.Println("p: ", p.name)
}

type T struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	p := person{"eddy", 123}
	fmt.Println(p)

	a := person{name: "eddy", height: 123}
	fmt.Println(a, a.name, a.height)

	g := group{"Line", person{"eddy", 123}}
	fmt.Println(g, g.height, g.name, g.person.name)

	p.Greeting1()
	g.Greeting2()
	greeting(p)
	//greeting(g)
}

// go build main.go
// ./main.exe
