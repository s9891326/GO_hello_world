package main

import "fmt"

func main() {
	b := map[string]int{"TY": 123}
	b["A"] = 1
	fmt.Println(b)

	delete(b, "A")
	fmt.Println(b)

	value1, exists1 := b["TY"]
	if exists1 {
		fmt.Println(value1)
	}

	if value, exists := b["TY"]; exists {
		fmt.Println(value)
	}

	var h = make(map[string]int)
	fmt.Println(h)

	a := make(map[string]int)
	fmt.Println(a)
}
