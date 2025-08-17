package hello_world

import "fmt"

func Day18() {
	showDotArray()
	showDotFunc()
	showSum()
	showAppendTwoArray()
}

func showAppendTwoArray() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	a1 := append(slice1, slice2[0], slice2[1], slice2[2])
	fmt.Println(a1, len(a1), cap(a1))

	a2 := append(slice1, slice2...)
	fmt.Println(a2, len(a2), cap(a2))

	a3 := slice1
	for _, num := range slice2 {
		a3 = append(a3, num)
	}
	fmt.Println(a3, len(a3), cap(a3))
}

func showSum() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(sumUnpacking(slice...))      // 把slice 解開、剝皮後傳入，同下
	fmt.Println(sumUnpacking(1, 2, 3, 4, 5)) // 可變參數函式
	fmt.Println(sumSlice(slice))             // 不曉得int長度，也可以直接包成一個slice型別來傳遞
}

func sumSlice(slice []int) (total int) {
	for _, num := range slice {
		total += num
	}
	return
}

func sumUnpacking(slice ...int) (total int) {
	for _, v := range slice {
		total += v
	}
	return
}

func showDotFunc() {
	test1("eddy", 1, 2, 3, 4, 5, 6, 7)
	test2(1, "a", "b", "c")

	s1 := sum()
	fmt.Println("s1:", s1)

	s2 := sum(1, 5, 9)
	fmt.Println("s2:", s2)
}

func showDotArray() {
	// [...] 省略長度是要編譯器幫我們計算Array陣列長度、自動將長度填入
	var a = [...]int{1, 2, 3}
	fmt.Println(a, len(a))
}

// ...在一個func參數中只能出現一次，而且是放在最後的參數位置上
func test1(a string, nums ...int) {
	fmt.Println(a, nums)
}

// 省略符號只能放在最後的位置上
func test2(a int, nums ...string) {
	fmt.Println(a, nums)
}

func sum(nums ...int) int {
	fmt.Printf("%T\n", nums)
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

// 這個會報錯 Can only use '...' as final argument in list
//func test3(...int, ...string) {
//
//}
