# 第十七章：JSON 處理

## 🎯 學習目標

- 掌握 Go 語言的 JSON 編碼和解碼
- 理解 JSON 標籤的使用和自定義
- 學會處理複雜的 JSON 結構
- 掌握 JSON 流處理和大數據處理
- 了解 JSON 驗證和轉換技巧
- 學會處理 JSON 錯誤和異常情況
- 掌握 JSON 性能優化技巧

## 📄 JSON 概述

JSON (JavaScript Object Notation) 是一種輕量級的數據交換格式，在 Go 語言中通過 `encoding/json` 包來處理。

### JSON 數據類型映射

```
Go Type          JSON Type
────────────────────────────
bool             boolean
string           string
int, float64     number
[]T              array
map[string]T     object
interface{}      any
nil              null
time.Time        string (RFC3339)
[]byte           string (base64)
```

## 🔄 基本 JSON 操作

### 1. JSON 編碼 (Marshal)

```go
import (
    "encoding/json"
    "fmt"
)

// 基本數據類型編碼
func basicMarshal() {
    // 字符串
    str := "Hello, JSON!"
    jsonStr, _ := json.Marshal(str)
    fmt.Printf("String: %s\n", jsonStr) // "Hello, JSON!"
    
    // 數字
    num := 42
    jsonNum, _ := json.Marshal(num)
    fmt.Printf("Number: %s\n", jsonNum) // 42
    
    // 布爾值
    flag := true
    jsonFlag, _ := json.Marshal(flag)
    fmt.Printf("Boolean: %s\n", jsonFlag) // true
    
    // 數組
    arr := []int{1, 2, 3, 4, 5}
    jsonArr, _ := json.Marshal(arr)
    fmt.Printf("Array: %s\n", jsonArr) // [1,2,3,4,5]
    
    // 映射
    m := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }
    jsonMap, _ := json.Marshal(m)
    fmt.Printf("Map: %s\n", jsonMap) // {"apple":5,"banana":3,"orange":8}
}
```

### 2. JSON 解碼 (Unmarshal)

```go
// 基本數據類型解碼
func basicUnmarshal() {
    // 解碼到基本類型
    var str string
    json.Unmarshal([]byte(`"Hello, JSON!"`), &str)
    fmt.Printf("Decoded string: %s\n", str)
    
    var num int
    json.Unmarshal([]byte(`42`), &num)
    fmt.Printf("Decoded number: %d\n", num)
    
    var arr []int
    json.Unmarshal([]byte(`[1,2,3,4,5]`), &arr)
    fmt.Printf("Decoded array: %v\n", arr)
    
    var m map[string]int
    json.Unmarshal([]byte(`{"apple":5,"banana":3}`), &m)
    fmt.Printf("Decoded map: %v\n", m)
}
```

## 🏗️ 結構體與 JSON

### 1. 基本結構體編碼解碼

```go
// 用戶結構體
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
    IsActive bool   `json:"is_active"`
}

func structMarshalExample() {
    user := User{
        ID:       1,
        Name:     "Alice",
        Email:    "alice@example.com",
        Age:      25,
        IsActive: true,
    }
    
    // 編碼為 JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Printf("編碼錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("JSON: %s\n", jsonData)
    
    // 解碼 JSON
    var decodedUser User
    err = json.Unmarshal(jsonData, &decodedUser)
    if err != nil {
        fmt.Printf("解碼錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("解碼用戶: %+v\n", decodedUser)
}
```

### 2. JSON 標籤詳解

```go
type Product struct {
    ID          int     `json:"id"`                          // 基本映射
    Name        string  `json:"name"`                        // 字段重命名
    Price       float64 `json:"price"`                       // 數值類型
    Description string  `json:"description,omitempty"`       // 空值時省略
    InStock     bool    `json:"in_stock"`                    // 布爾值
    Tags        []string `json:"tags,omitempty"`             // 數組，空時省略
    Metadata    map[string]interface{} `json:"metadata,omitempty"` // 映射
    CreatedAt   time.Time `json:"created_at"`                // 時間
    UpdatedAt   *time.Time `json:"updated_at,omitempty"`     // 指針時間
    Internal    string  `json:"-"`                           // 忽略字段
    Legacy      string  `json:"legacy_field,omitempty"`      // 遺留字段
}

// JSON 標籤選項說明
/*
- omitempty: 空值時省略該字段
- -: 完全忽略該字段
- string: 將數值類型編碼為字符串
- ,: 分隔標籤選項
*/

type TagExamples struct {
    Normal      string  `json:"normal"`                  // 普通字段
    Omit        string  `json:"omit,omitempty"`         // 空值省略
    Ignore      string  `json:"-"`                      // 忽略
    String      int     `json:"string_num,string"`      // 數值轉字符串
    Pointer     *string `json:"pointer,omitempty"`      // 指針字段
}
```

## 🌟 高級 JSON 處理

### 1. 自定義 JSON 編碼解碼

```go
import "time"

// 自定義時間格式
type CustomTime struct {
    time.Time
}

const customTimeFormat = "2006-01-02 15:04:05"

// 實現 json.Marshaler 接口
func (ct CustomTime) MarshalJSON() ([]byte, error) {
    return json.Marshal(ct.Time.Format(customTimeFormat))
}

// 實現 json.Unmarshaler 接口
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    var timeStr string
    if err := json.Unmarshal(data, &timeStr); err != nil {
        return err
    }
    
    t, err := time.Parse(customTimeFormat, timeStr)
    if err != nil {
        return err
    }
    
    ct.Time = t
    return nil
}

// 使用自定義時間的結構體
type Event struct {
    ID        int        `json:"id"`
    Title     string     `json:"title"`
    StartTime CustomTime `json:"start_time"`
    EndTime   CustomTime `json:"end_time"`
}
```

### 2. 處理未知結構的 JSON

```go
// 使用 interface{} 處理動態 JSON
func handleDynamicJSON() {
    jsonStr := `{
        "name": "Alice",
        "age": 25,
        "scores": [95, 87, 92],
        "address": {
            "city": "New York",
            "zipcode": "10001"
        }
    }`
    
    var data interface{}
    err := json.Unmarshal([]byte(jsonStr), &data)
    if err != nil {
        fmt.Printf("解碼錯誤: %v\n", err)
        return
    }
    
    // 類型斷言訪問數據
    m := data.(map[string]interface{})
    
    name := m["name"].(string)
    age := m["age"].(float64) // JSON 數字默認為 float64
    scores := m["scores"].([]interface{})
    address := m["address"].(map[string]interface{})
    
    fmt.Printf("姓名: %s\n", name)
    fmt.Printf("年齡: %.0f\n", age)
    fmt.Printf("成績: %v\n", scores)
    fmt.Printf("城市: %s\n", address["city"])
}

// 使用 map[string]interface{} 處理
func handleWithMap() {
    jsonStr := `{"name": "Bob", "age": 30, "active": true}`
    
    var data map[string]interface{}
    err := json.Unmarshal([]byte(jsonStr), &data)
    if err != nil {
        fmt.Printf("解碼錯誤: %v\n", err)
        return
    }
    
    for key, value := range data {
        fmt.Printf("%s: %v (類型: %T)\n", key, value, value)
    }
}
```

### 3. JSON 流處理

```go
import (
    "encoding/json"
    "strings"
)

// 使用 Decoder 流式解碼
func streamDecoding() {
    jsonStr := `{"name": "Alice"} {"name": "Bob"} {"name": "Charlie"}`
    
    decoder := json.NewDecoder(strings.NewReader(jsonStr))
    
    for decoder.More() {
        var user map[string]interface{}
        err := decoder.Decode(&user)
        if err != nil {
            fmt.Printf("解碼錯誤: %v\n", err)
            break
        }
        
        fmt.Printf("用戶: %v\n", user)
    }
}

// 使用 Encoder 流式編碼
func streamEncoding() {
    var buf strings.Builder
    encoder := json.NewEncoder(&buf)
    
    users := []map[string]interface{}{
        {"name": "Alice", "age": 25},
        {"name": "Bob", "age": 30},
        {"name": "Charlie", "age": 35},
    }
    
    for _, user := range users {
        err := encoder.Encode(user)
        if err != nil {
            fmt.Printf("編碼錯誤: %v\n", err)
            continue
        }
    }
    
    fmt.Printf("流式編碼結果:\n%s", buf.String())
}
```

## 🔧 複雜 JSON 結構處理

### 1. 嵌套結構體

```go
type Address struct {
    Street   string `json:"street"`
    City     string `json:"city"`
    State    string `json:"state"`
    ZipCode  string `json:"zip_code"`
    Country  string `json:"country"`
}

type Company struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
}

type Person struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Email    string  `json:"email"`
    Address  Address `json:"address"`
    Company  *Company `json:"company,omitempty"` // 可選字段
    Hobbies  []string `json:"hobbies,omitempty"`
}

func nestedStructExample() {
    person := Person{
        ID:    1,
        Name:  "Alice Johnson",
        Email: "alice@example.com",
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            State:   "NY",
            ZipCode: "10001",
            Country: "USA",
        },
        Company: &Company{
            Name: "Tech Corp",
            Address: Address{
                Street:  "456 Business Ave",
                City:    "New York",
                State:   "NY",
                ZipCode: "10002",
                Country: "USA",
            },
        },
        Hobbies: []string{"reading", "coding", "travel"},
    }
    
    // 編碼
    jsonData, _ := json.MarshalIndent(person, "", "  ")
    fmt.Printf("嵌套結構體 JSON:\n%s\n", jsonData)
    
    // 解碼
    var decodedPerson Person
    json.Unmarshal(jsonData, &decodedPerson)
    fmt.Printf("解碼結果: %+v\n", decodedPerson)
}
```

### 2. 處理數組和切片

```go
type OrderItem struct {
    ProductID int     `json:"product_id"`
    Name      string  `json:"name"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type Order struct {
    ID         int         `json:"id"`
    CustomerID int         `json:"customer_id"`
    Items      []OrderItem `json:"items"`
    Total      float64     `json:"total"`
    Status     string      `json:"status"`
    CreatedAt  time.Time   `json:"created_at"`
}

func arraySliceExample() {
    order := Order{
        ID:         1001,
        CustomerID: 123,
        Items: []OrderItem{
            {ProductID: 1, Name: "Laptop", Quantity: 1, Price: 999.99},
            {ProductID: 2, Name: "Mouse", Quantity: 2, Price: 29.99},
            {ProductID: 3, Name: "Keyboard", Quantity: 1, Price: 79.99},
        },
        Total:     1139.96,
        Status:    "pending",
        CreatedAt: time.Now(),
    }
    
    // 漂亮的 JSON 格式化
    jsonData, _ := json.MarshalIndent(order, "", "  ")
    fmt.Printf("訂單 JSON:\n%s\n", jsonData)
}
```

## ⚡ JSON 性能優化

### 1. 預分配和重用

```go
import "sync"

// JSON 編碼器池
var encoderPool = sync.Pool{
    New: func() interface{} {
        return json.NewEncoder(nil)
    },
}

// JSON 解碼器池
var decoderPool = sync.Pool{
    New: func() interface{} {
        return json.NewDecoder(nil)
    },
}

// 高效的 JSON 處理
func efficientJSONProcessing(data interface{}) ([]byte, error) {
    var buf bytes.Buffer
    
    encoder := encoderPool.Get().(*json.Encoder)
    defer encoderPool.Put(encoder)
    
    encoder.Reset(&buf)
    err := encoder.Encode(data)
    
    return buf.Bytes(), err
}
```

### 2. 使用 json.RawMessage

```go
// 延遲解碼
type Response struct {
    Status string          `json:"status"`
    Data   json.RawMessage `json:"data"` // 原始 JSON 數據
}

func rawMessageExample() {
    jsonStr := `{
        "status": "success",
        "data": {"name": "Alice", "age": 25}
    }`
    
    var resp Response
    err := json.Unmarshal([]byte(jsonStr), &resp)
    if err != nil {
        fmt.Printf("解碼錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("狀態: %s\n", resp.Status)
    fmt.Printf("原始數據: %s\n", resp.Data)
    
    // 根據需要解碼 Data 字段
    var user map[string]interface{}
    err = json.Unmarshal(resp.Data, &user)
    if err != nil {
        fmt.Printf("解碼數據錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("用戶數據: %v\n", user)
}
```

## 🛡️ JSON 驗證和錯誤處理

### 1. JSON 格式驗證

```go
// 驗證 JSON 格式
func validateJSON(jsonStr string) error {
    var js interface{}
    return json.Unmarshal([]byte(jsonStr), &js)
}

// 詳細的錯誤處理
func detailedErrorHandling() {
    invalidJSON := `{"name": "Alice", "age": 25,}` // 多餘的逗號
    
    var user User
    err := json.Unmarshal([]byte(invalidJSON), &user)
    if err != nil {
        if syntaxErr, ok := err.(*json.SyntaxError); ok {
            fmt.Printf("JSON 語法錯誤在位置 %d: %v\n", syntaxErr.Offset, syntaxErr)
        } else if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
            fmt.Printf("類型錯誤: 無法將 %s 轉換為 %s，字段 %s\n", 
                typeErr.Value, typeErr.Type, typeErr.Field)
        } else {
            fmt.Printf("其他錯誤: %v\n", err)
        }
    }
}
```

### 2. 自定義驗證

```go
type Email string

func (e Email) MarshalJSON() ([]byte, error) {
    return json.Marshal(string(e))
}

func (e *Email) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    
    // 簡單的郵箱驗證
    if !strings.Contains(s, "@") {
        return fmt.Errorf("無效的郵箱格式: %s", s)
    }
    
    *e = Email(s)
    return nil
}

type ValidatedUser struct {
    ID    int   `json:"id"`
    Name  string `json:"name"`
    Email Email  `json:"email"`
}
```

## 🔄 JSON 轉換工具

### 1. 結構體轉換

```go
// 源結構體
type SourceStruct struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 目標結構體
type TargetStruct struct {
    ID       int    `json:"id"`
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
    Status   string `json:"status"`
}

// 結構體轉換
func convertStruct(src SourceStruct) (TargetStruct, error) {
    // 通過 JSON 進行轉換
    jsonData, err := json.Marshal(src)
    if err != nil {
        return TargetStruct{}, err
    }
    
    var target TargetStruct
    err = json.Unmarshal(jsonData, &target)
    if err != nil {
        return TargetStruct{}, err
    }
    
    // 手動處理不匹配的字段
    target.FullName = src.Name
    target.Status = "active"
    
    return target, nil
}
```

## 🎯 實際應用場景

### 1. API 響應處理
### 2. 配置文件讀取
### 3. 數據序列化存儲
### 4. 日誌結構化輸出
### 5. 消息隊列數據格式

## ⚠️ 注意事項

1. **字段可見性**：只有公開字段（首字母大寫）才能被 JSON 編碼
2. **數字類型**：JSON 解碼時數字默認為 float64
3. **空值處理**：使用 omitempty 標籤處理空值
4. **循環引用**：避免結構體循環引用導致的無限遞歸
5. **性能考慮**：大量數據時考慮使用流式處理

---

**下一章：[HTTP 客戶端](../18-http-client/)**