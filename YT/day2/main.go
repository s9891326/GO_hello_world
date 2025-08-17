package main

import "fmt"

func main() {
	var x int
	var ans = 66

	for true {
		println("終極密碼(1~100):")
		_, err := fmt.Scan(x)
		if err != nil {
			print(err)
			return
		}

		if x == ans {
			println("答對了")
			break
		} else if x > ans {
			println("太大了")
		} else {
			println("太小了")
		}
	}
}
