package main

import "fmt"

func main4() {
	//if true {
	//	fmt.Println("hello world")
	//} else {
	//	fmt.Println("world")
	//}

	var money int
	fmt.Print("要領多少")
	fmt.Scan(&money)

	if money < 100 {
		fmt.Println("too few")
	} else if money < 100000 {
		fmt.Println("ok")
	} else {
		fmt.Println("too many")
	}
	fmt.Println("finish")
}
