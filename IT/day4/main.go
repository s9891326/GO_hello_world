package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [10]int
	fmt.Println(a, len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	x := [5]int{}
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	for _, n := range x {
		fmt.Println(n)
	}

	xx := [5]int{1, 2, 3}
	fmt.Println(xx)

}
