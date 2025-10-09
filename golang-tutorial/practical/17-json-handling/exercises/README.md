# 第十七章練習：JSON 處理

## 練習 1：基本 JSON 操作

### 題目
實現一個學生成績管理系統，支援 JSON 序列化和反序列化。

### 要求
```go
type Student struct {
    ID     int     `json:"id"`
    Name   string  `json:"name"`
    Age    int     `json:"age"`
    Email  string  `json:"email"`
    Scores []Score `json:"scores"`
}

type Score struct {
    Subject string  `json:"subject"`
    Points  float64 `json:"points"`
    Date    string  `json:"date"`
}

type StudentManager struct {
    students []Student
}

func NewStudentManager() *StudentManager
func (sm *StudentManager) AddStudent(student Student) error
func (sm *StudentManager) GetStudentByID(id int) (*Student, error)
func (sm *StudentManager) SaveToJSON(filename string) error
func (sm *StudentManager) LoadFromJSON(filename string) error
func (sm *StudentManager) GetTopStudents(limit int) []Student
```

### 測試數據
```json
{
  "id": 1,
  "name": "Alice Johnson",
  "age": 20,
  "email": "alice@university.edu",
  "scores": [
    {"subject": "Math", "points": 95.5, "date": "2023-10-15"},
    {"subject": "Physics", "points": 88.0, "date": "2023-10-20"}
  ]
}
```

---

## 練習 2：自定義 JSON 序列化

### 題目
實現一個支援自定義時間格式和加密字段的用戶系統。

### 要求
```go
type SecureUser struct {
    ID          int           `json:"id"`
    Username    string        `json:"username"`
    Password    EncryptedText `json:"password"`
    CreatedAt   CustomTime    `json:"created_at"`
    LastLoginAt *CustomTime   `json:"last_login_at,omitempty"`
    Profile     UserProfile   `json:"profile"`
}

type EncryptedText string
type CustomTime struct {
    time.Time
}

type UserProfile struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone,omitempty"`
}
```

### 自定義要求
1. `EncryptedText` 序列化時加密，反序列化時解密
2. `CustomTime` 使用 "YYYY-MM-DD HH:mm:ss" 格式
3. 實現 `json.Marshaler` 和 `json.Unmarshaler` 接口
4. 處理空時間和加密錯誤

---

## 練習 3：動態 JSON 處理

### 題目
創建一個 API 響應處理器，能夠處理不同格式的響應。

### 要求
```go
type APIResponse struct {
    Status  string          `json:"status"`
    Message string          `json:"message"`
    Data    json.RawMessage `json:"data"`
    Meta    ResponseMeta    `json:"meta,omitempty"`
}

type ResponseMeta struct {
    Page       int `json:"page,omitempty"`
    TotalPages int `json:"total_pages,omitempty"`
    Total      int `json:"total,omitempty"`
}

type ResponseProcessor struct{}

func NewResponseProcessor() *ResponseProcessor
func (rp *ResponseProcessor) ProcessUserResponse(response []byte) ([]User, error)
func (rp *ResponseProcessor) ProcessProductResponse(response []byte) ([]Product, error)
func (rp *ResponseProcessor) ProcessGenericResponse(response []byte) (interface{}, error)
```

### 測試場景
- 處理用戶列表響應
- 處理商品信息響應
- 處理錯誤響應
- 處理嵌套對象響應

---

## 練習 4：JSON 流處理

### 題目
實現一個大數據 JSON 處理器，支援流式讀取和寫入。

### 要求
```go
type StreamProcessor struct {
    batchSize int
}

func NewStreamProcessor(batchSize int) *StreamProcessor
func (sp *StreamProcessor) ProcessJSONStream(input io.Reader, output io.Writer, processor func(interface{}) interface{}) error
func (sp *StreamProcessor) ConvertCSVToJSON(csvFile, jsonFile string) error
func (sp *StreamProcessor) FilterJSONObjects(input io.Reader, output io.Writer, filter func(map[string]interface{}) bool) error
```

### 功能要求
1. 支援大文件流式處理
2. 內存使用量控制
3. 錯誤恢復機制
4. 進度報告

---

## 練習 5：JSON 配置管理器

### 題目
創建一個應用配置管理器，支援JSON配置文件的熱重載。

### 要求
```go
type Config struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Redis    RedisConfig    `json:"redis"`
    Logging  LoggingConfig  `json:"logging"`
}

type ConfigManager struct {
    configFile string
    config     *Config
    watchers   []func(*Config)
}

func NewConfigManager(configFile string) *ConfigManager
func (cm *ConfigManager) Load() error
func (cm *ConfigManager) Save() error
func (cm *ConfigManager) GetConfig() *Config
func (cm *ConfigManager) UpdateConfig(updater func(*Config)) error
func (cm *ConfigManager) AddWatcher(watcher func(*Config))
func (cm *ConfigManager) StartWatching() error
```

### 高級功能
- 配置文件變化監控
- 配置驗證
- 默認值處理
- 環境變數覆蓋

---

## 練習 6：JSON API 客戶端

### 題目
實現一個通用的 JSON API 客戶端，支援各種 REST 操作。

### 要求
```go
type APIClient struct {
    baseURL    string
    httpClient *http.Client
    headers    map[string]string
}

func NewAPIClient(baseURL string) *APIClient
func (ac *APIClient) SetHeader(key, value string)
func (ac *APIClient) Get(endpoint string, result interface{}) error
func (ac *APIClient) Post(endpoint string, payload interface{}, result interface{}) error
func (ac *APIClient) Put(endpoint string, payload interface{}, result interface{}) error
func (ac *APIClient) Delete(endpoint string) error
func (ac *APIClient) Upload(endpoint string, fieldName string, filename string, result interface{}) error
```

### 額外功能
- 自動重試機制
- 請求/響應日誌
- 錯誤處理統一化
- 認證支援

## 提交要求

1. 每個練習創建獨立的 `.go` 文件
2. 提供完整的測試用例
3. 包含詳細的註釋說明
4. 處理所有可能的錯誤情況
5. 遵循 Go 語言 JSON 處理最佳實踐

## 評分標準

- **功能完整性** (40%)：實現所有要求的功能
- **JSON 處理** (30%)：正確的序列化和反序列化
- **錯誤處理** (15%)：妥善處理各種錯誤情況
- **代碼品質** (10%)：代碼結構清晰，註釋充分
- **性能考慮** (5%)：高效的 JSON 處理實現