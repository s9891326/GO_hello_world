package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	// 系統信息
	fmt.Println("=== Go 程序系統信息 ===")
	fmt.Printf("Go 版本: %s\n", runtime.Version())
	fmt.Printf("作業系統: %s\n", runtime.GOOS)
	fmt.Printf("CPU 架構: %s\n", runtime.GOARCH)
	fmt.Printf("CPU 核心數: %d\n", runtime.NumCPU())

	// 程序信息
	fmt.Println("\n=== 程序信息 ===")
	fmt.Printf("程序名稱: %s\n", os.Args[0])
	fmt.Printf("啟動時間: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// 環境變數
	fmt.Println("\n=== 環境變數 ===")
	fmt.Printf("GOPATH: %s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT: %s\n", os.Getenv("GOROOT"))

	// 命令行參數
	fmt.Println("\n=== 命令行參數 ===")
	for i, arg := range os.Args {
		fmt.Printf("參數 %d: %s\n", i, arg)
	}
}
