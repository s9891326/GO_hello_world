# 第十六章練習：文件操作

## 練習 1：基本文件讀寫

### 題目
實現一個簡單的文本編輯器功能，包含文件的創建、讀取、寫入和追加操作。

### 要求
```go
type SimpleEditor struct {
    filename string
}

func NewSimpleEditor(filename string) *SimpleEditor
func (se *SimpleEditor) Create(content string) error
func (se *SimpleEditor) Read() (string, error)
func (se *SimpleEditor) Append(content string) error
func (se *SimpleEditor) WriteAt(line int, content string) error
func (se *SimpleEditor) GetLineCount() (int, error)
```

### 功能要求
1. 創建新文件並寫入內容
2. 讀取整個文件內容
3. 追加內容到文件末尾
4. 在指定行插入內容
5. 獲取文件總行數

### 測試用例
```go
editor := NewSimpleEditor("test.txt")
editor.Create("第一行\n第二行\n第三行")
content, _ := editor.Read()
editor.Append("\n第四行")
editor.WriteAt(2, "插入的新行")
lineCount, _ := editor.GetLineCount()
```

---

## 練習 2：文件系統瀏覽器

### 題目
創建一個命令行文件系統瀏覽器，可以列出目錄內容、顯示文件信息、搜索文件等。

### 要求
```go
type FileExplorer struct{}

func NewFileExplorer() *FileExplorer
func (fe *FileExplorer) ListDirectory(path string) ([]FileItem, error)
func (fe *FileExplorer) GetFileInfo(path string) (*DetailedFileInfo, error)
func (fe *FileExplorer) SearchFiles(root, pattern string) ([]string, error)
func (fe *FileExplorer) GetDirectoryTree(root string, maxDepth int) (*DirectoryNode, error)
```

### 數據結構
```go
type FileItem struct {
    Name        string
    Size        int64
    ModTime     time.Time
    IsDirectory bool
    Permissions string
}

type DetailedFileInfo struct {
    FileItem
    Extension   string
    MD5Hash     string
    Owner       string
    Group       string
}

type DirectoryNode struct {
    Name     string
    Path     string
    Children []*DirectoryNode
    Files    []FileItem
}
```

### 功能要求
1. 列出指定目錄的所有文件和子目錄
2. 顯示文件的詳細信息（大小、權限、修改時間等）
3. 支援通配符搜索文件
4. 生成目錄樹結構
5. 支援按不同條件排序（名稱、大小、時間）

---

## 練習 3：日誌文件管理器

### 題目
實現一個日誌文件管理器，支援日誌輪轉、壓縮、清理等功能。

### 要求
```go
type LogManager struct {
    logDir      string
    maxFileSize int64
    maxFiles    int
    compress    bool
}

func NewLogManager(logDir string, maxFileSize int64, maxFiles int) *LogManager
func (lm *LogManager) WriteLog(message string) error
func (lm *LogManager) RotateLog() error
func (lm *LogManager) CompressOldLogs() error
func (lm *LogManager) CleanupOldLogs() error
func (lm *LogManager) GetLogFiles() ([]string, error)
```

### 功能要求
1. 寫入日誌到當前日誌文件
2. 當文件大小超過限制時自動輪轉
3. 壓縮舊的日誌文件
4. 清理超過保留數量的舊日誌
5. 列出所有日誌文件

### 日誌格式
```
2023-12-07 15:30:45 [INFO] 這是一條日誌訊息
2023-12-07 15:30:46 [ERROR] 這是一條錯誤訊息
```

---

## 練習 4：文件同步工具

### 題目
創建一個簡單的文件同步工具，可以比較兩個目錄的差異並同步文件。

### 要求
```go
type FileSyncer struct{}

func NewFileSyncer() *FileSyncer
func (fs *FileSyncer) Compare(sourceDir, targetDir string) (*SyncReport, error)
func (fs *FileSyncer) Sync(sourceDir, targetDir string, options SyncOptions) error
func (fs *FileSyncer) Preview(sourceDir, targetDir string) (*SyncPlan, error)
```

### 數據結構
```go
type SyncReport struct {
    NewFiles     []string
    UpdatedFiles []string
    DeletedFiles []string
    Conflicts    []string
}

type SyncOptions struct {
    DeleteExtra bool
    OverwriteNewer bool
    BackupOriginal bool
}

type SyncPlan struct {
    Actions []SyncAction
}

type SyncAction struct {
    Type   string // "create", "update", "delete"
    Source string
    Target string
    Reason string
}
```

### 功能要求
1. 比較兩個目錄的文件差異
2. 根據選項執行同步操作
3. 顯示同步計劃預覽
4. 支援備份被覆蓋的文件
5. 處理衝突情況

---

## 練習 5：文件加密工具

### 題目
實現一個文件加密/解密工具，支援多種加密算法。

### 要求
```go
type FileEncryption struct {
    algorithm string
}

func NewFileEncryption(algorithm string) *FileEncryption
func (fe *FileEncryption) EncryptFile(inputFile, outputFile, password string) error
func (fe *FileEncryption) DecryptFile(inputFile, outputFile, password string) error
func (fe *FileEncryption) EncryptDirectory(inputDir, outputDir, password string) error
func (fe *FileEncryption) DecryptDirectory(inputDir, outputDir, password string) error
func (fe *FileEncryption) ValidatePassword(encryptedFile, password string) bool
```

### 功能要求
1. 支援 AES 加密算法
2. 加密/解密單個文件
3. 批量加密/解密目錄中的文件
4. 密碼驗證
5. 進度顯示（對於大文件）

### 加密文件格式
```
[4字節] 版本號
[32字節] 鹽值
[16字節] IV向量
[變長] 加密數據
[32字節] HMAC校驗
```

---

## 練習 6：文件監控系統

### 題目
創建一個文件系統監控工具，監控指定目錄的文件變化。

### 要求
```go
type FileWatcher struct {
    watchDir string
    events   chan FileEvent
    stop     chan bool
}

func NewFileWatcher(watchDir string) *FileWatcher
func (fw *FileWatcher) Start() error
func (fw *FileWatcher) Stop()
func (fw *FileWatcher) AddFilter(filter FileFilter)
func (fw *FileWatcher) GetEvents() <-chan FileEvent
```

### 數據結構
```go
type FileEvent struct {
    Type     EventType
    Path     string
    Time     time.Time
    FileInfo os.FileInfo
}

type EventType int
const (
    FileCreated EventType = iota
    FileModified
    FileDeleted
    FileRenamed
)

type FileFilter func(event FileEvent) bool
```

### 功能要求
1. 監控目錄中文件的創建、修改、刪除事件
2. 支援事件過濾器
3. 實時事件通知
4. 支援遞迴監控子目錄
5. 事件日誌記錄

---

## 提交要求

1. 每個練習創建獨立的 `.go` 文件
2. 包含完整的測試用例
3. 提供使用示例和說明文檔
4. 妥善處理錯誤情況
5. 遵循 Go 語言最佳實踐

## 評分標準

- **功能完整性** (40%)：實現所有要求的功能
- **錯誤處理** (25%)：妥善處理各種錯誤情況
- **代碼品質** (20%)：代碼結構清晰，註釋充分
- **性能優化** (10%)：高效的文件操作實現
- **用戶體驗** (5%)：友好的用戶界面和反饋

## 額外挑戰

### 挑戰 1：大文件處理
處理 GB 級別的大文件，要求：
- 記憶體使用量穩定
- 進度顯示
- 可中斷和恢復

### 挑戰 2：並發文件操作
實現並發文件處理，要求：
- 線程安全
- 限制並發數量
- 錯誤隔離

### 挑戰 3：跨平台兼容
確保代碼在 Windows、Linux、macOS 上都能正常運行，處理：
- 路徑分隔符差異
- 權限系統差異
- 文件鎖定機制