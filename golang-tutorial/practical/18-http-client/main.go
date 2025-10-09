package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 演示基本 GET 請求
func demonstrateBasicGET() {
	fmt.Println("=== 基本 GET 請求 ===")
	
	// 1. 簡單 GET 請求
	fmt.Println("\n--- 簡單 GET 請求 ---")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("請求錯誤: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("狀態碼: %d\n", resp.StatusCode)
	fmt.Printf("狀態: %s\n", resp.Status)
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("讀取響應錯誤: %v\n", err)
		return
	}
	
	fmt.Printf("響應長度: %d 字節\n", len(body))
	
	// 2. 帶參數的 GET 請求
	fmt.Println("\n--- 帶參數的 GET 請求 ---")
	baseURL := "https://httpbin.org/get"
	
	// 方法1: 直接拼接URL
	urlWithParams := baseURL + "?name=Alice&age=25&city=New York"
	resp1, err := http.Get(urlWithParams)
	if err == nil {
		defer resp1.Body.Close()
		fmt.Printf("方法1 - 狀態碼: %d\n", resp1.StatusCode)
	}
	
	// 方法2: 使用 url.Values
	params := url.Values{}
	params.Add("name", "Bob")
	params.Add("age", "30")
	params.Add("hobbies", "reading")
	params.Add("hobbies", "coding")
	
	fullURL := baseURL + "?" + params.Encode()
	resp2, err := http.Get(fullURL)
	if err == nil {
		defer resp2.Body.Close()
		fmt.Printf("方法2 - 狀態碼: %d\n", resp2.StatusCode)
		fmt.Printf("請求URL: %s\n", fullURL)
	}
}

// 演示 POST 請求
func demonstratePOST() {
	fmt.Println("\n=== POST 請求演示 ===")
	
	// 1. JSON POST 請求
	fmt.Println("\n--- JSON POST 請求 ---")
	userData := map[string]interface{}{
		"name":    "Alice Johnson",
		"email":   "alice@example.com",
		"age":     25,
		"active":  true,
		"hobbies": []string{"reading", "coding", "travel"},
	}
	
	jsonData, err := json.Marshal(userData)
	if err != nil {
		fmt.Printf("JSON 編碼錯誤: %v\n", err)
		return
	}
	
	resp, err := http.Post(
		"https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Printf("POST 請求錯誤: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("JSON POST 狀態碼: %d\n", resp.StatusCode)
	
	// 2. 表單 POST 請求
	fmt.Println("\n--- 表單 POST 請求 ---")
	formData := url.Values{}
	formData.Set("username", "bob")
	formData.Set("password", "secret123")
	formData.Set("remember", "true")
	
	resp2, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		fmt.Printf("表單 POST 錯誤: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	
	fmt.Printf("表單 POST 狀態碼: %d\n", resp2.StatusCode)
	
	// 3. 自定義 POST 請求
	fmt.Println("\n--- 自定義 POST 請求 ---")
	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader("custom data"))
	if err != nil {
		fmt.Printf("創建請求錯誤: %v\n", err)
		return
	}
	
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	
	client := &http.Client{}
	resp3, err := client.Do(req)
	if err != nil {
		fmt.Printf("自定義 POST 錯誤: %v\n", err)
		return
	}
	defer resp3.Body.Close()
	
	fmt.Printf("自定義 POST 狀態碼: %d\n", resp3.StatusCode)
}

// 演示請求頭操作
func demonstrateHeaders() {
	fmt.Println("\n=== 請求頭操作 ===")
	
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		fmt.Printf("創建請求錯誤: %v\n", err)
		return
	}
	
	// 設置各種請求頭
	req.Header.Set("User-Agent", "MyApp/1.0 (Go HTTP Client)")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Add("X-Custom-Header", "CustomValue1")
	req.Header.Add("X-Custom-Header", "CustomValue2") // 添加多個相同名稱的頭
	
	fmt.Println("設置的請求頭:")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("請求錯誤: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("\n響應狀態碼: %d\n", resp.StatusCode)
	
	// 讀取響應頭
	fmt.Println("響應頭:")
	for name, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}
	
	// 讀取響應體
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("讀取響應錯誤: %v\n", err)
		return
	}
	
	// 解析響應 JSON 並顯示服務器收到的頭信息
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err == nil {
		if headers, ok := response["headers"].(map[string]interface{}); ok {
			fmt.Println("\n服務器收到的頭信息:")
			for key, value := range headers {
				fmt.Printf("  %s: %v\n", key, value)
			}
		}
	}
}

// 演示基本認證
func demonstrateAuthentication() {
	fmt.Println("\n=== 認證演示 ===")
	
	// 1. Basic 認證
	fmt.Println("\n--- Basic 認證 ---")
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
		fmt.Printf("Basic 認證請求錯誤: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Basic 認證狀態碼: %d\n", resp.StatusCode)
	
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("認證成功響應: %s\n", body)
	}
	
	// 2. Bearer Token 認證
	fmt.Println("\n--- Bearer Token 認證 ---")
	req2, err := http.NewRequest("GET", "https://httpbin.org/bearer", nil)
	if err != nil {
		fmt.Printf("創建 Bearer 請求錯誤: %v\n", err)
		return
	}
	
	// 設置 Bearer Token
	req2.Header.Set("Authorization", "Bearer your-token-here")
	
	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Printf("Bearer 認證請求錯誤: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	
	fmt.Printf("Bearer 認證狀態碼: %d\n", resp2.StatusCode)
	
	// 3. API Key 認證
	fmt.Println("\n--- API Key 認證 ---")
	req3, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Printf("創建 API Key 請求錯誤: %v\n", err)
		return
	}
	
	// 方法1: 在頭中設置 API Key
	req3.Header.Set("X-API-Key", "your-api-key-here")
	
	// 方法2: 在查詢參數中設置 API Key
	q := req3.URL.Query()
	q.Add("api_key", "your-api-key-here")
	req3.URL.RawQuery = q.Encode()
	
	resp3, err := client.Do(req3)
	if err != nil {
		fmt.Printf("API Key 認證請求錯誤: %v\n", err)
		return
	}
	defer resp3.Body.Close()
	
	fmt.Printf("API Key 認證狀態碼: %d\n", resp3.StatusCode)
}

// 演示自定義客戶端
func demonstrateCustomClient() {
	fmt.Println("\n=== 自定義客戶端 ===")
	
	// 創建自定義客戶端
	client := &http.Client{
		Timeout: 10 * time.Second, // 設置超時
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 自定義重定向處理
			fmt.Printf("重定向到: %s\n", req.URL.String())
			if len(via) >= 3 {
				return fmt.Errorf("重定向次數過多")
			}
			return nil
		},
	}
	
	// 測試重定向
	fmt.Println("\n--- 測試重定向 ---")
	resp, err := client.Get("https://httpbin.org/redirect/2")
	if err != nil {
		fmt.Printf("重定向請求錯誤: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("最終狀態碼: %d\n", resp.StatusCode)
		fmt.Printf("最終 URL: %s\n", resp.Request.URL.String())
	}
	
	// 測試超時
	fmt.Println("\n--- 測試超時 ---")
	shortTimeoutClient := &http.Client{
		Timeout: 1 * time.Second, // 很短的超時時間
	}
	
	start := time.Now()
	resp2, err := shortTimeoutClient.Get("https://httpbin.org/delay/3") // 延遲3秒響應
	duration := time.Since(start)
	
	if err != nil {
		fmt.Printf("超時請求錯誤 (耗時 %v): %v\n", duration, err)
	} else {
		defer resp2.Body.Close()
		fmt.Printf("請求成功，耗時: %v\n", duration)
	}
}

// 演示錯誤處理
func demonstrateErrorHandling() {
	fmt.Println("\n=== 錯誤處理 ===")
	
	client := &http.Client{Timeout: 5 * time.Second}
	
	// 測試各種錯誤情況
	testCases := []struct {
		name string
		url  string
	}{
		{"無效 URL", "invalid-url"},
		{"不存在的域名", "https://nonexistent-domain-12345.com"},
		{"404 錯誤", "https://httpbin.org/status/404"},
		{"500 錯誤", "https://httpbin.org/status/500"},
		{"超時", "https://httpbin.org/delay/10"},
	}
	
	for _, tc := range testCases {
		fmt.Printf("\n--- %s ---\n", tc.name)
		
		resp, err := client.Get(tc.url)
		if err != nil {
			fmt.Printf("請求錯誤: %v\n", err)
			continue
		}
		
		fmt.Printf("狀態碼: %d\n", resp.StatusCode)
		
		// 檢查狀態碼
		if resp.StatusCode >= 400 {
			fmt.Printf("HTTP 錯誤: %s\n", resp.Status)
		} else {
			fmt.Printf("請求成功: %s\n", resp.Status)
		}
		
		resp.Body.Close()
	}
}

// 演示 Cookie 處理
func demonstrateCookies() {
	fmt.Println("\n=== Cookie 處理 ===")
	
	// 手動設置 Cookie
	fmt.Println("\n--- 手動設置 Cookie ---")
	req, err := http.NewRequest("GET", "https://httpbin.org/cookies", nil)
	if err != nil {
		fmt.Printf("創建請求錯誤: %v\n", err)
		return
	}
	
	// 添加 Cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: "abc123",
	})
	req.AddCookie(&http.Cookie{
		Name:  "user",
		Value: "alice",
	})
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cookie 請求錯誤: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Cookie 響應: %s\n", body)
	
	// 讀取響應中的 Cookie
	fmt.Println("\n--- 響應 Cookie ---")
	setCookieURL := "https://httpbin.org/cookies/set/test/value123"
	resp2, err := client.Get(setCookieURL)
	if err != nil {
		fmt.Printf("設置 Cookie 請求錯誤: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	
	fmt.Println("響應中的 Cookie:")
	for _, cookie := range resp2.Cookies() {
		fmt.Printf("  %s = %s\n", cookie.Name, cookie.Value)
	}
}

// 主函數
func main() {
	fmt.Println("===== Go HTTP 客戶端示例 =====")
	
	demonstrateBasicGET()
	demonstratePOST()
	demonstrateHeaders()
	demonstrateAuthentication()
	demonstrateCustomClient()
	demonstrateErrorHandling()
	demonstrateCookies()
	
	fmt.Println("\n===== 示例完成 =====")
}