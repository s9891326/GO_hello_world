package main

import "fmt"

func main2() {
	fmt.Println("hello world")

	//var x int
	//_, err := fmt.Scanln(&x)
	//if err != nil {
	//	return
	//}
	//fmt.Println(x)

	var x int
	var y int

	//fmt.Print("第一個")
	//fmt.Scan(&x)
	//fmt.Print("第二個")
	//fmt.Scan(&y)
	//fmt.Println(x + y)

	fmt.Print("第兩個數字，用空格隔開")
	fmt.Scan(&x, &y)
	fmt.Println(x + y)
}
