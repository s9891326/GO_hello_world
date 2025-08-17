package day1

import "fmt"

//package main

var EddyName = "eddy"
var eddyName = "eddy2"

func foo(name string) string {
	return name + "asdf"
}

func main() {
	fmt.Println(foo("name test"))
}
