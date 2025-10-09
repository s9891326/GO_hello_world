# 第十五章練習：包管理

## 練習 1：創建工具包

### 題目
創建一個完整的工具包，包含常用的功能模組。

### 要求
創建以下包結構：
```
utils/
├── math/
│   └── calculator.go
├── string/
│   └── processor.go
├── time/
│   └── helper.go
└── validator/
    └── rules.go
```

### 具體實現

#### math/calculator.go
```go
package math

// Calculator 提供基本數學運算
type Calculator struct{}

func NewCalculator() *Calculator
func (c *Calculator) Add(a, b float64) float64
func (c *Calculator) Subtract(a, b float64) float64
func (c *Calculator) Multiply(a, b float64) float64
func (c *Calculator) Divide(a, b float64) (float64, error)
func (c *Calculator) Power(base, exponent float64) float64
func (c *Calculator) Sqrt(x float64) (float64, error)
```

#### string/processor.go
```go
package string

// Processor 提供字符串處理功能
type Processor struct{}

func NewProcessor() *Processor
func (p *Processor) Reverse(s string) string
func (p *Processor) IsPalindrome(s string) bool
func (p *Processor) RemoveSpaces(s string) string
func (p *Processor) CountWords(s string) int
func (p *Processor) ToSlug(s string) string // 轉換為 URL 友好格式
```

---

## 練習 2：Go Modules 實戰

### 題目
創建一個完整的 Go 模組，包含多個包和外部依賴。

### 要求
1. 初始化 Go 模組
2. 創建多個內部包
3. 添加外部依賴
4. 實現包之間的相互調用

### 目錄結構
```
myapp/
├── go.mod
├── go.sum
├── main.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── connection.go
│   └── handlers/
│       └── user.go
└── internal/
    ├── models/
    │   └── user.go
    └── services/
        └── user_service.go
```

### 依賴要求
- 使用 `github.com/gin-gonic/gin` 作為 Web 框架
- 使用 `github.com/spf13/viper` 進行配置管理
- 實現用戶管理 API

---

## 練習 3：包的可見性設計

### 題目
設計一個銀行系統包，演示 Go 語言的可見性規則。

### 要求
```go
package bank

// 公共接口
type Bank interface {
    CreateAccount(customerID string, initialBalance float64) (Account, error)
    GetAccount(accountID string) (Account, error)
    Transfer(fromID, toID string, amount float64) error
}

// 公共類型
type Account interface {
    GetID() string
    GetBalance() float64
    GetCustomerID() string
}

// 內部實現（不導出）
type bankImpl struct { /* ... */ }
type accountImpl struct { /* ... */ }

// 工廠函數
func NewBank() Bank
```

### 設計要求
- 只導出接口，不導出具體實現
- 使用工廠模式創建實例
- 內部驗證邏輯不對外暴露
- 提供清晰的錯誤處理

---

## 練習 4：包的初始化順序

### 題目
創建多個相互依賴的包，觀察和控制初始化順序。

### 要求
創建包結構：
```
initdemo/
├── main.go
├── pkg/
│   ├── a/
│   │   └── a.go
│   ├── b/
│   │   └── b.go
│   └── c/
│       └── c.go
```

### 依賴關係
- main 依賴 a, b, c
- a 依賴 b
- b 依賴 c
- 每個包都有 init 函數

### 觀察要求
- 記錄初始化順序
- 理解包級變數初始化
- 演示 init 函數的執行時機

---

## 練習 5：創建可發布的包

### 題目
創建一個可以發布到 GitHub 的完整包，包含文檔、測試和示例。

### 要求
```
github.com/yourname/mathlib/
├── README.md
├── LICENSE
├── go.mod
├── mathlib.go          # 主要實現
├── mathlib_test.go     # 測試文件
├── examples/
│   └── basic/
│       └── main.go     # 使用示例
└── docs/
    └── api.md          # API 文檔
```

### 功能要求
1. **基本數學運算**：加減乘除、冪運算
2. **統計函數**：平均值、中位數、標準差
3. **幾何計算**：圓形、矩形面積計算
4. **完整測試**：單元測試、基準測試、示例測試
5. **清晰文檔**：包文檔、函數註釋、使用示例

### 發布要求
- 使用語義化版本控制
- 創建 git 標籤
- 編寫清晰的 README
- 提供使用示例

## 提交要求

1. 每個練習創建獨立的目錄
2. 包含完整的 go.mod 文件
3. 提供測試文件
4. 編寫清晰的文檔
5. 遵循 Go 語言包設計最佳實踐

## 評分標準

- **包設計** (30%)：合理的包結構和接口設計
- **可見性控制** (25%)：正確使用導出/未導出規則
- **依賴管理** (20%)：正確使用 Go Modules
- **代碼品質** (15%)：代碼清晰，註釋充分
- **測試覆蓋** (10%)：完整的測試用例