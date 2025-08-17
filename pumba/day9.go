package main

import "fmt"

func add2(x int) {
	x += 10
	fmt.Println("add function", x)
}

func add_pointer(x *int) {
	*x += 10
	fmt.Println("add pointer", *x)
}

func main9() {
	//var x int = 10
	//add2(x)
	//fmt.Println(x)

	var a int = 10
	add_pointer(&a)
	fmt.Println(a)

	var msg string
	var msgPtr *string = &msg
	//fmt.Scan(&msg)
	fmt.Scan(msgPtr)
	fmt.Println(msg)
}
