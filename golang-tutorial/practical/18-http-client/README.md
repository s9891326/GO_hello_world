# 第十八章：HTTP 客戶端

## 🎯 學習目標

- 掌握 Go 語言的 HTTP 客戶端操作
- 理解 HTTP 請求方法和響應處理
- 學會處理請求頭、參數和身份驗證
- 掌握文件上傳和下載
- 了解超時、重試和錯誤處理
- 學會使用 HTTP 中間件和攔截器
- 掌握並發 HTTP 請求處理

## 🌐 HTTP 客戶端概述

Go 語言的 `net/http` 包提供了強大的 HTTP 客戶端功能，支持各種 HTTP 操作。

### 核心組件

```
HTTP 客戶端架構：
┌─────────────────────────────────────┐
│ http.Client                          │
├─────────────────────────────────────┤
│ • Transport (傳輸層)                 │
│ • Timeout (超時設置)                 │
│ • CheckRedirect (重定向處理)         │
│ • Jar (Cookie 管理)                  │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ http.Request                         │
├─────────────────────────────────────┤
│ • Method (請求方法)                  │
│ • URL (請求地址)                     │
│ • Header (請求頭)                    │
│ • Body (請求體)                      │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ http.Response                        │
├─────────────────────────────────────┤
│ • StatusCode (狀態碼)                │
│ • Header (響應頭)                    │
│ • Body (響應體)                      │
│ • Cookies (Cookie)                   │
└─────────────────────────────────────┘
```

## 🚀 基本 HTTP 請求

### 1. GET 請求

```go
import (
    "fmt"
    "io"
    "net/http"
)

// 簡單 GET 請求
func simpleGet() {
    resp, err := http.Get("https://api.github.com/users/octocat")
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("讀取響應錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("狀態碼: %d\n", resp.StatusCode)
    fmt.Printf("響應體: %s\n", body)
}

// 帶參數的 GET 請求
func getWithParams() {
    baseURL := "https://api.github.com/search/repositories"
    
    // 創建請求
    req, err := http.NewRequest("GET", baseURL, nil)
    if err != nil {
        fmt.Printf("創建請求錯誤: %v\n", err)
        return
    }
    
    // 添加查詢參數
    q := req.URL.Query()
    q.Add("q", "golang")
    q.Add("sort", "stars")
    q.Add("order", "desc")
    req.URL.RawQuery = q.Encode()
    
    // 發送請求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("請求 URL: %s\n", req.URL.String())
    fmt.Printf("狀態碼: %d\n", resp.StatusCode)
}
```

### 2. POST 請求

```go
import (
    "bytes"
    "encoding/json"
    "net/http"
    "strings"
)

// JSON POST 請求
func postJSON() {
    // 準備 JSON 數據
    data := map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
        "age":   30,
    }
    
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Printf("JSON 編碼錯誤: %v\n", err)
        return
    }
    
    // 發送 POST 請求
    resp, err := http.Post(
        "https://httpbin.org/post",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("響應: %s\n", body)
}

// 表單 POST 請求
func postForm() {
    // 準備表單數據
    formData := "name=Alice&email=alice@example.com&age=25"
    
    resp, err := http.Post(
        "https://httpbin.org/post",
        "application/x-www-form-urlencoded",
        strings.NewReader(formData),
    )
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("表單響應: %s\n", body)
}
```

## 🔧 自定義 HTTP 客戶端

### 1. 客戶端配置

```go
import (
    "crypto/tls"
    "net/http"
    "time"
)

// 創建自定義客戶端
func createCustomClient() *http.Client {
    // 自定義傳輸層
    transport := &http.Transport{
        MaxIdleConns:        100,                // 最大空閒連接數
        MaxIdleConnsPerHost: 10,                 // 每個主機的最大空閒連接數
        IdleConnTimeout:     90 * time.Second,   // 空閒連接超時
        DisableCompression:  false,              // 啟用壓縮
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: false, // 驗證 SSL 證書
        },
    }
    
    // 創建客戶端
    client := &http.Client{
        Transport: transport,
        Timeout:   30 * time.Second, // 總超時時間
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            // 自定義重定向邏輯
            if len(via) >= 10 {
                return fmt.Errorf("重定向次數過多")
            }
            return nil
        },
    }
    
    return client
}
```

### 2. 請求頭和認證

```go
// 設置請求頭
func requestWithHeaders() {
    req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    if err != nil {
        fmt.Printf("創建請求錯誤: %v\n", err)
        return
    }
    
    // 設置請求頭
    req.Header.Set("User-Agent", "MyApp/1.0")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Bearer YOUR_TOKEN")
    req.Header.Add("X-Custom-Header", "custom-value")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    // 讀取響應頭
    fmt.Println("響應頭:")
    for key, values := range resp.Header {
        for _, value := range values {
            fmt.Printf("  %s: %s\n", key, value)
        }
    }
}

// Basic 認證
func basicAuth() {
    req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/pass", nil)
    if err != nil {
        fmt.Printf("創建請求錯誤: %v\n", err)
        return
    }
    
    // 設置 Basic 認證
    req.SetBasicAuth("user", "pass")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("Basic 認證狀態碼: %d\n", resp.StatusCode)
}
```

## 📁 文件操作

### 1. 文件下載

```go
import (
    "io"
    "os"
)

// 下載文件
func downloadFile(url, filepath string) error {
    // 發送 GET 請求
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // 檢查狀態碼
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("下載失敗，狀態碼: %d", resp.StatusCode)
    }
    
    // 創建文件
    file, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 複製響應體到文件
    _, err = io.Copy(file, resp.Body)
    return err
}

// 帶進度的文件下載
type ProgressReader struct {
    io.Reader
    Total    int64
    Current  int64
    Callback func(current, total int64)
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
    n, err := pr.Reader.Read(p)
    pr.Current += int64(n)
    
    if pr.Callback != nil {
        pr.Callback(pr.Current, pr.Total)
    }
    
    return n, err
}

func downloadWithProgress(url, filepath string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    file, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 創建進度讀取器
    progressReader := &ProgressReader{
        Reader: resp.Body,
        Total:  resp.ContentLength,
        Callback: func(current, total int64) {
            if total > 0 {
                percentage := float64(current) / float64(total) * 100
                fmt.Printf("\r下載進度: %.2f%%", percentage)
            }
        },
    }
    
    _, err = io.Copy(file, progressReader)
    fmt.Println() // 換行
    return err
}
```

### 2. 文件上傳

```go
import (
    "bytes"
    "mime/multipart"
)

// 上傳文件
func uploadFile(url, fieldname, filename string) error {
    // 讀取文件
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 創建 multipart 表單
    var buf bytes.Buffer
    writer := multipart.NewWriter(&buf)
    
    // 添加文件字段
    part, err := writer.CreateFormFile(fieldname, filename)
    if err != nil {
        return err
    }
    
    _, err = io.Copy(part, file)
    if err != nil {
        return err
    }
    
    // 添加其他表單字段
    writer.WriteField("description", "文件上傳測試")
    
    // 關閉 writer
    writer.Close()
    
    // 創建請求
    req, err := http.NewRequest("POST", url, &buf)
    if err != nil {
        return err
    }
    
    req.Header.Set("Content-Type", writer.FormDataContentType())
    
    // 發送請求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    fmt.Printf("上傳狀態碼: %d\n", resp.StatusCode)
    return nil
}
```

## ⚡ 高級功能

### 1. 並發請求

```go
import (
    "sync"
)

// 並發請求結構
type ConcurrentRequest struct {
    URL    string
    Method string
    Data   interface{}
}

type RequestResult struct {
    URL        string
    StatusCode int
    Body       []byte
    Error      error
}

// 並發執行 HTTP 請求
func concurrentRequests(requests []ConcurrentRequest) []RequestResult {
    var wg sync.WaitGroup
    results := make([]RequestResult, len(requests))
    
    for i, req := range requests {
        wg.Add(1)
        go func(index int, request ConcurrentRequest) {
            defer wg.Done()
            
            result := RequestResult{URL: request.URL}
            
            resp, err := http.Get(request.URL)
            if err != nil {
                result.Error = err
                results[index] = result
                return
            }
            defer resp.Body.Close()
            
            result.StatusCode = resp.StatusCode
            result.Body, result.Error = io.ReadAll(resp.Body)
            results[index] = result
        }(i, req)
    }
    
    wg.Wait()
    return results
}
```

### 2. 重試機制

```go
import (
    "math"
    "time"
)

// HTTP 客戶端帶重試
type RetryClient struct {
    client      *http.Client
    maxRetries  int
    baseDelay   time.Duration
    maxDelay    time.Duration
}

func NewRetryClient(maxRetries int) *RetryClient {
    return &RetryClient{
        client:     &http.Client{Timeout: 30 * time.Second},
        maxRetries: maxRetries,
        baseDelay:  1 * time.Second,
        maxDelay:   30 * time.Second,
    }
}

func (rc *RetryClient) Do(req *http.Request) (*http.Response, error) {
    var resp *http.Response
    var err error
    
    for attempt := 0; attempt <= rc.maxRetries; attempt++ {
        resp, err = rc.client.Do(req)
        
        // 如果成功或不可重試的錯誤，直接返回
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }
        
        // 最後一次嘗試，不再重試
        if attempt == rc.maxRetries {
            break
        }
        
        // 計算延遲時間（指數退避）
        delay := time.Duration(math.Pow(2, float64(attempt))) * rc.baseDelay
        if delay > rc.maxDelay {
            delay = rc.maxDelay
        }
        
        fmt.Printf("請求失敗，%v 後重試 (嘗試 %d/%d)\n", delay, attempt+1, rc.maxRetries)
        time.Sleep(delay)
    }
    
    return resp, err
}
```

## 🍪 Cookie 和 Session 管理

### Cookie 處理

```go
import (
    "net/http/cookiejar"
    "net/url"
)

// 使用 Cookie Jar
func cookieExample() {
    // 創建 Cookie Jar
    jar, err := cookiejar.New(nil)
    if err != nil {
        fmt.Printf("創建 Cookie Jar 錯誤: %v\n", err)
        return
    }
    
    // 創建帶 Cookie 的客戶端
    client := &http.Client{
        Jar: jar,
    }
    
    // 第一個請求設置 Cookie
    resp1, err := client.Get("https://httpbin.org/cookies/set/session/abc123")
    if err != nil {
        fmt.Printf("第一個請求錯誤: %v\n", err)
        return
    }
    resp1.Body.Close()
    
    // 第二個請求會自動帶上 Cookie
    resp2, err := client.Get("https://httpbin.org/cookies")
    if err != nil {
        fmt.Printf("第二個請求錯誤: %v\n", err)
        return
    }
    defer resp2.Body.Close()
    
    body, _ := io.ReadAll(resp2.Body)
    fmt.Printf("Cookie 響應: %s\n", body)
}

// 手動設置 Cookie
func manualCookie() {
    req, _ := http.NewRequest("GET", "https://httpbin.org/cookies", nil)
    
    // 添加 Cookie
    cookie := &http.Cookie{
        Name:  "session",
        Value: "xyz789",
    }
    req.AddCookie(cookie)
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("手動 Cookie 響應: %s\n", body)
}
```

## 🔒 HTTPS 和 TLS

### TLS 配置

```go
import (
    "crypto/tls"
    "crypto/x509"
)

// 自定義 TLS 配置
func customTLS() {
    // 自定義 TLS 配置
    tlsConfig := &tls.Config{
        InsecureSkipVerify: false,                    // 驗證證書
        MinVersion:         tls.VersionTLS12,         // 最低 TLS 版本
        CipherSuites: []uint16{                       // 加密套件
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }
    
    transport := &http.Transport{
        TLSClientConfig: tlsConfig,
    }
    
    client := &http.Client{
        Transport: transport,
        Timeout:   30 * time.Second,
    }
    
    resp, err := client.Get("https://httpbin.org/get")
    if err != nil {
        fmt.Printf("HTTPS 請求錯誤: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("HTTPS 響應狀態: %d\n", resp.StatusCode)
}
```

## 🎯 最佳實踐

### 1. 資源管理
- 總是關閉響應體 `defer resp.Body.Close()`
- 設置合適的超時時間
- 重用 HTTP 客戶端實例

### 2. 錯誤處理
- 檢查 HTTP 狀態碼
- 處理網路錯誤和超時
- 實現重試機制

### 3. 性能優化
- 使用連接池
- 啟用 Keep-Alive
- 合理設置並發數

---

**下一章：[Web 伺服器](../19-web-server/)**