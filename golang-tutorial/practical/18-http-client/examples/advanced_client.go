package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"sync"
	"time"
)

// HTTPClient 高級HTTP客戶端
type HTTPClient struct {
	client     *http.Client
	baseURL    string
	headers    map[string]string
	retryCount int
	timeout    time.Duration
	mutex      sync.RWMutex
}

// NewHTTPClient 創建新的HTTP客戶端
func NewHTTPClient(baseURL string) *HTTPClient {
	// 創建Cookie Jar
	jar, _ := cookiejar.New(nil)
	
	// 自定義Transport
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	
	return &HTTPClient{
		client: &http.Client{
			Transport: transport,
			Jar:       jar,
			Timeout:   30 * time.Second,
		},
		baseURL:    strings.TrimRight(baseURL, "/"),
		headers:    make(map[string]string),
		retryCount: 3,
		timeout:    30 * time.Second,
	}
}

// SetHeader 設置默認請求頭
func (c *HTTPClient) SetHeader(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.headers[key] = value
}

// SetTimeout 設置超時時間
func (c *HTTPClient) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
	c.timeout = timeout
}

// SetRetryCount 設置重試次數
func (c *HTTPClient) SetRetryCount(count int) {
	c.retryCount = count
}

// buildURL 構建完整URL
func (c *HTTPClient) buildURL(endpoint string) string {
	endpoint = strings.TrimLeft(endpoint, "/")
	return fmt.Sprintf("%s/%s", c.baseURL, endpoint)
}

// addHeaders 添加默認頭部
func (c *HTTPClient) addHeaders(req *http.Request) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

// doWithRetry 帶重試的請求執行
func (c *HTTPClient) doWithRetry(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		// 創建請求副本（因為Body可能被消費）
		var bodyReader io.Reader
		if req.Body != nil {
			bodyBytes, _ := io.ReadAll(req.Body)
			req.Body.Close()
			bodyReader = bytes.NewReader(bodyBytes)
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		
		resp, err = c.client.Do(req)
		
		// 成功或不可重試的錯誤
		if err == nil || attempt == c.retryCount {
			break
		}
		
		// 重試延遲
		delay := time.Duration(attempt+1) * time.Second
		fmt.Printf("請求失敗，%v 後重試 (嘗試 %d/%d): %v\n", delay, attempt+1, c.retryCount, err)
		time.Sleep(delay)
		
		// 重置Body用於重試
		if bodyReader != nil {
			req.Body = io.NopCloser(bodyReader)
		}
	}
	
	return resp, err
}

// GET 執行GET請求
func (c *HTTPClient) GET(endpoint string) (*http.Response, error) {
	url := c.buildURL(endpoint)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	c.addHeaders(req)
	return c.doWithRetry(req)
}

// POST 執行POST請求
func (c *HTTPClient) POST(endpoint string, data interface{}) (*http.Response, error) {
	url := c.buildURL(endpoint)
	
	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}
	
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	c.addHeaders(req)
	
	return c.doWithRetry(req)
}

// PUT 執行PUT請求
func (c *HTTPClient) PUT(endpoint string, data interface{}) (*http.Response, error) {
	url := c.buildURL(endpoint)
	
	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}
	
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	c.addHeaders(req)
	
	return c.doWithRetry(req)
}

// DELETE 執行DELETE請求
func (c *HTTPClient) DELETE(endpoint string) (*http.Response, error) {
	url := c.buildURL(endpoint)
	
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	
	c.addHeaders(req)
	return c.doWithRetry(req)
}

// UploadFile 上傳文件
func (c *HTTPClient) UploadFile(endpoint string, fieldName string, filePath string, additionalFields map[string]string) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	// 創建multipart表單
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	
	// 添加文件字段
	part, err := writer.CreateFormFile(fieldName, filePath)
	if err != nil {
		return nil, err
	}
	
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	
	// 添加其他字段
	for key, value := range additionalFields {
		writer.WriteField(key, value)
	}
	
	writer.Close()
	
	url := c.buildURL(endpoint)
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", writer.FormDataContentType())
	c.addHeaders(req)
	
	return c.doWithRetry(req)
}

// DownloadFile 下載文件
func (c *HTTPClient) DownloadFile(endpoint string, savePath string) error {
	resp, err := c.GET(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下載失敗，狀態碼: %d", resp.StatusCode)
	}
	
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = io.Copy(file, resp.Body)
	return err
}

// 帶進度的下載
type ProgressReader struct {
	io.Reader
	total    int64
	current  int64
	callback func(current, total int64)
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.current += int64(n)
	
	if pr.callback != nil {
		pr.callback(pr.current, pr.total)
	}
	
	return n, err
}

// DownloadFileWithProgress 帶進度的文件下載
func (c *HTTPClient) DownloadFileWithProgress(endpoint string, savePath string, progressCallback func(current, total int64)) error {
	resp, err := c.GET(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下載失敗，狀態碼: %d", resp.StatusCode)
	}
	
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	progressReader := &ProgressReader{
		Reader:   resp.Body,
		total:    resp.ContentLength,
		callback: progressCallback,
	}
	
	_, err = io.Copy(file, progressReader)
	return err
}

// GetJSON 獲取JSON響應
func (c *HTTPClient) GetJSON(endpoint string, result interface{}) error {
	resp, err := c.GET(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("請求失敗，狀態碼: %d", resp.StatusCode)
	}
	
	return json.NewDecoder(resp.Body).Decode(result)
}

// PostJSON 發送JSON並獲取JSON響應
func (c *HTTPClient) PostJSON(endpoint string, data interface{}, result interface{}) error {
	resp, err := c.POST(endpoint, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("請求失敗，狀態碼: %d", resp.StatusCode)
	}
	
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	
	return nil
}

// 並發請求管理器
type ConcurrentRequestManager struct {
	client   *HTTPClient
	maxConcurrency int
	semaphore chan struct{}
}

// NewConcurrentRequestManager 創建並發請求管理器
func NewConcurrentRequestManager(client *HTTPClient, maxConcurrency int) *ConcurrentRequestManager {
	return &ConcurrentRequestManager{
		client:         client,
		maxConcurrency: maxConcurrency,
		semaphore:      make(chan struct{}, maxConcurrency),
	}
}

// RequestResult 請求結果
type RequestResult struct {
	URL        string
	StatusCode int
	Body       []byte
	Error      error
	Duration   time.Duration
}

// ConcurrentGET 並發執行多個GET請求
func (crm *ConcurrentRequestManager) ConcurrentGET(endpoints []string) []RequestResult {
	results := make([]RequestResult, len(endpoints))
	var wg sync.WaitGroup
	
	for i, endpoint := range endpoints {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()
			
			// 獲取信號量
			crm.semaphore <- struct{}{}
			defer func() { <-crm.semaphore }()
			
			start := time.Now()
			result := RequestResult{URL: url}
			
			resp, err := crm.client.GET(url)
			result.Duration = time.Since(start)
			
			if err != nil {
				result.Error = err
			} else {
				result.StatusCode = resp.StatusCode
				result.Body, result.Error = io.ReadAll(resp.Body)
				resp.Body.Close()
			}
			
			results[index] = result
		}(i, endpoint)
	}
	
	wg.Wait()
	return results
}

// 演示高級HTTP客戶端功能
func demonstrateAdvancedClient() {
	fmt.Println("=== 高級HTTP客戶端演示 ===")
	
	// 創建客戶端
	client := NewHTTPClient("https://httpbin.org")
	
	// 設置默認頭部
	client.SetHeader("User-Agent", "Advanced-Go-Client/1.0")
	client.SetHeader("Accept", "application/json")
	
	// 1. 基本請求
	fmt.Println("\n--- 基本GET請求 ---")
	resp, err := client.GET("/get")
	if err != nil {
		fmt.Printf("GET請求錯誤: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("狀態碼: %d\n", resp.StatusCode)
	}
	
	// 2. JSON操作
	fmt.Println("\n--- JSON操作 ---")
	data := map[string]interface{}{
		"name":  "Alice",
		"age":   25,
		"email": "alice@example.com",
	}
	
	var result map[string]interface{}
	err = client.PostJSON("/post", data, &result)
	if err != nil {
		fmt.Printf("PostJSON錯誤: %v\n", err)
	} else {
		fmt.Printf("POST響應成功\n")
	}
	
	// 3. 文件下載演示
	fmt.Println("\n--- 文件下載演示 ---")
	err = client.DownloadFileWithProgress("/bytes/1024", "downloaded_file.bin", 
		func(current, total int64) {
			if total > 0 {
				percentage := float64(current) / float64(total) * 100
				fmt.Printf("\r下載進度: %.2f%%", percentage)
			}
		})
	
	if err != nil {
		fmt.Printf("\n下載錯誤: %v\n", err)
	} else {
		fmt.Printf("\n文件下載完成\n")
	}
	
	// 4. 並發請求演示
	fmt.Println("\n--- 並發請求演示 ---")
	crm := NewConcurrentRequestManager(client, 3)
	
	endpoints := []string{
		"/delay/1",
		"/delay/2", 
		"/delay/3",
		"/status/200",
		"/status/404",
	}
	
	results := crm.ConcurrentGET(endpoints)
	
	fmt.Println("並發請求結果:")
	for i, result := range results {
		if result.Error != nil {
			fmt.Printf("  %d. %s - 錯誤: %v (耗時: %v)\n", 
				i+1, result.URL, result.Error, result.Duration)
		} else {
			fmt.Printf("  %d. %s - 狀態碼: %d (耗時: %v)\n", 
				i+1, result.URL, result.StatusCode, result.Duration)
		}
	}
}

// 演示錯誤處理和重試
func demonstrateErrorHandling() {
	fmt.Println("\n=== 錯誤處理和重試演示 ===")
	
	client := NewHTTPClient("https://httpbin.org")
	client.SetRetryCount(3)
	client.SetTimeout(5 * time.Second)
	
	// 測試超時
	fmt.Println("\n--- 測試超時 ---")
	client.SetTimeout(1 * time.Second)
	_, err := client.GET("/delay/3")
	if err != nil {
		fmt.Printf("超時錯誤: %v\n", err)
	}
	
	// 恢復正常超時
	client.SetTimeout(10 * time.Second)
	
	// 測試狀態碼錯誤
	fmt.Println("\n--- 測試HTTP錯誤狀態 ---")
	errorCodes := []int{400, 404, 500, 503}
	
	for _, code := range errorCodes {
		endpoint := fmt.Sprintf("/status/%d", code)
		resp, err := client.GET(endpoint)
		if err != nil {
			fmt.Printf("狀態碼 %d 請求錯誤: %v\n", code, err)
		} else {
			fmt.Printf("狀態碼 %d 響應: %d\n", code, resp.StatusCode)
			resp.Body.Close()
		}
	}
}

// 清理測試文件
func cleanup() {
	os.Remove("downloaded_file.bin")
	fmt.Println("清理完成")
}

func main() {
	demonstrateAdvancedClient()
	demonstrateErrorHandling()
	
	fmt.Println("\n按 Enter 鍵清理測試文件...")
	fmt.Scanln()
	
	cleanup()
	
	fmt.Println("===== 演示完成 =====")
}