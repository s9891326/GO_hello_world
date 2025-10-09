# 第十一章：接口

## 🎯 學習目標

- 理解接口的概念和作用
- 掌握接口的定義和實現
- 學會接口的組合和嵌入
- 了解空接口和類型斷言
- 掌握接口的多態應用
- 學會接口的設計模式和最佳實踐

## 🔌 接口基礎

接口（Interface）定義了一組方法簽名的集合。在 Go 中，接口是隱式實現的，任何類型只要實現了接口定義的所有方法，就自動實現了該接口。

### 接口的特點

```
Go 接口的關鍵特性：
┌─────────────────────────────────────┐
│ • 隱式實現（Duck Typing）             │
│ • 接口是類型                          │
│ • 可以作為變數、參數、返回值           │
│ • 支援接口組合                        │
│ • 零值是 nil                        │
│ • 面向接口編程                        │
└─────────────────────────────────────┘
```

### 接口的定義

```go
package main

import "fmt"

// 定義接口
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// 組合接口
type ReadWriter interface {
    Reader
    Writer
}

// 更複雜的接口
type Shape interface {
    Area() float64
    Perimeter() float64
    String() string
}

// 定義具體類型
type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

// Rectangle 實現 Shape 接口
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f×%.2f)", r.Width, r.Height)
}

// Circle 實現 Shape 接口
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

func demonstrateBasicInterface() {
    fmt.Println("--- 基本接口演示 ---")
    
    // 創建具體類型實例
    rect := Rectangle{Width: 5.0, Height: 3.0}
    circle := Circle{Radius: 2.0}
    
    // 使用接口變數
    var shape Shape
    
    shape = rect
    fmt.Printf("矩形: %s, 面積: %.2f, 周長: %.2f\n", 
        shape.String(), shape.Area(), shape.Perimeter())
    
    shape = circle
    fmt.Printf("圓形: %s, 面積: %.2f, 周長: %.2f\n", 
        shape.String(), shape.Area(), shape.Perimeter())
    
    // 接口切片
    shapes := []Shape{rect, circle}
    fmt.Println("\n形狀集合:")
    for i, s := range shapes {
        fmt.Printf("  %d. %s - 面積: %.2f\n", i+1, s.String(), s.Area())
    }
}
```

## 🏗️ 接口的實現

### 隱式實現

```go
// 定義接口
type Stringer interface {
    String() string
}

type Closer interface {
    Close() error
}

// 定義類型
type File struct {
    Name string
    Data []byte
}

type Database struct {
    Host string
    Port int
}

// File 實現 Stringer 接口
func (f File) String() string {
    return fmt.Sprintf("File: %s (%d bytes)", f.Name, len(f.Data))
}

// File 實現 Closer 接口
func (f File) Close() error {
    fmt.Printf("關閉文件: %s\n", f.Name)
    return nil
}

// Database 實現 Stringer 接口
func (db Database) String() string {
    return fmt.Sprintf("Database: %s:%d", db.Host, db.Port)
}

// Database 實現 Closer 接口
func (db Database) Close() error {
    fmt.Printf("關閉數據庫連接: %s:%d\n", db.Host, db.Port)
    return nil
}

func demonstrateImplicitImplementation() {
    fmt.Println("\n--- 隱式實現演示 ---")
    
    file := File{Name: "config.txt", Data: []byte("configuration data")}
    db := Database{Host: "localhost", Port: 5432}
    
    // 使用 Stringer 接口
    var s Stringer
    s = file
    fmt.Printf("文件信息: %s\n", s.String())
    
    s = db
    fmt.Printf("數據庫信息: %s\n", s.String())
    
    // 使用 Closer 接口
    var c Closer
    c = file
    c.Close()
    
    c = db
    c.Close()
    
    // 同時實現多個接口
    resources := []Closer{file, db}
    fmt.Println("\n關閉所有資源:")
    for _, resource := range resources {
        resource.Close()
    }
}
```

### 接口作為參數

```go
// 使用接口作為參數實現多態
func printInfo(s Stringer) {
    fmt.Printf("對象信息: %s\n", s.String())
}

func closeResource(c Closer) error {
    fmt.Printf("正在關閉資源...\n")
    return c.Close()
}

func calculateTotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func demonstrateInterfaceParameters() {
    fmt.Println("\n--- 接口作為參數演示 ---")
    
    rect := Rectangle{Width: 4.0, Height: 3.0}
    circle := Circle{Radius: 2.5}
    file := File{Name: "data.txt", Data: []byte("some data")}
    
    // 多態調用
    printInfo(rect)
    printInfo(circle)
    printInfo(file)
    
    // 計算總面積
    shapes := []Shape{rect, circle}
    total := calculateTotalArea(shapes)
    fmt.Printf("所有形狀的總面積: %.2f\n", total)
    
    // 關閉資源
    closeResource(file)
}
```

## 🔍 空接口和類型斷言

### 空接口 interface{}

```go
func demonstrateEmptyInterface() {
    fmt.Println("\n--- 空接口演示 ---")
    
    // 空接口可以存儲任何類型的值
    var anything interface{}
    
    anything = 42
    fmt.Printf("存儲整數: %v (類型: %T)\n", anything, anything)
    
    anything = "Hello, World!"
    fmt.Printf("存儲字符串: %v (類型: %T)\n", anything, anything)
    
    anything = []int{1, 2, 3, 4, 5}
    fmt.Printf("存儲切片: %v (類型: %T)\n", anything, anything)
    
    anything = Rectangle{Width: 10, Height: 5}
    fmt.Printf("存儲結構體: %v (類型: %T)\n", anything, anything)
    
    // 空接口切片
    items := []interface{}{
        42,
        "hello",
        3.14,
        true,
        []string{"a", "b", "c"},
    }
    
    fmt.Println("混合類型切片:")
    for i, item := range items {
        fmt.Printf("  [%d] %v (類型: %T)\n", i, item, item)
    }
}
```

### 類型斷言

```go
func demonstrateTypeAssertion() {
    fmt.Println("\n--- 類型斷言演示 ---")
    
    var value interface{} = "Hello, Go!"
    
    // 基本類型斷言
    if str, ok := value.(string); ok {
        fmt.Printf("字符串值: %s (長度: %d)\n", str, len(str))
    } else {
        fmt.Println("不是字符串類型")
    }
    
    // 錯誤的類型斷言
    if num, ok := value.(int); ok {
        fmt.Printf("整數值: %d\n", num)
    } else {
        fmt.Println("不是整數類型")
    }
    
    // 類型斷言 panic 示例（註釋掉以避免程序崩潰）
    // str := value.(int) // 這會導致 panic
    
    // 處理多種類型
    values := []interface{}{
        42,
        "hello",
        3.14159,
        true,
        Rectangle{Width: 5, Height: 3},
    }
    
    fmt.Println("處理多種類型:")
    for i, v := range values {
        fmt.Printf("  [%d] ", i)
        processValue(v)
    }
}

func processValue(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("整數: %d (平方: %d)\n", v, v*v)
    case string:
        fmt.Printf("字符串: %s (長度: %d)\n", v, len(v))
    case float64:
        fmt.Printf("浮點數: %.3f (開方: %.3f)\n", v, math.Sqrt(v))
    case bool:
        fmt.Printf("布爾值: %t (非: %t)\n", v, !v)
    case Rectangle:
        fmt.Printf("矩形: %s (面積: %.2f)\n", v.String(), v.Area())
    default:
        fmt.Printf("未知類型: %T = %v\n", v, v)
    }
}
```

## 🧩 接口組合

### 接口嵌入

```go
// 基礎接口
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// 組合接口
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// 或者更簡潔的方式
type ReadWriteCloser2 interface {
    ReadWriter
    Closer
}

// 實現組合接口的類型
type Buffer struct {
    data []byte
    pos  int
}

func (b *Buffer) Read(p []byte) (int, error) {
    if b.pos >= len(b.data) {
        return 0, fmt.Errorf("EOF")
    }
    
    n := copy(p, b.data[b.pos:])
    b.pos += n
    fmt.Printf("讀取 %d 字節\n", n)
    return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
    b.data = append(b.data, p...)
    fmt.Printf("寫入 %d 字節\n", len(p))
    return len(p), nil
}

func (b *Buffer) Close() error {
    fmt.Println("關閉緩衝區")
    b.data = nil
    b.pos = 0
    return nil
}

func demonstrateInterfaceComposition() {
    fmt.Println("\n--- 接口組合演示 ---")
    
    buffer := &Buffer{}
    
    // 作為 Writer 使用
    var w Writer = buffer
    w.Write([]byte("Hello, "))
    w.Write([]byte("Interface!"))
    
    // 作為 Reader 使用
    var r Reader = buffer
    data := make([]byte, 5)
    r.Read(data)
    fmt.Printf("讀取的數據: %s\n", string(data))
    
    // 作為組合接口使用
    var rwc ReadWriteCloser = buffer
    rwc.Write([]byte(" More data"))
    
    data2 := make([]byte, 10)
    rwc.Read(data2)
    fmt.Printf("讀取的數據: %s\n", string(data2))
    
    rwc.Close()
}
```

## 🎭 接口的多態應用

### 策略模式

```go
// 定義策略接口
type SortStrategy interface {
    Sort([]int)
    Name() string
}

// 冒泡排序策略
type BubbleSort struct{}

func (bs BubbleSort) Sort(data []int) {
    n := len(data)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }
}

func (bs BubbleSort) Name() string {
    return "Bubble Sort"
}

// 快速排序策略
type QuickSort struct{}

func (qs QuickSort) Sort(data []int) {
    if len(data) < 2 {
        return
    }
    quicksort(data, 0, len(data)-1)
}

func (qs QuickSort) Name() string {
    return "Quick Sort"
}

func quicksort(data []int, low, high int) {
    if low < high {
        pi := partition(data, low, high)
        quicksort(data, low, pi-1)
        quicksort(data, pi+1, high)
    }
}

func partition(data []int, low, high int) int {
    pivot := data[high]
    i := low - 1
    
    for j := low; j < high; j++ {
        if data[j] < pivot {
            i++
            data[i], data[j] = data[j], data[i]
        }
    }
    data[i+1], data[high] = data[high], data[i+1]
    return i + 1
}

// 排序器上下文
type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(data []int) {
    if s.strategy != nil {
        fmt.Printf("使用 %s 排序\n", s.strategy.Name())
        s.strategy.Sort(data)
    }
}

func demonstrateStrategyPattern() {
    fmt.Println("\n--- 策略模式演示 ---")
    
    data1 := []int{64, 34, 25, 12, 22, 11, 90}
    data2 := make([]int, len(data1))
    copy(data2, data1)
    
    sorter := &Sorter{}
    
    // 使用冒泡排序
    sorter.SetStrategy(BubbleSort{})
    fmt.Printf("排序前: %v\n", data1)
    sorter.Sort(data1)
    fmt.Printf("排序後: %v\n", data1)
    
    fmt.Println()
    
    // 使用快速排序
    sorter.SetStrategy(QuickSort{})
    fmt.Printf("排序前: %v\n", data2)
    sorter.Sort(data2)
    fmt.Printf("排序後: %v\n", data2)
}
```

### 工廠模式

```go
// 產品接口
type Animal interface {
    Speak() string
    Type() string
}

// 具體產品
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof! I'm " + d.Name
}

func (d Dog) Type() string {
    return "Dog"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow! I'm " + c.Name
}

func (c Cat) Type() string {
    return "Cat"
}

type Bird struct {
    Name string
}

func (b Bird) Speak() string {
    return "Tweet! I'm " + b.Name
}

func (b Bird) Type() string {
    return "Bird"
}

// 工廠函數
func CreateAnimal(animalType, name string) Animal {
    switch animalType {
    case "dog":
        return Dog{Name: name}
    case "cat":
        return Cat{Name: name}
    case "bird":
        return Bird{Name: name}
    default:
        return nil
    }
}

func demonstrateFactoryPattern() {
    fmt.Println("\n--- 工廠模式演示 ---")
    
    animals := []Animal{
        CreateAnimal("dog", "Buddy"),
        CreateAnimal("cat", "Whiskers"),
        CreateAnimal("bird", "Tweety"),
    }
    
    fmt.Println("動物園:")
    for i, animal := range animals {
        if animal != nil {
            fmt.Printf("  %d. %s: %s\n", i+1, animal.Type(), animal.Speak())
        }
    }
}
```

## 💡 接口設計原則

### 1. 接口隔離原則

```go
// 不好的設計：接口太大
type BadWorker interface {
    Work()
    Eat()
    Sleep()
    Code()
    Debug()
    Test()
}

// 好的設計：小而專注的接口
type Worker interface {
    Work()
}

type Eater interface {
    Eat()
}

type Sleeper interface {
    Sleep()
}

type Developer interface {
    Worker
    Code()
    Debug()
}

type Tester interface {
    Worker
    Test()
}

// 具體實現
type Programmer struct {
    Name string
}

func (p Programmer) Work() {
    fmt.Printf("%s 正在工作\n", p.Name)
}

func (p Programmer) Code() {
    fmt.Printf("%s 正在編程\n", p.Name)
}

func (p Programmer) Debug() {
    fmt.Printf("%s 正在調試\n", p.Name)
}

func (p Programmer) Eat() {
    fmt.Printf("%s 正在吃飯\n", p.Name)
}

func (p Programmer) Sleep() {
    fmt.Printf("%s 正在睡覺\n", p.Name)
}
```

### 2. 依賴倒置原則

```go
// 高層模組不應該依賴低層模組，兩者都應該依賴抽象

// 抽象接口
type Logger interface {
    Log(message string)
}

type Database interface {
    Save(data interface{}) error
}

// 低層模組實現
type FileLogger struct {
    filename string
}

func (fl FileLogger) Log(message string) {
    fmt.Printf("寫入文件 %s: %s\n", fl.filename, message)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
    fmt.Printf("控制台輸出: %s\n", message)
}

type MySQLDatabase struct {
    host string
}

func (db MySQLDatabase) Save(data interface{}) error {
    fmt.Printf("保存到 MySQL (%s): %v\n", db.host, data)
    return nil
}

// 高層模組
type UserService struct {
    logger Logger
    db     Database
}

func NewUserService(logger Logger, db Database) *UserService {
    return &UserService{
        logger: logger,
        db:     db,
    }
}

func (us *UserService) CreateUser(name string) error {
    us.logger.Log("開始創建用戶: " + name)
    
    user := map[string]string{"name": name}
    if err := us.db.Save(user); err != nil {
        us.logger.Log("創建用戶失敗: " + err.Error())
        return err
    }
    
    us.logger.Log("用戶創建成功: " + name)
    return nil
}

func demonstrateDependencyInversion() {
    fmt.Println("\n--- 依賴倒置原則演示 ---")
    
    // 可以靈活切換實現
    fileLogger := FileLogger{filename: "app.log"}
    consoleLogger := ConsoleLogger{}
    mysql := MySQLDatabase{host: "localhost:3306"}
    
    // 使用文件日誌
    service1 := NewUserService(fileLogger, mysql)
    service1.CreateUser("Alice")
    
    fmt.Println()
    
    // 使用控制台日誌
    service2 := NewUserService(consoleLogger, mysql)
    service2.CreateUser("Bob")
}
```

## 🎯 接口最佳實踐

### 1. 小接口原則

```go
// 保持接口小而專注
type Stringer interface {
    String() string
}

type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// 通過組合創建大接口
type ReadWriter interface {
    Reader
    Writer
}
```

### 2. 接受接口，返回結構體

```go
// 好的實踐：參數使用接口，返回具體類型
func ProcessData(r Reader) *DataProcessor {
    return &DataProcessor{reader: r}
}

type DataProcessor struct {
    reader Reader
}

func (dp *DataProcessor) Process() error {
    // 處理邏輯
    return nil
}
```

### 3. 零值友好的接口

```go
type SafeWriter interface {
    Write([]byte) (int, error)
    IsReady() bool
}

type NullWriter struct{}

func (nw NullWriter) Write(p []byte) (int, error) {
    return len(p), nil // 丟棄所有數據
}

func (nw NullWriter) IsReady() bool {
    return true
}

// 零值友好的使用
func WriteData(w SafeWriter, data []byte) error {
    if w == nil {
        w = NullWriter{} // 提供默認實現
    }
    
    if !w.IsReady() {
        return fmt.Errorf("writer not ready")
    }
    
    _, err := w.Write(data)
    return err
}
```

## 🎯 本章練習

1. 實現圖形計算器
2. 創建插件系統
3. 實現數據處理管道
4. 設計通知系統

---

**下一章：[協程](../12-goroutines/)**