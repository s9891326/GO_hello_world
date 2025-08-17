package hello_world

import "fmt"

func Day13() {
	showIntPoint()
	testSwap()
	showCat()
	newFunc()
}

func newFunc() {
	var c = &Cat{name: "始祖貓"}
	fmt.Println(c, &c, *c)

	n1 := newCat("")
	n2 := newCat("複製貓三號")
	fmt.Println(n1, &n1, *n1)
	fmt.Println(n2, &n2, *n2)

	c2 := new(Cat)
	fmt.Println(c2, &c2, *c2)
}

func newCat(s string) *Cat {
	return &Cat{name: s}
}

func (c *Cat) rename(n string) {
	c.name = n
}

func showCat() {
	var cat1 = Cat{"肥貓一號"}
	fmt.Println(cat1)
	cat1.rename("肥貓一號2")
	cat1.eat()

	var cat2 = &Cat{"笨貓二號"}
	fmt.Println(cat2)
	cat2.rename("聰明貓三號") // 奇怪，怎麼改名失敗了
	cat2.eat()
}

func testSwap() {
	a, b := 1, 2
	//swap(a, b)
	swap(&a, &b)
	fmt.Printf("a:%d, b:%d\n", a, b)
}

//func swap(a int, b int) {
//	fmt.Println(a, b)
//	temp := a
//	a = b
//	b = temp
//}

func swap(a *int, b *int) {
	fmt.Println(a, b)
	temp := *a
	*a = *b
	*b = temp
}

func showIntPoint() {
	/*
		& 取變數的位址
		* 取變數的數值
	*/
	var a int = 100
	var b *int = &a
	fmt.Println(a, b, *b)

	var c string = "hello"
	var d *string = &c
	*d = "world" // 透過`*向址取值`的方式來改變變數裡面的內容值。
	fmt.Println(c, d, *d)

	x := 1
	p := &x                  //p(type: *int)指向x
	fmt.Println("x:", x)     //1
	fmt.Println("p:", p)     //p指向的位址
	fmt.Println("*p:", *p)   //p指向的位址的值，意即變數x
	fmt.Println("&p:", &p)   //p本身的位址
	y := &p                  //y(type:**int)y存放了 p本身的位址
	fmt.Println("**y:", **y) //到y取值(到p本身的位址取值) 再取值，意即變數x
	**y = 100
	fmt.Println("x:", x)
}
