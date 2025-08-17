package main

import "fmt"

func main5() {
	var x int = 0
	for x < 3 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
}
