# 第十四章：錯誤處理

## 🎯 學習目標

- 理解 Go 語言的錯誤處理哲學
- 掌握 error 接口的使用和實現
- 學會自定義錯誤類型
- 了解錯誤包裝和鏈式錯誤
- 掌握錯誤處理的最佳實踐
- 學會錯誤監控和日誌記錄

## 🚨 錯誤處理哲學

Go 語言的錯誤處理採用顯式錯誤返回的方式，而不是異常機制。這種設計哲學強調：

### 核心原則

```
Go 錯誤處理原則：
┌─────────────────────────────────────┐
│ • 錯誤是值，可以被程序處理            │
│ • 顯式處理，不隱藏錯誤               │
│ • 錯誤應該被檢查，不應該被忽略        │
│ • 錯誤信息應該有用且有上下文          │
│ • 在適當的層級處理錯誤               │
│ • 失敗時快速返回                    │
└─────────────────────────────────────┘
```

### error 接口

```go
// Go 內建的 error 接口
type error interface {
    Error() string
}
```

## 🔍 基本錯誤處理

### 錯誤返回和檢查

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

func demonstrateBasicErrorHandling() {
    fmt.Println("--- 基本錯誤處理演示 ---")
    
    // 示例1：字符串轉換
    str := "123"
    if num, err := strconv.Atoi(str); err != nil {
        fmt.Printf("轉換失敗: %v\n", err)
    } else {
        fmt.Printf("轉換成功: %d\n", num)
    }
    
    // 示例2：無效轉換
    invalidStr := "abc"
    if num, err := strconv.Atoi(invalidStr); err != nil {
        fmt.Printf("轉換失敗: %v\n", err)
    } else {
        fmt.Printf("轉換成功: %d\n", num)
    }
    
    // 示例3：自定義函數錯誤處理
    result, err := divide(10, 2)
    if err != nil {
        fmt.Printf("除法錯誤: %v\n", err)
    } else {
        fmt.Printf("除法結果: %.2f\n", result)
    }
    
    result, err = divide(10, 0)
    if err != nil {
        fmt.Printf("除法錯誤: %v\n", err)
    } else {
        fmt.Printf("除法結果: %.2f\n", result)
    }
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除數不能為零")
    }
    return a / b, nil
}
```

### 錯誤創建方式

```go
import (
    "errors"
    "fmt"
)

func demonstrateErrorCreation() {
    fmt.Println("\n--- 錯誤創建方式 ---")
    
    // 方式1：使用 errors.New
    err1 := errors.New("這是一個簡單錯誤")
    fmt.Printf("errors.New: %v\n", err1)
    
    // 方式2：使用 fmt.Errorf
    name := "用戶"
    err2 := fmt.Errorf("找不到 %s: ID=%d", name, 123)
    fmt.Printf("fmt.Errorf: %v\n", err2)
    
    // 方式3：預定義錯誤
    var ErrNotFound = errors.New("記錄未找到")
    err3 := ErrNotFound
    fmt.Printf("預定義錯誤: %v\n", err3)
    
    // 方式4：錯誤比較
    if err3 == ErrNotFound {
        fmt.Println("這是 NotFound 錯誤")
    }
}
```

## 🎭 自定義錯誤類型

### 實現 error 接口

```go
// 自定義錯誤類型
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("驗證失敗 [%s]: %s (值: %v)", e.Field, e.Message, e.Value)
}

type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("數據庫操作失敗 [%s.%s]: %v", e.Table, e.Operation, e.Err)
}

// 實現 Unwrap 方法以支援錯誤鏈
func (e DatabaseError) Unwrap() error {
    return e.Err
}

func demonstrateCustomErrors() {
    fmt.Println("\n--- 自定義錯誤類型 ---")
    
    // 驗證錯誤
    validateUser := func(age int) error {
        if age < 0 {
            return ValidationError{
                Field:   "age",
                Value:   age,
                Message: "年齡不能為負數",
            }
        }
        if age > 150 {
            return ValidationError{
                Field:   "age",
                Value:   age,
                Message: "年齡不能超過150",
            }
        }
        return nil
    }
    
    // 測試驗證
    if err := validateUser(-5); err != nil {
        fmt.Printf("驗證錯誤: %v\n", err)
        
        // 類型斷言獲取詳細信息
        if ve, ok := err.(ValidationError); ok {
            fmt.Printf("錯誤字段: %s, 錯誤值: %v\n", ve.Field, ve.Value)
        }
    }
    
    // 數據庫錯誤
    dbErr := DatabaseError{
        Operation: "INSERT",
        Table:     "users",
        Err:       errors.New("連接超時"),
    }
    
    fmt.Printf("數據庫錯誤: %v\n", dbErr)
    fmt.Printf("底層錯誤: %v\n", dbErr.Unwrap())
}
```

### 錯誤類型檢查

```go
func demonstrateErrorTypeChecking() {
    fmt.Println("\n--- 錯誤類型檢查 ---")
    
    errors := []error{
        ValidationError{Field: "email", Value: "invalid", Message: "格式錯誤"},
        DatabaseError{Operation: "SELECT", Table: "products", Err: errors.New("表不存在")},
        fmt.Errorf("一般錯誤: %s", "系統繁忙"),
    }
    
    for i, err := range errors {
        fmt.Printf("\n錯誤 %d: %v\n", i+1, err)
        
        // 方式1：類型斷言
        switch e := err.(type) {
        case ValidationError:
            fmt.Printf("  這是驗證錯誤，字段: %s\n", e.Field)
        case DatabaseError:
            fmt.Printf("  這是數據庫錯誤，操作: %s，表: %s\n", e.Operation, e.Table)
        default:
            fmt.Printf("  這是其他類型錯誤: %T\n", e)
        }
        
        // 方式2：使用 errors.As (Go 1.13+)
        var ve ValidationError
        if errors.As(err, &ve) {
            fmt.Printf("  通過 errors.As 檢測到驗證錯誤\n")
        }
        
        var de DatabaseError
        if errors.As(err, &de) {
            fmt.Printf("  通過 errors.As 檢測到數據庫錯誤\n")
        }
    }
}
```

## 🔗 錯誤包裝和鏈式錯誤

### 錯誤包裝 (Go 1.13+)

```go
import (
    "errors"
    "fmt"
)

func demonstrateErrorWrapping() {
    fmt.Println("\n--- 錯誤包裝演示 ---")
    
    // 原始錯誤
    originalErr := errors.New("網路連接失敗")
    
    // 包裝錯誤
    wrappedErr := fmt.Errorf("處理用戶請求失敗: %w", originalErr)
    
    // 再次包裝
    finalErr := fmt.Errorf("API 調用失敗: %w", wrappedErr)
    
    fmt.Printf("最終錯誤: %v\n", finalErr)
    
    // 檢查錯誤鏈
    fmt.Printf("是否包含原始錯誤: %t\n", errors.Is(finalErr, originalErr))
    
    // 解包錯誤
    fmt.Printf("直接底層錯誤: %v\n", errors.Unwrap(finalErr))
    fmt.Printf("最底層錯誤: %v\n", errors.Unwrap(errors.Unwrap(finalErr)))
    
    // 遍歷錯誤鏈
    fmt.Println("錯誤鏈:")
    err := finalErr
    level := 0
    for err != nil {
        fmt.Printf("  層級 %d: %v\n", level, err)
        err = errors.Unwrap(err)
        level++
    }
}

// 實際應用：服務調用鏈錯誤
func callExternalAPI() error {
    return errors.New("外部 API 返回 500 錯誤")
}

func processData() error {
    if err := callExternalAPI(); err != nil {
        return fmt.Errorf("數據處理失敗: %w", err)
    }
    return nil
}

func handleRequest() error {
    if err := processData(); err != nil {
        return fmt.Errorf("處理請求失敗: %w", err)
    }
    return nil
}

func demonstrateServiceErrorChain() {
    fmt.Println("\n--- 服務調用鏈錯誤 ---")
    
    if err := handleRequest(); err != nil {
        fmt.Printf("頂層錯誤: %v\n", err)
        
        // 檢查特定錯誤
        var apiErr error
        if errors.As(err, &apiErr) {
            fmt.Printf("包含 API 錯誤: %v\n", apiErr)
        }
        
        // 檢查錯誤消息
        if strings.Contains(err.Error(), "500 錯誤") {
            fmt.Println("檢測到 500 錯誤，執行重試邏輯")
        }
    }
}
```

## 🛡️ 錯誤處理模式

### 提前返回模式

```go
func demonstrateEarlyReturn() {
    fmt.Println("\n--- 提前返回模式 ---")
    
    processUser := func(userID int) error {
        // 驗證用戶ID
        if userID <= 0 {
            return fmt.Errorf("無效的用戶ID: %d", userID)
        }
        
        // 獲取用戶信息
        user, err := getUser(userID)
        if err != nil {
            return fmt.Errorf("獲取用戶失敗: %w", err)
        }
        
        // 驗證用戶狀態
        if err := validateUserStatus(user); err != nil {
            return fmt.Errorf("用戶狀態驗證失敗: %w", err)
        }
        
        // 更新用戶信息
        if err := updateUser(user); err != nil {
            return fmt.Errorf("更新用戶失敗: %w", err)
        }
        
        return nil
    }
    
    // 測試不同情況
    testCases := []int{-1, 0, 999, 1}
    
    for _, userID := range testCases {
        if err := processUser(userID); err != nil {
            fmt.Printf("處理用戶 %d 失敗: %v\n", userID, err)
        } else {
            fmt.Printf("處理用戶 %d 成功\n", userID)
        }
    }
}

// 模擬函數
func getUser(id int) (User, error) {
    if id == 999 {
        return User{}, errors.New("用戶不存在")
    }
    return User{ID: id, Name: "測試用戶", Active: id != 2}, nil
}

func validateUserStatus(user User) error {
    if !user.Active {
        return errors.New("用戶已被停用")
    }
    return nil
}

func updateUser(user User) error {
    if user.ID == 3 {
        return errors.New("數據庫更新失敗")
    }
    return nil
}

type User struct {
    ID     int
    Name   string
    Active bool
}
```

### 錯誤聚合模式

```go
type MultiError struct {
    Errors []error
}

func (me MultiError) Error() string {
    if len(me.Errors) == 0 {
        return "無錯誤"
    }
    if len(me.Errors) == 1 {
        return me.Errors[0].Error()
    }
    
    var result strings.Builder
    result.WriteString(fmt.Sprintf("發生 %d 個錯誤:", len(me.Errors)))
    for i, err := range me.Errors {
        result.WriteString(fmt.Sprintf("\n  %d. %v", i+1, err))
    }
    return result.String()
}

func (me *MultiError) Add(err error) {
    if err != nil {
        me.Errors = append(me.Errors, err)
    }
}

func (me MultiError) HasErrors() bool {
    return len(me.Errors) > 0
}

func demonstrateMultiError() {
    fmt.Println("\n--- 錯誤聚合模式 ---")
    
    validateForm := func(name, email string, age int) error {
        var multiErr MultiError
        
        // 驗證姓名
        if name == "" {
            multiErr.Add(ValidationError{
                Field:   "name",
                Value:   name,
                Message: "姓名不能為空",
            })
        }
        
        // 驗證郵箱
        if !strings.Contains(email, "@") {
            multiErr.Add(ValidationError{
                Field:   "email",
                Value:   email,
                Message: "郵箱格式無效",
            })
        }
        
        // 驗證年齡
        if age < 0 || age > 150 {
            multiErr.Add(ValidationError{
                Field:   "age",
                Value:   age,
                Message: "年齡必須在 0-150 之間",
            })
        }
        
        if multiErr.HasErrors() {
            return multiErr
        }
        return nil
    }
    
    // 測試表單驗證
    testCases := []struct {
        name, email string
        age         int
    }{
        {"張三", "zhang@example.com", 25},     // 有效
        {"", "invalid-email", -5},            // 全部無效
        {"李四", "li@example.com", 200},       // 年齡無效
    }
    
    for i, tc := range testCases {
        fmt.Printf("\n測試案例 %d:\n", i+1)
        if err := validateForm(tc.name, tc.email, tc.age); err != nil {
            fmt.Printf("驗證失敗:\n%v\n", err)
        } else {
            fmt.Println("驗證成功")
        }
    }
}
```

## 🔄 錯誤重試機制

### 指數退避重試

```go
import (
    "context"
    "math"
    "math/rand"
    "time"
)

type RetryConfig struct {
    MaxAttempts int
    BaseDelay   time.Duration
    MaxDelay    time.Duration
    Multiplier  float64
}

func WithRetry(ctx context.Context, config RetryConfig, operation func() error) error {
    var lastErr error
    
    for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
        if err := operation(); err == nil {
            return nil // 成功
        } else {
            lastErr = err
        }
        
        if attempt == config.MaxAttempts {
            break // 最後一次嘗試
        }
        
        // 計算延遲時間
        delay := time.Duration(float64(config.BaseDelay) * math.Pow(config.Multiplier, float64(attempt-1)))
        if delay > config.MaxDelay {
            delay = config.MaxDelay
        }
        
        // 添加抖動
        jitter := time.Duration(rand.Int63n(int64(delay / 4)))
        delay += jitter
        
        fmt.Printf("嘗試 %d 失敗: %v，%v 後重試\n", attempt, lastErr, delay)
        
        // 等待重試
        select {
        case <-ctx.Done():
            return fmt.Errorf("重試被取消: %w", ctx.Err())
        case <-time.After(delay):
            // 繼續重試
        }
    }
    
    return fmt.Errorf("重試 %d 次後仍然失敗: %w", config.MaxAttempts, lastErr)
}

func demonstrateRetryMechanism() {
    fmt.Println("\n--- 錯誤重試機制 ---")
    
    // 模擬不穩定的服務
    attemptCount := 0
    unstableService := func() error {
        attemptCount++
        if attemptCount < 3 {
            return fmt.Errorf("服務暫時不可用 (嘗試 %d)", attemptCount)
        }
        return nil // 第3次成功
    }
    
    config := RetryConfig{
        MaxAttempts: 5,
        BaseDelay:   100 * time.Millisecond,
        MaxDelay:    2 * time.Second,
        Multiplier:  2.0,
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    if err := WithRetry(ctx, config, unstableService); err != nil {
        fmt.Printf("最終失敗: %v\n", err)
    } else {
        fmt.Println("重試成功！")
    }
}
```

## 📊 錯誤監控和指標

### 錯誤統計

```go
import (
    "sync"
    "time"
)

type ErrorStats struct {
    mutex      sync.RWMutex
    counters   map[string]int
    lastErrors map[string]time.Time
}

func NewErrorStats() *ErrorStats {
    return &ErrorStats{
        counters:   make(map[string]int),
        lastErrors: make(map[string]time.Time),
    }
}

func (es *ErrorStats) Record(errorType string) {
    es.mutex.Lock()
    defer es.mutex.Unlock()
    
    es.counters[errorType]++
    es.lastErrors[errorType] = time.Now()
}

func (es *ErrorStats) GetStats() map[string]ErrorStat {
    es.mutex.RLock()
    defer es.mutex.RUnlock()
    
    stats := make(map[string]ErrorStat)
    for errorType, count := range es.counters {
        stats[errorType] = ErrorStat{
            Count:     count,
            LastSeen:  es.lastErrors[errorType],
            ErrorType: errorType,
        }
    }
    return stats
}

type ErrorStat struct {
    ErrorType string
    Count     int
    LastSeen  time.Time
}

func demonstrateErrorMonitoring() {
    fmt.Println("\n--- 錯誤監控演示 ---")
    
    stats := NewErrorStats()
    
    // 模擬各種錯誤
    errors := []string{
        "DatabaseConnection",
        "ValidationError",
        "APITimeout",
        "DatabaseConnection",
        "ValidationError",
        "NetworkError",
        "DatabaseConnection",
    }
    
    for _, errType := range errors {
        stats.Record(errType)
        time.Sleep(10 * time.Millisecond)
    }
    
    // 顯示統計信息
    fmt.Println("錯誤統計:")
    for errorType, stat := range stats.GetStats() {
        fmt.Printf("  %s: %d 次 (最後發生: %s)\n", 
            stat.ErrorType, stat.Count, stat.LastSeen.Format("15:04:05"))
    }
}
```

## 💡 錯誤處理最佳實踐

### 1. 錯誤消息設計

```go
func demonstrateErrorMessageBestPractices() {
    fmt.Println("\n--- 錯誤消息最佳實踐 ---")
    
    // ❌ 不好的錯誤消息
    badExamples := []error{
        errors.New("錯誤"),                    // 太簡略
        errors.New("something went wrong"),    // 不明確
        errors.New("failed"),                 // 沒有上下文
    }
    
    // ✅ 好的錯誤消息
    goodExamples := []error{
        fmt.Errorf("無法連接到數據庫: %s", "connection timeout"),
        fmt.Errorf("用戶驗證失敗: 用戶 %s 不存在", "john_doe"),
        fmt.Errorf("文件操作失敗: 無法讀取 %s (權限不足)", "/etc/passwd"),
    }
    
    fmt.Println("❌ 不好的錯誤消息:")
    for _, err := range badExamples {
        fmt.Printf("  %v\n", err)
    }
    
    fmt.Println("\n✅ 好的錯誤消息:")
    for _, err := range goodExamples {
        fmt.Printf("  %v\n", err)
    }
}
```

### 2. 錯誤處理策略

```go
func demonstrateErrorHandlingStrategies() {
    fmt.Println("\n--- 錯誤處理策略 ---")
    
    // 策略1: 記錄並返回
    logAndReturn := func() error {
        err := errors.New("數據庫連接失敗")
        fmt.Printf("LOG: %v\n", err)
        return fmt.Errorf("服務暫時不可用: %w", err)
    }
    
    // 策略2: 重試
    retryOperation := func() error {
        for i := 0; i < 3; i++ {
            if err := riskyOperation(); err == nil {
                return nil
            }
            fmt.Printf("重試 %d/3\n", i+1)
            time.Sleep(100 * time.Millisecond)
        }
        return errors.New("操作失敗：已重試3次")
    }
    
    // 策略3: 降級
    fallbackOperation := func() (string, error) {
        if err := primaryService(); err != nil {
            fmt.Println("主服務失敗，使用備用服務")
            return "備用結果", nil
        }
        return "主服務結果", nil
    }
    
    // 策略4: 斷路器
    circuitBreaker := NewCircuitBreaker(3, time.Minute)
    protectedOperation := func() error {
        return circuitBreaker.Execute(func() error {
            return unreliableService()
        })
    }
    
    fmt.Println("測試不同策略:")
    
    // 測試記錄並返回
    if err := logAndReturn(); err != nil {
        fmt.Printf("記錄並返回: %v\n", err)
    }
    
    // 測試重試
    if err := retryOperation(); err != nil {
        fmt.Printf("重試策略: %v\n", err)
    }
    
    // 測試降級
    if result, err := fallbackOperation(); err != nil {
        fmt.Printf("降級策略失敗: %v\n", err)
    } else {
        fmt.Printf("降級策略成功: %s\n", result)
    }
    
    // 測試斷路器
    for i := 0; i < 5; i++ {
        if err := protectedOperation(); err != nil {
            fmt.Printf("斷路器保護 %d: %v\n", i+1, err)
        }
    }
}

// 模擬函數
func riskyOperation() error {
    if rand.Float32() < 0.7 {
        return errors.New("操作失敗")
    }
    return nil
}

func primaryService() error {
    return errors.New("主服務不可用")
}

func unreliableService() error {
    if rand.Float32() < 0.8 {
        return errors.New("服務失敗")
    }
    return nil
}

// 簡單的斷路器實現
type CircuitBreaker struct {
    failures    int
    maxFailures int
    resetTime   time.Time
    state       string
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        maxFailures: maxFailures,
        state:       "closed",
    }
}

func (cb *CircuitBreaker) Execute(operation func() error) error {
    if cb.state == "open" && time.Now().Before(cb.resetTime) {
        return errors.New("斷路器開啟：服務不可用")
    }
    
    if cb.state == "open" {
        cb.state = "half-open"
    }
    
    err := operation()
    if err != nil {
        cb.failures++
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
            cb.resetTime = time.Now().Add(time.Minute)
        }
        return err
    }
    
    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

## 🎯 本章練習

1. 實現完整的錯誤處理系統
2. 創建自定義錯誤類型庫
3. 實現錯誤重試機制
4. 創建錯誤監控系統

---

**下一章：[包管理](../15-packages/)**