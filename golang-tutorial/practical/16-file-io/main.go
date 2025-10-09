package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 演示基本文件操作
func demonstrateBasicFileOperations() {
	fmt.Println("=== 基本文件操作 ===")
	
	// 1. 創建和寫入文件
	fmt.Println("\n--- 創建和寫入文件 ---")
	filename := "example.txt"
	content := "Hello, Go File I/O!\n這是第二行文本。\n這是第三行文本。"
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("寫入文件錯誤: %v\n", err)
		return
	}
	fmt.Printf("成功創建文件: %s\n", filename)
	
	// 2. 讀取整個文件
	fmt.Println("\n--- 讀取整個文件 ---")
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("讀取文件錯誤: %v\n", err)
		return
	}
	fmt.Printf("文件內容:\n%s\n", string(data))
	
	// 3. 獲取文件信息
	fmt.Println("\n--- 文件信息 ---")
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("獲取文件信息錯誤: %v\n", err)
		return
	}
	
	fmt.Printf("文件名: %s\n", info.Name())
	fmt.Printf("文件大小: %d 字節\n", info.Size())
	fmt.Printf("文件權限: %s\n", info.Mode())
	fmt.Printf("修改時間: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("是否目錄: %t\n", info.IsDir())
}

// 演示按行讀取文件
func demonstrateLineByLineReading() {
	fmt.Println("\n=== 按行讀取文件 ===")
	
	filename := "example.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打開文件錯誤: %v\n", err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("行 %d: %s\n", lineNumber, line)
		lineNumber++
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("讀取文件錯誤: %v\n", err)
	}
}

// 演示追加寫入文件
func demonstrateAppendFile() {
	fmt.Println("\n=== 追加寫入文件 ===")
	
	filename := "example.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打開文件錯誤: %v\n", err)
		return
	}
	defer file.Close()
	
	appendContent := "\n這是追加的內容"
	_, err = file.WriteString(appendContent)
	if err != nil {
		fmt.Printf("追加寫入錯誤: %v\n", err)
		return
	}
	
	fmt.Println("成功追加內容到文件")
	
	// 重新讀取文件顯示結果
	data, _ := os.ReadFile(filename)
	fmt.Printf("更新後的文件內容:\n%s\n", string(data))
}

// 演示緩衝寫入
func demonstrateBufferedWrite() {
	fmt.Println("\n=== 緩衝寫入 ===")
	
	filename := "buffered_example.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("創建文件錯誤: %v\n", err)
		return
	}
	defer file.Close()
	
	// 創建緩衝寫入器
	writer := bufio.NewWriter(file)
	defer writer.Flush() // 確保緩衝區內容被寫入
	
	// 寫入多行數據
	lines := []string{
		"第一行數據",
		"第二行數據",
		"第三行數據",
		"第四行數據",
		"第五行數據",
	}
	
	for i, line := range lines {
		_, err := writer.WriteString(fmt.Sprintf("%d. %s\n", i+1, line))
		if err != nil {
			fmt.Printf("寫入錯誤: %v\n", err)
			return
		}
	}
	
	fmt.Printf("成功使用緩衝寫入創建文件: %s\n", filename)
}

// 演示目錄操作
func demonstrateDirectoryOperations() {
	fmt.Println("\n=== 目錄操作 ===")
	
	// 1. 創建目錄
	fmt.Println("\n--- 創建目錄 ---")
	dirName := "testdir"
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("創建目錄錯誤: %v\n", err)
		return
	}
	fmt.Printf("創建目錄: %s\n", dirName)
	
	// 創建嵌套目錄
	nestedDir := "testdir/nested/deep"
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("創建嵌套目錄錯誤: %v\n", err)
		return
	}
	fmt.Printf("創建嵌套目錄: %s\n", nestedDir)
	
	// 2. 在目錄中創建文件
	fmt.Println("\n--- 在目錄中創建文件 ---")
	testFile := filepath.Join(dirName, "test.txt")
	err = os.WriteFile(testFile, []byte("這是測試目錄中的文件"), 0644)
	if err != nil {
		fmt.Printf("創建文件錯誤: %v\n", err)
		return
	}
	fmt.Printf("在目錄中創建文件: %s\n", testFile)
	
	// 3. 列出目錄內容
	fmt.Println("\n--- 列出目錄內容 ---")
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("讀取目錄錯誤: %v\n", err)
		return
	}
	
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		
		fileType := "文件"
		if entry.IsDir() {
			fileType = "目錄"
		}
		
		fmt.Printf("  %s: %s (%d 字節)\n", fileType, entry.Name(), info.Size())
	}
}

// 演示文件複製
func demonstrateFileCopy() {
	fmt.Println("\n=== 文件複製 ===")
	
	sourceFile := "example.txt"
	destFile := "copied_example.txt"
	
	// 打開源文件
	src, err := os.Open(sourceFile)
	if err != nil {
		fmt.Printf("打開源文件錯誤: %v\n", err)
		return
	}
	defer src.Close()
	
	// 創建目標文件
	dst, err := os.Create(destFile)
	if err != nil {
		fmt.Printf("創建目標文件錯誤: %v\n", err)
		return
	}
	defer dst.Close()
	
	// 複製文件內容
	bytesWritten, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("複製文件錯誤: %v\n", err)
		return
	}
	
	fmt.Printf("成功複製 %d 字節從 %s 到 %s\n", bytesWritten, sourceFile, destFile)
}

// 演示文件查找和遍歷
func demonstrateFileWalking() {
	fmt.Println("\n=== 文件遍歷 ===")
	
	fmt.Println("遍歷當前目錄及子目錄中的所有文件:")
	
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// 跳過隱藏文件和目錄
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		
		fileType := "文件"
		if info.IsDir() {
			fileType = "目錄"
		}
		
		fmt.Printf("  %s: %s (%d 字節) - %s\n", 
			fileType, path, info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))
		
		return nil
	})
	
	if err != nil {
		fmt.Printf("遍歷目錄錯誤: %v\n", err)
	}
}

// 演示文件權限檢查
func demonstrateFilePermissions() {
	fmt.Println("\n=== 文件權限檢查 ===")
	
	filename := "example.txt"
	
	// 檢查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("文件不存在: %s\n", filename)
		return
	}
	
	fmt.Printf("文件 %s 存在\n", filename)
	
	// 嘗試以不同模式打開文件
	modes := []struct {
		flag int
		name string
	}{
		{os.O_RDONLY, "只讀"},
		{os.O_WRONLY, "只寫"},
		{os.O_RDWR, "讀寫"},
	}
	
	for _, mode := range modes {
		file, err := os.OpenFile(filename, mode.flag, 0)
		if err != nil {
			fmt.Printf("  %s模式: 無法打開 (%v)\n", mode.name, err)
		} else {
			fmt.Printf("  %s模式: 可以打開\n", mode.name)
			file.Close()
		}
	}
}

// 清理函數
func cleanup() {
	fmt.Println("\n=== 清理臨時文件 ===")
	
	// 清理創建的文件和目錄
	filesToRemove := []string{
		"example.txt",
		"buffered_example.txt",
		"copied_example.txt",
	}
	
	for _, file := range filesToRemove {
		if err := os.Remove(file); err != nil {
			fmt.Printf("刪除文件 %s 錯誤: %v\n", file, err)
		} else {
			fmt.Printf("已刪除文件: %s\n", file)
		}
	}
	
	// 刪除測試目錄
	if err := os.RemoveAll("testdir"); err != nil {
		fmt.Printf("刪除目錄 testdir 錯誤: %v\n", err)
	} else {
		fmt.Println("已刪除目錄: testdir")
	}
}

// 主函數
func main() {
	fmt.Println("===== Go 文件操作示例 =====")
	
	// 演示各種文件操作
	demonstrateBasicFileOperations()
	demonstrateLineByLineReading()
	demonstrateAppendFile()
	demonstrateBufferedWrite()
	demonstrateDirectoryOperations()
	demonstrateFileCopy()
	demonstrateFileWalking()
	demonstrateFilePermissions()
	
	// 等待用戶輸入再清理
	fmt.Println("\n按 Enter 鍵繼續清理臨時文件...")
	fmt.Scanln()
	
	cleanup()
	
	fmt.Println("\n===== 示例完成 =====")
}