package main

import "fmt"

func visit(friends []string, callback func(string)) {
	for _, n := range friends {
		callback(n)
	}
}

func double(x int) int {
	defer fmt.Println("hello defer")
	x = x * 2
	if x > 64 {
		return x
	}
	return double(x)
}

func hello_defer() {
	defer fmt.Println("hello defer")
	fmt.Println("hello")
}

func change(x *string) {
	*x = "Tom"
}

type stuff struct {
	name  string
	price int
}

func main() {
	//visit([]string{"Tina", "James", "Mary"}, func(s string) {
	//	fmt.Println(s)
	//})

	//fmt.Println(double(5))

	//hello_defer()

	//name := "a"
	//fmt.Println(&name)
	//fmt.Println(name)
	//change(&name)
	//fmt.Println(&name)
	//fmt.Println(name)

	p := stuff{"eddy", 55}
	fmt.Println(p)
	inprice(&p, 15)
	fmt.Println(p)

	//func() {
	//	fmt.Println("asdf")
	//}()
}

func inprice(p *stuff, variable int) {
	p.price += variable
}
