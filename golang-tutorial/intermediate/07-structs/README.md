# 第七章：結構體

## 🎯 學習目標

- 理解結構體的概念和用途
- 掌握結構體的定義和初始化
- 學會為結構體添加方法
- 了解結構體的嵌入和組合
- 掌握結構體標籤的使用
- 學會結構體的最佳實踐

## 📦 結構體基礎

結構體（struct）是 Go 語言中用於創建自定義數據類型的重要工具，它將相關的數據組合在一起。

### 結構體定義

```go
// 基本結構體定義
type Person struct {
    Name string
    Age  int
    City string
}

// 帶有不同類型字段的結構體
type Product struct {
    ID          int
    Name        string
    Price       float64
    InStock     bool
    Categories  []string
    CreatedAt   time.Time
}
```

### 結構體初始化

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateStructInitialization() {
    // 方式 1: 零值初始化
    var p1 Person
    fmt.Printf("零值初始化: %+v\n", p1)
    
    // 方式 2: 字段名初始化
    p2 := Person{
        Name: "Alice",
        Age:  25,
        City: "台北",
    }
    fmt.Printf("字段名初始化: %+v\n", p2)
    
    // 方式 3: 位置初始化（不推薦）
    p3 := Person{"Bob", 30, "高雄"}
    fmt.Printf("位置初始化: %+v\n", p3)
    
    // 方式 4: 部分初始化
    p4 := Person{
        Name: "Charlie",
        Age:  35,
        // City 將使用零值 ""
    }
    fmt.Printf("部分初始化: %+v\n", p4)
    
    // 方式 5: 使用 new 函數
    p5 := new(Person)
    p5.Name = "David"
    p5.Age = 40
    fmt.Printf("new 函數: %+v\n", *p5)
}
```

### 結構體字段訪問

```go
func demonstrateFieldAccess() {
    person := Person{
        Name: "Emily",
        Age:  28,
        City: "台中",
    }
    
    // 讀取字段
    fmt.Printf("姓名: %s\n", person.Name)
    fmt.Printf("年齡: %d\n", person.Age)
    
    // 修改字段
    person.Age = 29
    person.City = "台南"
    fmt.Printf("修改後: %+v\n", person)
    
    // 通過指針訪問
    ptr := &person
    ptr.Name = "Emily Chen"  // Go 自動解引用
    fmt.Printf("通過指針修改: %+v\n", person)
}
```

## 🔧 結構體方法

Go 語言通過方法（method）為結構體添加行為。

### 值接收者方法

```go
// 值接收者方法
func (p Person) GetFullInfo() string {
    return fmt.Sprintf("%s, %d歲, 住在%s", p.Name, p.Age, p.City)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

// 值接收者方法不能修改結構體
func (p Person) TryToAge() {
    p.Age++  // 這不會影響原始結構體
}
```

### 指針接收者方法

```go
// 指針接收者方法可以修改結構體
func (p *Person) SetAge(age int) {
    if age >= 0 && age <= 150 {
        p.Age = age
    }
}

func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("生日快樂！%s 現在 %d 歲了\n", p.Name, p.Age)
}

func (p *Person) MoveTo(city string) {
    oldCity := p.City
    p.City = city
    fmt.Printf("%s 從 %s 搬到了 %s\n", p.Name, oldCity, city)
}
```

### 方法演示

```go
func demonstrateMethods() {
    person := Person{
        Name: "Frank",
        Age:  25,
        City: "新竹",
    }
    
    // 調用值接收者方法
    fmt.Println("個人信息:", person.GetFullInfo())
    fmt.Printf("是否成年: %t\n", person.IsAdult())
    
    // 值接收者方法不會修改原結構體
    person.TryToAge()
    fmt.Printf("嘗試增加年齡後: %d\n", person.Age)  // 仍然是 25
    
    // 調用指針接收者方法
    person.SetAge(26)
    fmt.Printf("設置年齡後: %d\n", person.Age)
    
    person.HaveBirthday()
    person.MoveTo("桃園")
}
```

## 🎭 結構體組合和嵌入

Go 語言通過組合而非繼承來實現代碼重用。

### 結構體組合

```go
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

// 員工結構體，包含其他結構體
type Employee struct {
    ID       int
    Person   Person    // 組合
    Address  Address   // 組合
    Contact  Contact   // 組合
    Salary   float64
    Position string
}

func (e Employee) GetContactInfo() string {
    return fmt.Sprintf("Email: %s, Phone: %s", e.Contact.Email, e.Contact.Phone)
}
```

### 結構體嵌入（匿名字段）

```go
// 使用嵌入的員工結構體
type Manager struct {
    Person           // 嵌入，可以直接訪問 Person 的字段
    Address          // 嵌入
    Contact          // 嵌入
    EmployeeID   int
    Department   string
    TeamSize     int
}

// 為嵌入結構體添加方法
func (m Manager) GetManagementInfo() string {
    return fmt.Sprintf("Manager %s manages %d people in %s", 
        m.Name, m.TeamSize, m.Department)
}

func demonstrateEmbedding() {
    manager := Manager{
        Person: Person{
            Name: "Grace",
            Age:  35,
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
        EmployeeID: 1001,
        Department: "工程部",
        TeamSize:   8,
    }
    
    // 直接訪問嵌入結構體的字段
    fmt.Printf("經理姓名: %s\n", manager.Name)  // 來自 Person
    fmt.Printf("經理年齡: %d\n", manager.Age)   // 來自 Person
    fmt.Printf("經理城市: %s\n", manager.City)  // 來自 Address.City（會有衝突）
    
    // 訪問嵌入結構體的方法
    fmt.Println("個人信息:", manager.GetFullInfo())  // 來自 Person
    fmt.Println("管理信息:", manager.GetManagementInfo())
}
```

### 處理字段名衝突

```go
type Student struct {
    Person
    Address
    StudentID string
    Grade     int
}

func demonstrateFieldConflicts() {
    student := Student{
        Person: Person{
            Name: "Henry",
            Age:  20,
            City: "學校城市",  // Person.City
        },
        Address: Address{
            Street:  "大學路200號",
            City:    "地址城市",  // Address.City
            ZipCode: "300",
            Country: "台灣",
        },
        StudentID: "S2024001",
        Grade:     3,
    }
    
    // 當有字段名衝突時，必須明確指定
    fmt.Printf("學生居住城市: %s\n", student.Person.City)
    fmt.Printf("學校地址城市: %s\n", student.Address.City)
    
    // 無衝突的字段可以直接訪問
    fmt.Printf("學生姓名: %s\n", student.Name)  // 來自 Person
    fmt.Printf("學生年齡: %d\n", student.Age)   // 來自 Person
}
```

## 🏷️ 結構體標籤

結構體標籤用於為字段提供元數據，常用於JSON序列化、數據庫映射等。

### JSON 標籤

```go
import (
    "encoding/json"
    "fmt"
)

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
    } `json:"profile"`
}

func demonstrateJSONTags() {
    user := User{
        ID:       1,
        Name:     "John Doe",
        Email:    "john@example.com",
        Password: "secret123",
        Age:      0,  // 零值，會被 omitempty 忽略
        IsActive: true,
    }
    user.Profile.Bio = "Software Developer"
    
    // 序列化為 JSON
    jsonData, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        fmt.Printf("JSON 序列化錯誤: %v\n", err)
        return
    }
    
    fmt.Println("JSON 輸出:")
    fmt.Println(string(jsonData))
    
    // 從 JSON 反序列化
    jsonStr := `{
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@example.com",
        "age": 28,
        "is_active": true,
        "profile": {
            "bio": "Product Manager",
            "website": "https://jane.example.com"
        }
    }`
    
    var newUser User
    err = json.Unmarshal([]byte(jsonStr), &newUser)
    if err != nil {
        fmt.Printf("JSON 反序列化錯誤: %v\n", err)
        return
    }
    
    fmt.Printf("反序列化結果: %+v\n", newUser)
}
```

### 自定義標籤

```go
// 數據庫映射標籤示例
type Product struct {
    ID          int     `db:"id" json:"id"`
    Name        string  `db:"product_name" json:"name" validate:"required"`
    Description string  `db:"description" json:"description,omitempty"`
    Price       float64 `db:"price" json:"price" validate:"gt=0"`
    Category    string  `db:"category" json:"category" validate:"required"`
    InStock     bool    `db:"in_stock" json:"in_stock"`
}

// 使用反射讀取標籤
import "reflect"

func demonstrateCustomTags() {
    product := Product{}
    productType := reflect.TypeOf(product)
    
    fmt.Println("結構體字段標籤:")
    for i := 0; i < productType.NumField(); i++ {
        field := productType.Field(i)
        
        dbTag := field.Tag.Get("db")
        jsonTag := field.Tag.Get("json")
        validateTag := field.Tag.Get("validate")
        
        fmt.Printf("字段: %s\n", field.Name)
        if dbTag != "" {
            fmt.Printf("  數據庫欄位: %s\n", dbTag)
        }
        if jsonTag != "" {
            fmt.Printf("  JSON 標籤: %s\n", jsonTag)
        }
        if validateTag != "" {
            fmt.Printf("  驗證規則: %s\n", validateTag)
        }
        fmt.Println()
    }
}
```

## 🏗️ 結構體構造函數

Go 沒有內建的構造函數，但可以使用工廠函數模式。

### 簡單構造函數

```go
// 構造函數模式
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
```

### 選項模式（Options Pattern）

```go
// 選項函數類型
type PersonOption func(*Person)

// 選項函數
func WithAge(age int) PersonOption {
    return func(p *Person) {
        p.Age = age
    }
}

func WithCity(city string) PersonOption {
    return func(p *Person) {
        p.City = city
    }
}

// 使用選項模式的構造函數
func NewPersonWithOptions(name string, options ...PersonOption) *Person {
    person := &Person{
        Name: name,
        Age:  0,    // 默認值
        City: "未知", // 默認值
    }
    
    // 應用所有選項
    for _, option := range options {
        option(person)
    }
    
    return person
}

func demonstrateConstructors() {
    // 簡單構造函數
    person1 := NewPerson("Alice", 25, "台北")
    fmt.Printf("簡單構造: %+v\n", person1)
    
    // 帶驗證的構造函數
    person2, err := NewPersonValidated("Bob", 30, "高雄")
    if err != nil {
        fmt.Printf("構造錯誤: %v\n", err)
    } else {
        fmt.Printf("驗證構造: %+v\n", person2)
    }
    
    // 選項模式
    person3 := NewPersonWithOptions("Charlie")
    fmt.Printf("僅姓名: %+v\n", person3)
    
    person4 := NewPersonWithOptions("David", 
        WithAge(35), 
        WithCity("台中"))
    fmt.Printf("選項模式: %+v\n", person4)
}
```

## 🔍 結構體比較和複製

### 結構體比較

```go
func demonstrateStructComparison() {
    person1 := Person{Name: "Alice", Age: 25, City: "台北"}
    person2 := Person{Name: "Alice", Age: 25, City: "台北"}
    person3 := Person{Name: "Bob", Age: 30, City: "高雄"}
    
    // 結構體可以直接比較（如果所有字段都可比較）
    fmt.Printf("person1 == person2: %t\n", person1 == person2)  // true
    fmt.Printf("person1 == person3: %t\n", person1 == person3)  // false
    
    // 包含不可比較字段的結構體無法比較
    type PersonWithSlice struct {
        Name     string
        Age      int
        Hobbies  []string  // slice 不可比較
    }
    
    // p1 := PersonWithSlice{Name: "Alice", Hobbies: []string{"reading"}}
    // p2 := PersonWithSlice{Name: "Alice", Hobbies: []string{"reading"}}
    // fmt.Println(p1 == p2)  // 編譯錯誤！
}
```

### 結構體複製

```go
func demonstrateStructCopy() {
    original := Person{Name: "Alice", Age: 25, City: "台北"}
    
    // 值複製（深複製）
    copied := original
    copied.Age = 30
    
    fmt.Printf("原始: %+v\n", original)  // Age 仍然是 25
    fmt.Printf("複製: %+v\n", copied)    // Age 是 30
    
    // 指針複製（淺複製）
    ptr1 := &original
    ptr2 := ptr1
    ptr2.Age = 35
    
    fmt.Printf("通過指針修改後: %+v\n", original)  // Age 變成 35
}
```

## 💡 結構體最佳實踐

### 1. 字段順序和對齊

```go
// 不好的字段順序（占用更多內存）
type BadStruct struct {
    a bool    // 1 byte
    b int64   // 8 bytes  
    c bool    // 1 byte
    d int64   // 8 bytes
}

// 好的字段順序（內存對齊）
type GoodStruct struct {
    b int64   // 8 bytes
    d int64   // 8 bytes
    a bool    // 1 byte
    c bool    // 1 byte
}

func demonstrateMemoryAlignment() {
    fmt.Printf("BadStruct 大小: %d bytes\n", unsafe.Sizeof(BadStruct{}))
    fmt.Printf("GoodStruct 大小: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
}
```

### 2. 零值友好設計

```go
// 零值友好的結構體
type Counter struct {
    value int
    mutex sync.Mutex
}

func (c *Counter) Increment() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.value
}

// 可以直接使用零值
func demonstrateZeroValueFriendly() {
    var counter Counter  // 零值可以直接使用
    counter.Increment()
    fmt.Printf("Counter value: %d\n", counter.Value())
}
```

### 3. 接口設計

```go
// 定義行為接口
type Speaker interface {
    Speak() string
}

type Mover interface {
    Move(destination string)
}

// 實現接口
func (p Person) Speak() string {
    return fmt.Sprintf("我是 %s", p.Name)
}

func (p *Person) Move(destination string) {
    p.City = destination
}

func demonstrateInterfaces() {
    person := NewPerson("Eve", 28, "台南")
    
    // 作為 Speaker 使用
    var speaker Speaker = *person
    fmt.Println(speaker.Speak())
    
    // 作為 Mover 使用
    var mover Mover = person
    mover.Move("嘉義")
    fmt.Printf("移動後: %+v\n", person)
}
```

## 🎯 本章練習

1. 創建學生管理系統的結構體
2. 實現圖書館管理系統
3. 設計員工薪資計算系統
4. 創建電商產品管理結構體

---

**下一章：[指針](../08-pointers/)**