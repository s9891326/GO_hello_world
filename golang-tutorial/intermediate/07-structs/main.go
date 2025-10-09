package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

// 基本結構體定義
type Person struct {
	Name string
	Age  int
	City string
}

// 帶有多種類型字段的結構體
type Product struct {
	ID          int
	Name        string
	Price       float64
	InStock     bool
	Categories  []string
	CreatedAt   time.Time
}

// 地址結構體
type Address struct {
	Street   string
	City     string
	ZipCode  string
	Country  string
}

// 聯繫信息結構體
type Contact struct {
	Email string
	Phone string
}

// 員工結構體（組合）
type Employee struct {
	ID       int
	Person   Person
	Address  Address
	Contact  Contact
	Salary   float64
	Position string
}

// 經理結構體（嵌入）
type Manager struct {
	Person
	Address
	Contact
	EmployeeID int
	Department string
	TeamSize   int
}

func main() {
	fmt.Println("=== Go 結構體示例 ===")
	
	// 1. 基本結構體操作
	demonstrateBasicStruct()
	
	// 2. 結構體方法
	demonstrateMethods()
	
	// 3. 結構體組合和嵌入
	demonstrateComposition()
	
	// 4. 結構體標籤
	demonstrateStructTags()
	
	// 5. 構造函數模式
	demonstrateConstructors()
	
	// 6. 內存對齊和性能
	demonstrateMemoryAlignment()
	
	// 7. 實際應用示例
	demonstrateRealWorldExamples()
}

// 添加缺少的方法定義
func (p Person) GetFullInfo() string {
	return fmt.Sprintf("%s, %d歲, 住在%s", p.Name, p.Age, p.City)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) TryToAge() {
	p.Age++
}

func (p *Person) SetAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("年齡必須在 0-150 之間")
	}
	p.Age = age
	return nil
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("🎉 生日快樂！%s 現在 %d 歲了\n", p.Name, p.Age)
}

func (p *Person) MoveTo(city string) {
	oldCity := p.City
	p.City = city
	fmt.Printf("📍 %s 從 %s 搬到了 %s\n", p.Name, oldCity, city)
}

func (m Manager) GetManagementInfo() string {
	return fmt.Sprintf("經理 %s 管理 %s 部門的 %d 個人", 
		m.Name, m.Department, m.TeamSize)
}

func (e Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("員工ID: %d, %s, 職位: %s, 薪資: %.2f", 
		e.ID, e.Person.GetFullInfo(), e.Position, e.Salary)
}

func demonstrateBasicStruct() {
	fmt.Println("\n--- 基本結構體操作 ---")
	
	// 零值初始化
	var p1 Person
	fmt.Printf("零值初始化: %+v\n", p1)
	
	// 字段名初始化
	p2 := Person{
		Name: "Alice",
		Age:  25,
		City: "台北",
	}
	fmt.Printf("字段名初始化: %+v\n", p2)
	
	// 位置初始化
	p3 := Person{"Bob", 30, "高雄"}
	fmt.Printf("位置初始化: %+v\n", p3)
	
	// 部分初始化
	p4 := Person{
		Name: "Charlie",
		Age:  35,
	}
	fmt.Printf("部分初始化: %+v\n", p4)
	
	// 字段訪問和修改
	fmt.Printf("修改前 - 姓名: %s, 年齡: %d\n", p2.Name, p2.Age)
	p2.Age = 26
	p2.City = "台南"
	fmt.Printf("修改後: %+v\n", p2)
	
	// 通過指針訪問
	ptr := &p2
	ptr.Name = "Alice Chen"
	fmt.Printf("通過指針修改: %+v\n", p2)
}

func demonstrateMethods() {
	fmt.Println("\n--- 結構體方法演示 ---")
	
	person := Person{
		Name: "Frank",
		Age:  25,
		City: "新竹",
	}
	
	// 調用值接收者方法
	fmt.Println("📝 個人信息:", person.GetFullInfo())
	fmt.Printf("🔞 是否成年: %t\n", person.IsAdult())
	
	// 值接收者方法不會修改原結構體
	fmt.Printf("🎂 嘗試增加年齡前: %d\n", person.Age)
	person.TryToAge()
	fmt.Printf("🎂 嘗試增加年齡後: %d (原結構體未改變)\n", person.Age)
	
	// 調用指針接收者方法
	err := person.SetAge(26)
	if err != nil {
		fmt.Printf("❌ 設置年齡失敗: %v\n", err)
	} else {
		fmt.Printf("✅ 設置年齡成功: %d\n", person.Age)
	}
	
	person.HaveBirthday()
	person.MoveTo("桃園")
}

func demonstrateComposition() {
	fmt.Println("\n--- 結構體組合和嵌入演示 ---")
	
	// 創建組合結構體（Employee）
	employee := Employee{
		ID: 1001,
		Person: Person{
			Name: "Grace",
			Age:  28,
			City: "台北",
		},
		Address: Address{
			Street:  "信義路100號",
			City:    "台北市",
			ZipCode: "110",
			Country: "台灣",
		},
		Contact: Contact{
			Email: "grace@company.com",
			Phone: "02-1234-5678",
		},
		Salary:   60000,
		Position: "軟體工程師",
	}
	
	fmt.Println("👩‍💼 員工信息:")
	fmt.Println("  ", employee.GetEmployeeInfo())
	
	// 創建嵌入結構體（Manager）
	manager := Manager{
		Person: Person{
			Name: "Henry",
			Age:  35,
			City: "台北",
		},
		Address: Address{
			Street:  "敦化南路200號",
			City:    "台北市",
			ZipCode: "106",
			Country: "台灣",
		},
		Contact: Contact{
			Email: "henry@company.com",
			Phone: "02-5678-9012",
		},
		EmployeeID: 2001,
		Department: "工程部",
		TeamSize:   5,
	}
	
	fmt.Println("\n👨‍💼 經理信息:")
	fmt.Printf("   姓名: %s (來自 Person)\n", manager.Name)
	fmt.Printf("   年齡: %d (來自 Person)\n", manager.Age)
	fmt.Printf("   郵箱: %s (來自 Contact)\n", manager.Email)
	fmt.Println("   管理信息:", manager.GetManagementInfo())
}

func demonstrateStructTags() {
	fmt.Println("\n--- 結構體標籤演示 ---")
	
	type User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
		Age      int    `json:"age,omitempty"`
		IsActive bool   `json:"is_active"`
	}
	
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123",
		Age:      0,
		IsActive: true,
	}
	
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("❌ JSON 序列化錯誤: %v\n", err)
		return
	}
	
	fmt.Println("📤 JSON 序列化輸出:")
	fmt.Println(string(jsonData))
}

func demonstrateConstructors() {
	fmt.Println("\n--- 構造函數模式演示 ---")
	
	// 簡單構造函數
	person1 := NewPerson("Alice", 25, "台北")
	fmt.Printf("🏗️ 簡單構造: %+v\n", person1)
	
	// 帶驗證的構造函數
	person2, err := NewPersonValidated("Bob", 30, "高雄")
	if err != nil {
		fmt.Printf("❌ 創建失敗: %v\n", err)
	} else {
		fmt.Printf("✅ 驗證構造: %+v\n", person2)
	}
}

// 簡單構造函數
func NewPerson(name string, age int, city string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// 帶驗證的構造函數
func NewPersonValidated(name string, age int, city string) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("姓名不能為空")
	}
	if age < 0 || age > 150 {
		return nil, fmt.Errorf("年齡必須在 0-150 之間")
	}
	if city == "" {
		return nil, fmt.Errorf("城市不能為空")
	}
	
	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}, nil
}

func demonstrateMemoryAlignment() {
	fmt.Println("\n--- 內存對齊演示 ---")
	
	type BadStruct struct {
		a bool
		b int64
		c bool
		d int64
	}
	
	type GoodStruct struct {
		b int64
		d int64
		a bool
		c bool
	}
	
	fmt.Printf("💾 BadStruct 大小: %d bytes\n", unsafe.Sizeof(BadStruct{}))
	fmt.Printf("💾 GoodStruct 大小: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
}

func demonstrateRealWorldExamples() {
	fmt.Println("\n--- 實際應用示例 ---")
	
	// 零值友好的計數器
	type Counter struct {
		value int
		mutex sync.Mutex
	}
	
	counter := Counter{} // 零值可以直接使用
	fmt.Printf("🔢 計數器初始值: %d\n", counter.value)
	
	// 銀行帳戶示例
	type BankAccount struct {
		AccountNumber string
		Balance       float64
		Owner         Person
		IsActive      bool
	}
	
	account := BankAccount{
		AccountNumber: "ACC-2024-001",
		Balance:       1000.00,
		Owner:         *NewPerson("David", 35, "台中"),
		IsActive:      true,
	}
	
	fmt.Printf("🏦 銀行帳戶: %s, 餘額: %.2f, 持有人: %s\n", 
		account.AccountNumber, account.Balance, account.Owner.Name)
}