package main

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// FileManager 高級文件管理器
type FileManager struct {
	mutex sync.RWMutex
}

// NewFileManager 創建文件管理器
func NewFileManager() *FileManager {
	return &FileManager{}
}

// 大文件處理示例
func (fm *FileManager) ProcessLargeFile(filename string, processor func(string) string) error {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// 打開源文件
	inputFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("無法打開文件: %w", err)
	}
	defer inputFile.Close()
	
	// 創建輸出文件
	outputFile, err := os.Create(filename + ".processed")
	if err != nil {
		return fmt.Errorf("無法創建輸出文件: %w", err)
	}
	defer outputFile.Close()
	
	// 使用緩衝讀取器和寫入器
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processor(line)
		
		_, err := writer.WriteString(processedLine + "\n")
		if err != nil {
			return fmt.Errorf("寫入錯誤: %w", err)
		}
		
		lineCount++
		if lineCount%1000 == 0 {
			fmt.Printf("已處理 %d 行\n", lineCount)
		}
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("讀取錯誤: %w", err)
	}
	
	fmt.Printf("完成處理，總共 %d 行\n", lineCount)
	return nil
}

// 帶進度顯示的文件複製
type ProgressReader struct {
	reader   io.Reader
	total    int64
	progress int64
	lastShow time.Time
}

func NewProgressReader(reader io.Reader, total int64) *ProgressReader {
	return &ProgressReader{
		reader:   reader,
		total:    total,
		progress: 0,
		lastShow: time.Now(),
	}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.progress += int64(n)
	
	// 每100ms更新一次進度
	if time.Since(pr.lastShow) > 100*time.Millisecond {
		pr.showProgress()
		pr.lastShow = time.Now()
	}
	
	return n, err
}

func (pr *ProgressReader) showProgress() {
	if pr.total > 0 {
		percentage := float64(pr.progress) / float64(pr.total) * 100
		fmt.Printf("\r複製進度: %.2f%% (%d/%d 字節)", percentage, pr.progress, pr.total)
	}
}

// 帶進度的文件複製
func (fm *FileManager) CopyFileWithProgress(src, dst string) error {
	// 獲取源文件信息
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("無法獲取源文件信息: %w", err)
	}
	
	// 打開源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("無法打開源文件: %w", err)
	}
	defer srcFile.Close()
	
	// 創建目標文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("無法創建目標文件: %w", err)
	}
	defer dstFile.Close()
	
	// 創建進度讀取器
	progressReader := NewProgressReader(srcFile, srcInfo.Size())
	
	fmt.Printf("開始複製文件 %s -> %s\n", src, dst)
	startTime := time.Now()
	
	// 執行複製
	_, err = io.Copy(dstFile, progressReader)
	if err != nil {
		return fmt.Errorf("複製文件錯誤: %w", err)
	}
	
	duration := time.Since(startTime)
	fmt.Printf("\n複製完成，耗時: %v\n", duration)
	
	return nil
}

// 計算文件 MD5 校驗值
func (fm *FileManager) CalculateFileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("無法打開文件: %w", err)
	}
	defer file.Close()
	
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("計算 MD5 錯誤: %w", err)
	}
	
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 文件壓縮
func (fm *FileManager) CompressFiles(files []string, zipFile string) error {
	// 創建 zip 文件
	zipFile_, err := os.Create(zipFile)
	if err != nil {
		return fmt.Errorf("無法創建 zip 文件: %w", err)
	}
	defer zipFile_.Close()
	
	// 創建 zip 寫入器
	zipWriter := zip.NewWriter(zipFile_)
	defer zipWriter.Close()
	
	fmt.Printf("開始壓縮 %d 個文件到 %s\n", len(files), zipFile)
	
	for i, file := range files {
		err := fm.addFileToZip(zipWriter, file)
		if err != nil {
			return fmt.Errorf("添加文件 %s 到壓縮包失敗: %w", file, err)
		}
		
		fmt.Printf("已添加 %d/%d: %s\n", i+1, len(files), file)
	}
	
	fmt.Println("壓縮完成")
	return nil
}

// 添加文件到 zip
func (fm *FileManager) addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// 獲取文件信息
	info, err := file.Stat()
	if err != nil {
		return err
	}
	
	// 創建文件頭
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	
	// 設置壓縮方法
	header.Method = zip.Deflate
	
	// 創建寫入器
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	
	// 複製文件內容
	_, err = io.Copy(writer, file)
	return err
}

// 解壓縮文件
func (fm *FileManager) ExtractZip(zipFile, destDir string) error {
	// 打開 zip 文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("無法打開 zip 文件: %w", err)
	}
	defer reader.Close()
	
	// 創建目標目錄
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return fmt.Errorf("無法創建目標目錄: %w", err)
	}
	
	fmt.Printf("開始解壓縮 %s 到 %s\n", zipFile, destDir)
	
	// 解壓縮每個文件
	for i, file := range reader.File {
		err := fm.extractFile(file, destDir)
		if err != nil {
			return fmt.Errorf("解壓縮文件 %s 失敗: %w", file.Name, err)
		}
		
		fmt.Printf("已解壓縮 %d/%d: %s\n", i+1, len(reader.File), file.Name)
	}
	
	fmt.Println("解壓縮完成")
	return nil
}

// 解壓縮單個文件
func (fm *FileManager) extractFile(file *zip.File, destDir string) error {
	// 構建完整路徑
	path := filepath.Join(destDir, file.Name)
	
	// 確保路徑安全（防止 Zip Slip 攻擊）
	if !strings.HasPrefix(path, filepath.Clean(destDir)+string(os.PathSeparator)) {
		return fmt.Errorf("無效的文件路徑: %s", file.Name)
	}
	
	// 如果是目錄，創建目錄
	if file.FileInfo().IsDir() {
		return os.MkdirAll(path, file.FileInfo().Mode())
	}
	
	// 創建文件所在目錄
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	
	// 打開 zip 中的文件
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()
	
	// 創建目標文件
	targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.FileInfo().Mode())
	if err != nil {
		return err
	}
	defer targetFile.Close()
	
	// 複製內容
	_, err = io.Copy(targetFile, fileReader)
	return err
}

// 文件監控器
type FileWatcher struct {
	files    map[string]time.Time
	interval time.Duration
	mutex    sync.RWMutex
	stopCh   chan bool
}

// NewFileWatcher 創建文件監控器
func NewFileWatcher(interval time.Duration) *FileWatcher {
	return &FileWatcher{
		files:    make(map[string]time.Time),
		interval: interval,
		stopCh:   make(chan bool),
	}
}

// AddFile 添加監控文件
func (fw *FileWatcher) AddFile(filename string) error {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()
	
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}
	
	fw.files[filename] = info.ModTime()
	return nil
}

// Start 開始監控
func (fw *FileWatcher) Start(callback func(string)) {
	ticker := time.NewTicker(fw.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			fw.checkFiles(callback)
		case <-fw.stopCh:
			return
		}
	}
}

// Stop 停止監控
func (fw *FileWatcher) Stop() {
	fw.stopCh <- true
}

// 檢查文件變化
func (fw *FileWatcher) checkFiles(callback func(string)) {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()
	
	for filename, lastMod := range fw.files {
		info, err := os.Stat(filename)
		if err != nil {
			continue
		}
		
		if info.ModTime().After(lastMod) {
			fw.files[filename] = info.ModTime()
			callback(filename)
		}
	}
}

// 演示函數
func demonstrateAdvancedOperations() {
	fm := NewFileManager()
	
	// 創建測試文件
	createTestFiles()
	
	// 1. 處理大文件
	fmt.Println("=== 大文件處理 ===")
	processor := func(line string) string {
		return strings.ToUpper(line)
	}
	
	err := fm.ProcessLargeFile("large_test.txt", processor)
	if err != nil {
		fmt.Printf("處理大文件錯誤: %v\n", err)
	}
	
	// 2. 帶進度的文件複製
	fmt.Println("\n=== 帶進度的文件複製 ===")
	err = fm.CopyFileWithProgress("large_test.txt", "copied_large_test.txt")
	if err != nil {
		fmt.Printf("複製文件錯誤: %v\n", err)
	}
	
	// 3. 計算 MD5
	fmt.Println("\n=== 計算文件 MD5 ===")
	md5Hash, err := fm.CalculateFileMD5("large_test.txt")
	if err != nil {
		fmt.Printf("計算 MD5 錯誤: %v\n", err)
	} else {
		fmt.Printf("文件 MD5: %s\n", md5Hash)
	}
	
	// 4. 文件壓縮
	fmt.Println("\n=== 文件壓縮 ===")
	files := []string{"large_test.txt", "copied_large_test.txt"}
	err = fm.CompressFiles(files, "test_archive.zip")
	if err != nil {
		fmt.Printf("壓縮文件錯誤: %v\n", err)
	}
	
	// 5. 文件解壓縮
	fmt.Println("\n=== 文件解壓縮 ===")
	err = fm.ExtractZip("test_archive.zip", "extracted")
	if err != nil {
		fmt.Printf("解壓縮文件錯誤: %v\n", err)
	}
	
	// 6. 文件監控演示
	fmt.Println("\n=== 文件監控 ===")
	watcher := NewFileWatcher(1 * time.Second)
	watcher.AddFile("large_test.txt")
	
	// 在 goroutine 中啟動監控
	go watcher.Start(func(filename string) {
		fmt.Printf("檢測到文件變化: %s\n", filename)
	})
	
	// 模擬文件修改
	time.Sleep(2 * time.Second)
	os.WriteFile("large_test.txt", []byte("修改後的內容"), 0644)
	time.Sleep(3 * time.Second)
	
	watcher.Stop()
}

// 創建測試文件
func createTestFiles() {
	// 創建一個較大的測試文件
	file, err := os.Create("large_test.txt")
	if err != nil {
		fmt.Printf("創建測試文件錯誤: %v\n", err)
		return
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	
	// 寫入 1000 行數據
	for i := 1; i <= 1000; i++ {
		line := fmt.Sprintf("這是第 %d 行測試數據，包含一些中文字符和數字 %d\n", i, i*i)
		writer.WriteString(line)
	}
	
	fmt.Println("創建測試文件完成")
}

// 清理測試文件
func cleanupTestFiles() {
	files := []string{
		"large_test.txt",
		"large_test.txt.processed",
		"copied_large_test.txt",
		"test_archive.zip",
	}
	
	for _, file := range files {
		os.Remove(file)
	}
	
	os.RemoveAll("extracted")
	fmt.Println("清理完成")
}

func main() {
	fmt.Println("===== 高級文件操作示例 =====")
	
	demonstrateAdvancedOperations()
	
	fmt.Println("\n按 Enter 鍵清理測試文件...")
	fmt.Scanln()
	
	cleanupTestFiles()
	
	fmt.Println("===== 示例完成 =====")
}