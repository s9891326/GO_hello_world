package main

import (
	"fmt"
	"time"
)

// 包級別常數
const (
	AppName    = "Golang Tutorial"
	AppVersion = "1.0.0"
)

// 使用 iota 定義枚舉
const (
	StatusPending = iota
	StatusRunning
	StatusCompleted
	StatusFailed
)

// 包級別變數
var globalCounter int

func main() {
	fmt.Println("=== Go 變數和常數示例 ===")

	// 1. 不同的變數聲明方式
	demonstrateVariableDeclaration()

	// 2. 零值演示
	demonstrateZeroValues()

	// 3. 常數使用
	demonstrateConstants()

	// 4. 作用域演示
	demonstrateScope()

	// 5. 變數命名演示
	demonstrateNaming()
}

func demonstrateVariableDeclaration() {
	fmt.Println("\n--- 變數聲明方式 ---")

	// 方式 1: 標準聲明
	var name string
	name = "張三"

	// 方式 2: 聲明並初始化
	var age int = 25

	// 方式 3: 類型推導
	var isStudent = true

	// 方式 4: 短變數聲明
	city := "台北"

	// 多變數聲明
	var x, y, z = 1, 2, 3
	a, b, c := "Hello", 42, true

	fmt.Printf("姓名: %s, 年齡: %d, 學生: %t, 城市: %s\n", name, age, isStudent, city)
	fmt.Printf("xyz: %d %d %d\n", x, y, z)
	fmt.Printf("abc: %s %d %t\n", a, b, c)
}

func demonstrateZeroValues() {
	fmt.Println("\n--- 零值演示 ---")

	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("int 零值: %d\n", i)
	fmt.Printf("float64 零值: %f\n", f)
	fmt.Printf("bool 零值: %t\n", b)
	fmt.Printf("string 零值: '%s' (長度: %d)\n", s, len(s))
}

func demonstrateConstants() {
	fmt.Println("\n--- 常數演示 ---")

	fmt.Printf("應用名稱: %s\n", AppName)
	fmt.Printf("應用版本: %s\n", AppVersion)

	// 本地常數
	const pi = 3.14159
	const greeting = "歡迎使用 Go!"

	fmt.Printf("圓周率: %.10f\n", pi)
	fmt.Println(greeting)

	// 枚舉狀態
	status := StatusRunning
	fmt.Printf("當前狀態: %d\n", status)

	// 文件大小常數
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
}

func demonstrateScope() {
	fmt.Println("\n--- 作用域演示 ---")

	// 函數級別變數
	localVar := "我是局部變數"
	globalCounter++

	fmt.Printf("局部變數: %s\n", localVar)
	fmt.Printf("全局計數器: %d\n", globalCounter)

	// 塊級作用域
	if true {
		blockVar := "我是塊級變數"
		localVar = "局部變數被修改了"
		fmt.Printf("塊內變數: %s\n", blockVar)
		fmt.Printf("修改後的局部變數: %s\n", localVar)
	}

	// blockVar 在此處無法訪問
	fmt.Printf("塊外的局部變數: %s\n", localVar)
}

func demonstrateNaming() {
	fmt.Println("\n--- 變數命名演示 ---")

	// 好的命名示例
	userName := "alice"
	userAge := 30
	isUserActive := true
	maxRetryCount := 5
	serverTimeout := 30 * time.Second

	fmt.Printf("用戶名: %s\n", userName)
	fmt.Printf("用戶年齡: %d\n", userAge)
	fmt.Printf("用戶活躍: %t\n", isUserActive)
	fmt.Printf("最大重試次數: %d\n", maxRetryCount)
	fmt.Printf("服務器超時: %v\n", serverTimeout)

	// 縮寫在上下文明確時是可以的
	var (
		id   = 12345
		url  = "https://example.com"
		json = `{"name": "test"}`
	)

	fmt.Printf("ID: %d, URL: %s\n", id, url)
	fmt.Printf("JSON: %s\n", json)
}
