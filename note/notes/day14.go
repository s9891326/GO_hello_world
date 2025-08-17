package note

import "fmt"

func IfElse() {
	var age int8
	fmt.Println("輸入年齡")
	fmt.Scanln(&age)

	if age < 13 {
		fmt.Println("<13")
	} else if age < 25 {
		fmt.Println("<25")
	} else {
		fmt.Println(">=25")
	}

	if l := 3; l < 2 {
		fmt.Println("<2")
	} else {
		fmt.Println(">=2")
	}
}


func Label() {
	outside:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if i == 5 && j == 4 {
					break outside
				}
				fmt.Print(" + ")
			}
			fmt.Println()
		}

	// 如果單純只用break會只往外跳了一層，使用label的break才能真正地跳到最外層
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		if i == 5 && j == 4 {
	//			break
	//		}
	//		fmt.Print(" - ")
	//	}
	//	fmt.Println()
	//}
}

func main14() {
	//IfElse()
	Label()
}
