package hello_world

import (
	"fmt"
)

func Day9() {
	showArray()
	showSlice()
}

func showArray() {
	// 陣列的長度在宣告之後無法改變
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	b := [5]int{1, 2, 3} //單行宣告
	fmt.Println(b)

	c := [5]int{
		10,
		20,
		30,
		55, //使用多行宣告的話，最後一個元素要逗號
	}
	fmt.Println(c[1:3])

	d := [...]int{4, 6, 8} //用...省略符號，讓go判斷長度
	fmt.Printf("%T, %v\n", d, d)

	e := []int{1, 2} //初始化slice
	fmt.Printf("%T, %v\n", e, e)
}

func showSlice() {
	{
		fmt.Println("hello world")
	}
	{
		fmt.Println("hello world2")
	}
	/*
		Variable := make([]Type, Len, Cap)
		b := make([]int, 5, 10)
		宣告變數b為len:5、cap:10的整數切片。
	*/
	a := make([]int, 10) //設定 len:10。現在長度10了，容量雖然沒給，但最大容納長度當然不可能小於10吧，所以就是10了
	fmt.Println(a, len(a), cap(a), len(a) == 0, a == nil)

	b := make([]int, 5, 10) //設定 len:5、cap:10
	fmt.Println(b, len(b), cap(b), len(b) == 0, b == nil)

	var c = []int{} //初始化slice
	fmt.Println(c, len(c), cap(c), len(c) == 0, c == nil)

	var d []int //尚未實體化，此時等於nil
	fmt.Println(d, len(d), cap(d), len(d) == 0, d == nil)

	e := []string{"youtube", "golang"} //直接賦值
	fmt.Println(e, len(e), cap(e), len(e) == 0, e == nil)

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	// 因為空間夠，所以bb、cc共用aa後面的三個位子，所以當創建出cc的時候bb也會跟著更改
	aa := make([]int, 0, 10) //給足夠容量
	bb := append(aa, 1, 2, 3, 4)
	fmt.Println("aa:", aa, ", bb:", bb)
	cc := append(aa, 99, 88, 77)
	fmt.Println("aa:", aa, ", b:", bb, ", cc:", cc)
	fmt.Printf("%p, %p\n", &bb[0], &cc[0]) // 0xc00001a1e0, 0xc00001a1e0

	// 因為空間不夠，所以當append的時候就會觸發 growslice() 進行空間的擴充依照2的次方進行擴充最大到256
	// 超過256的時候會依照1.25倍來擴充容量
	// 以至於aaa bbb ccc三者的所指到的array是不同一個，並不會相互影響數值
	aaa := make([]int, 0, 2) //空間不夠的時候
	bbb := append(aaa, 1, 2, 3, 4)
	fmt.Println("aaa:", aaa, ", bbb:", bbb)
	ccc := append(aaa, 99, 88, 77)
	fmt.Println("aaa:", aaa, ", bbb:", bbb, ", ccc:", ccc)
	fmt.Printf("%p, %p\n", &bbb[0], &ccc[0]) // 0xc0000141e0, 0xc000014200
}
