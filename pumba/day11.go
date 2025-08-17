package main

import "fmt"

func main11() {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)
	fmt.Println(numbers[1] * 10)
	fmt.Println(len(numbers))

	var names [2]string = [2]string{"A", "B"}
	fmt.Println(names)
	fmt.Println(len(names))

	var grades [3]int = [3]int{60, 90, 75}
	var sum int
	for i := 0; i < len(grades); i++ {
		sum += grades[i]
	}
	fmt.Println(sum)

	var sum2 int
	for i, grade := range grades {
		fmt.Println(i, grade)
		sum2 += grade
	}
	fmt.Println(sum2)
}
