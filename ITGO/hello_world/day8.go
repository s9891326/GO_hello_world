package hello_world

import "fmt"

func Day8() {
	var a int8 = 127 // ~128 ~ 127
	fmt.Println(a)

	var b uint8 = 255 // 少了負數 0 ~ 255
	fmt.Println(b)

	/*
		string中的字無法改變
		那怎摸辦餒？常見辦法有以下：

		用string[0:1]直接串文字
		用fmt.Sprint()串文字
		用Slice[] Append方法來接文字
	*/
	var s string = "hello world"
	s2 := "h2"
	fmt.Println("s:", s)
	fmt.Println("s:", s, ", s[0:5]:", s[0:8])

	s3 := s + "@" + s2
	fmt.Println(s3)

	s4 := fmt.Sprintf("%s@%s", s2, s3)
	fmt.Println(s4)

	// Byte用法非常的多，登場通常都會搭配Slice[]，
	var c = []byte("hello world")
	fmt.Println(c)

	c = []byte("爆肝工程師的異世界安魂曲")
	fmt.Println(c)

	/*
		rune意思是符號，簡單理解為UTF-8中的一個字符。
		因為Go語言預設使用UTF-8編碼，
		而UTF-8編碼中的每個字元長度不是固定的，以8bit為基本單位長度，範圍為1 ~ 4個字節(1 byte ~ 4 bytes)。
	*/
	var r = []rune("0")
	fmt.Println(r)

	r = []rune("爆肝工程師的異世界安魂曲")
	fmt.Println(r)

	a1 := "Hi,世界"
	fmt.Println(a1)
	n := 0
	for range a1 {
		n++
	}
	fmt.Println(len(a1)) // 9 (1 + 1 + 1 + 3 + 3) 英文算1 中文算3
	fmt.Println(n)       // 5
}
