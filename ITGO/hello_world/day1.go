package hello_world

import "fmt"

// 變數: 能更改
var a = 123
var b int = 456

//b := 456  // 全域變數不能使用:=來宣告，只能用var、let

// 常數: 無法更改，單字通常為全大寫
const C = 789

/*
*
iota是希臘字符，在Golang中是關鍵字之一，用在宣告常數中，
效果為數字遞增，iota本身數值從0開始，
便於工程師不用手動打數字0、1、2、3...重複且無聊的事情。
*/
const (
	A = iota
	B
	CC
	DD
	E = iota * 0.1
	F
	G
	H
)

const (
	b1 = 1 << iota // 1  右側被塞入0個bit (2^0 二的零次方)
	b2             // 2  右側被塞入1個bit (2^1 二的一次方)
	b3             // 4  右側被塞入2個bit
	b4             // 8
	b5             // 16
)

func Day1() {
	fmt.Println("a", a)
	fmt.Println("Hello World!!")
	fmt.Println("b", b)
	b = 444
	fmt.Println("b", b)

	fmt.Println("C", C)
}
