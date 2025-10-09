# 第十六章：文件操作

## 🎯 學習目標

- 掌握 Go 語言的文件讀寫操作
- 理解不同的文件操作方式和使用場景
- 學會處理文件權限和錯誤
- 掌握目錄操作和文件系統遍歷
- 了解文件監控和高級文件操作
- 學會處理大文件和流式操作
- 掌握文件壓縮和解壓縮

## 📁 文件操作概述

Go 語言提供了豐富的文件操作功能，主要通過以下包實現：

### 核心包

```
文件操作相關包：
┌─────────────────────────────────────┐
│ os 包                                │
├─────────────────────────────────────┤
│ • 基礎文件操作                       │
│ • 文件創建、打開、刪除               │
│ • 文件權限和屬性                     │
│ • 目錄操作                          │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ io 包                                │
├─────────────────────────────────────┤
│ • 通用 I/O 原語                      │
│ • Reader/Writer 接口                 │
│ • 複製和管道操作                     │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ bufio 包                             │
├─────────────────────────────────────┤
│ • 緩衝 I/O 操作                      │
│ • 按行讀取                          │
│ • 高效的讀寫                        │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ filepath 包                          │
├─────────────────────────────────────┤
│ • 文件路徑操作                       │
│ • 路徑拼接和清理                     │
│ • 模式匹配                          │
└─────────────────────────────────────┘
```

## 📖 基礎文件操作

### 1. 文件創建和打開

```go
import (
    "os"
    "fmt"
)

// 創建文件
func createFile() {
    // 創建新文件（如果存在會截斷）
    file, err := os.Create("example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    // 寫入內容
    file.WriteString("Hello, World!")
}

// 打開文件
func openFile() {
    // 只讀模式打開
    file, err := os.Open("example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    // 讀取內容
    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("讀取了 %d 字節: %s\n", count, data[:count])
}

// 以特定模式打開文件
func openFileWithFlag() {
    // 讀寫模式，如果不存在則創建
    file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()
}
```

### 2. 文件權限和標誌

```go
// 文件打開標誌
const (
    O_RDONLY int = syscall.O_RDONLY // 只讀
    O_WRONLY int = syscall.O_WRONLY // 只寫
    O_RDWR   int = syscall.O_RDWR   // 讀寫
    O_APPEND int = syscall.O_APPEND // 追加
    O_CREATE int = syscall.O_CREAT  // 創建
    O_EXCL   int = syscall.O_EXCL   // 與 O_CREATE 一起使用，文件必須不存在
    O_SYNC   int = syscall.O_SYNC   // 同步 I/O
    O_TRUNC  int = syscall.O_TRUNC  // 截斷
)

// 權限模式
const (
    // 文件權限（Unix 風格）
    ModeDir        = fs.ModeDir        // 目錄
    ModeAppend     = fs.ModeAppend     // 只能追加
    ModeExclusive  = fs.ModeExclusive  // 獨占訪問
    ModeTemporary  = fs.ModeTemporary  // 臨時文件
    ModeSymlink    = fs.ModeSymlink    // 符號鏈接
    ModeDevice     = fs.ModeDevice     // 設備文件
    ModeNamedPipe  = fs.ModeNamedPipe  // 命名管道
    ModeSocket     = fs.ModeSocket     // Unix 域套接字
    ModeSetuid     = fs.ModeSetuid     // setuid
    ModeSetgid     = fs.ModeSetgid     // setgid
    ModeCharDevice = fs.ModeCharDevice // 字符設備
    ModeSticky     = fs.ModeSticky     // 粘滯位
    ModeIrregular  = fs.ModeIrregular  // 非常規文件
)

// 常用權限組合
const (
    // 所有者：讀寫執行，組：讀執行，其他：讀執行
    FileMode0755 = 0755
    // 所有者：讀寫，組：讀，其他：讀
    FileMode0644 = 0644
    // 所有者：讀寫執行，組和其他：無權限
    FileMode0700 = 0700
)
```

## 📝 文件讀取方法

### 1. 一次性讀取整個文件

```go
import (
    "io"
    "os"
)

// 方法 1: 使用 os.ReadFile（推薦）
func readEntireFile1(filename string) ([]byte, error) {
    return os.ReadFile(filename)
}

// 方法 2: 使用 io.ReadAll
func readEntireFile2(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    return io.ReadAll(file)
}

// 使用示例
func demonstrateFileReading() {
    // 讀取整個文件
    content, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("讀取文件錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("文件內容: %s\n", string(content))
}
```

### 2. 按塊讀取文件

```go
import "bufio"

// 按固定大小讀取
func readFileInChunks(filename string, chunkSize int) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    buffer := make([]byte, chunkSize)
    for {
        bytesRead, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break // 文件結束
            }
            return err
        }
        
        // 處理讀取的數據
        fmt.Printf("讀取 %d 字節: %s\n", bytesRead, string(buffer[:bytesRead]))
    }
    
    return nil
}

// 按行讀取
func readFileByLines(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNumber := 1
    
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("行 %d: %s\n", lineNumber, line)
        lineNumber++
    }
    
    return scanner.Err()
}
```

## ✏️ 文件寫入方法

### 1. 一次性寫入

```go
// 方法 1: 使用 os.WriteFile（推薦）
func writeEntireFile1(filename string, content []byte) error {
    return os.WriteFile(filename, content, 0644)
}

// 方法 2: 使用 file.Write
func writeEntireFile2(filename string, content string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(content)
    return err
}
```

### 2. 追加寫入

```go
// 追加內容到文件
func appendToFile(filename string, content string) error {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(content)
    return err
}
```

### 3. 緩衝寫入

```go
// 使用緩衝寫入提高性能
func bufferedWrite(filename string, lines []string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    defer writer.Flush() // 確保緩衝區內容被寫入
    
    for _, line := range lines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

## 📂 目錄操作

### 1. 目錄創建和刪除

```go
import "path/filepath"

// 創建目錄
func createDirectory() {
    // 創建單個目錄
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Printf("創建目錄錯誤: %v\n", err)
    }
    
    // 創建多層目錄
    err = os.MkdirAll("path/to/nested/dir", 0755)
    if err != nil {
        fmt.Printf("創建嵌套目錄錯誤: %v\n", err)
    }
}

// 刪除目錄
func removeDirectory() {
    // 刪除空目錄
    err := os.Remove("testdir")
    if err != nil {
        fmt.Printf("刪除目錄錯誤: %v\n", err)
    }
    
    // 遞歸刪除目錄及其內容
    err = os.RemoveAll("path/to/nested")
    if err != nil {
        fmt.Printf("遞歸刪除錯誤: %v\n", err)
    }
}
```

### 2. 目錄遍歷

```go
// 讀取目錄內容
func listDirectory(dirPath string) error {
    entries, err := os.ReadDir(dirPath)
    if err != nil {
        return err
    }
    
    for _, entry := range entries {
        info, err := entry.Info()
        if err != nil {
            continue
        }
        
        fmt.Printf("名稱: %s, 大小: %d, 是否目錄: %t\n", 
            entry.Name(), info.Size(), entry.IsDir())
    }
    
    return nil
}

// 遞歸遍歷目錄
func walkDirectory(root string) error {
    return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        fmt.Printf("路徑: %s, 大小: %d\n", path, info.Size())
        return nil
    })
}
```

## 🔍 文件信息和屬性

### 獲取文件信息

```go
func getFileInfo(filename string) error {
    info, err := os.Stat(filename)
    if err != nil {
        return err
    }
    
    fmt.Printf("文件名: %s\n", info.Name())
    fmt.Printf("大小: %d 字節\n", info.Size())
    fmt.Printf("權限: %s\n", info.Mode())
    fmt.Printf("修改時間: %s\n", info.ModTime())
    fmt.Printf("是否目錄: %t\n", info.IsDir())
    
    return nil
}

// 檢查文件是否存在
func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

// 檢查是否為目錄
func isDirectory(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}
```

## 🚀 高級文件操作

### 1. 文件複製

```go
// 複製文件
func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()
    
    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()
    
    _, err = io.Copy(destFile, sourceFile)
    return err
}

// 帶進度的文件複製
func copyFileWithProgress(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()
    
    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()
    
    // 獲取文件大小
    fileInfo, _ := sourceFile.Stat()
    fileSize := fileInfo.Size()
    
    // 創建進度讀取器
    progressReader := &ProgressReader{
        reader:   sourceFile,
        total:    fileSize,
        progress: 0,
    }
    
    _, err = io.Copy(destFile, progressReader)
    return err
}

// 進度讀取器
type ProgressReader struct {
    reader   io.Reader
    total    int64
    progress int64
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
    n, err := pr.reader.Read(p)
    pr.progress += int64(n)
    
    // 顯示進度
    if pr.total > 0 {
        percentage := float64(pr.progress) / float64(pr.total) * 100
        fmt.Printf("\r複製進度: %.2f%%", percentage)
    }
    
    return n, err
}
```

### 2. 文件監控

```go
import (
    "time"
    "path/filepath"
)

// 簡單的文件監控
func watchFile(filename string, interval time.Duration) {
    var lastModTime time.Time
    
    // 獲取初始修改時間
    if info, err := os.Stat(filename); err == nil {
        lastModTime = info.ModTime()
    }
    
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    
    for range ticker.C {
        info, err := os.Stat(filename)
        if err != nil {
            continue
        }
        
        if info.ModTime().After(lastModTime) {
            fmt.Printf("文件 %s 已被修改\n", filename)
            lastModTime = info.ModTime()
        }
    }
}
```

## 🛠️ 實用工具函數

### 文件工具集

```go
package fileutils

import (
    "crypto/md5"
    "encoding/hex"
    "io"
    "os"
    "path/filepath"
    "strings"
)

// FileUtils 文件工具集
type FileUtils struct{}

// NewFileUtils 創建文件工具實例
func NewFileUtils() *FileUtils {
    return &FileUtils{}
}

// GetFileExtension 獲取文件擴展名
func (fu *FileUtils) GetFileExtension(filename string) string {
    return strings.ToLower(filepath.Ext(filename))
}

// GetFileSize 獲取文件大小
func (fu *FileUtils) GetFileSize(filename string) (int64, error) {
    info, err := os.Stat(filename)
    if err != nil {
        return 0, err
    }
    return info.Size(), nil
}

// CalculateMD5 計算文件 MD5 值
func (fu *FileUtils) CalculateMD5(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()
    
    hash := md5.New()
    _, err = io.Copy(hash, file)
    if err != nil {
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
        
        if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
            matches = append(matches, path)
        }
        
        return nil
    })
    
    return matches, err
}
```

## 🎯 最佳實踐

### 1. 錯誤處理

```go
// 安全的文件操作
func safeFileOperation(filename string) error {
    // 檢查文件是否存在
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return fmt.Errorf("文件不存在: %s", filename)
    }
    
    // 檢查權限
    file, err := os.OpenFile(filename, os.O_RDWR, 0)
    if err != nil {
        return fmt.Errorf("無法打開文件: %w", err)
    }
    defer file.Close()
    
    // 進行文件操作...
    return nil
}
```

### 2. 資源管理

```go
// 正確的資源管理
func properResourceManagement() {
    file, err := os.Open("example.txt")
    if err != nil {
        // 處理錯誤
        return
    }
    defer file.Close() // 確保文件被關閉
    
    // 使用文件...
}
```

### 3. 並發安全

```go
import "sync"

// 線程安全的文件操作
type SafeFileWriter struct {
    filename string
    mutex    sync.Mutex
}

func (sfw *SafeFileWriter) WriteString(content string) error {
    sfw.mutex.Lock()
    defer sfw.mutex.Unlock()
    
    return os.WriteFile(sfw.filename, []byte(content), 0644)
}
```

## 📋 常見使用場景

### 1. 配置文件讀取
### 2. 日誌文件寫入
### 3. 數據備份和恢復
### 4. 文件批量處理
### 5. 臨時文件管理

## 🚨 注意事項

1. **記住關閉文件**：使用 `defer file.Close()`
2. **處理錯誤**：文件操作容易出錯，要妥善處理
3. **檢查權限**：確保有足夠的文件權限
4. **大文件處理**：對於大文件使用流式處理
5. **跨平台兼容**：注意不同操作系統的路徑差異

---

**下一章：[JSON 處理](../17-json-handling/)**