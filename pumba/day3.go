package main

import "fmt"

func main3() {
	var x int
	x = 3*3 + 10
	fmt.Println(x)

	x = 5
	x -= 2
	fmt.Println(x)

	x++
	fmt.Println(x)

	var result bool = 4 > 3
	fmt.Println(result)
}
