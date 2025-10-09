package main

import (
	"fmt"
	"time"
)

// 構造函數模式示例

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
		return nil, fmt.Errorf("年齡必須在 0-150 之間，得到: %d", age)
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

// 選項模式相關結構
type PersonOption func(*Person)

// 選項函數
func WithAge(age int) PersonOption {
	return func(p *Person) {
		if age >= 0 && age <= 150 {
			p.Age = age
		}
	}
}

func WithCity(city string) PersonOption {
	return func(p *Person) {
		if city != "" {
			p.City = city
		}
	}
}

// 使用選項模式的構造函數
func NewPersonWithOptions(name string, options ...PersonOption) *Person {
	if name == "" {
		name = "未知"
	}
	
	person := &Person{
		Name: name,
		Age:  0,
		City: "未指定",
	}
	
	// 應用所有選項
	for _, option := range options {
		option(person)
	}
	
	return person
}

// 複雜結構體的構造函數
type Config struct {
	Host         string
	Port         int
	Timeout      time.Duration
	MaxRetries   int
	EnableSSL    bool
	DatabaseURL  string
	LogLevel     string
	Features     map[string]bool
}

// 配置選項類型
type ConfigOption func(*Config)

// 配置選項函數
func WithHost(host string) ConfigOption {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) ConfigOption {
	return func(c *Config) {
		if port > 0 && port <= 65535 {
			c.Port = port
		}
	}
}

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(c *Config) {
		if timeout > 0 {
			c.Timeout = timeout
		}
	}
}

func WithSSL(enable bool) ConfigOption {
	return func(c *Config) {
		c.EnableSSL = enable
	}
}

func WithDatabase(url string) ConfigOption {
	return func(c *Config) {
		c.DatabaseURL = url
	}
}

func WithLogLevel(level string) ConfigOption {
	return func(c *Config) {
		validLevels := map[string]bool{
			"debug": true, "info": true, "warn": true, "error": true,
		}
		if validLevels[level] {
			c.LogLevel = level
		}
	}
}

func WithFeature(name string, enabled bool) ConfigOption {
	return func(c *Config) {
		if c.Features == nil {
			c.Features = make(map[string]bool)
		}
		c.Features[name] = enabled
	}
}

// 帶有默認值的配置構造函數
func NewConfig(options ...ConfigOption) *Config {
	// 設置默認值
	config := &Config{
		Host:        "localhost",
		Port:        8080,
		Timeout:     30 * time.Second,
		MaxRetries:  3,
		EnableSSL:   false,
		DatabaseURL: "",
		LogLevel:    "info",
		Features:    make(map[string]bool),
	}
	
	// 應用選項
	for _, option := range options {
		option(config)
	}
	
	return config
}

// 建造者模式
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) Age(age int) *PersonBuilder {
	if age >= 0 && age <= 150 {
		pb.person.Age = age
	}
	return pb
}

func (pb *PersonBuilder) City(city string) *PersonBuilder {
	pb.person.City = city
	return pb
}

func (pb *PersonBuilder) Build() (*Person, error) {
	if pb.person.Name == "" {
		return nil, fmt.Errorf("姓名是必需的")
	}
	
	// 創建副本
	result := &Person{
		Name: pb.person.Name,
		Age:  pb.person.Age,
		City: pb.person.City,
	}
	
	return result, nil
}

func demonstrateConstructors() {
	fmt.Println("\n--- 構造函數模式演示 ---")
	
	// 簡單構造函數
	fmt.Println("\n🏗️ 簡單構造函數:")
	person1 := NewPerson("Alice", 25, "台北")
	fmt.Printf("   創建的人員: %+v\n", person1)
	
	// 帶驗證的構造函數
	fmt.Println("\n✅ 帶驗證的構造函數:")
	person2, err := NewPersonValidated("Bob", 30, "高雄")
	if err != nil {
		fmt.Printf("   ❌ 創建失敗: %v\n", err)
	} else {
		fmt.Printf("   ✅ 創建成功: %+v\n", person2)
	}
	
	// 測試驗證失敗的情況
	_, err = NewPersonValidated("", 25, "台中")
	if err != nil {
		fmt.Printf("   ❌ 驗證失敗（姓名為空）: %v\n", err)
	}
	
	_, err = NewPersonValidated("Charlie", -5, "台南")
	if err != nil {
		fmt.Printf("   ❌ 驗證失敗（年齡無效）: %v\n", err)
	}
	
	// 選項模式
	fmt.Println("\n⚙️ 選項模式:")
	person3 := NewPersonWithOptions("David")
	fmt.Printf("   僅姓名: %+v\n", person3)
	
	person4 := NewPersonWithOptions("Emily", 
		WithAge(28), 
		WithCity("桃園"))
	fmt.Printf("   使用選項: %+v\n", person4)
	
	person5 := NewPersonWithOptions("Frank", 
		WithAge(35), 
		WithCity("新竹"))
	fmt.Printf("   多個選項: %+v\n", person5)
	
	// 複雜配置構造
	fmt.Println("\n🔧 複雜配置構造:")
	config1 := NewConfig()
	fmt.Printf("   默認配置: %+v\n", *config1)
	
	config2 := NewConfig(
		WithHost("api.example.com"),
		WithPort(443),
		WithSSL(true),
		WithTimeout(60*time.Second),
		WithDatabase("postgres://localhost:5432/myapp"),
		WithLogLevel("debug"),
		WithFeature("cache", true),
		WithFeature("metrics", true),
	)
	fmt.Printf("   自定義配置:\n")
	fmt.Printf("      主機: %s:%d (SSL: %t)\n", config2.Host, config2.Port, config2.EnableSSL)
	fmt.Printf("      超時: %v\n", config2.Timeout)
	fmt.Printf("      數據庫: %s\n", config2.DatabaseURL)
	fmt.Printf("      日誌級別: %s\n", config2.LogLevel)
	fmt.Printf("      功能: %v\n", config2.Features)
	
	// 建造者模式
	fmt.Println("\n🏭 建造者模式:")
	person6, err := NewPersonBuilder().
		Name("Grace").
		Age(32).
		City("台中").
		Build()
	
	if err != nil {
		fmt.Printf("   ❌ 建造失敗: %v\n", err)
	} else {
		fmt.Printf("   ✅ 建造成功: %+v\n", person6)
	}
	
	// 建造者模式失敗案例
	_, err = NewPersonBuilder().
		Age(25).
		City("嘉義").
		Build() // 缺少姓名
	
	if err != nil {
		fmt.Printf("   ❌ 建造失敗（缺少姓名）: %v\n", err)
	}
}