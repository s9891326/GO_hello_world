package main

import "fmt"

func main6() {
	var x int
	for x < 5 {
		if x == 3 {
			break
		}
		fmt.Println(x)
		x++
	}

	for x := 0; x < 5; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}

	var result int = 0
	for true {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		result += n
	}
	fmt.Println(result)
}
