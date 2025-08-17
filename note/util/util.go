package util

import "fmt"

var i = 0
var F = func(s string) int {
	fmt.Printf("被%s: 使用\n", s)
	i++
	fmt.Println("第幾次使用: ", i)
	return i
}

func SelectByKey(text ...string) (key int) {
	for i, s := range text {
		fmt.Printf("i=%d: %s\n", i+1, s)
	}
	fmt.Println("請選擇")
	fmt.Scanln(&key)
	return
}
