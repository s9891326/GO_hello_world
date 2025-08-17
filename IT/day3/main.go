package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//var p *int
	//a := 10
	//p = &a
	//
	//fmt.Println(p)
	//fmt.Println(*p)
	//fmt.Println(a)
	//fmt.Println(&a)

	a := 10
	fmt.Println(&a) // main裡面a的記憶體位置
	foo(a)
	fmt.Println(a)

	// &是取地址符，取到Person類別對象的地址
	//Bob := Person{"bob", 20}
	//fmt.Println("Bob: ", Bob, " &Bob: ", &Bob)
	//
	//// Person類型的指針
	//Lisa := &Person{"Lisa", 30}
	//fmt.Println("Lisa: ", Lisa, " &Lisa: ", &Lisa)
	//
	//// * 可以表達一個變量是指針類型
	//var John *Person = &Person{name: "john", age: 10}
	//fmt.Println("John: ", John, " &John: ", &John)
	//
	//// * 也可以表示指針類型變量所指向的存儲單元，也就是這個地址所指向的值
	//fmt.Println("*John: ", *John)
	//
	//changeName(John)
	//fmt.Println("John: ", John)
}

func foo(x int) {
	fmt.Println(&x) // function內x的記憶體位置
	x += 10
}

func changeName(p *Person) {
	// p 是一個指針
	// 基本類型的話 p.name = "C" == (*p).name = "C"
	(*p).name = "C"
	p.age = 100
}
