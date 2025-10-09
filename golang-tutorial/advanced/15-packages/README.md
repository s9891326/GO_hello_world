# 第十五章：包管理

## 🎯 學習目標

- 理解 Go 包（package）的概念和作用
- 掌握包的創建、導入和使用
- 學會使用 Go Modules 進行依賴管理
- 了解包的可見性規則
- 掌握包的初始化機制
- 學會創建和發布自己的包
- 了解包的版本管理和語義化版本

## 📦 Go 包系統概述

Go 的包系統是代碼組織和重用的基礎，它提供了：

### 核心概念

```
Go 包系統架構：
┌─────────────────────────────────────┐
│ Go Modules (模組)                    │
├─────────────────────────────────────┤
│ • go.mod - 模組定義文件              │
│ • go.sum - 依賴校驗文件              │
│ • 語義化版本控制                     │
│ • 依賴管理                          │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ Packages (包)                       │
├─────────────────────────────────────┤
│ • 代碼組織單元                       │
│ • 命名空間                          │
│ • 可見性控制                        │
│ • 初始化機制                        │
└─────────────────────────────────────┘
```

## 🏗️ 包的基本概念

### 包的定義

```go
// 每個 Go 文件都必須屬於一個包
package main    // 可執行程序的入口包
package utils   // 庫包
package myapp   // 應用程序包
```

### 包的導入

```go
package main

import (
    "fmt"                    // 標準庫包
    "net/http"              // 標準庫子包
    "github.com/gin-gonic/gin" // 第三方包
    "./utils"               // 相對路徑（不推薦）
    "myapp/internal/config" // 模組內部包
)
```

## 📁 包的組織結構

### 標準項目結構

```
myproject/
├── go.mod                 # 模組定義
├── go.sum                 # 依賴鎖定
├── main.go                # 主程序入口
├── README.md              # 項目說明
├── cmd/                   # 可執行程序
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── pkg/                   # 公共庫代碼
│   ├── utils/
│   │   ├── string.go
│   │   └── math.go
│   └── models/
│       └── user.go
├── internal/              # 私有代碼
│   ├── handlers/
│   │   └── user.go
│   └── database/
│       └── connection.go
├── api/                   # API 定義
│   └── openapi.yaml
├── web/                   # Web 資源
│   ├── static/
│   └── templates/
├── scripts/               # 構建和部署腳本
├── docs/                  # 文檔
└── tests/                 # 測試文件
```

## 🔒 可見性規則

### 導出和未導出

```go
package utils

import "fmt"

// 導出的（公共的）- 首字母大寫
type User struct {
    Name  string // 導出字段
    Email string // 導出字段
    age   int    // 未導出字段（私有）
}

// 導出的函數
func NewUser(name, email string, age int) *User {
    return &User{
        Name:  name,
        Email: email,
        age:   age,
    }
}

// 導出的方法
func (u *User) GetAge() int {
    return u.age
}

// 未導出的函數（私有）
func validateEmail(email string) bool {
    // 內部實現
    return len(email) > 0
}

// 導出的常數
const MaxUsers = 1000

// 未導出的常數
const defaultTimeout = 30

// 導出的變數
var DefaultConfig = Config{
    Host: "localhost",
    Port: 8080,
}

// 未導出的變數
var internalCounter int
```

## 🔄 包的初始化

### init 函數

```go
package config

import (
    "log"
    "os"
)

var (
    DatabaseURL string
    APIKey      string
)

// init 函數在包被導入時自動執行
func init() {
    log.Println("初始化 config 包")
    
    // 從環境變數讀取配置
    DatabaseURL = os.Getenv("DATABASE_URL")
    if DatabaseURL == "" {
        DatabaseURL = "localhost:5432"
    }
    
    APIKey = os.Getenv("API_KEY")
    if APIKey == "" {
        log.Fatal("API_KEY 環境變數未設置")
    }
}

// 可以有多個 init 函數，按順序執行
func init() {
    log.Println("第二個 init 函數")
}
```

### 初始化順序

```go
package main

import (
    "fmt"
    _ "myapp/config" // 僅執行初始化，不使用包內容
)

/*
初始化順序：
1. 計算包的依賴圖
2. 按依賴順序初始化包
3. 在每個包內：
   a. 初始化包級別變數
   b. 執行所有 init 函數（按出現順序）
4. 最後執行 main 函數
*/

func main() {
    fmt.Println("main 函數執行")
}
```

## 📋 Go Modules 詳解

### 創建模組

```bash
# 初始化新模組
go mod init github.com/username/myproject

# 這會創建 go.mod 文件
```

### go.mod 文件結構

```go
module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/golang-jwt/jwt/v5 v5.0.0
)

require (
    // 間接依賴
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
)

exclude (
    // 排除特定版本
    github.com/some/problematic v1.2.3
)

replace (
    // 替換依賴
    github.com/old/package => github.com/new/package v1.0.0
    github.com/local/package => ./local/path
)

retract (
    // 撤回已發布的版本
    v1.0.1 // 包含安全漏洞
    [v1.1.0, v1.2.0] // 範圍撤回
)
```

### 依賴管理命令

```bash
# 添加依賴
go get github.com/gin-gonic/gin

# 添加特定版本
go get github.com/gin-gonic/gin@v1.9.1

# 添加最新版本
go get github.com/gin-gonic/gin@latest

# 升級依賴
go get -u github.com/gin-gonic/gin

# 升級所有依賴
go get -u ./...

# 移除依賴
go mod tidy

# 下載依賴到本地緩存
go mod download

# 驗證依賴
go mod verify

# 查看依賴圖
go mod graph

# 解釋依賴
go mod why github.com/gin-gonic/gin
```

## 🛠️ 實用包管理技巧

### 1. 內部包

```go
// internal/ 目錄下的包只能被父目錄及其子目錄導入
myproject/
├── internal/
│   ├── auth/
│   │   └── jwt.go      // 只能被 myproject 內部使用
│   └── database/
│       └── conn.go
└── cmd/
    └── server/
        └── main.go     // 可以導入 internal/auth
```

### 2. 包別名

```go
package main

import (
    "database/sql"
    
    // 包別名
    mysql "github.com/go-sql-driver/mysql"
    postgres "github.com/lib/pq"
    
    // 點導入（不推薦）
    . "fmt"
    
    // 空白導入（僅執行 init）
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // 使用別名
    db, err := sql.Open("mysql", mysql.Config{}.FormatDSN())
    
    // 點導入後可直接使用
    Println("Hello")
}
```

### 3. 包文檔

```go
// Package utils 提供常用的工具函數和類型。
//
// 這個包包含了字符串處理、數學計算和文件操作等實用工具。
//
// 基本使用示例：
//
//	import "myapp/pkg/utils"
//
//	result := utils.StringReverse("hello")
//	fmt.Println(result) // "olleh"
//
// 更多信息請參考: https://github.com/username/myproject
package utils

import "fmt"

// StringReverse 反轉字符串。
//
// 參數 s 是要反轉的字符串。
// 返回反轉後的字符串。
//
// 示例：
//
//	reversed := StringReverse("hello")
//	fmt.Println(reversed) // 輸出: "olleh"
func StringReverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

## 🚀 創建和發布包

### 1. 包的設計原則

```go
package calculator

import "errors"

// 定義清晰的公共 API
type Calculator interface {
    Add(a, b float64) float64
    Subtract(a, b float64) float64
    Multiply(a, b float64) float64
    Divide(a, b float64) (float64, error)
}

// 實現接口
type basicCalculator struct{}

// NewCalculator 創建新的計算器實例
func NewCalculator() Calculator {
    return &basicCalculator{}
}

func (c *basicCalculator) Add(a, b float64) float64 {
    return a + b
}

func (c *basicCalculator) Subtract(a, b float64) float64 {
    return a - b
}

func (c *basicCalculator) Multiply(a, b float64) float64 {
    return a * b
}

func (c *basicCalculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除數不能為零")
    }
    return a / b, nil
}
```

### 2. 版本管理

```bash
# 創建 git 標籤進行版本發布
git tag v1.0.0
git push origin v1.0.0

# 語義化版本控制
# v主版本.次版本.修復版本
# v1.2.3
# - 主版本：不兼容的 API 變更
# - 次版本：向後兼容的功能新增
# - 修復版本：向後兼容的錯誤修復
```

### 3. 包測試

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    calc := NewCalculator()
    
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"正數相加", 2, 3, 5},
        {"負數相加", -2, -3, -5},
        {"正負數相加", 5, -3, 2},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calc.Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// 基準測試
func BenchmarkAdd(b *testing.B) {
    calc := NewCalculator()
    for i := 0; i < b.N; i++ {
        calc.Add(1.23, 4.56)
    }
}

// 示例測試
func ExampleCalculator_Add() {
    calc := NewCalculator()
    result := calc.Add(2, 3)
    fmt.Println(result)
    // Output: 5
}
```

## 🔧 最佳實踐

### 1. 包設計原則

- **單一職責**：每個包應該有明確的單一職責
- **最小化導出**：只導出必要的類型和函數
- **一致的命名**：使用一致的命名約定
- **清晰的文檔**：為公共 API 提供清晰的文檔

### 2. 依賴管理

- **明確版本**：使用具體版本而不是 latest
- **定期更新**：定期檢查和更新依賴
- **安全審核**：定期檢查依賴的安全漏洞
- **最小化依賴**：避免不必要的依賴

### 3. 項目結構

- **標準布局**：遵循 Go 社區的標準項目布局
- **清晰分層**：明確區分業務邏輯、數據訪問、API 層
- **內部包**：使用 internal/ 目錄保護內部實現

## 📚 常用標準庫包

```go
import (
    // 系統操作
    "os"           // 操作系統接口
    "os/exec"      // 執行外部命令
    "path/filepath" // 文件路徑操作
    
    // 網路和 HTTP
    "net"          // 網路相關
    "net/http"     // HTTP 客戶端和服務端
    "net/url"      // URL 解析
    
    // 數據處理
    "encoding/json" // JSON 編碼解碼
    "encoding/xml"  // XML 處理
    "encoding/csv"  // CSV 處理
    
    // 字符串和正則
    "strings"      // 字符串操作
    "regexp"       // 正則表達式
    "strconv"      // 字符串轉換
    
    // 時間和數學
    "time"         // 時間處理
    "math"         // 數學函數
    "math/rand"    // 隨機數
    
    // 併發
    "sync"         // 同步原語
    "context"      // 上下文
    
    // 錯誤和日誌
    "errors"       // 錯誤處理
    "log"          // 日誌記錄
    
    // 測試
    "testing"      // 測試框架
)
```

## 🎯 總結

包管理是 Go 語言的核心特性之一，它提供了：

1. **代碼組織**：通過包來組織和結構化代碼
2. **命名空間**：避免命名衝突
3. **可見性控制**：控制 API 的公開程度
4. **依賴管理**：使用 Go Modules 管理外部依賴
5. **版本控制**：支援語義化版本控制

掌握這些概念對於構建可維護、可擴展的 Go 應用程序至關重要。

---

**下一章：[文件操作](../../practical/16-file-io/)**