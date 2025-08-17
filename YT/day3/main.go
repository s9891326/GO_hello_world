package main

import "fmt"

func main() {
	var result = add(3, 4)
	fmt.Println(result)

	var r1 int
	var r2 int
	r1, r2 = add2(10, 45)
	fmt.Println(r1, r2)
}

func add(i int, i2 int) int {
	return i + i2
}

func add2(i int, i2 int) (int, int) {
	return i + i2, i - i2
}
