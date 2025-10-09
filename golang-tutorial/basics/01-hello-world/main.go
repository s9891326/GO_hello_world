package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本的 Hello World
	fmt.Println("Hello, World!")
	
	// 使用不同的 fmt 函數
	fmt.Print("這是 Print: ")
	fmt.Print("不會自動換行 ")
	fmt.Println("這是 Println: 會自動換行")
	
	// 格式化輸出
	name := "Go 語言學習者"
	fmt.Printf("歡迎 %s！\n", name)
	
	// 顯示當前時間
	now := time.Now()
	fmt.Printf("現在時間：%s\n", now.Format("2006-01-02 15:04:05"))
	
	// 多行字符串
	message := `
	恭喜您成功運行了第一個 Go 程序！
	Go 語言的學習之旅正式開始。
	`
	fmt.Println(message)
}