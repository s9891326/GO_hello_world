package note

import "fmt"

func Array() {
	//var a [3]int = [3]int{}
	//fmt.Println(a)

	//var a = [...]int{1,2,3}
	//fmt.Println(a, len(a))
	//
	//for i := 0; i < len(a); i++ {
	//	fmt.Println("i: ", i)
	//}
	//
	//for i, j := range a {
	//	fmt.Println("i: ", i , " j: ", j)
	//}

	var twoDimensionArray [3][3]int = [3][3]int{
		{1, 2, 3},
		{0, 2, 5},
	}

	for i, v := range twoDimensionArray {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]: %v ", i, i2, v2)
		}
		fmt.Println()
	}
}

func main18() {
	Array()
}
