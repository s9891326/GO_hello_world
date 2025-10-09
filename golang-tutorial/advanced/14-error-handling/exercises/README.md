# 第十四章練習：錯誤處理

## 練習 1：基本錯誤處理

### 題目
實現一個簡單的計算器，包含以下功能：
- 加法、減法、乘法、除法
- 對於除法，需要檢查除數是否為零
- 對於所有操作，需要檢查輸入是否為有效數字

### 要求
```go
type Calculator struct{}

func (c *Calculator) Add(a, b string) (float64, error) {
    // 實現加法，返回結果和可能的錯誤
}

func (c *Calculator) Subtract(a, b string) (float64, error) {
    // 實現減法
}

func (c *Calculator) Multiply(a, b string) (float64, error) {
    // 實現乘法
}

func (c *Calculator) Divide(a, b string) (float64, error) {
    // 實現除法，注意除零錯誤
}
```

### 測試用例
```go
calc := &Calculator{}

// 正常情況
result, err := calc.Add("10", "5")
// result: 15, err: nil

// 錯誤情況
result, err := calc.Add("abc", "5")
// result: 0, err: "無效的數字: abc"

result, err := calc.Divide("10", "0")
// result: 0, err: "除數不能為零"
```

---

## 練習 2：自定義錯誤類型

### 題目
創建一個用戶管理系統，實現自定義錯誤類型來處理不同的錯誤情況。

### 要求
```go
// 定義自定義錯誤類型
type UserError struct {
    Code    int
    Field   string
    Message string
}

type User struct {
    ID       int
    Username string
    Email    string
    Age      int
}

type UserManager struct {
    users map[int]User
}

func (um *UserManager) CreateUser(user User) error {
    // 驗證用戶數據
    // 檢查用戶名是否已存在
    // 返回適當的自定義錯誤
}

func (um *UserManager) GetUser(id int) (User, error) {
    // 獲取用戶，如果不存在返回自定義錯誤
}

func (um *UserManager) UpdateUser(id int, user User) error {
    // 更新用戶信息
}
```

### 錯誤代碼定義
- `1001`: 用戶名不能為空
- `1002`: 郵箱格式無效
- `1003`: 年齡必須在 18-100 之間
- `1004`: 用戶名已存在
- `1005`: 用戶不存在

---

## 練習 3：錯誤包裝和鏈式錯誤

### 題目
實現一個文件處理服務，包含多層錯誤包裝。

### 要求
```go
type FileProcessor struct{}

func (fp *FileProcessor) ReadFile(filename string) ([]byte, error) {
    // 讀取文件，包裝可能的錯誤
}

func (fp *FileProcessor) ParseJSON(data []byte) (map[string]interface{}, error) {
    // 解析 JSON，包裝可能的錯誤
}

func (fp *FileProcessor) ValidateData(data map[string]interface{}) error {
    // 驗證數據，包裝可能的錯誤
}

func (fp *FileProcessor) ProcessFile(filename string) error {
    // 組合上述三個方法，創建錯誤鏈
    // 使用 fmt.Errorf 和 %w 動詞進行錯誤包裝
}
```

### 錯誤處理要求
- 每一層都要添加有用的上下文信息
- 保留原始錯誤信息
- 最終錯誤應該能夠追溯到根本原因

---

## 練習 4：錯誤恢復和重試機制

### 題目
實現一個網路請求客戶端，包含重試機制和錯誤恢復。

### 要求
```go
type HTTPClient struct {
    MaxRetries int
    RetryDelay time.Duration
}

type RequestError struct {
    StatusCode int
    Message    string
    Retryable  bool
}

func (c *HTTPClient) MakeRequest(url string) ([]byte, error) {
    // 模擬網路請求
    // 對於可重試的錯誤（5xx），進行重試
    // 對於不可重試的錯誤（4xx），直接返回
}

func (c *HTTPClient) shouldRetry(err error) bool {
    // 判斷錯誤是否可重試
}
```

### 重試邏輯
- 5xx 錯誤：可重試
- 4xx 錯誤：不可重試
- 網路超時：可重試
- 連接錯誤：可重試

---

## 練習 5：綜合練習 - 銀行轉帳系統

### 題目
實現一個銀行轉帳系統，包含完整的錯誤處理機制。

### 要求
```go
type Account struct {
    ID      string
    Balance decimal.Decimal
    Status  string // active, frozen, closed
}

type TransferError struct {
    Code    string
    Message string
    Account string
}

type Bank struct {
    accounts map[string]*Account
}

func (b *Bank) Transfer(fromID, toID string, amount decimal.Decimal) error {
    // 實現轉帳邏輯
    // 包含所有必要的驗證和錯誤處理
}

func (b *Bank) GetAccount(id string) (*Account, error) {
    // 獲取帳戶信息
}

func (b *Bank) ValidateTransfer(from, to *Account, amount decimal.Decimal) error {
    // 驗證轉帳條件
}
```

### 錯誤情況
- 帳戶不存在
- 帳戶已凍結或關閉
- 餘額不足
- 轉帳金額無效（小於等於 0）
- 不能向自己轉帳

### 錯誤處理要求
- 使用自定義錯誤類型
- 提供詳細的錯誤信息
- 包含錯誤代碼以便客戶端處理
- 支援錯誤分類和不同的處理策略

---

## 提交要求

1. 每個練習創建獨立的 `.go` 文件
2. 包含完整的測試用例
3. 添加適當的註釋說明
4. 確保錯誤信息清晰有用
5. 遵循 Go 語言錯誤處理最佳實踐

## 評分標準

- **正確性** (40%)：程式功能正確實現
- **錯誤處理** (30%)：適當的錯誤檢查和處理
- **代碼品質** (20%)：代碼結構清晰，註釋充分
- **最佳實踐** (10%)：遵循 Go 語言慣例