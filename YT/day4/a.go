package main

import "fmt"

func add(x int)  {
	x += 10
	fmt.Println(x)
}

func add2(x *int) {
	*x += 10
	fmt.Println(*x)
}

func main() {
	/*
	// 1. 建立存放資料的變數
	var x int = 3
	fmt.Println("原始的資料: ", x)

	// 2. 取得記憶體位置: &變數名稱
	// 3. 存放到指標變數， 型態: *資料型態
	var xx *int = &x
	fmt.Println("記憶體位置: ", xx)

	// 4.反解指標變數: *指標變數名稱
	fmt.Println("反解指標變數", *xx)
	*/

	fmt.Println("透過一般的參數進行加減")
	var a = 10
	add(a)
	fmt.Println("結果: ", a)

	fmt.Println("利用指標的方式進行加減")
	var a2 = 10
	add2(&a2)  // pass by pointer
	fmt.Println("結果: ", a2)
}
