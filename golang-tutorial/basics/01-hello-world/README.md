# 第一章：Hello World 和環境設置

## 🎯 學習目標

- 了解 Go 語言的特點和優勢
- 設置 Go 開發環境
- 編寫和運行第一個 Go 程序
- 理解 Go 程序的基本結構
- 學會使用 `go` 命令工具

## 📖 Go 語言簡介

Go（也稱為 Golang）是由 Google 開發的開源程式語言，於 2009 年首次發布。

### 主要特點：
- **簡潔** - 語法簡單，易於學習
- **高效** - 編譯速度快，執行效率高
- **併發** - 內建協程（goroutine）支援
- **跨平台** - 支援多種作業系統
- **靜態類型** - 編譯時類型檢查
- **垃圾回收** - 自動內存管理

### 適用場景：
- Web 後端開發
- 微服務架構
- 容器和雲原生應用
- 命令行工具
- 系統程式設計

## 🛠️ 環境設置

### 1. 安裝 Go

#### macOS
```bash
# 使用 Homebrew
brew install go

# 或下載安裝包
# https://golang.org/dl/
```

#### Windows
```bash
# 下載 MSI 安裝包
# https://golang.org/dl/

# 或使用 Chocolatey
choco install golang
```

#### Linux (Ubuntu/Debian)
```bash
# 使用包管理器
sudo apt update
sudo apt install golang-go

# 或下載二進制包
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. 驗證安裝
```bash
go version
# 輸出：go version go1.21.0 darwin/amd64
```

### 3. 設置環境變數

在 `~/.bashrc` 或 `~/.zshrc` 中添加：
```bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

### 4. 驗證環境
```bash
go env GOPATH
go env GOROOT
```

## 💻 第一個 Go 程序

### 基本程序結構

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 程序結構說明：

1. **package main** - 定義包名，`main` 是可執行程序的入口包
2. **import "fmt"** - 導入格式化輸出包
3. **func main()** - 主函數，程序執行的入口點
4. **fmt.Println()** - 輸出文字到控制台

## 🚀 運行程序

### 方法一：直接運行
```bash
go run main.go
```

### 方法二：編譯後運行
```bash
go build main.go
./main  # Linux/macOS
main.exe  # Windows
```

### 方法三：安裝到 GOBIN
```bash
go install main.go
```

## 📁 Go 專案結構

### 基本專案結構：
```
my-project/
├── go.mod          # 模組定義文件
├── go.sum          # 依賴校驗文件
├── main.go         # 主程序文件
├── pkg/            # 可重用的庫代碼
├── cmd/            # 應用程序入口
├── internal/       # 私有應用程序代碼
├── api/            # API 定義文件
├── web/            # Web 應用程序資源
├── configs/        # 配置文件
├── scripts/        # 腳本文件
├── test/           # 測試相關文件
├── docs/           # 文檔
└── README.md       # 專案說明
```

## 🔧 Go 工具鏈

### 常用 go 命令：

```bash
# 運行程序
go run main.go

# 編譯程序
go build

# 安裝程序
go install

# 下載依賴
go mod download

# 整理依賴
go mod tidy

# 格式化代碼
go fmt

# 檢查代碼
go vet

# 運行測試
go test

# 查看文檔
go doc fmt.Println
```

## 💡 實用技巧

### 1. 代碼格式化
Go 有內建的代碼格式化工具：
```bash
go fmt ./...
```

### 2. 代碼檢查
```bash
go vet ./...
```

### 3. 依賴管理
```bash
# 初始化模組
go mod init myproject

# 添加依賴
go get github.com/gin-gonic/gin

# 移除未使用的依賴
go mod tidy
```

## 🎯 本章練習

完成以下練習來鞏固學習：

1. 創建一個輸出您名字的程序
2. 創建一個顯示當前時間的程序
3. 嘗試使用不同的 `fmt` 函數（`Print`, `Printf`, `Println`）

## 🔗 相關資源

- [Go 官方安裝指南](https://golang.org/doc/install)
- [Go 工具鏈文檔](https://golang.org/cmd/go/)
- [有效的 Go 程式設計](https://golang.org/doc/effective_go.html)

---

**下一章：[變數和常數](../02-variables/)**