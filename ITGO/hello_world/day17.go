package hello_world

import "fmt"

var aa = 50

func Day17() {
	/*
		『自始至終，就是想拖』—— Golang.Defer
		defer: 推遲執行defer的敘述式，盡可能地拖拖拖等等等，直到目前所在的func打烊結束要回傳時才做事。
		常用在: 關閉channel 及 關閉DB、讀檔與關檔寫在一起，這樣比較不容易忘記要關
	*/

	showDefer()
	showDeferValue()

	fmt.Println("a:", func2())
	fmt.Println("a:", func3())
}

func func2() int {
	var a int
	defer func() {
		a = 100
	}()
	return a // defer：『喔 要回傳a了喔，可是func還沒退出所以我不想做事，反正回上司也沒有規定要回傳哪個a，所以擺爛。』
}

func func3() (a int) {
	defer func() {
		a = 100
	}()
	return a //defer：『蛤，要回傳了喔？雖然想擺爛，但上司一開始指名規定要回傳a，先趕一下進度好了。』
}

func showDeferValue() {
	// defer的值: 全看上頭交代時的參數
	fmt.Println("assign1:", assign1(50)) // 100
	fmt.Println("assign2:", assign2(50)) // 100
}

func assign1(a int) int {
	defer fmt.Println(a) // 50
	a = 100
	return a
}

func assign2(a int) int {
	a = 100
	defer fmt.Println(a) // 100
	return a
}

func showDefer() {
	/*
		為了要貫徹拖延的行為，越早分派的任務要越晚達成才行
		hi
		aa:  50
		defer aa: 150
		退出main才執行
	*/
	defer fmt.Println("退出main才執行")
	fmt.Println("hi")

	defer func() {
		aa += 100
		fmt.Println("defer aa:", aa)
	}()
	fmt.Println("aa: ", aa)
}
