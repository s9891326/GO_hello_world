// 練習 3 解答：九九乘法表
package main

import "fmt"

func main() {
	fmt.Println("=== 九九乘法表 ===")
	
	for {
		fmt.Println("\n請選擇顯示模式：")
		fmt.Println("1. 完整九九乘法表")
		fmt.Println("2. 指定行的乘法表")
		fmt.Println("3. 三角形格式")
		fmt.Println("4. 表格格式")
		fmt.Println("0. 退出")
		fmt.Print("請選擇 (0-4): ")
		
		var choice int
		fmt.Scanf("%d", &choice)
		
		switch choice {
		case 1:
			printFullTable()
		case 2:
			printSpecificRow()
		case 3:
			printTriangleFormat()
		case 4:
			printTableFormat()
		case 0:
			fmt.Println("再見！")
			return
		default:
			fmt.Println("無效選擇，請重試！")
		}
	}
}

// 完整九九乘法表
func printFullTable() {
	fmt.Println("\n=== 完整九九乘法表 ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d×%d=%2d  ", i, j, i*j)
		}
		fmt.Println()
	}
}

// 指定行的乘法表
func printSpecificRow() {
	var row int
	fmt.Print("請輸入要顯示的行數 (1-9): ")
	fmt.Scanf("%d", &row)
	
	if row < 1 || row > 9 {
		fmt.Println("請輸入 1-9 之間的數字！")
		return
	}
	
	fmt.Printf("\n=== %d 的乘法表 ===\n", row)
	for j := 1; j <= 9; j++ {
		fmt.Printf("%d × %d = %d\n", row, j, row*j)
	}
}

// 三角形格式（下三角）
func printTriangleFormat() {
	fmt.Println("\n=== 三角形格式九九乘法表 ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%2d  ", j, i, j*i)
		}
		fmt.Println()
	}
}

// 表格格式
func printTableFormat() {
	fmt.Println("\n=== 表格格式九九乘法表 ===")
	
	// 打印表頭
	fmt.Print("   ")
	for j := 1; j <= 9; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Println()
	
	// 打印分隔線
	fmt.Print("   ")
	for j := 1; j <= 9; j++ {
		fmt.Print("----")
	}
	fmt.Println()
	
	// 打印表格內容
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d |", i)
		for j := 1; j <= 9; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}