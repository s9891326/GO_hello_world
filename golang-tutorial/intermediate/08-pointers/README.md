# 第八章：指針

## 🎯 學習目標

- 理解指針的概念和用途
- 掌握指針的聲明和使用
- 學會指針與函數的配合
- 了解指針與結構體的關係
- 掌握指針的常見應用場景
- 學會指針的最佳實踐和安全使用

## 📍 指針基礎

指針是存儲另一個變數內存地址的變數。Go 語言的指針相比 C/C++ 更加安全，不支援指針運算，但保持了高效的內存訪問能力。

### 指針的基本概念

```
內存地址示意圖：
┌─────────────────────────────────────┐
│ 內存地址  │  變數名  │    值        │
├─────────────────────────────────────┤
│ 0x1040a124 │   num   │     42      │
│ 0x1040a128 │   ptr   │  0x1040a124 │
│ 0x1040a12c │   ...   │    ...      │
└─────────────────────────────────────┘

num: 普通變數，存儲值 42
ptr: 指針變數，存儲 num 的地址 0x1040a124
```

### 指針的聲明和初始化

```go
package main

import "fmt"

func demonstratePointerBasics() {
    // 聲明普通變數
    var num int = 42
    var name string = "Go語言"
    
    // 聲明指針變數
    var ptr *int        // 指向 int 的指針，零值為 nil
    var strPtr *string  // 指向 string 的指針
    
    // 獲取變數的地址
    ptr = &num       // & 是取地址運算符
    strPtr = &name
    
    fmt.Printf("num 的值: %d\n", num)
    fmt.Printf("num 的地址: %p\n", &num)
    fmt.Printf("ptr 的值（即 num 的地址）: %p\n", ptr)
    fmt.Printf("ptr 指向的值: %d\n", *ptr)  // * 是解引用運算符
    
    fmt.Printf("name 的值: %s\n", name)
    fmt.Printf("name 的地址: %p\n", &name)
    fmt.Printf("strPtr 指向的值: %s\n", *strPtr)
    
    // 檢查指針是否為 nil
    var nilPtr *int
    fmt.Printf("nilPtr 是否為 nil: %t\n", nilPtr == nil)
}
```

### 指針的基本操作

```go
func demonstratePointerOperations() {
    fmt.Println("\n--- 指針基本操作 ---")
    
    // 創建變數
    x := 100
    fmt.Printf("原始值 x: %d\n", x)
    
    // 創建指針
    ptr := &x
    fmt.Printf("指針地址: %p\n", ptr)
    fmt.Printf("指針指向的值: %d\n", *ptr)
    
    // 通過指針修改值
    *ptr = 200
    fmt.Printf("通過指針修改後 x: %d\n", x)
    
    // 指針的指針
    ptrPtr := &ptr
    fmt.Printf("指針的指針地址: %p\n", ptrPtr)
    fmt.Printf("指針的指針指向的地址: %p\n", *ptrPtr)
    fmt.Printf("指針的指針最終指向的值: %d\n", **ptrPtr)
    
    // 修改指針指向
    y := 300
    ptr = &y  // ptr 現在指向 y
    fmt.Printf("ptr 現在指向 y，值為: %d\n", *ptr)
    fmt.Printf("x 仍然是: %d\n", x)
}
```

## 🔄 指針與函數

指針在函數中的應用是 Go 語言中非常重要的概念。

### 值傳遞 vs 指針傳遞

```go
// 值傳遞：函數接收變數的副本
func doubleValue(x int) int {
    x = x * 2
    return x  // 需要返回修改後的值
}

// 指針傳遞：函數接收變數的地址
func doublePointer(x *int) {
    *x = *x * 2  // 直接修改原始變數
}

// 交換兩個值（值傳遞版本 - 無效）
func swapValues(a, b int) {
    a, b = b, a  // 只交換了副本
}

// 交換兩個值（指針版本 - 有效）
func swapPointers(a, b *int) {
    *a, *b = *b, *a  // 交換了原始變數
}

func demonstratePointerFunctions() {
    fmt.Println("\n--- 指針與函數 ---")
    
    // 值傳遞示例
    num1 := 10
    doubled := doubleValue(num1)
    fmt.Printf("值傳遞 - 原始值: %d, 加倍後: %d\n", num1, doubled)
    
    // 指針傳遞示例
    num2 := 10
    doublePointer(&num2)
    fmt.Printf("指針傳遞 - 修改後的值: %d\n", num2)
    
    // 交換值示例
    a, b := 100, 200
    fmt.Printf("交換前: a=%d, b=%d\n", a, b)
    
    swapValues(a, b)  // 無效的交換
    fmt.Printf("值交換後: a=%d, b=%d\n", a, b)
    
    swapPointers(&a, &b)  // 有效的交換
    fmt.Printf("指針交換後: a=%d, b=%d\n", a, b)
}
```

### 函數返回指針

```go
// 返回局部變數的指針（Go 中是安全的）
func createInt(value int) *int {
    x := value  // 局部變數
    return &x   // 返回局部變數的地址（Go 會自動處理內存管理）
}

// 工廠函數
func newPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// 創建指定大小的切片
func makeSlice(size int) *[]int {
    slice := make([]int, size)
    return &slice
}

func demonstrateReturnPointers() {
    fmt.Println("\n--- 函數返回指針 ---")
    
    // 獲取指向新建整數的指針
    intPtr := createInt(42)
    fmt.Printf("創建的整數: %d, 地址: %p\n", *intPtr, intPtr)
    
    // 創建 Person 實例
    person := newPerson("Alice", 25)
    fmt.Printf("創建的人員: %+v\n", *person)
    
    // 創建切片指針
    slicePtr := makeSlice(5)
    (*slicePtr)[0] = 100
    fmt.Printf("創建的切片: %v\n", *slicePtr)
}
```

## 🏗️ 指針與結構體

指針與結構體的組合是 Go 語言中非常強大的特性。

### 結構體指針

```go
type Person struct {
    Name string
    Age  int
    City string
}

// 結構體方法：指針接收者
func (p *Person) SetAge(age int) {
    p.Age = age
}

func (p *Person) MoveTo(city string) {
    p.City = city
}

// 結構體方法：值接收者
func (p Person) GetInfo() string {
    return fmt.Sprintf("%s (%d歲) 住在 %s", p.Name, p.Age, p.City)
}

func demonstrateStructPointers() {
    fmt.Println("\n--- 指針與結構體 ---")
    
    // 創建結構體實例
    person1 := Person{Name: "Bob", Age: 30, City: "台北"}
    fmt.Printf("person1: %+v\n", person1)
    
    // 創建指向結構體的指針
    personPtr := &person1
    fmt.Printf("指針地址: %p\n", personPtr)
    fmt.Printf("通過指針訪問: %+v\n", *personPtr)
    
    // Go 語言的語法糖：自動解引用
    fmt.Printf("姓名: %s\n", personPtr.Name)  // 等同於 (*personPtr).Name
    fmt.Printf("年齡: %d\n", personPtr.Age)   // 等同於 (*personPtr).Age
    
    // 通過指針修改結構體
    personPtr.Age = 31
    personPtr.City = "高雄"
    fmt.Printf("修改後: %+v\n", person1)
    
    // 使用 new 創建結構體指針
    person2 := new(Person)
    person2.Name = "Charlie"
    person2.Age = 28
    person2.City = "台中"
    fmt.Printf("new 創建: %+v\n", *person2)
    
    // 調用方法
    person2.SetAge(29)
    person2.MoveTo("台南")
    fmt.Printf("方法調用後: %s\n", person2.GetInfo())
}
```

### 結構體中的指針字段

```go
type Node struct {
    Value int
    Next  *Node  // 指向下一個節點的指針
}

type LinkedList struct {
    Head *Node
    Size int
}

func (ll *LinkedList) Append(value int) {
    newNode := &Node{Value: value, Next: nil}
    
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        current := ll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    ll.Size++
}

func (ll *LinkedList) Display() {
    fmt.Print("鏈表: ")
    current := ll.Head
    for current != nil {
        fmt.Printf("%d", current.Value)
        if current.Next != nil {
            fmt.Print(" -> ")
        }
        current = current.Next
    }
    fmt.Printf(" (大小: %d)\n", ll.Size)
}

func demonstrateStructWithPointers() {
    fmt.Println("\n--- 結構體中的指針字段 ---")
    
    // 創建鏈表
    list := &LinkedList{}
    
    // 添加元素
    list.Append(1)
    list.Append(2)
    list.Append(3)
    list.Append(4)
    
    // 顯示鏈表
    list.Display()
    
    // 直接操作節點
    if list.Head != nil {
        fmt.Printf("第一個節點的值: %d\n", list.Head.Value)
        if list.Head.Next != nil {
            fmt.Printf("第二個節點的值: %d\n", list.Head.Next.Value)
        }
    }
}
```

## 🧮 指針與數組、切片

### 數組指針 vs 指針數組

```go
func demonstrateArrayPointers() {
    fmt.Println("\n--- 指針與數組 ---")
    
    // 數組
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Printf("原始數組: %v\n", arr)
    
    // 數組指針：指向整個數組的指針
    arrPtr := &arr
    fmt.Printf("數組指針指向的數組: %v\n", *arrPtr)
    
    // 通過數組指針修改數組
    (*arrPtr)[0] = 100
    fmt.Printf("修改後的數組: %v\n", arr)
    
    // 指針數組：存儲指針的數組
    var ptrArray [3]*int
    a, b, c := 10, 20, 30
    ptrArray[0] = &a
    ptrArray[1] = &b
    ptrArray[2] = &c
    
    fmt.Printf("指針數組: [%p, %p, %p]\n", ptrArray[0], ptrArray[1], ptrArray[2])
    fmt.Printf("指針數組指向的值: [%d, %d, %d]\n", *ptrArray[0], *ptrArray[1], *ptrArray[2])
    
    // 修改指針數組指向的值
    *ptrArray[0] = 100
    fmt.Printf("修改後 a 的值: %d\n", a)
}
```

### 切片與指針

```go
func demonstrateSlicePointers() {
    fmt.Println("\n--- 指針與切片 ---")
    
    // 切片本身就是引用類型
    slice := []int{1, 2, 3, 4, 5}
    fmt.Printf("原始切片: %v\n", slice)
    
    // 切片指針
    slicePtr := &slice
    fmt.Printf("切片指針指向的切片: %v\n", *slicePtr)
    
    // 通過切片指針修改
    (*slicePtr)[0] = 100
    fmt.Printf("修改後的切片: %v\n", slice)
    
    // 切片元素的指針
    elementPtr := &slice[1]
    fmt.Printf("第二個元素的地址: %p, 值: %d\n", elementPtr, *elementPtr)
    
    *elementPtr = 200
    fmt.Printf("修改元素後的切片: %v\n", slice)
    
    // 指針切片：存儲指針的切片
    var ptrSlice []*int
    for i := range slice {
        ptrSlice = append(ptrSlice, &slice[i])
    }
    
    fmt.Printf("指針切片長度: %d\n", len(ptrSlice))
    fmt.Print("指針切片指向的值: [")
    for i, ptr := range ptrSlice {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Printf("%d", *ptr)
    }
    fmt.Println("]")
}
```

## 🗺️ 指針與映射

```go
func demonstrateMapPointers() {
    fmt.Println("\n--- 指針與映射 ---")
    
    // 創建映射
    m := map[string]int{
        "apple":  10,
        "banana": 20,
        "orange": 30,
    }
    fmt.Printf("原始映射: %v\n", m)
    
    // 映射指針
    mapPtr := &m
    (*mapPtr)["apple"] = 100
    fmt.Printf("通過指針修改後: %v\n", m)
    
    // 注意：無法獲取映射值的地址
    // valuePtr := &m["apple"]  // 編譯錯誤！
    
    // 但可以創建指向值的指針的映射
    ptrMap := make(map[string]*int)
    
    values := map[string]int{"x": 100, "y": 200, "z": 300}
    for k, v := range values {
        temp := v  // 重要：需要創建新變數
        ptrMap[k] = &temp
    }
    
    fmt.Printf("指針映射: ")
    for k, ptr := range ptrMap {
        fmt.Printf("%s->%d ", k, *ptr)
    }
    fmt.Println()
    
    // 修改指針映射中的值
    *ptrMap["x"] = 1000
    fmt.Printf("修改後指針映射中 x 的值: %d\n", *ptrMap["x"])
}
```

## ⚠️ 指針的安全使用

### 避免常見錯誤

```go
func demonstratePointerSafety() {
    fmt.Println("\n--- 指針安全使用 ---")
    
    // 1. 空指針檢查
    var ptr *int
    if ptr != nil {
        fmt.Printf("指針值: %d\n", *ptr)
    } else {
        fmt.Println("指針為 nil，不能解引用")
    }
    
    // 2. 正確的指針初始化
    num := 42
    ptr = &num
    if ptr != nil {
        fmt.Printf("安全的指針值: %d\n", *ptr)
    }
    
    // 3. 避免懸空指針（Go 的 GC 會處理）
    createAndUsePointer()
    
    // 4. 循環引用的處理
    demonstrateCircularReference()
}

func createAndUsePointer() {
    ptr := createInt(100)  // 局部變數的指針
    fmt.Printf("函數返回的指針值: %d\n", *ptr)
    // Go 的垃圾回收器會自動管理內存
}

type CircularNode struct {
    Value int
    Ref   *CircularNode
}

func demonstrateCircularReference() {
    fmt.Println("循環引用示例:")
    
    node1 := &CircularNode{Value: 1}
    node2 := &CircularNode{Value: 2}
    
    // 創建循環引用
    node1.Ref = node2
    node2.Ref = node1
    
    fmt.Printf("Node1 -> Node2: %d -> %d\n", node1.Value, node1.Ref.Value)
    fmt.Printf("Node2 -> Node1: %d -> %d\n", node2.Value, node2.Ref.Value)
    
    // Go 的垃圾回收器可以處理循環引用
}
```

### 指針的性能考量

```go
import "time"

func demonstratePointerPerformance() {
    fmt.Println("\n--- 指針性能考量 ---")
    
    // 大結構體
    type LargeStruct struct {
        Data [1000]int
        Name string
        Age  int
    }
    
    largeStruct := LargeStruct{Name: "Test", Age: 25}
    
    // 測試值傳遞的性能
    start := time.Now()
    for i := 0; i < 10000; i++ {
        processLargeStructByValue(largeStruct)
    }
    valueDuration := time.Since(start)
    
    // 測試指針傳遞的性能
    start = time.Now()
    for i := 0; i < 10000; i++ {
        processLargeStructByPointer(&largeStruct)
    }
    pointerDuration := time.Since(start)
    
    fmt.Printf("值傳遞耗時: %v\n", valueDuration)
    fmt.Printf("指針傳遞耗時: %v\n", pointerDuration)
    fmt.Printf("性能提升: %.2fx\n", float64(valueDuration)/float64(pointerDuration))
}

func processLargeStructByValue(ls LargeStruct) {
    // 模擬處理
    _ = ls.Name
}

func processLargeStructByPointer(ls *LargeStruct) {
    // 模擬處理
    _ = ls.Name
}
```

## 💡 指針最佳實踐

### 1. 何時使用指針

```go
// ✅ 好的使用場景
func goodPointerUsage() {
    // 1. 需要修改原始數據
    var count int
    increment(&count)
    
    // 2. 避免大結構體的複製
    person := &Person{Name: "Alice", Age: 25}
    updatePerson(person)
    
    // 3. 可選值（可能為 nil）
    var optionalValue *int
    if someCondition() {
        value := 42
        optionalValue = &value
    }
    processOptionalValue(optionalValue)
    
    // 4. 實現鏈表、樹等數據結構
    node := &Node{Value: 1, Next: nil}
    buildLinkedList(node)
}

func increment(n *int) {
    *n++
}

func updatePerson(p *Person) {
    p.Age++
}

func someCondition() bool {
    return true
}

func processOptionalValue(val *int) {
    if val != nil {
        fmt.Printf("可選值: %d\n", *val)
    } else {
        fmt.Println("無值")
    }
}

func buildLinkedList(head *Node) {
    // 構建鏈表邏輯
}
```

### 2. 指針設計模式

```go
// 單例模式
type Config struct {
    DatabaseURL string
    APIKey      string
}

var configInstance *Config

func GetConfig() *Config {
    if configInstance == nil {
        configInstance = &Config{
            DatabaseURL: "localhost:5432",
            APIKey:      "secret-key",
        }
    }
    return configInstance
}

// 建造者模式
type RequestBuilder struct {
    request *HttpRequest
}

type HttpRequest struct {
    Method  string
    URL     string
    Headers map[string]string
    Body    string
}

func NewRequestBuilder() *RequestBuilder {
    return &RequestBuilder{
        request: &HttpRequest{
            Headers: make(map[string]string),
        },
    }
}

func (rb *RequestBuilder) Method(method string) *RequestBuilder {
    rb.request.Method = method
    return rb
}

func (rb *RequestBuilder) URL(url string) *RequestBuilder {
    rb.request.URL = url
    return rb
}

func (rb *RequestBuilder) Header(key, value string) *RequestBuilder {
    rb.request.Headers[key] = value
    return rb
}

func (rb *RequestBuilder) Body(body string) *RequestBuilder {
    rb.request.Body = body
    return rb
}

func (rb *RequestBuilder) Build() *HttpRequest {
    return rb.request
}

func demonstratePointerPatterns() {
    fmt.Println("\n--- 指針設計模式 ---")
    
    // 單例模式
    config1 := GetConfig()
    config2 := GetConfig()
    fmt.Printf("單例模式 - 同一實例: %t\n", config1 == config2)
    
    // 建造者模式
    request := NewRequestBuilder().
        Method("POST").
        URL("https://api.example.com/users").
        Header("Content-Type", "application/json").
        Header("Authorization", "Bearer token").
        Body(`{"name": "Alice", "email": "alice@example.com"}`).
        Build()
    
    fmt.Printf("建造者模式創建的請求:\n")
    fmt.Printf("  方法: %s\n", request.Method)
    fmt.Printf("  URL: %s\n", request.URL)
    fmt.Printf("  頭部數量: %d\n", len(request.Headers))
}
```

## 🎯 本章練習

1. 實現雙向鏈表
2. 創建二叉樹結構
3. 實現對象池模式
4. 創建內存緩存系統

---

**下一章：[數組和切片](../09-arrays-slices/)**