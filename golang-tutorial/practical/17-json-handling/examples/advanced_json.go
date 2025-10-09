package main

import (
	"encoding/json"
	"fmt"
	"time"
	"bytes"
	"strings"
)

// 自定義時間格式
type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05"

// 實現 json.Marshaler 接口
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format(customTimeFormat))
}

// 實現 json.Unmarshaler 接口
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}
	
	t, err := time.Parse(customTimeFormat, timeStr)
	if err != nil {
		return err
	}
	
	ct.Time = t
	return nil
}

// 自定義 Email 類型
type Email string

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(e))
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	
	// 簡單的郵箱驗證
	if !strings.Contains(s, "@") {
		return fmt.Errorf("無效的郵箱格式: %s", s)
	}
	
	*e = Email(s)
	return nil
}

// 使用自定義類型的結構體
type Event struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	StartTime CustomTime `json:"start_time"`
	EndTime   CustomTime `json:"end_time"`
	Organizer Email      `json:"organizer"`
}

// 演示自定義 JSON 序列化
func demonstrateCustomSerialization() {
	fmt.Println("=== 自定義 JSON 序列化 ===")
	
	event := Event{
		ID:        1,
		Title:     "Go 語言學習會議",
		StartTime: CustomTime{time.Now()},
		EndTime:   CustomTime{time.Now().Add(2 * time.Hour)},
		Organizer: Email("organizer@example.com"),
	}
	
	// 編碼
	jsonData, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
		return
	}
	
	fmt.Printf("自定義序列化事件:\n%s\n", jsonData)
	
	// 解碼
	var decodedEvent Event
	err = json.Unmarshal(jsonData, &decodedEvent)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}
	
	fmt.Printf("解碼事件: %+v\n", decodedEvent)
	fmt.Printf("開始時間: %s\n", decodedEvent.StartTime.Time.Format("2006-01-02 15:04:05"))
}

// 處理任意 JSON 結構
type FlexibleData struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func demonstrateFlexibleJSON() {
	fmt.Println("\n=== 靈活 JSON 處理 ===")
	
	// 不同類型的數據
	jsonSamples := []string{
		`{"type": "user", "data": {"name": "Alice", "age": 25}}`,
		`{"type": "product", "data": {"name": "Laptop", "price": 999.99, "inStock": true}}`,
		`{"type": "order", "data": {"id": 123, "items": ["item1", "item2"], "total": 150.50}}`,
	}
	
	for i, jsonStr := range jsonSamples {
		fmt.Printf("\n--- 處理樣本 %d ---\n", i+1)
		
		var flexData FlexibleData
		err := json.Unmarshal([]byte(jsonStr), &flexData)
		if err != nil {
			fmt.Printf("解碼錯誤: %v\n", err)
			continue
		}
		
		fmt.Printf("類型: %s\n", flexData.Type)
		fmt.Println("數據:")
		
		for key, value := range flexData.Data {
			fmt.Printf("  %s: %v (類型: %T)\n", key, value, value)
		}
		
		// 根據類型進行特定處理
		switch flexData.Type {
		case "user":
			if name, ok := flexData.Data["name"].(string); ok {
				fmt.Printf("用戶名: %s\n", name)
			}
		case "product":
			if price, ok := flexData.Data["price"].(float64); ok {
				fmt.Printf("產品價格: %.2f\n", price)
			}
		case "order":
			if id, ok := flexData.Data["id"].(float64); ok {
				fmt.Printf("訂單 ID: %.0f\n", id)
			}
		}
	}
}

// 使用 json.RawMessage 延遲解碼
type Response struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type ProductData struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func demonstrateRawMessage() {
	fmt.Println("\n=== json.RawMessage 延遲解碼 ===")
	
	responses := []string{
		`{"status": "success", "message": "user data", "data": {"name": "Alice", "email": "alice@example.com", "age": 25}}`,
		`{"status": "success", "message": "product data", "data": {"name": "Laptop", "price": 999.99}}`,
	}
	
	for i, respStr := range responses {
		fmt.Printf("\n--- 響應 %d ---\n", i+1)
		
		var resp Response
		err := json.Unmarshal([]byte(respStr), &resp)
		if err != nil {
			fmt.Printf("解碼響應錯誤: %v\n", err)
			continue
		}
		
		fmt.Printf("狀態: %s\n", resp.Status)
		fmt.Printf("消息: %s\n", resp.Message)
		fmt.Printf("原始數據: %s\n", resp.Data)
		
		// 根據消息類型解碼數據
		switch resp.Message {
		case "user data":
			var userData UserData
			if err := json.Unmarshal(resp.Data, &userData); err == nil {
				fmt.Printf("用戶數據: %+v\n", userData)
			}
		case "product data":
			var productData ProductData
			if err := json.Unmarshal(resp.Data, &productData); err == nil {
				fmt.Printf("產品數據: %+v\n", productData)
			}
		}
	}
}

// 處理大型 JSON 數據
func demonstrateLargeJSONProcessing() {
	fmt.Println("\n=== 大型 JSON 處理 ===")
	
	// 模擬大型 JSON 數據
	largeData := make([]map[string]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		largeData[i] = map[string]interface{}{
			"id":   i,
			"name": fmt.Sprintf("User_%d", i),
			"data": fmt.Sprintf("Data for user %d", i),
		}
	}
	
	fmt.Printf("處理 %d 條記錄\n", len(largeData))
	
	// 編碼大數據
	startTime := time.Now()
	jsonData, err := json.Marshal(largeData)
	if err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
		return
	}
	encodingTime := time.Since(startTime)
	
	fmt.Printf("編碼時間: %v\n", encodingTime)
	fmt.Printf("JSON 大小: %d 字節\n", len(jsonData))
	
	// 解碼大數據
	startTime = time.Now()
	var decodedData []map[string]interface{}
	err = json.Unmarshal(jsonData, &decodedData)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}
	decodingTime := time.Since(startTime)
	
	fmt.Printf("解碼時間: %v\n", decodingTime)
	fmt.Printf("解碼記錄數: %d\n", len(decodedData))
}

// JSON 數據轉換
type SourceStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TargetStruct struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
}

func demonstrateJSONConversion() {
	fmt.Println("\n=== JSON 數據轉換 ===")
	
	source := SourceStruct{
		ID:   1,
		Name: "Alice Johnson",
		Age:  25,
	}
	
	fmt.Printf("源結構體: %+v\n", source)
	
	// 通過 JSON 進行結構體轉換
	jsonData, err := json.Marshal(source)
	if err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
		return
	}
	
	var target TargetStruct
	err = json.Unmarshal(jsonData, &target)
	if err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}
	
	// 手動處理不匹配的字段
	target.FullName = source.Name
	target.Status = "active"
	
	fmt.Printf("目標結構體: %+v\n", target)
	
	// 輸出轉換後的 JSON
	targetJSON, _ := json.MarshalIndent(target, "", "  ")
	fmt.Printf("轉換後 JSON:\n%s\n", targetJSON)
}

// JSON 校驗和美化
func demonstrateJSONValidation() {
	fmt.Println("\n=== JSON 校驗和美化 ===")
	
	jsonSamples := []string{
		`{"name":"Alice","age":25,"active":true}`,                    // 有效 JSON
		`{"name": "Bob", "age": 30,}`,                               // 無效 JSON (多餘逗號)
		`{"name":"Charlie","details":{"city":"NY","zip":"10001"}}`,  // 嵌套 JSON
	}
	
	for i, jsonStr := range jsonSamples {
		fmt.Printf("\n--- 樣本 %d ---\n", i+1)
		fmt.Printf("原始: %s\n", jsonStr)
		
		// 校驗 JSON
		var temp interface{}
		err := json.Unmarshal([]byte(jsonStr), &temp)
		if err != nil {
			fmt.Printf("無效 JSON: %v\n", err)
			continue
		}
		
		fmt.Println("JSON 校驗: 有效")
		
		// 美化 JSON
		var buf bytes.Buffer
		err = json.Indent(&buf, []byte(jsonStr), "", "  ")
		if err != nil {
			fmt.Printf("美化錯誤: %v\n", err)
			continue
		}
		
		fmt.Printf("美化後:\n%s\n", buf.String())
	}
}

func main() {
	fmt.Println("===== 高級 JSON 處理示例 =====")
	
	demonstrateCustomSerialization()
	demonstrateFlexibleJSON()
	demonstrateRawMessage()
	demonstrateLargeJSONProcessing()
	demonstrateJSONConversion()
	demonstrateJSONValidation()
	
	fmt.Println("\n===== 示例完成 =====")
}