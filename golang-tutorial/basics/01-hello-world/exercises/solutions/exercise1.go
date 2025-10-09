// 練習 1 解答：個人介紹程序
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func exercise1() {
	name := "張小明"
	age := 28
	job := "Go 開發工程師"
	hobbies := []string{"程式設計", "閱讀技術書籍", "爬山", "攝影"}

	// 格式化輸出個人信息
	fmt.Printf("=== 個人介紹 ===\n")
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("年齡：%d 歲\n", age)
	fmt.Printf("職業：%s\n", job)
	fmt.Print("興趣：")

	// 輸出興趣愛好
	for i, hobby := range hobbies {
		if i == len(hobbies)-1 {
			fmt.Printf("%s\n", hobby)
		} else {
			fmt.Printf("%s、", hobby)
		}
	}
}

func exercise2() {
	now := time.Now()

	// 顯示當前日期和時間
	fmt.Printf("當前時間：%s\n", now.Format("2006-01-02 15:04:05"))

	// 顯示星期幾
	weekdays := map[time.Weekday]string{
		time.Sunday:    "星期日",
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
	}
	fmt.Printf("今天是：%s\n", weekdays[now.Weekday()])

	// 計算距離新年還有多少天
	currentYear := now.Year()
	nextNewYear := time.Date(currentYear+1, 1, 1, 0, 0, 0, 0, now.Location())
	daysUntilNewYear := int(nextNewYear.Sub(now).Hours() / 24)

	fmt.Printf("距離新年還有：%d 天\n", daysUntilNewYear)
}

func exercise3() {
	/*
		創建一個簡單的計算器：
		- 定義兩個數字變數
		- 計算並輸出加、減、乘、除的結果
		- 使用格式化輸出顯示計算過程
	*/

	a := 10
	b := 50

	fmt.Println("創建一個簡單的計算器： a = ", a, ", b = ", b)
	fmt.Println("加法計算 a + b ：", a+b)
	fmt.Println("減法計算 a - b：", a-b)
	fmt.Println("乘法計算 a * b：", a*b)
	fmt.Println("除法計算 a / b：", a/b)
	fmt.Println("餘數計算 a % b：", a%b)
}

func exercise4() {
	/*
		創建一個程序顯示：
		- Go 版本信息
		- 作業系統信息
		- 當前工作目錄
		- 程序執行時的命令行參數
	*/
	fmt.Println("Go 版本信息：", runtime.Version()) // 假設的版本信息
	fmt.Println("作業系統信息：", runtime.GOOS)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("獲取當前工作目錄失敗:", err)
	}
	fmt.Println("當前工作目錄:", dir)

	fmt.Println("程序執行時的命令行參數：")
	for i, arg := range os.Args {
		fmt.Printf("參數 %d: %s\n", i, arg)
	}
}

func main() {
	// 定義個人信息變數
	fmt.Println("練習 1：個人介紹程序")
	exercise1()

	fmt.Println("練習 2：時間顯示程序")
	exercise2()

	fmt.Println("練習 3：計算器程序")
	exercise3()

	fmt.Println("練習 4：系統信息程序")
	exercise4()
}
