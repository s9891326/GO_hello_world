package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 基本數據結構
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
	InStock     bool    `json:"in_stock"`
	Tags        []string `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// 演示基本 JSON 編碼
func demonstrateBasicMarshal() {
	fmt.Println("=== 基本 JSON 編碼 ===")
	
	// 1. 基本類型編碼
	fmt.Println("\n--- 基本類型編碼 ---")
	
	str := "Hello, JSON!"
	jsonStr, _ := json.Marshal(str)
	fmt.Printf("字符串: %s\n", jsonStr)
	
	num := 42
	jsonNum, _ := json.Marshal(num)
	fmt.Printf("數字: %s\n", jsonNum)
	
	flag := true
	jsonFlag, _ := json.Marshal(flag)
	fmt.Printf("布爾值: %s\n", jsonFlag)
	
	// 2. 數組編碼
	fmt.Println("\n--- 數組編碼 ---")
	arr := []int{1, 2, 3, 4, 5}
	jsonArr, _ := json.Marshal(arr)
	fmt.Printf("數組: %s\n", jsonArr)
	
	// 3. 映射編碼
	fmt.Println("\n--- 映射編碼 ---")
	m := map[string]interface{}{
		"name":    "Alice",
		"age":     25,
		"active":  true,
		"scores":  []int{95, 87, 92},
	}
	jsonMap, _ := json.Marshal(m)
	fmt.Printf("映射: %s\n", jsonMap)
	
	// 4. 結構體編碼
	fmt.Println("\n--- 結構體編碼 ---")
	user := User{
		ID:       1,
		Name:     "Alice Johnson",
		Email:    "alice@example.com",
		Age:      25,
		IsActive: true,
	}
	
	jsonUser, _ := json.Marshal(user)
	fmt.Printf("用戶結構體: %s\n", jsonUser)
	
	// 5. 漂亮格式化
	fmt.Println("\n--- 格式化 JSON ---")
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("格式化用戶:\n%s\n", prettyJSON)
}

// 演示基本 JSON 解碼
func demonstrateBasicUnmarshal() {
	fmt.Println("\n=== 基本 JSON 解碼 ===")
	
	// 1. 基本類型解碼
	fmt.Println("\n--- 基本類型解碼 ---")
	
	var str string
	json.Unmarshal([]byte(`"Hello, JSON!"`), &str)
	fmt.Printf("解碼字符串: %s\n", str)
	
	var num int
	json.Unmarshal([]byte(`42`), &num)
	fmt.Printf("解碼數字: %d\n", num)
	
	var arr []int
	json.Unmarshal([]byte(`[1,2,3,4,5]`), &arr)
	fmt.Printf("解碼數組: %v\n", arr)
	
	// 2. 結構體解碼
	fmt.Println("\n--- 結構體解碼 ---")
	jsonData := `{
		"id": 1,
		"name": "Bob Smith",
		"email": "bob@example.com",
		"age": 30,
		"is_active": true
	}`
	
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
	} else {
		fmt.Printf("解碼用戶: %+v\n", user)
	}
	
	// 3. 部分字段解碼
	fmt.Println("\n--- 部分字段解碼 ---")
	partialJSON := `{"id": 2, "name": "Charlie", "extra_field": "ignored"}`
	var partialUser User
	json.Unmarshal([]byte(partialJSON), &partialUser)
	fmt.Printf("部分解碼: %+v\n", partialUser)
}

// 演示 JSON 標籤的使用
func demonstrateJSONTags() {
	fmt.Println("\n=== JSON 標籤演示 ===")
	
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Description: "High-performance laptop",
		InStock:     true,
		Tags:        []string{"electronics", "computers"},
		CreatedAt:   time.Now(),
	}
	
	// 完整產品信息
	fmt.Println("\n--- 完整產品信息 ---")
	fullJSON, _ := json.MarshalIndent(product, "", "  ")
	fmt.Printf("完整產品:\n%s\n", fullJSON)
	
	// 空字段測試
	fmt.Println("\n--- 空字段測試 ---")
	emptyProduct := Product{
		ID:      2,
		Name:    "Empty Product",
		Price:   0,
		InStock: false,
		// Description 和 Tags 為空，應該被 omitempty 省略
	}
	
	emptyJSON, _ := json.MarshalIndent(emptyProduct, "", "  ")
	fmt.Printf("空字段產品:\n%s\n", emptyJSON)
}

// 演示處理動態 JSON
func demonstrateDynamicJSON() {
	fmt.Println("\n=== 動態 JSON 處理 ===")
	
	// 使用 interface{} 處理未知結構
	fmt.Println("\n--- 處理未知結構 ---")
	jsonStr := `{
		"name": "Alice",
		"age": 25,
		"scores": [95, 87, 92],
		"address": {
			"city": "New York",
			"zipcode": "10001"
		},
		"active": true
	}`
	
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}
	
	// 類型斷言訪問數據
	m := data.(map[string]interface{})
	
	name := m["name"].(string)
	age := m["age"].(float64) // JSON 數字默認為 float64
	scores := m["scores"].([]interface{})
	address := m["address"].(map[string]interface{})
	active := m["active"].(bool)
	
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年齡: %.0f\n", age)
	fmt.Printf("成績: %v\n", scores)
	fmt.Printf("城市: %s\n", address["city"])
	fmt.Printf("活躍: %t\n", active)
	
	// 使用 map[string]interface{} 直接處理
	fmt.Println("\n--- 使用 map 處理 ---")
	var mapData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &mapData)
	
	for key, value := range mapData {
		fmt.Printf("%s: %v (類型: %T)\n", key, value, value)
	}
}

// 演示嵌套結構體
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Company struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Employee struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address Address  `json:"address"`
	Company *Company `json:"company,omitempty"`
}

func demonstrateNestedStructs() {
	fmt.Println("\n=== 嵌套結構體處理 ===")
	
	employee := Employee{
		ID:   1,
		Name: "John Doe",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Company: &Company{
			Name: "Tech Corp",
			Address: Address{
				Street:  "456 Business Ave",
				City:    "San Francisco",
				Country: "USA",
			},
		},
	}
	
	// 編碼嵌套結構體
	jsonData, _ := json.MarshalIndent(employee, "", "  ")
	fmt.Printf("嵌套結構體:\n%s\n", jsonData)
	
	// 解碼嵌套結構體
	var decodedEmployee Employee
	json.Unmarshal(jsonData, &decodedEmployee)
	fmt.Printf("解碼結果: %+v\n", decodedEmployee)
	fmt.Printf("公司地址: %+v\n", decodedEmployee.Company.Address)
}

// 演示 JSON 數組處理
func demonstrateJSONArrays() {
	fmt.Println("\n=== JSON 數組處理 ===")
	
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25, IsActive: true},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30, IsActive: false},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35, IsActive: true},
	}
	
	// 編碼用戶數組
	jsonData, _ := json.MarshalIndent(users, "", "  ")
	fmt.Printf("用戶數組:\n%s\n", jsonData)
	
	// 解碼用戶數組
	var decodedUsers []User
	json.Unmarshal(jsonData, &decodedUsers)
	
	fmt.Println("解碼的用戶:")
	for i, user := range decodedUsers {
		fmt.Printf("  %d. %s (%s) - 活躍: %t\n", i+1, user.Name, user.Email, user.IsActive)
	}
}

// 演示錯誤處理
func demonstrateErrorHandling() {
	fmt.Println("\n=== JSON 錯誤處理 ===")
	
	// 1. 語法錯誤
	fmt.Println("\n--- 語法錯誤 ---")
	invalidJSON := `{"name": "Alice", "age": 25,}` // 多餘的逗號
	
	var user User
	err := json.Unmarshal([]byte(invalidJSON), &user)
	if err != nil {
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("JSON 語法錯誤在位置 %d: %v\n", syntaxErr.Offset, syntaxErr)
		} else {
			fmt.Printf("其他錯誤: %v\n", err)
		}
	}
	
	// 2. 類型錯誤
	fmt.Println("\n--- 類型錯誤 ---")
	typeErrorJSON := `{"id": "not_a_number", "name": "Alice"}`
	
	err = json.Unmarshal([]byte(typeErrorJSON), &user)
	if err != nil {
		if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
			fmt.Printf("類型錯誤: 無法將 %s 轉換為 %s (字段: %s)\n", 
				typeErr.Value, typeErr.Type, typeErr.Field)
		} else {
			fmt.Printf("其他錯誤: %v\n", err)
		}
	}
	
	// 3. 字段缺失
	fmt.Println("\n--- 字段缺失處理 ---")
	partialJSON := `{"name": "Bob"}` // 缺少其他字段
	
	var partialUser User
	err = json.Unmarshal([]byte(partialJSON), &partialUser)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
	} else {
		fmt.Printf("部分數據用戶: %+v\n", partialUser)
		fmt.Println("注意: 缺失的字段使用默認值")
	}
}

// 演示 JSON 流處理
func demonstrateJSONStreaming() {
	fmt.Println("\n=== JSON 流處理 ===")
	
	// 多個 JSON 對象的字符串
	jsonStream := `{"name": "Alice", "age": 25}
{"name": "Bob", "age": 30}
{"name": "Charlie", "age": 35}`
	
	fmt.Println("--- 流式解碼 ---")
	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	
	var count int
	for decoder.More() {
		var user map[string]interface{}
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Printf("解碼錯誤: %v\n", err)
			break
		}
		
		count++
		fmt.Printf("用戶 %d: %v\n", count, user)
	}
	
	// 流式編碼
	fmt.Println("\n--- 流式編碼 ---")
	var buf strings.Builder
	encoder := json.NewEncoder(&buf)
	
	users := []map[string]interface{}{
		{"name": "David", "age": 28},
		{"name": "Eve", "age": 32},
		{"name": "Frank", "age": 45},
	}
	
	for _, user := range users {
		err := encoder.Encode(user)
		if err != nil {
			fmt.Printf("編碼錯誤: %v\n", err)
			continue
		}
	}
	
	fmt.Printf("流式編碼結果:\n%s", buf.String())
}

// 主函數
func main() {
	fmt.Println("===== Go JSON 處理示例 =====")
	
	demonstrateBasicMarshal()
	demonstrateBasicUnmarshal()
	demonstrateJSONTags()
	demonstrateDynamicJSON()
	demonstrateNestedStructs()
	demonstrateJSONArrays()
	demonstrateErrorHandling()
	demonstrateJSONStreaming()
	
	fmt.Println("\n===== 示例完成 =====")
}