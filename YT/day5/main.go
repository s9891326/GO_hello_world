package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var p1 Person = Person{"eddy", 12}
	fmt.Println(p1)

	var p2 Person = Person{name:"eddy", age: 5555}
	fmt.Println(p2)

	// 整數陣列
	var num [3]int
	fmt.Println(num)

	num[0] = 1
	num[1] = 2
	num[2] = 3
	fmt.Println(num)

	var names [2]string = [2]string{"a", "b"}
	fmt.Println(names)
	fmt.Println(len(names))

	var grades [3]int = [3]int{61, 1, 3}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
