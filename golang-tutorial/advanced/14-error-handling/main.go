package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 基本錯誤處理示例
func basicErrorHandling() {
	fmt.Println("=== 基本錯誤處理 ===")
	
	// 使用 errors.New 創建錯誤
	err := errors.New("這是一個基本錯誤")
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	}
	
	// 使用 fmt.Errorf 創建格式化錯誤
	user := "Alice"
	age := -5
	err = fmt.Errorf("用戶 %s 的年齡無效: %d", user, age)
	if err != nil {
		fmt.Printf("格式化錯誤: %v\n", err)
	}
}

// 函數返回錯誤示例
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不能為零")
	}
	return a / b, nil
}

func parseAndCalculate(str1, str2 string) (float64, error) {
	num1, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		return 0, fmt.Errorf("解析第一個數字失敗: %w", err)
	}
	
	num2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		return 0, fmt.Errorf("解析第二個數字失敗: %w", err)
	}
	
	result, err := divide(num1, num2)
	if err != nil {
		return 0, fmt.Errorf("計算失敗: %w", err)
	}
	
	return result, nil
}

// 自定義錯誤類型
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("驗證錯誤 - 字段: %s, 值: %v, 訊息: %s", e.Field, e.Value, e.Message)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func validateUser(user User) error {
	if user.Name == "" {
		return &ValidationError{
			Field:   "name",
			Value:   user.Name,
			Message: "姓名不能為空",
		}
	}
	
	if user.Age < 0 || user.Age > 120 {
		return &ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "年齡必須在 0-120 之間",
		}
	}
	
	if user.Email == "" {
		return &ValidationError{
			Field:   "email",
			Value:   user.Email,
			Message: "郵箱不能為空",
		}
	}
	
	return nil
}

// 錯誤類型檢查
func handleValidationError(err error) {
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("發現驗證錯誤: %s\n", validationErr.Error())
		fmt.Printf("  字段: %s\n", validationErr.Field)
		fmt.Printf("  值: %v\n", validationErr.Value)
	} else {
		fmt.Printf("其他錯誤: %v\n", err)
	}
}

// 錯誤包裝和解包
func demonstrateErrorWrapping() {
	fmt.Println("\n=== 錯誤包裝和解包 ===")
	
	// 創建一層層包裝的錯誤
	originalErr := errors.New("原始錯誤")
	wrappedErr := fmt.Errorf("業務邏輯錯誤: %w", originalErr)
	finalErr := fmt.Errorf("API 調用失敗: %w", wrappedErr)
	
	fmt.Printf("最終錯誤: %v\n", finalErr)
	
	// 檢查是否包含特定錯誤
	if errors.Is(finalErr, originalErr) {
		fmt.Println("錯誤鏈中包含原始錯誤")
	}
	
	// 解包錯誤
	unwrapped := errors.Unwrap(finalErr)
	fmt.Printf("解包一層: %v\n", unwrapped)
}

// 重試機制示例
func unstableOperation() error {
	// 模擬不穩定的操作
	if time.Now().UnixNano()%3 == 0 {
		return nil // 成功
	}
	return errors.New("操作失敗")
}

func retryOperation(maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := unstableOperation()
		if err == nil {
			fmt.Printf("操作在第 %d 次嘗試後成功\n", i+1)
			return nil
		}
		
		if i < maxRetries-1 {
			fmt.Printf("第 %d 次嘗試失敗，重試中...\n", i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	return fmt.Errorf("操作在 %d 次重試後仍然失敗", maxRetries)
}

func main() {
	// 基本錯誤處理
	basicErrorHandling()
	
	// 函數錯誤返回
	fmt.Println("\n=== 函數錯誤返回 ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	} else {
		fmt.Printf("結果: %.2f\n", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	}
	
	// 錯誤鏈
	fmt.Println("\n=== 錯誤鏈 ===")
	result, err = parseAndCalculate("10", "0")
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	}
	
	result, err = parseAndCalculate("abc", "2")
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	}
	
	// 自定義錯誤類型
	fmt.Println("\n=== 自定義錯誤類型 ===")
	users := []User{
		{"Alice", "alice@example.com", 25},
		{"", "bob@example.com", 30},
		{"Charlie", "", 35},
		{"David", "david@example.com", -5},
	}
	
	for _, user := range users {
		err := validateUser(user)
		if err != nil {
			handleValidationError(err)
		} else {
			fmt.Printf("用戶 %s 驗證通過\n", user.Name)
		}
	}
	
	// 錯誤包裝
	demonstrateErrorWrapping()
	
	// 重試機制
	fmt.Println("\n=== 重試機制 ===")
	err = retryOperation(3)
	if err != nil {
		fmt.Printf("最終錯誤: %v\n", err)
	}
}