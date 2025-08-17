package note

import "fmt"

func variableAndConstant() {
	fmt.Println("變量")
	var v1 int
	var v2 int = 2
	var v3 = 3
	v1 = 1
	v4 := 4
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v, v7=%v\n", v1, v2, v3, v4, v5, v6, v7)
	fmt.Println("常量")
	const (
		c1 = 8
		c2 = iota // 當前行數
		c3 = iota //
		c4        // 默認依照上面的形式(當前行數)
		c5 = 12
		c6
	)
	fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v\n", c1, c2, c3, c4, c5, c6)
}

func basicDataType() {
	fmt.Println("整數")
	var (
		n1      = 5
		n2 int8 = 127
		n3 uint16
		n4        = 0b111
		n5 int8   = 0o20
		n6 uint16 = 0xAF
	)
	fmt.Printf("n1=%v, type: %T\n", n1, n1)
	fmt.Printf("n2=%v, type: %T\n", n2, n2)
	fmt.Printf("n3=%v, type: %T\n", n3, n3)
	fmt.Printf("n4=%v, type: %T\n", n4, n4)
	fmt.Printf("n5=%v, type: %T\n", n5, n5)
	fmt.Printf("n6=%v, type: %T\n", n6, n6)

	fmt.Println("浮點數")
	var (
		f1         = 1.5
		f2 float32 = 1
	)
	fmt.Printf("f1=%v, type: %T\n", f1, f1)
	fmt.Printf("f2=%v, type: %T\n", f2, f2)

	fmt.Println("轉換格式")
	n5 = int8(n6)
	fmt.Printf("n5=%v, type: %T\n", n5, n5)

	fmt.Println("字串")
	var (
		c1 byte // uint8的別名
		c2      = '0'
		c3 rune = 456 // int32的別名
	)
	fmt.Printf("c1=%v, 直:%c type: %T\n", c1, c1, c1)
	fmt.Printf("c2=%v, 直:%c type: %T\n", c2, c2, c2)
	fmt.Printf("c3=%v, 直:%c type: %T\n", c3, c3, c3)
	c4 := 'A' - 'a'
	c5 := 'x'
	c6 := c5 + c4
	fmt.Printf("c6=%v, 直:%c type: %T\n", c6, c6, c6)

}

func increase(b *int) {
	*b++
	fmt.Printf("結束:%v, 內存地址:%v, 指向的值:%v\n", b, &b, *b)
}

func pointer() {
	var src = 111
	increase(&src)
	fmt.Printf("使用後:%v, 內存地址:%v\n", src, &src)
	var ptr = new(int)
	fmt.Printf("使用後:%v, 內存地址:%v, 指向的值:%v\n", ptr, &ptr, *ptr)
}

func main() {
	//variableAndConstant()
	//basicDataType()
	pointer()
}
