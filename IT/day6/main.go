package main

import "fmt"

func main() {
	//s := make([]int, 5, 5)
	// make(T, length, capacity)
	s := make([]int, 0, 5)
	fmt.Println(s, len(s), cap(s), s[1:2])

	//s[0] = 12 // 無法填入元素，因為 s 的 length 不足

	s = append(s, 5)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 6)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 7)
	s = append(s, 8)
	s = append(s, 9)
	s = append(s, 10)

	fmt.Println(s, len(s), cap(s))
}
