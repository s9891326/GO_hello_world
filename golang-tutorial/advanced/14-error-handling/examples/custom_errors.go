package main

import (
	"fmt"
	"net"
	"time"
)

// 1. 基本自定義錯誤
type NetworkError struct {
	Op   string    // 操作類型
	Addr string    // 網路地址
	Time time.Time // 錯誤發生時間
	Err  error     // 原始錯誤
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("網路錯誤 [%s] %s 在 %v: %v", 
		e.Op, e.Addr, e.Time.Format("15:04:05"), e.Err)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// 2. 業務邏輯錯誤
type BusinessError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *BusinessError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("業務錯誤 [%d]: %s - %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("業務錯誤 [%d]: %s", e.Code, e.Message)
}

// 預定義業務錯誤
var (
	ErrUserNotFound     = &BusinessError{Code: 404, Message: "用戶不存在"}
	ErrInvalidPassword  = &BusinessError{Code: 401, Message: "密碼錯誤"}
	ErrInsufficientFund = &BusinessError{Code: 400, Message: "餘額不足"}
	ErrAccountLocked    = &BusinessError{Code: 423, Message: "帳戶已鎖定"}
)

// 3. 多重錯誤（收集多個錯誤）
type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	if len(e.Errors) == 0 {
		return "無錯誤"
	}
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	return fmt.Sprintf("發生 %d 個錯誤: %v", len(e.Errors), e.Errors[0])
}

func (e *MultiError) Add(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
}

func (e *MultiError) HasErrors() bool {
	return len(e.Errors) > 0
}

// 4. 帶有建議的錯誤
type HelpfulError struct {
	Err         error
	Suggestion  string
	DocumentURL string
}

func (e *HelpfulError) Error() string {
	msg := e.Err.Error()
	if e.Suggestion != "" {
		msg += fmt.Sprintf("\n建議: %s", e.Suggestion)
	}
	if e.DocumentURL != "" {
		msg += fmt.Sprintf("\n參考文檔: %s", e.DocumentURL)
	}
	return msg
}

func (e *HelpfulError) Unwrap() error {
	return e.Err
}

// 使用示例

// 模擬網路操作
func connectToServer(address string) error {
	// 模擬連接失敗
	originalErr := &net.OpError{
		Op:   "dial",
		Net:  "tcp",
		Addr: &net.TCPAddr{IP: net.ParseIP("192.168.1.1"), Port: 8080},
		Err:  fmt.Errorf("connection refused"),
	}
	
	return &NetworkError{
		Op:   "connect",
		Addr: address,
		Time: time.Now(),
		Err:  originalErr,
	}
}

// 模擬用戶認證
func authenticateUser(username, password string) error {
	if username == "" {
		return &HelpfulError{
			Err:         fmt.Errorf("用戶名不能為空"),
			Suggestion:  "請提供有效的用戶名",
			DocumentURL: "https://example.com/docs/auth",
		}
	}
	
	if username == "locked_user" {
		return ErrAccountLocked
	}
	
	if password != "correct_password" {
		return ErrInvalidPassword
	}
	
	return nil
}

// 模擬批量操作
func batchProcess(items []string) error {
	var multiErr MultiError
	
	for i, item := range items {
		if item == "" {
			multiErr.Add(fmt.Errorf("項目 %d: 不能為空", i))
		}
		if len(item) > 10 {
			multiErr.Add(fmt.Errorf("項目 %d: 長度超過限制", i))
		}
	}
	
	if multiErr.HasErrors() {
		return &multiErr
	}
	
	return nil
}

func main() {
	fmt.Println("=== 自定義錯誤類型示例 ===")
	
	// 1. 網路錯誤
	fmt.Println("\n1. 網路錯誤:")
	err := connectToServer("192.168.1.1:8080")
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
		
		// 檢查錯誤類型
		var netErr *NetworkError
		if errors.As(err, &netErr) {
			fmt.Printf("操作: %s, 地址: %s\n", netErr.Op, netErr.Addr)
		}
	}
	
	// 2. 業務邏輯錯誤
	fmt.Println("\n2. 業務邏輯錯誤:")
	testUsers := []struct {
		username string
		password string
	}{
		{"alice", "correct_password"},
		{"", "password"},
		{"locked_user", "password"},
		{"bob", "wrong_password"},
	}
	
	for _, user := range testUsers {
		err := authenticateUser(user.username, user.password)
		if err != nil {
			fmt.Printf("用戶 '%s': %v\n", user.username, err)
			
			// 檢查特定業務錯誤
			var businessErr *BusinessError
			if errors.As(err, &businessErr) {
				fmt.Printf("  錯誤代碼: %d\n", businessErr.Code)
			}
		} else {
			fmt.Printf("用戶 '%s': 認證成功\n", user.username)
		}
	}
	
	// 3. 多重錯誤
	fmt.Println("\n3. 多重錯誤:")
	items := []string{"valid", "", "this_is_too_long_item", "ok", ""}
	err = batchProcess(items)
	if err != nil {
		fmt.Printf("批量處理錯誤: %v\n", err)
		
		var multiErr *MultiError
		if errors.As(err, &multiErr) {
			fmt.Printf("總共 %d 個錯誤:\n", len(multiErr.Errors))
			for i, e := range multiErr.Errors {
				fmt.Printf("  %d. %v\n", i+1, e)
			}
		}
	}
}