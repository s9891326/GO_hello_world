# 第十九章練習：Web 伺服器

## 練習 1：RESTful API 伺服器

### 題目
實現一個完整的 RESTful API 伺服器，支援用戶和文章管理。

### 要求
```go
type User struct {
    ID       int       `json:"id"`
    Username string    `json:"username"`
    Email    string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type Post struct {
    ID       int       `json:"id"`
    Title    string    `json:"title"`
    Content  string    `json:"content"`
    AuthorID int       `json:"author_id"`
    CreatedAt time.Time `json:"created_at"`
}

type APIServer struct {
    users []User
    posts []Post
}

func NewAPIServer() *APIServer
func (s *APIServer) SetupRoutes() *http.ServeMux
func (s *APIServer) Start(addr string) error
```

### API 端點
- `GET /api/users` - 獲取所有用戶
- `POST /api/users` - 創建用戶
- `GET /api/users/{id}` - 獲取特定用戶
- `PUT /api/users/{id}` - 更新用戶
- `DELETE /api/users/{id}` - 刪除用戶
- `GET /api/posts` - 獲取所有文章
- `POST /api/posts` - 創建文章
- `GET /api/users/{id}/posts` - 獲取用戶的文章

### 功能要求
1. JSON 格式的請求和響應
2. 適當的 HTTP 狀態碼
3. 輸入驗證和錯誤處理
4. 支援查詢參數（分頁、排序）

---

## 練習 2：模板引擎和表單處理

### 題目
創建一個完整的 Web 應用，包含用戶註冊、登錄和管理界面。

### 要求
```go
type WebApp struct {
    templates *template.Template
    sessions  map[string]Session
}

type Session struct {
    UserID   int
    Username string
    ExpiresAt time.Time
}

func NewWebApp() *WebApp
func (app *WebApp) LoadTemplates() error
func (app *WebApp) SetupRoutes() *http.ServeMux
func (app *WebApp) RequireAuth(next http.HandlerFunc) http.HandlerFunc
```

### 頁面和功能
1. **首頁** (`/`) - 顯示歡迎信息
2. **註冊頁** (`/register`) - 用戶註冊表單
3. **登錄頁** (`/login`) - 用戶登錄表單
4. **用戶儀表板** (`/dashboard`) - 需要登錄
5. **用戶列表** (`/admin/users`) - 管理員功能

### 模板要求
- 使用 HTML 模板
- 支援佈局模板繼承
- 表單數據驗證
- Flash 消息顯示
- CSS 樣式美化

---

## 練習 3：文件上傳和處理

### 題目
實現一個文件上傳和管理系統。

### 要求
```go
type FileUploadServer struct {
    uploadDir   string
    maxFileSize int64
}

type UploadedFile struct {
    ID       string    `json:"id"`
    Filename string    `json:"filename"`
    Size     int64     `json:"size"`
    MimeType string    `json:"mime_type"`
    UploadedAt time.Time `json:"uploaded_at"`
    URL      string    `json:"url"`
}

func NewFileUploadServer(uploadDir string, maxFileSize int64) *FileUploadServer
func (s *FileUploadServer) HandleUpload(w http.ResponseWriter, r *http.Request)
func (s *FileUploadServer) HandleDownload(w http.ResponseWriter, r *http.Request)
func (s *FileUploadServer) HandleList(w http.ResponseWriter, r *http.Request)
func (s *FileUploadServer) HandleDelete(w http.ResponseWriter, r *http.Request)
```

### 功能特性
1. 多文件上傳支援
2. 文件類型驗證
3. 文件大小限制
4. 縮略圖生成（圖片文件）
5. 上傳進度顯示
6. 文件列表和搜索
7. 文件下載和刪除

---

## 練習 4：WebSocket 實時聊天

### 題目
實現一個 WebSocket 實時聊天應用。

### 要求
```go
type ChatServer struct {
    clients   map[*websocket.Conn]*Client
    broadcast chan []byte
    register  chan *Client
    unregister chan *Client
    rooms     map[string]*Room
}

type Client struct {
    conn     *websocket.Conn
    send     chan []byte
    username string
    room     string
}

type Room struct {
    name    string
    clients map[*Client]bool
}

type Message struct {
    Type     string    `json:"type"`
    Username string    `json:"username"`
    Content  string    `json:"content"`
    Room     string    `json:"room"`
    Timestamp time.Time `json:"timestamp"`
}

func NewChatServer() *ChatServer
func (s *ChatServer) Run()
func (s *ChatServer) HandleWebSocket(w http.ResponseWriter, r *http.Request)
```

### 功能要求
1. 多房間聊天支援
2. 用戶上線/下線通知
3. 私人消息
4. 消息歷史記錄
5. 在線用戶列表
6. 表情符號支援

---

## 練習 5：中間件系統

### 題目
設計一個靈活的中間件系統，支援各種功能。

### 要求
```go
type Middleware func(http.Handler) http.Handler

type Server struct {
    mux         *http.ServeMux
    middlewares []Middleware
}

func NewServer() *Server
func (s *Server) Use(middleware Middleware)
func (s *Server) Handle(pattern string, handler http.Handler)
func (s *Server) HandleFunc(pattern string, handler http.HandlerFunc)
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

### 內建中間件
```go
func LoggingMiddleware() Middleware
func AuthMiddleware(secretKey string) Middleware
func CORSMiddleware(origins []string) Middleware
func RateLimitMiddleware(rate int, burst int) Middleware
func CompressionMiddleware() Middleware
func SecurityHeadersMiddleware() Middleware
func CacheMiddleware(ttl time.Duration) Middleware
func MetricsMiddleware() Middleware
```

### 中間件功能
1. **日誌記錄** - 記錄所有請求
2. **身份驗證** - JWT token 驗證
3. **跨域支援** - CORS 處理
4. **請求限流** - 防止濫用
5. **響應壓縮** - Gzip 壓縮
6. **安全頭部** - 設置安全相關頭部
7. **緩存控制** - HTTP 緩存
8. **監控指標** - 收集性能數據

---

## 練習 6：微服務 API 網關

### 題目
實現一個簡單的 API 網關，支援路由、負載均衡和服務發現。

### 要求
```go
type APIGateway struct {
    routes   map[string]*Route
    backends map[string]*Backend
    lb       LoadBalancer
}

type Route struct {
    Pattern     string
    Methods     []string
    Backend     string
    Middlewares []Middleware
}

type Backend struct {
    Name      string
    Instances []string
    Health    HealthChecker
}

type LoadBalancer interface {
    NextInstance(backend string) string
}

func NewAPIGateway() *APIGateway
func (gw *APIGateway) AddRoute(route *Route)
func (gw *APIGateway) AddBackend(backend *Backend)
func (gw *APIGateway) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

### 功能特性
1. 路由匹配和轉發
2. 負載均衡算法（輪詢、權重、最少連接）
3. 健康檢查
4. 請求重試
5. 斷路器模式
6. 限流和熔斷
7. 監控和日誌

## 提交要求

1. 每個練習創建獨立的目錄和文件
2. 提供完整的 HTML 模板（如適用）
3. 包含測試用例和示例數據
4. 添加詳細的註釋和文檔
5. 遵循 Web 開發最佳實踐

## 評分標準

- **功能完整性** (35%)：實現所有要求的功能
- **API 設計** (25%)：RESTful 設計和響應格式
- **用戶體驗** (20%)：界面設計和交互流程
- **代碼品質** (15%)：代碼結構和可維護性
- **安全性** (5%)：輸入驗證和安全措施

## 額外挑戰

### 挑戰 1：GraphQL API
實現 GraphQL API 支援：
- Schema 定義
- Query 和 Mutation
- Subscription 支援
- 數據加載器

### 挑戰 2：gRPC 服務
添加 gRPC 服務支援：
- Protocol Buffers 定義
- 服務實現
- 客戶端生成
- 流式處理

### 挑戰 3：Docker 部署
容器化部署：
- Dockerfile 編寫
- Docker Compose 配置
- 多階段構建
- 健康檢查