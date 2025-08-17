package main

import "fmt"

func main() {
	var x int = 10
	var y int

	fmt.Println("輸入兩個數字，並用空格隔開")
	_scanln, err := fmt.Scanln(&x, &y)
	if err != nil {
		return
	}
	println(_scanln)
	println(x + y)
}

