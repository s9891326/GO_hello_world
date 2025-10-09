package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// 1. 錯誤處理模式 - Guard Clauses
func processUser(userID string, email string, age int) error {
	// Guard clauses - 早期返回錯誤條件
	if userID == "" {
		return errors.New("用戶ID不能為空")
	}
	
	if email == "" {
		return errors.New("郵箱不能為空")
	}
	
	if age < 0 || age > 120 {
		return errors.New("年齡必須在0-120之間")
	}
	
	// 主要業務邏輯
	fmt.Printf("處理用戶: ID=%s, Email=%s, Age=%d\n", userID, email, age)
	return nil
}

// 2. 錯誤聚合模式
func validateUserData(users []map[string]interface{}) []error {
	var errors []error
	
	for i, user := range users {
		if name, ok := user["name"].(string); !ok || name == "" {
			errors = append(errors, fmt.Errorf("用戶 %d: 姓名無效", i))
		}
		
		if age, ok := user["age"].(float64); !ok || age < 0 {
			errors = append(errors, fmt.Errorf("用戶 %d: 年齡無效", i))
		}
		
		if email, ok := user["email"].(string); !ok || email == "" {
			errors = append(errors, fmt.Errorf("用戶 %d: 郵箱無效", i))
		}
	}
	
	return errors
}

// 3. 重試模式
type RetryConfig struct {
	MaxAttempts int
	Delay       time.Duration
	BackoffRate float64
}

func retryWithBackoff(operation func() error, config RetryConfig) error {
	var lastErr error
	delay := config.Delay
	
	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		err := operation()
		if err == nil {
			if attempt > 1 {
				fmt.Printf("操作在第 %d 次嘗試後成功\n", attempt)
			}
			return nil
		}
		
		lastErr = err
		
		if attempt == config.MaxAttempts {
			break
		}
		
		fmt.Printf("第 %d 次嘗試失敗: %v, %v 後重試\n", attempt, err, delay)
		time.Sleep(delay)
		delay = time.Duration(float64(delay) * config.BackoffRate)
	}
	
	return fmt.Errorf("操作在 %d 次嘗試後失敗: %w", config.MaxAttempts, lastErr)
}

// 4. 帶有超時的錯誤處理
func operationWithTimeout(ctx context.Context, duration time.Duration) error {
	// 創建帶超時的上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	
	// 模擬長時間運行的操作
	resultCh := make(chan error, 1)
	go func() {
		// 模擬工作
		time.Sleep(2 * time.Second)
		resultCh <- nil // 成功
	}()
	
	select {
	case err := <-resultCh:
		return err
	case <-timeoutCtx.Done():
		return fmt.Errorf("操作超時: %w", timeoutCtx.Err())
	}
}

// 5. 錯誤恢復模式
func safeOperation(data string) (result string, err error) {
	// 使用 defer 進行錯誤恢復
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("操作發生 panic: %v", r)
		}
	}()
	
	// 可能會 panic 的操作
	if data == "panic" {
		panic("模擬 panic")
	}
	
	if data == "error" {
		return "", errors.New("模擬錯誤")
	}
	
	return "處理成功: " + data, nil
}

// 6. 錯誤上下文包裝
type OperationContext struct {
	UserID    string
	Operation string
	RequestID string
}

func (ctx *OperationContext) WrapError(err error, message string) error {
	return fmt.Errorf("%s [User: %s, Op: %s, ReqID: %s]: %w", 
		message, ctx.UserID, ctx.Operation, ctx.RequestID, err)
}

func businessOperation(ctx *OperationContext, data string) error {
	// 模擬數據驗證錯誤
	if data == "" {
		return ctx.WrapError(errors.New("數據為空"), "數據驗證失敗")
	}
	
	// 模擬數據庫錯誤
	if data == "db_error" {
		dbErr := errors.New("連接數據庫失敗")
		return ctx.WrapError(dbErr, "數據庫操作失敗")
	}
	
	return nil
}

// 7. 錯誤分類和處理
type ErrorType int

const (
	ErrorTypeValidation ErrorType = iota
	ErrorTypeNetwork
	ErrorTypeDatabase
	ErrorTypePermission
	ErrorTypeInternal
)

type ClassifiedError struct {
	Type    ErrorType
	Message string
	Code    string
	Err     error
}

func (e *ClassifiedError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
}

func classifyAndHandle(err error) {
	var classifiedErr *ClassifiedError
	if errors.As(err, &classifiedErr) {
		switch classifiedErr.Type {
		case ErrorTypeValidation:
			log.Printf("驗證錯誤: %v", err)
			// 返回 400 Bad Request
		case ErrorTypeNetwork:
			log.Printf("網路錯誤: %v", err)
			// 重試或返回 503 Service Unavailable
		case ErrorTypeDatabase:
			log.Printf("數據庫錯誤: %v", err)
			// 重試或返回 500 Internal Server Error
		case ErrorTypePermission:
			log.Printf("權限錯誤: %v", err)
			// 返回 403 Forbidden
		default:
			log.Printf("未知錯誤: %v", err)
			// 返回 500 Internal Server Error
		}
	} else {
		log.Printf("未分類錯誤: %v", err)
	}
}

func main() {
	fmt.Println("=== 錯誤處理模式示例 ===")
	
	// 1. Guard Clauses 模式
	fmt.Println("\n1. Guard Clauses 模式:")
	testCases := []struct {
		userID string
		email  string
		age    int
	}{
		{"user1", "user1@example.com", 25},
		{"", "user2@example.com", 30},
		{"user3", "", 35},
		{"user4", "user4@example.com", -5},
	}
	
	for _, tc := range testCases {
		err := processUser(tc.userID, tc.email, tc.age)
		if err != nil {
			fmt.Printf("錯誤: %v\n", err)
		}
	}
	
	// 2. 錯誤聚合模式
	fmt.Println("\n2. 錯誤聚合模式:")
	users := []map[string]interface{}{
		{"name": "Alice", "age": 25.0, "email": "alice@example.com"},
		{"name": "", "age": 30.0, "email": "bob@example.com"},
		{"name": "Charlie", "age": -5.0, "email": ""},
	}
	
	validationErrors := validateUserData(users)
	if len(validationErrors) > 0 {
		fmt.Printf("發現 %d 個驗證錯誤:\n", len(validationErrors))
		for _, err := range validationErrors {
			fmt.Printf("  - %v\n", err)
		}
	}
	
	// 3. 重試模式
	fmt.Println("\n3. 重試模式:")
	attempts := 0
	unstableOp := func() error {
		attempts++
		if attempts < 3 {
			return fmt.Errorf("操作失敗 (嘗試 %d)", attempts)
		}
		return nil
	}
	
	config := RetryConfig{
		MaxAttempts: 5,
		Delay:       100 * time.Millisecond,
		BackoffRate: 2.0,
	}
	
	err := retryWithBackoff(unstableOp, config)
	if err != nil {
		fmt.Printf("重試失敗: %v\n", err)
	}
	
	// 4. 帶超時的操作
	fmt.Println("\n4. 帶超時的操作:")
	ctx := context.Background()
	err = operationWithTimeout(ctx, 1*time.Second)
	if err != nil {
		fmt.Printf("操作錯誤: %v\n", err)
	}
	
	// 5. 錯誤恢復模式
	fmt.Println("\n5. 錯誤恢復模式:")
	testData := []string{"normal", "error", "panic"}
	for _, data := range testData {
		result, err := safeOperation(data)
		if err != nil {
			fmt.Printf("數據 '%s': 錯誤 - %v\n", data, err)
		} else {
			fmt.Printf("數據 '%s': %s\n", data, result)
		}
	}
	
	// 6. 錯誤上下文包裝
	fmt.Println("\n6. 錯誤上下文包裝:")
	opCtx := &OperationContext{
		UserID:    "user123",
		Operation: "UpdateProfile",
		RequestID: "req-456",
	}
	
	testOperations := []string{"valid_data", "", "db_error"}
	for _, data := range testOperations {
		err := businessOperation(opCtx, data)
		if err != nil {
			fmt.Printf("業務操作錯誤: %v\n", err)
		} else {
			fmt.Printf("操作成功: %s\n", data)
		}
	}
	
	// 7. 錯誤分類
	fmt.Println("\n7. 錯誤分類和處理:")
	classifiedErrors := []*ClassifiedError{
		{ErrorTypeValidation, "無效的輸入", "VALIDATION_001", errors.New("年齡不能為負數")},
		{ErrorTypeNetwork, "網路連接失敗", "NETWORK_001", errors.New("連接超時")},
		{ErrorTypeDatabase, "數據庫錯誤", "DB_001", errors.New("查詢失敗")},
	}
	
	for _, err := range classifiedErrors {
		classifyAndHandle(err)
	}
}