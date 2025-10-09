package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// FileInfo 文件信息結構
type FileInfo struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	ModTime      time.Time `json:"mod_time"`
	IsDir        bool      `json:"is_dir"`
	Extension    string    `json:"extension"`
	Permissions  string    `json:"permissions"`
	MD5Hash      string    `json:"md5_hash,omitempty"`
	SHA256Hash   string    `json:"sha256_hash,omitempty"`
}

// FileUtils 文件工具集合
type FileUtils struct{}

// NewFileUtils 創建文件工具實例
func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

// GetFileInfo 獲取詳細文件信息
func (fu *FileUtils) GetFileInfo(path string) (*FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	
	fileInfo := &FileInfo{
		Name:        info.Name(),
		Size:        info.Size(),
		ModTime:     info.ModTime(),
		IsDir:       info.IsDir(),
		Extension:   strings.ToLower(filepath.Ext(path)),
		Permissions: info.Mode().String(),
	}
	
	return fileInfo, nil
}

// GetFileExtension 獲取文件擴展名
func (fu *FileUtils) GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

// GetFileNameWithoutExt 獲取不帶擴展名的文件名
func (fu *FileUtils) GetFileNameWithoutExt(filename string) string {
	name := filepath.Base(filename)
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext)
}

// IsTextFile 判斷是否為文本文件
func (fu *FileUtils) IsTextFile(filename string) bool {
	textExtensions := []string{
		".txt", ".md", ".json", ".xml", ".html", ".css", ".js",
		".go", ".py", ".java", ".c", ".cpp", ".h", ".hpp",
		".sql", ".log", ".csv", ".yaml", ".yml", ".ini",
	}
	
	ext := fu.GetFileExtension(filename)
	for _, textExt := range textExtensions {
		if ext == textExt {
			return true
		}
	}
	return false
}

// IsImageFile 判斷是否為圖片文件
func (fu *FileUtils) IsImageFile(filename string) bool {
	imageExtensions := []string{
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp", ".svg",
	}
	
	ext := fu.GetFileExtension(filename)
	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}

// CalculateSHA256 計算文件 SHA256 值
func (fu *FileUtils) CalculateSHA256(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// FindFiles 查找匹配模式的文件
func (fu *FileUtils) FindFiles(root, pattern string) ([]string, error) {
	var matches []string
	
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		matched, _ := filepath.Match(pattern, filepath.Base(path))
		if matched {
			matches = append(matches, path)
		}
		
		return nil
	})
	
	return matches, err
}

// FindFilesByRegex 使用正則表達式查找文件
func (fu *FileUtils) FindFilesByRegex(root, regexPattern string) ([]string, error) {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}
	
	var matches []string
	
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		if regex.MatchString(filepath.Base(path)) {
			matches = append(matches, path)
		}
		
		return nil
	})
	
	return matches, err
}

// GetDirectorySize 計算目錄大小
func (fu *FileUtils) GetDirectorySize(dirPath string) (int64, error) {
	var size int64
	
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() {
			size += info.Size()
		}
		
		return nil
	})
	
	return size, err
}

// GetFileCount 計算目錄中的文件數量
func (fu *FileUtils) GetFileCount(dirPath string) (int, int, error) {
	var fileCount, dirCount int
	
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
		
		return nil
	})
	
	return fileCount, dirCount, err
}

// ReadFileLines 讀取文件所有行
func (fu *FileUtils) ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
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

// WriteFileLines 寫入文件行
func (fu *FileUtils) WriteFileLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	
	return nil
}

// MergeFiles 合並多個文件
func (fu *FileUtils) MergeFiles(inputFiles []string, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()
	
	writer := bufio.NewWriter(output)
	defer writer.Flush()
	
	for _, inputFile := range inputFiles {
		err := fu.appendFileContent(writer, inputFile)
		if err != nil {
			return fmt.Errorf("合並文件 %s 失敗: %w", inputFile, err)
		}
	}
	
	return nil
}

// appendFileContent 追加文件內容到寫入器
func (fu *FileUtils) appendFileContent(writer *bufio.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	
	return scanner.Err()
}

// SplitFile 分割文件
func (fu *FileUtils) SplitFile(filename string, linesPerFile int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	fileCount := 1
	lineCount := 0
	
	baseFilename := fu.GetFileNameWithoutExt(filename)
	ext := fu.GetFileExtension(filename)
	
	var currentFile *os.File
	var writer *bufio.Writer
	
	for scanner.Scan() {
		// 如果需要創建新文件
		if lineCount%linesPerFile == 0 {
			if currentFile != nil {
				writer.Flush()
				currentFile.Close()
			}
			
			newFilename := fmt.Sprintf("%s_part%d%s", baseFilename, fileCount, ext)
			currentFile, err = os.Create(newFilename)
			if err != nil {
				return err
			}
			
			writer = bufio.NewWriter(currentFile)
			fileCount++
		}
		
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
		
		lineCount++
	}
	
	if currentFile != nil {
		writer.Flush()
		currentFile.Close()
	}
	
	return scanner.Err()
}

// CreateFileIndex 創建文件索引
func (fu *FileUtils) CreateFileIndex(rootDir, indexFile string) error {
	var files []*FileInfo
	
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		fileInfo := &FileInfo{
			Name:        info.Name(),
			Size:        info.Size(),
			ModTime:     info.ModTime(),
			IsDir:       info.IsDir(),
			Extension:   strings.ToLower(filepath.Ext(path)),
			Permissions: info.Mode().String(),
		}
		
		// 計算 SHA256（可選，對大文件可能很慢）
		if info.Size() < 10*1024*1024 { // 小於 10MB 的文件才計算 hash
			if hash, err := fu.CalculateSHA256(path); err == nil {
				fileInfo.SHA256Hash = hash
			}
		}
		
		files = append(files, fileInfo)
		return nil
	})
	
	if err != nil {
		return err
	}
	
	// 按名稱排序
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
	
	// 保存為 JSON 文件
	jsonData, err := json.MarshalIndent(files, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(indexFile, jsonData, 0644)
}

// 文件統計信息
type FileStats struct {
	TotalFiles      int               `json:"total_files"`
	TotalDirs       int               `json:"total_dirs"`
	TotalSize       int64             `json:"total_size"`
	ExtensionStats  map[string]int    `json:"extension_stats"`
	SizeDistribution map[string]int   `json:"size_distribution"`
	LargestFiles    []*FileInfo       `json:"largest_files"`
}

// GetDirectoryStats 獲取目錄統計信息
func (fu *FileUtils) GetDirectoryStats(dirPath string) (*FileStats, error) {
	stats := &FileStats{
		ExtensionStats:  make(map[string]int),
		SizeDistribution: make(map[string]int),
		LargestFiles:    make([]*FileInfo, 0),
	}
	
	var allFiles []*FileInfo
	
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			stats.TotalDirs++
			return nil
		}
		
		stats.TotalFiles++
		stats.TotalSize += info.Size()
		
		// 統計擴展名
		ext := strings.ToLower(filepath.Ext(path))
		if ext == "" {
			ext = "(無擴展名)"
		}
		stats.ExtensionStats[ext]++
		
		// 統計文件大小分布
		sizeCategory := fu.getSizeCategory(info.Size())
		stats.SizeDistribution[sizeCategory]++
		
		// 收集文件信息用於查找最大文件
		fileInfo := &FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
		}
		allFiles = append(allFiles, fileInfo)
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// 找出最大的 10 個文件
	sort.Slice(allFiles, func(i, j int) bool {
		return allFiles[i].Size > allFiles[j].Size
	})
	
	maxFiles := 10
	if len(allFiles) < maxFiles {
		maxFiles = len(allFiles)
	}
	stats.LargestFiles = allFiles[:maxFiles]
	
	return stats, nil
}

// getSizeCategory 獲取文件大小分類
func (fu *FileUtils) getSizeCategory(size int64) string {
	if size < 1024 {
		return "< 1KB"
	} else if size < 1024*1024 {
		return "1KB - 1MB"
	} else if size < 10*1024*1024 {
		return "1MB - 10MB"
	} else if size < 100*1024*1024 {
		return "10MB - 100MB"
	} else {
		return "> 100MB"
	}
}

// FormatFileSize 格式化文件大小
func (fu *FileUtils) FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	units := []string{"KB", "MB", "GB", "TB", "PB"}
	return fmt.Sprintf("%.1f %s", float64(size)/float64(div), units[exp])
}

// 演示函數
func demonstrateFileUtils() {
	fu := NewFileUtils()
	
	// 創建測試文件
	createTestFiles()
	
	fmt.Println("=== 文件工具演示 ===")
	
	// 1. 文件信息
	fmt.Println("\n--- 文件信息 ---")
	if info, err := fu.GetFileInfo("test1.txt"); err == nil {
		fmt.Printf("文件: %s\n", info.Name)
		fmt.Printf("大小: %s\n", fu.FormatFileSize(info.Size))
		fmt.Printf("修改時間: %s\n", info.ModTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("權限: %s\n", info.Permissions)
	}
	
	// 2. 查找文件
	fmt.Println("\n--- 查找文件 ---")
	if files, err := fu.FindFiles(".", "*.txt"); err == nil {
		fmt.Printf("找到 %d 個 .txt 文件:\n", len(files))
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}
	
	// 3. 目錄統計
	fmt.Println("\n--- 目錄統計 ---")
	if stats, err := fu.GetDirectoryStats("."); err == nil {
		fmt.Printf("總文件數: %d\n", stats.TotalFiles)
		fmt.Printf("總目錄數: %d\n", stats.TotalDirs)
		fmt.Printf("總大小: %s\n", fu.FormatFileSize(stats.TotalSize))
		
		fmt.Println("擴展名統計:")
		for ext, count := range stats.ExtensionStats {
			fmt.Printf("  %s: %d 個\n", ext, count)
		}
	}
	
	// 4. 文件合並
	fmt.Println("\n--- 文件合並 ---")
	files := []string{"test1.txt", "test2.txt", "test3.txt"}
	err := fu.MergeFiles(files, "merged.txt")
	if err != nil {
		fmt.Printf("合並文件錯誤: %v\n", err)
	} else {
		fmt.Println("文件合並完成: merged.txt")
	}
	
	// 5. 文件分割
	fmt.Println("\n--- 文件分割 ---")
	err = fu.SplitFile("merged.txt", 2)
	if err != nil {
		fmt.Printf("分割文件錯誤: %v\n", err)
	} else {
		fmt.Println("文件分割完成")
	}
	
	// 6. 創建文件索引
	fmt.Println("\n--- 創建文件索引 ---")
	err = fu.CreateFileIndex(".", "file_index.json")
	if err != nil {
		fmt.Printf("創建索引錯誤: %v\n", err)
	} else {
		fmt.Println("文件索引創建完成: file_index.json")
	}
}

// 創建測試文件
func createTestFiles() {
	testFiles := map[string]string{
		"test1.txt": "這是測試文件1的內容\n包含多行文本\n用於演示文件操作",
		"test2.txt": "這是測試文件2的內容\n也有多行\n用於合並操作",
		"test3.txt": "這是測試文件3的內容\n最後一個測試文件\n完成演示",
	}
	
	for filename, content := range testFiles {
		os.WriteFile(filename, []byte(content), 0644)
	}
}

// 清理測試文件
func cleanupTestFiles() {
	files := []string{
		"test1.txt", "test2.txt", "test3.txt",
		"merged.txt", "merged_part1.txt", "merged_part2.txt",
		"file_index.json",
	}
	
	for _, file := range files {
		os.Remove(file)
	}
}

func main() {
	fmt.Println("===== 文件工具演示 =====")
	
	demonstrateFileUtils()
	
	fmt.Println("\n按 Enter 鍵清理測試文件...")
	fmt.Scanln()
	
	cleanupTestFiles()
	
	fmt.Println("===== 演示完成 =====")
}