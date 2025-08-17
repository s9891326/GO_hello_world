package main

import "fmt"

func main8() {
	var x int = 5
	fmt.Println("原來的資料:", x)

	var xPtr *int = &x
	fmt.Println("記憶體位置:", xPtr)
	fmt.Println("x記憶體位置的記憶體位置:", &xPtr)

	fmt.Println("反解指標變數:", *xPtr)
}
