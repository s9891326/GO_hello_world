# 第十八章練習：HTTP 客戶端

## 練習 1：REST API 客戶端

### 題目
實現一個完整的 REST API 客戶端，支援 CRUD 操作。

### 要求
```go
type APIClient struct {
    baseURL    string
    httpClient *http.Client
    apiKey     string
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func NewAPIClient(baseURL, apiKey string) *APIClient
func (c *APIClient) GetUsers() ([]User, error)
func (c *APIClient) GetUser(id int) (*User, error)
func (c *APIClient) CreateUser(user User) (*User, error)
func (c *APIClient) UpdateUser(id int, user User) (*User, error)
func (c *APIClient) DeleteUser(id int) error
func (c *APIClient) SearchUsers(query string) ([]User, error)
```

### 功能要求
1. 支援 API Key 認證
2. 自動處理 JSON 序列化/反序列化
3. 適當的錯誤處理
4. 請求/響應日誌記錄
5. 支援查詢參數

### 測試用例
- 測試所有 CRUD 操作
- 測試錯誤響應處理
- 測試認證失敗情況

---

## 練習 2：文件上傳下載客戶端

### 題目
創建一個支援文件上傳和下載的 HTTP 客戶端。

### 要求
```go
type FileClient struct {
    client  *http.Client
    baseURL string
}

func NewFileClient(baseURL string) *FileClient
func (fc *FileClient) UploadFile(endpoint string, filePath string, additionalFields map[string]string) (*UploadResponse, error)
func (fc *FileClient) DownloadFile(url string, savePath string) error
func (fc *FileClient) DownloadFileWithProgress(url string, savePath string, progressCallback func(int64, int64)) error
func (fc *FileClient) UploadMultipleFiles(endpoint string, files []string) ([]UploadResponse, error)
```

### 進階功能
1. 支援大文件分塊上傳
2. 斷點續傳下載
3. 上傳/下載進度顯示
4. 並發文件處理
5. MD5 校驗

### 數據結構
```go
type UploadResponse struct {
    FileID   string `json:"file_id"`
    Filename string `json:"filename"`
    Size     int64  `json:"size"`
    URL      string `json:"url"`
}
```

---

## 練習 3：重試機制和錯誤處理

### 題目
實現一個具有智能重試機制的 HTTP 客戶端。

### 要求
```go
type RetryConfig struct {
    MaxRetries  int
    BaseDelay   time.Duration
    MaxDelay    time.Duration
    BackoffRate float64
    RetryableStatusCodes []int
}

type RetryClient struct {
    client *http.Client
    config RetryConfig
}

func NewRetryClient(config RetryConfig) *RetryClient
func (rc *RetryClient) Do(req *http.Request) (*http.Response, error)
func (rc *RetryClient) Get(url string) (*http.Response, error)
func (rc *RetryClient) Post(url string, body io.Reader) (*http.Response, error)
```

### 重試策略
1. 指數退避算法
2. 抖動機制避免雷群效應
3. 可配置的重試條件
4. 斷路器模式
5. 請求超時處理

### 錯誤分類
- 網路錯誤：重試
- 4xx 錯誤：不重試
- 5xx 錯誤：重試
- 超時錯誤：重試

---

## 練習 4：並發 HTTP 請求池

### 題目
實現一個支援並發請求的 HTTP 客戶端池。

### 要求
```go
type RequestPool struct {
    maxConcurrency int
    semaphore     chan struct{}
    client        *http.Client
}

type BatchRequest struct {
    Method   string
    URL      string
    Headers  map[string]string
    Body     []byte
    Metadata interface{}
}

type BatchResponse struct {
    Request    BatchRequest
    Response   *http.Response
    Body       []byte
    Error      error
    Duration   time.Duration
    StatusCode int
}

func NewRequestPool(maxConcurrency int) *RequestPool
func (rp *RequestPool) ExecuteBatch(requests []BatchRequest) []BatchResponse
func (rp *RequestPool) ExecuteWithCallback(requests []BatchRequest, callback func(BatchResponse))
```

### 功能特性
1. 控制並發數量
2. 支援批量請求處理
3. 請求執行統計
4. 失敗請求重新排隊
5. 動態調整並發數

---

## 練習 5：HTTP 客戶端中間件

### 題目
設計一個支援中間件的 HTTP 客戶端架構。

### 要求
```go
type Middleware func(next RoundTripper) RoundTripper

type RoundTripper interface {
    RoundTrip(*http.Request) (*http.Response, error)
}

type MiddlewareClient struct {
    client      *http.Client
    middlewares []Middleware
}

func NewMiddlewareClient() *MiddlewareClient
func (mc *MiddlewareClient) Use(middleware Middleware)
func (mc *MiddlewareClient) Do(req *http.Request) (*http.Response, error)
```

### 內建中間件
```go
func LoggingMiddleware() Middleware
func AuthMiddleware(token string) Middleware
func RetryMiddleware(maxRetries int) Middleware
func RateLimitMiddleware(rate int) Middleware
func CacheMiddleware(cache Cache) Middleware
func MetricsMiddleware(metrics MetricsCollector) Middleware
```

### 中間件功能
1. 請求/響應日誌記錄
2. 自動添加認證頭
3. 請求重試
4. 請求限速
5. 響應緩存
6. 性能指標收集

---

## 練習 6：WebSocket 客戶端

### 題目
實現一個 WebSocket 客戶端，支援實時通信。

### 要求
```go
type WebSocketClient struct {
    conn        *websocket.Conn
    url         string
    messageHandler func([]byte)
    errorHandler   func(error)
}

func NewWebSocketClient(url string) *WebSocketClient
func (wsc *WebSocketClient) Connect() error
func (wsc *WebSocketClient) Disconnect() error
func (wsc *WebSocketClient) SendMessage(message []byte) error
func (wsc *WebSocketClient) SendJSON(data interface{}) error
func (wsc *WebSocketClient) SetMessageHandler(handler func([]byte))
func (wsc *WebSocketClient) SetErrorHandler(handler func(error))
func (wsc *WebSocketClient) StartListening()
```

### 功能特性
1. 自動重連機制
2. 心跳檢測
3. 消息隊列
4. 二進制和文本消息支援
5. 連接狀態管理

## 提交要求

1. 每個練習創建獨立的 `.go` 文件
2. 提供完整的測試用例
3. 包含錯誤處理和日誌記錄
4. 遵循 HTTP 客戶端最佳實踐
5. 提供使用示例和文檔

## 評分標準

- **功能完整性** (35%)：實現所有要求的功能
- **錯誤處理** (25%)：妥善處理各種網路錯誤
- **並發安全** (20%)：正確處理並發請求
- **代碼品質** (15%)：代碼結構清晰，註釋充分
- **性能優化** (5%)：高效的網路請求處理

## 額外挑戰

### 挑戰 1：HTTP/2 支援
實現 HTTP/2 協議支援，包括：
- 服務器推送
- 多路復用
- 頭部壓縮

### 挑戰 2：代理支援
添加代理服務器支援：
- HTTP 代理
- SOCKS 代理
- 代理認證
- 代理鏈

### 挑戰 3：性能監控
實現詳細的性能監控：
- 請求延遲統計
- 成功率監控
- 連接池狀態
- 內存使用監控