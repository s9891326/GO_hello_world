package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sum(max int) {
	result := 0
	for i := 0; i <= max; i++ {
		result += i
	}
	fmt.Println(result)
}

func main7() {
	fmt.Println("hello world")
	fmt.Println(add(3, 5))
	fmt.Println(add(-2, 55))
	sum(20)
}
