// 練習 1 解答：基本文件讀寫 - 簡單文本編輯器
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SimpleEditor 簡單文本編輯器
type SimpleEditor struct {
	filename string
}

// NewSimpleEditor 創建新的文本編輯器
func NewSimpleEditor(filename string) *SimpleEditor {
	return &SimpleEditor{
		filename: filename,
	}
}

// Create 創建新文件並寫入內容
func (se *SimpleEditor) Create(content string) error {
	return os.WriteFile(se.filename, []byte(content), 0644)
}

// Read 讀取整個文件內容
func (se *SimpleEditor) Read() (string, error) {
	data, err := os.ReadFile(se.filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Append 追加內容到文件末尾
func (se *SimpleEditor) Append(content string) error {
	file, err := os.OpenFile(se.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = file.WriteString(content)
	return err
}

// WriteAt 在指定行插入內容
func (se *SimpleEditor) WriteAt(line int, content string) error {
	// 讀取所有行
	lines, err := se.readLines()
	if err != nil {
		return err
	}
	
	// 檢查行號是否有效
	if line < 1 || line > len(lines)+1 {
		return fmt.Errorf("行號 %d 無效，文件共有 %d 行", line, len(lines))
	}
	
	// 插入新行
	newLines := make([]string, 0, len(lines)+1)
	
	// 添加插入位置之前的行
	for i := 0; i < line-1 && i < len(lines); i++ {
		newLines = append(newLines, lines[i])
	}
	
	// 添加新行
	newLines = append(newLines, content)
	
	// 添加插入位置之後的行
	for i := line - 1; i < len(lines); i++ {
		newLines = append(newLines, lines[i])
	}
	
	// 寫回文件
	return se.writeLines(newLines)
}

// GetLineCount 獲取文件總行數
func (se *SimpleEditor) GetLineCount() (int, error) {
	lines, err := se.readLines()
	if err != nil {
		return 0, err
	}
	return len(lines), nil
}

// readLines 讀取文件所有行
func (se *SimpleEditor) readLines() ([]string, error) {
	file, err := os.Open(se.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	return lines, scanner.Err()
}

// writeLines 寫入所有行到文件
func (se *SimpleEditor) writeLines(lines []string) error {
	content := strings.Join(lines, "\n")
	return os.WriteFile(se.filename, []byte(content), 0644)
}

// 額外功能方法

// Replace 替換文件中的文本
func (se *SimpleEditor) Replace(oldText, newText string) error {
	content, err := se.Read()
	if err != nil {
		return err
	}
	
	newContent := strings.ReplaceAll(content, oldText, newText)
	return os.WriteFile(se.filename, []byte(newContent), 0644)
}

// DeleteLine 刪除指定行
func (se *SimpleEditor) DeleteLine(line int) error {
	lines, err := se.readLines()
	if err != nil {
		return err
	}
	
	if line < 1 || line > len(lines) {
		return fmt.Errorf("行號 %d 無效，文件共有 %d 行", line, len(lines))
	}
	
	// 刪除指定行
	newLines := make([]string, 0, len(lines)-1)
	for i, l := range lines {
		if i+1 != line {
			newLines = append(newLines, l)
		}
	}
	
	return se.writeLines(newLines)
}

// GetLine 獲取指定行的內容
func (se *SimpleEditor) GetLine(line int) (string, error) {
	lines, err := se.readLines()
	if err != nil {
		return "", err
	}
	
	if line < 1 || line > len(lines) {
		return "", fmt.Errorf("行號 %d 無效，文件共有 %d 行", line, len(lines))
	}
	
	return lines[line-1], nil
}

// UpdateLine 更新指定行的內容
func (se *SimpleEditor) UpdateLine(line int, content string) error {
	lines, err := se.readLines()
	if err != nil {
		return err
	}
	
	if line < 1 || line > len(lines) {
		return fmt.Errorf("行號 %d 無效，文件共有 %d 行", line, len(lines))
	}
	
	lines[line-1] = content
	return se.writeLines(lines)
}

// Search 搜索包含指定文本的行
func (se *SimpleEditor) Search(searchText string) ([]int, error) {
	lines, err := se.readLines()
	if err != nil {
		return nil, err
	}
	
	var matchingLines []int
	for i, line := range lines {
		if strings.Contains(line, searchText) {
			matchingLines = append(matchingLines, i+1)
		}
	}
	
	return matchingLines, nil
}

// Backup 創建文件備份
func (se *SimpleEditor) Backup() error {
	content, err := se.Read()
	if err != nil {
		return err
	}
	
	backupFilename := se.filename + ".backup"
	return os.WriteFile(backupFilename, []byte(content), 0644)
}

// 演示函數
func demonstrateSimpleEditor() {
	fmt.Println("=== 簡單文本編輯器演示 ===")
	
	// 創建編輯器實例
	editor := NewSimpleEditor("demo.txt")
	
	// 1. 創建文件
	fmt.Println("\n--- 創建文件 ---")
	initialContent := "第一行內容\n第二行內容\n第三行內容"
	err := editor.Create(initialContent)
	if err != nil {
		fmt.Printf("創建文件錯誤: %v\n", err)
		return
	}
	fmt.Println("文件創建成功")
	
	// 2. 讀取文件
	fmt.Println("\n--- 讀取文件 ---")
	content, err := editor.Read()
	if err != nil {
		fmt.Printf("讀取文件錯誤: %v\n", err)
		return
	}
	fmt.Printf("文件內容:\n%s\n", content)
	
	// 3. 獲取行數
	fmt.Println("\n--- 獲取行數 ---")
	lineCount, err := editor.GetLineCount()
	if err != nil {
		fmt.Printf("獲取行數錯誤: %v\n", err)
		return
	}
	fmt.Printf("文件共有 %d 行\n", lineCount)
	
	// 4. 追加內容
	fmt.Println("\n--- 追加內容 ---")
	err = editor.Append("\n第四行內容")
	if err != nil {
		fmt.Printf("追加內容錯誤: %v\n", err)
		return
	}
	fmt.Println("內容追加成功")
	
	// 5. 在指定行插入內容
	fmt.Println("\n--- 在第2行插入內容 ---")
	err = editor.WriteAt(2, "新插入的第二行")
	if err != nil {
		fmt.Printf("插入內容錯誤: %v\n", err)
		return
	}
	fmt.Println("內容插入成功")
	
	// 6. 讀取更新後的內容
	fmt.Println("\n--- 讀取更新後的內容 ---")
	content, _ = editor.Read()
	fmt.Printf("更新後內容:\n%s\n", content)
	
	// 7. 獲取指定行
	fmt.Println("\n--- 獲取第3行內容 ---")
	line3, err := editor.GetLine(3)
	if err != nil {
		fmt.Printf("獲取行內容錯誤: %v\n", err)
	} else {
		fmt.Printf("第3行: %s\n", line3)
	}
	
	// 8. 更新指定行
	fmt.Println("\n--- 更新第1行內容 ---")
	err = editor.UpdateLine(1, "更新後的第一行")
	if err != nil {
		fmt.Printf("更新行錯誤: %v\n", err)
	} else {
		fmt.Println("第1行更新成功")
	}
	
	// 9. 搜索文本
	fmt.Println("\n--- 搜索包含'第'的行 ---")
	matches, err := editor.Search("第")
	if err != nil {
		fmt.Printf("搜索錯誤: %v\n", err)
	} else {
		fmt.Printf("找到匹配的行: %v\n", matches)
	}
	
	// 10. 替換文本
	fmt.Println("\n--- 替換文本 ---")
	err = editor.Replace("內容", "文本")
	if err != nil {
		fmt.Printf("替換文本錯誤: %v\n", err)
	} else {
		fmt.Println("文本替換成功")
	}
	
	// 11. 創建備份
	fmt.Println("\n--- 創建備份 ---")
	err = editor.Backup()
	if err != nil {
		fmt.Printf("創建備份錯誤: %v\n", err)
	} else {
		fmt.Println("備份創建成功: demo.txt.backup")
	}
	
	// 12. 顯示最終內容
	fmt.Println("\n--- 最終內容 ---")
	finalContent, _ := editor.Read()
	fmt.Printf("最終內容:\n%s\n", finalContent)
	
	// 13. 獲取最終行數
	finalLineCount, _ := editor.GetLineCount()
	fmt.Printf("最終行數: %d\n", finalLineCount)
}

// 測試錯誤處理
func testErrorHandling() {
	fmt.Println("\n=== 錯誤處理測試 ===")
	
	editor := NewSimpleEditor("demo.txt")
	
	// 測試無效行號
	fmt.Println("\n--- 測試無效行號 ---")
	err := editor.WriteAt(100, "無效行號測試")
	if err != nil {
		fmt.Printf("預期錯誤: %v\n", err)
	}
	
	err = editor.UpdateLine(100, "無效行號更新")
	if err != nil {
		fmt.Printf("預期錯誤: %v\n", err)
	}
	
	_, err = editor.GetLine(100)
	if err != nil {
		fmt.Printf("預期錯誤: %v\n", err)
	}
	
	// 測試讀取不存在的文件
	fmt.Println("\n--- 測試不存在的文件 ---")
	nonExistentEditor := NewSimpleEditor("nonexistent.txt")
	_, err = nonExistentEditor.Read()
	if err != nil {
		fmt.Printf("預期錯誤: %v\n", err)
	}
}

// 清理測試文件
func cleanup() {
	files := []string{"demo.txt", "demo.txt.backup"}
	for _, file := range files {
		os.Remove(file)
	}
	fmt.Println("清理完成")
}

func main() {
	demonstrateSimpleEditor()
	testErrorHandling()
	
	fmt.Println("\n按 Enter 鍵清理測試文件...")
	fmt.Scanln()
	
	cleanup()
	
	fmt.Println("===== 演示完成 =====")
}