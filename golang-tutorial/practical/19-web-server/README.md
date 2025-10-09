# 第十九章：Web 伺服器

## 🎯 學習目標

- 掌握 Go 語言 Web 伺服器開發
- 理解 HTTP 路由和處理器
- 學會處理中間件和攔截器
- 掌握模板引擎和靜態文件服務
- 了解 RESTful API 設計和實現
- 學會處理 JSON、表單和文件上傳
- 掌握 Web 安全和性能優化

## 🌐 Web 伺服器基礎

Go 語言的 `net/http` 包提供了完整的 HTTP 伺服器功能。

### 核心概念

```
Web 伺服器架構：
┌─────────────────────────────────────┐
│ http.Server                          │
├─────────────────────────────────────┤
│ • Addr (監聽地址)                    │
│ • Handler (處理器)                   │
│ • ReadTimeout (讀取超時)             │
│ • WriteTimeout (寫入超時)            │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ http.ServeMux (路由多路復用器)        │
├─────────────────────────────────────┤
│ • HandleFunc (註冊處理函數)          │
│ • Handle (註冊處理器)                │
│ • ServeHTTP (處理請求)               │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ http.Handler (處理器接口)            │
├─────────────────────────────────────┤
│ • ServeHTTP(ResponseWriter, Request) │
└─────────────────────────────────────┘
```

## 🚀 基本 Web 伺服器

### 1. 最簡單的伺服器

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // 註冊處理函數
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    fmt.Println("伺服器啟動在 :8080")
    http.ListenAndServe(":8080", nil)
}
```

### 2. 多路由處理

```go
func setupRoutes() {
    // 首頁
    http.HandleFunc("/", homeHandler)
    
    // 用戶相關路由
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler) // 帶參數
    
    // API 路由
    http.HandleFunc("/api/status", statusHandler)
    http.HandleFunc("/api/users", apiUsersHandler)
    
    // 靜態文件
    http.Handle("/static/", http.StripPrefix("/static/", 
        http.FileServer(http.Dir("./static/"))))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    fmt.Fprintf(w, "歡迎來到首頁！")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        fmt.Fprintf(w, "獲取用戶列表")
    case "POST":
        fmt.Fprintf(w, "創建新用戶")
    default:
        http.Error(w, "方法不允許", http.StatusMethodNotAllowed)
    }
}
```

## 🛠️ 請求處理

### 1. 處理不同 HTTP 方法

```go
func methodHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleGET(w, r)
    case "POST":
        handlePOST(w, r)
    case "PUT":
        handlePUT(w, r)
    case "DELETE":
        handleDELETE(w, r)
    case "PATCH":
        handlePATCH(w, r)
    default:
        http.Error(w, "方法不支援", http.StatusMethodNotAllowed)
    }
}

func handleGET(w http.ResponseWriter, r *http.Request) {
    // 獲取查詢參數
    id := r.URL.Query().Get("id")
    name := r.URL.Query().Get("name")
    
    response := fmt.Sprintf("GET 請求 - ID: %s, Name: %s", id, name)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintf(w, response)
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
    // 解析表單數據
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "解析表單錯誤", http.StatusBadRequest)
        return
    }
    
    name := r.FormValue("name")
    email := r.FormValue("email")
    
    response := fmt.Sprintf("POST 請求 - Name: %s, Email: %s", name, email)
    fmt.Fprintf(w, response)
}
```

### 2. JSON API 處理

```go
import (
    "encoding/json"
    "io"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func jsonAPIHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case "GET":
        users := []User{
            {ID: 1, Name: "Alice", Email: "alice@example.com"},
            {ID: 2, Name: "Bob", Email: "bob@example.com"},
        }
        json.NewEncoder(w).Encode(users)
        
    case "POST":
        body, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "讀取請求體錯誤", http.StatusBadRequest)
            return
        }
        
        var user User
        err = json.Unmarshal(body, &user)
        if err != nil {
            http.Error(w, "JSON 解析錯誤", http.StatusBadRequest)
            return
        }
        
        // 模擬創建用戶
        user.ID = 123
        
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
        
    default:
        http.Error(w, "方法不支援", http.StatusMethodNotAllowed)
    }
}
```

## 🎨 模板引擎

### 1. HTML 模板

```go
import (
    "html/template"
)

type PageData struct {
    Title    string
    Username string
    Items    []string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>歡迎, {{.Username}}!</h1>
    <ul>
    {{range .Items}}
        <li>{{.}}</li>
    {{end}}
    </ul>
</body>
</html>`

    t, err := template.New("page").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    data := PageData{
        Title:    "我的頁面",
        Username: "Alice",
        Items:    []string{"項目1", "項目2", "項目3"},
    }
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    t.Execute(w, data)
}

// 從文件載入模板
func fileTemplateHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    data := PageData{
        Title:    "文件模板",
        Username: "Bob",
        Items:    []string{"讀書", "編程", "旅遊"},
    }
    
    tmpl.Execute(w, data)
}
```

### 2. 模板功能函數

```go
func customTemplateHandler(w http.ResponseWriter, r *http.Request) {
    // 自定義函數
    funcMap := template.FuncMap{
        "upper": strings.ToUpper,
        "lower": strings.ToLower,
        "add": func(a, b int) int {
            return a + b
        },
        "formatDate": func(t time.Time) string {
            return t.Format("2006-01-02 15:04:05")
        },
    }
    
    tmpl := `
<h1>{{.Title | upper}}</h1>
<p>當前時間: {{.Now | formatDate}}</p>
<p>計算結果: {{add 5 3}}</p>
<p>小寫標題: {{.Title | lower}}</p>
`

    t, err := template.New("custom").Funcs(funcMap).Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    data := map[string]interface{}{
        "Title": "Custom Template",
        "Now":   time.Now(),
    }
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    t.Execute(w, data)
}
```

## 🔧 中間件

### 1. 日誌中間件

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 創建響應記錄器
        recorder := &responseRecorder{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }
        
        // 執行下一個處理器
        next.ServeHTTP(recorder, r)
        
        // 記錄請求信息
        duration := time.Since(start)
        fmt.Printf("[%s] %s %s %d %v\n",
            start.Format("2006-01-02 15:04:05"),
            r.Method,
            r.URL.Path,
            recorder.statusCode,
            duration,
        )
    })
}

type responseRecorder struct {
    http.ResponseWriter
    statusCode int
}

func (rec *responseRecorder) WriteHeader(code int) {
    rec.statusCode = code
    rec.ResponseWriter.WriteHeader(code)
}
```

### 2. 認證中間件

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 檢查 Authorization 頭
        auth := r.Header.Get("Authorization")
        if auth == "" {
            http.Error(w, "未授權", http.StatusUnauthorized)
            return
        }
        
        // 簡單的 token 驗證
        if !strings.HasPrefix(auth, "Bearer ") {
            http.Error(w, "無效的授權格式", http.StatusUnauthorized)
            return
        }
        
        token := strings.TrimPrefix(auth, "Bearer ")
        if !isValidToken(token) {
            http.Error(w, "無效的 token", http.StatusUnauthorized)
            return
        }
        
        // 驗證通過，繼續處理
        next.ServeHTTP(w, r)
    })
}

func isValidToken(token string) bool {
    // 實際應用中應該驗證 JWT 或查詢數據庫
    return token == "valid-token-123"
}
```

### 3. CORS 中間件

```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 設置 CORS 頭
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // 處理預檢請求
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

## 📁 文件處理

### 1. 文件上傳

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "只支援 POST 方法", http.StatusMethodNotAllowed)
        return
    }
    
    // 解析 multipart 表單，限制大小為 10MB
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "解析表單錯誤", http.StatusBadRequest)
        return
    }
    
    // 獲取上傳的文件
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "獲取文件錯誤", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    // 創建保存文件
    dst, err := os.Create("uploads/" + handler.Filename)
    if err != nil {
        http.Error(w, "創建文件錯誤", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    
    // 複製文件內容
    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "保存文件錯誤", http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, "文件 %s 上傳成功！", handler.Filename)
}
```

### 2. 文件下載

```go
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("file")
    if filename == "" {
        http.Error(w, "缺少文件名參數", http.StatusBadRequest)
        return
    }
    
    // 安全檢查，防止路徑遍歷攻擊
    if strings.Contains(filename, "..") {
        http.Error(w, "無效的文件名", http.StatusBadRequest)
        return
    }
    
    filepath := "uploads/" + filename
    
    // 檢查文件是否存在
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        http.NotFound(w, r)
        return
    }
    
    // 設置響應頭
    w.Header().Set("Content-Disposition", "attachment; filename="+filename)
    w.Header().Set("Content-Type", "application/octet-stream")
    
    // 服務文件
    http.ServeFile(w, r, filepath)
}
```

## 🏗️ RESTful API

### 完整的 RESTful 用戶 API

```go
type UserAPI struct {
    users map[int]User
    nextID int
}

func NewUserAPI() *UserAPI {
    return &UserAPI{
        users: make(map[int]User),
        nextID: 1,
    }
}

func (api *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/api/users")
    
    switch {
    case path == "" || path == "/":
        api.handleUsers(w, r)
    case strings.HasPrefix(path, "/"):
        api.handleUser(w, r, strings.TrimPrefix(path, "/"))
    default:
        http.NotFound(w, r)
    }
}

func (api *UserAPI) handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        api.getUsers(w, r)
    case "POST":
        api.createUser(w, r)
    default:
        http.Error(w, "方法不支援", http.StatusMethodNotAllowed)
    }
}

func (api *UserAPI) handleUser(w http.ResponseWriter, r *http.Request, idStr string) {
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "無效的用戶 ID", http.StatusBadRequest)
        return
    }
    
    switch r.Method {
    case "GET":
        api.getUser(w, r, id)
    case "PUT":
        api.updateUser(w, r, id)
    case "DELETE":
        api.deleteUser(w, r, id)
    default:
        http.Error(w, "方法不支援", http.StatusMethodNotAllowed)
    }
}

func (api *UserAPI) getUsers(w http.ResponseWriter, r *http.Request) {
    users := make([]User, 0, len(api.users))
    for _, user := range api.users {
        users = append(users, user)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (api *UserAPI) createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "JSON 解析錯誤", http.StatusBadRequest)
        return
    }
    
    user.ID = api.nextID
    api.nextID++
    api.users[user.ID] = user
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (api *UserAPI) getUser(w http.ResponseWriter, r *http.Request, id int) {
    user, exists := api.users[id]
    if !exists {
        http.Error(w, "用戶不存在", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

## 🔒 安全考慮

### 1. 輸入驗證和清理

```go
import (
    "html"
    "regexp"
)

func sanitizeInput(input string) string {
    // HTML 轉義
    return html.EscapeString(input)
}

func validateEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
    // 輸入驗證
    name := r.FormValue("name")
    email := r.FormValue("email")
    
    if name == "" {
        http.Error(w, "姓名不能為空", http.StatusBadRequest)
        return
    }
    
    if !validateEmail(email) {
        http.Error(w, "無效的郵箱格式", http.StatusBadRequest)
        return
    }
    
    // 輸入清理
    safeName := sanitizeInput(name)
    safeEmail := sanitizeInput(email)
    
    response := fmt.Sprintf("姓名: %s, 郵箱: %s", safeName, safeEmail)
    fmt.Fprintf(w, response)
}
```

### 2. 速率限制

```go
import "sync"

type RateLimiter struct {
    requests map[string][]time.Time
    mutex    sync.RWMutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (rl *RateLimiter) Allow(clientIP string) bool {
    rl.mutex.Lock()
    defer rl.mutex.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // 清理過期請求
    requests := rl.requests[clientIP]
    validRequests := make([]time.Time, 0)
    for _, reqTime := range requests {
        if reqTime.After(cutoff) {
            validRequests = append(validRequests, reqTime)
        }
    }
    
    // 檢查是否超過限制
    if len(validRequests) >= rl.limit {
        return false
    }
    
    // 添加當前請求
    validRequests = append(validRequests, now)
    rl.requests[clientIP] = validRequests
    
    return true
}

func rateLimitMiddleware(limiter *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            clientIP := r.RemoteAddr
            
            if !limiter.Allow(clientIP) {
                http.Error(w, "請求過於頻繁", http.StatusTooManyRequests)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

## 🚀 性能優化

### 1. Gzip 壓縮

```go
import (
    "compress/gzip"
    "strings"
)

func gzipMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 檢查客戶端是否支援 gzip
        if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
            next.ServeHTTP(w, r)
            return
        }
        
        // 設置響應頭
        w.Header().Set("Content-Encoding", "gzip")
        w.Header().Set("Vary", "Accept-Encoding")
        
        // 創建 gzip 寫入器
        gz := gzip.NewWriter(w)
        defer gz.Close()
        
        // 包裝響應寫入器
        gzw := &gzipResponseWriter{
            ResponseWriter: w,
            Writer:         gz,
        }
        
        next.ServeHTTP(gzw, r)
    })
}

type gzipResponseWriter struct {
    http.ResponseWriter
    io.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
    return w.Writer.Write(b)
}
```

---

**下一章：[數據庫操作](../20-database/)**