package main

import (
	"fmt"
	"log"
	"os"
	
	// 標準庫包
	"encoding/json"
	"net/http"
	"strings"
	"time"
	
	// 模擬的本地包導入（在實際項目中這些會是真實的包）
	// "myproject/pkg/utils"
	// "myproject/internal/config"
)

// 演示包的基本使用
func demonstrateStandardPackages() {
	fmt.Println("=== 標準庫包使用示例 ===")
	
	// strings 包
	fmt.Println("\n--- strings 包 ---")
	text := "Hello, Go Packages!"
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("大寫: %s\n", strings.ToUpper(text))
	fmt.Printf("小寫: %s\n", strings.ToLower(text))
	fmt.Printf("包含 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("替換: %s\n", strings.ReplaceAll(text, "Go", "Golang"))
	
	// time 包
	fmt.Println("\n--- time 包 ---")
	now := time.Now()
	fmt.Printf("當前時間: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix 時間戳: %d\n", now.Unix())
	
	// 創建特定時間
	birthday := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("千禧年: %s\n", birthday.Format("2006-01-02"))
	
	// json 包
	fmt.Println("\n--- json 包 ---")
	user := map[string]interface{}{
		"name":  "Alice",
		"age":   25,
		"email": "alice@example.com",
	}
	
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Printf("JSON 編碼錯誤: %v", err)
	} else {
		fmt.Printf("JSON 編碼: %s\n", jsonData)
	}
	
	// JSON 解碼
	var decoded map[string]interface{}
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		log.Printf("JSON 解碼錯誤: %v", err)
	} else {
		fmt.Printf("JSON 解碼: %+v\n", decoded)
	}
}

// 演示自定義包結構
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 導出的函數 - 首字母大寫
func NewUser(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// 導出的方法
func (u *User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}

// 導出的方法
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

// 未導出的函數 - 首字母小寫
func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// 導出的函數，內部使用未導出函數
func CreateUser(id int, name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("姓名不能為空")
	}
	
	if !validateEmail(email) {
		return nil, fmt.Errorf("無效的郵箱格式: %s", email)
	}
	
	return NewUser(id, name, email), nil
}

// 演示包級別變數和常數
var (
	// 導出的變數
	DefaultTimeout = 30 * time.Second
	MaxRetries     = 3
	
	// 未導出的變數
	internalCounter int
)

// 導出的常數
const (
	Version     = "1.0.0"
	MaxUsers    = 1000
	DefaultPort = 8080
)

// 未導出的常數
const (
	bufferSize = 1024
	maxRetries = 5
)

// 模擬 init 函數的使用
func init() {
	fmt.Println("package main: init 函數執行")
	
	// 初始化包級別變數
	internalCounter = 0
	
	// 檢查環境變數
	if port := os.Getenv("PORT"); port != "" {
		fmt.Printf("從環境變數讀取端口: %s\n", port)
	}
}

// 演示 HTTP 包的使用
func demonstrateHTTPPackage() {
	fmt.Println("\n=== HTTP 包使用示例 ===")
	
	// 創建簡單的 HTTP 處理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"message":   "Hello from Go HTTP server!",
			"timestamp": time.Now().Unix(),
			"path":      r.URL.Path,
			"method":    r.Method,
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []*User{
			NewUser(1, "Alice", "alice@example.com"),
			NewUser(2, "Bob", "bob@example.com"),
			NewUser(3, "Charlie", "charlie@example.com"),
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
	
	fmt.Printf("HTTP 服務器已配置，可以在端口 %d 上啟動\n", DefaultPort)
	fmt.Println("路由:")
	fmt.Println("  GET / - 基本信息")
	fmt.Println("  GET /users - 用戶列表")
	
	// 注意：在示例中我們不實際啟動服務器
	// 實際使用時可以用: log.Fatal(http.ListenAndServe(":8080", nil))
}

// 演示錯誤處理和包的結合
func demonstrateErrorHandling() {
	fmt.Println("\n=== 錯誤處理與包 ===")
	
	// 測試用戶創建
	testCases := []struct {
		id    int
		name  string
		email string
	}{
		{1, "Alice", "alice@example.com"},
		{2, "", "bob@example.com"},
		{3, "Charlie", "invalid-email"},
		{4, "David", "david@example.com"},
	}
	
	for _, tc := range testCases {
		user, err := CreateUser(tc.id, tc.name, tc.email)
		if err != nil {
			fmt.Printf("創建用戶失敗 [ID: %d]: %v\n", tc.id, err)
		} else {
			fmt.Printf("創建用戶成功: %s\n", user.String())
			
			// 演示 JSON 序列化
			jsonData, _ := user.ToJSON()
			fmt.Printf("  JSON: %s\n", jsonData)
		}
	}
}

// 演示文件操作包
func demonstrateFileOperations() {
	fmt.Println("\n=== 文件操作包示例 ===")
	
	// 使用 os 包
	fmt.Println("--- os 包 ---")
	
	// 獲取環境變數
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") // Windows
	}
	fmt.Printf("用戶主目錄: %s\n", home)
	
	// 獲取當前工作目錄
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("獲取當前目錄錯誤: %v", err)
	} else {
		fmt.Printf("當前工作目錄: %s\n", pwd)
	}
	
	// 獲取程序參數
	fmt.Printf("程序參數: %v\n", os.Args)
}

// 主函數
func main() {
	fmt.Printf("=== Go 包管理示例程序 ===\n")
	fmt.Printf("版本: %s\n", Version)
	fmt.Printf("最大用戶數: %d\n", MaxUsers)
	fmt.Printf("默認端口: %d\n", DefaultPort)
	fmt.Printf("默認超時: %v\n", DefaultTimeout)
	
	// 演示各種包的使用
	demonstrateStandardPackages()
	demonstrateErrorHandling()
	demonstrateHTTPPackage()
	demonstrateFileOperations()
	
	fmt.Println("\n=== 程序執行完成 ===")
}