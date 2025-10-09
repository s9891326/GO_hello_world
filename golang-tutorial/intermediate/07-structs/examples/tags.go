package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 帶有 JSON 標籤的用戶結構體
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`                    // 忽略此字段
	Age      int    `json:"age,omitempty"`        // 零值時忽略
	IsActive bool   `json:"is_active"`
	Profile  struct {
		Bio     string `json:"bio"`
		Website string `json:"website,omitempty"`
		Avatar  string `json:"avatar,omitempty"`
	} `json:"profile"`
}

// 帶有多種標籤的產品結構體
type ProductWithTags struct {
	ID          int     `json:"id" db:"product_id" validate:"required"`
	Name        string  `json:"name" db:"product_name" validate:"required,min=1,max=100"`
	Description string  `json:"description,omitempty" db:"description"`
	Price       float64 `json:"price" db:"price" validate:"gt=0"`
	Category    string  `json:"category" db:"category" validate:"required"`
	InStock     bool    `json:"in_stock" db:"in_stock"`
	Tags        []string `json:"tags,omitempty" db:"tags"`
}

// 數據庫映射結構體
type DBUser struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email_address" json:"email"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func demonstrateStructTags() {
	fmt.Println("\n--- 結構體標籤演示 ---")
	
	// JSON 序列化和反序列化
	demonstrateJSONTags()
	
	// 自定義標籤讀取
	demonstrateCustomTags()
	
	// 標籤驗證示例
	demonstrateTagValidation()
}

func demonstrateJSONTags() {
	fmt.Println("\n🏷️ JSON 標籤示例:")
	
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123", // 這個字段會被忽略
		Age:      0,           // 零值，會被 omitempty 忽略
		IsActive: true,
	}
	user.Profile.Bio = "Passionate software developer"
	user.Profile.Website = "https://johndoe.dev"
	
	// 序列化為 JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("❌ JSON 序列化錯誤: %v\n", err)
		return
	}
	
	fmt.Println("📤 JSON 序列化輸出:")
	fmt.Println(string(jsonData))
	
	// 從 JSON 反序列化
	jsonStr := `{
		"id": 2,
		"name": "Jane Smith",
		"email": "jane@example.com",
		"age": 28,
		"is_active": true,
		"profile": {
			"bio": "Product Manager with 5 years experience",
			"website": "https://janesmith.com",
			"avatar": "https://example.com/avatar.jpg"
		}
	}`
	
	var newUser User
	err = json.Unmarshal([]byte(jsonStr), &newUser)
	if err != nil {
		fmt.Printf("❌ JSON 反序列化錯誤: %v\n", err)
		return
	}
	
	fmt.Println("\n📥 JSON 反序列化結果:")
	fmt.Printf("   用戶: %s (%s)\n", newUser.Name, newUser.Email)
	fmt.Printf("   年齡: %d\n", newUser.Age)
	fmt.Printf("   簡介: %s\n", newUser.Profile.Bio)
	fmt.Printf("   網站: %s\n", newUser.Profile.Website)
}

func demonstrateCustomTags() {
	fmt.Println("\n🔍 自定義標籤讀取:")
	
	product := ProductWithTags{}
	productType := reflect.TypeOf(product)
	
	fmt.Println("📋 結構體字段標籤分析:")
	for i := 0; i < productType.NumField(); i++ {
		field := productType.Field(i)
		
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		validateTag := field.Tag.Get("validate")
		
		fmt.Printf("\n   字段: %s (%s)\n", field.Name, field.Type)
		if jsonTag != "" {
			fmt.Printf("      JSON: %s\n", jsonTag)
		}
		if dbTag != "" {
			fmt.Printf("      DB: %s\n", dbTag)
		}
		if validateTag != "" {
			fmt.Printf("      驗證: %s\n", validateTag)
		}
	}
}

func demonstrateTagValidation() {
	fmt.Println("\n✅ 標籤驗證示例:")
	
	// 模擬驗證邏輯
	products := []ProductWithTags{
		{
			ID:          1,
			Name:        "筆記本電腦",
			Description: "高性能筆記本電腦",
			Price:       25000.00,
			Category:    "電子產品",
			InStock:     true,
			Tags:        []string{"電腦", "辦公", "便攜"},
		},
		{
			ID:       2,
			Name:     "", // 無效：名稱為空
			Price:    -100, // 無效：價格為負
			Category: "電子產品",
			InStock:  false,
		},
	}
	
	for i, product := range products {
		fmt.Printf("\n   產品 %d 驗證結果:\n", i+1)
		errors := validateProduct(product)
		if len(errors) == 0 {
			fmt.Println("      ✅ 驗證通過")
			fmt.Printf("      產品: %s, 價格: %.2f\n", product.Name, product.Price)
		} else {
			fmt.Println("      ❌ 驗證失敗:")
			for _, err := range errors {
				fmt.Printf("         %s\n", err)
			}
		}
	}
}

// 簡單的驗證函數（模擬根據標籤進行驗證）
func validateProduct(product ProductWithTags) []string {
	var errors []string
	
	// 根據 validate 標籤進行驗證
	if product.ID == 0 {
		errors = append(errors, "ID 是必需的")
	}
	
	if product.Name == "" {
		errors = append(errors, "產品名稱是必需的")
	}
	
	if len(product.Name) > 100 {
		errors = append(errors, "產品名稱不能超過100個字符")
	}
	
	if product.Price <= 0 {
		errors = append(errors, "價格必須大於0")
	}
	
	if product.Category == "" {
		errors = append(errors, "分類是必需的")
	}
	
	return errors
}

// 展示標籤在 ORM 中的應用
func demonstrateORMTags() {
	fmt.Println("\n🗄️ ORM 標籤應用示例:")
	
	user := DBUser{
		ID:        1,
		Username:  "johndoe",
		Email:     "john@example.com",
		CreatedAt: "2024-01-15 10:30:00",
		UpdatedAt: "2024-01-15 10:30:00",
	}
	
	// 模擬生成 SQL 查詢
	fmt.Println("   根據標籤生成的 SQL:")
	fmt.Printf("   SELECT id, username, email_address, created_at, updated_at\n")
	fmt.Printf("   FROM users WHERE id = %d;\n", user.ID)
	
	fmt.Printf("   INSERT INTO users (username, email_address, created_at, updated_at)\n")
	fmt.Printf("   VALUES ('%s', '%s', '%s', '%s');\n", 
		user.Username, user.Email, user.CreatedAt, user.UpdatedAt)
}