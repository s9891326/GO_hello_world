# 第十章：映射（Map）

## 🎯 學習目標

- 理解映射的概念和特點
- 掌握映射的創建和初始化
- 學會映射的基本操作
- 了解映射的內部實現原理
- 掌握映射的高級用法
- 學會映射的最佳實踐

## 🗺️ 映射基礎

映射（Map）是一種鍵值對的集合，類似於其他語言中的字典、哈希表或關聯數組。在 Go 中，映射是引用類型。

### 映射的特點

```
映射的關鍵特性：
┌─────────────────────────────────────┐
│ • 無序存儲鍵值對                      │
│ • 鍵必須是可比較的類型                 │
│ • 值可以是任意類型                    │
│ • 零值是 nil                        │
│ • 引用類型                          │
│ • 線程不安全                        │
└─────────────────────────────────────┘
```

### 映射的聲明和初始化

```go
package main

import "fmt"

func demonstrateMapBasics() {
    // 1. 聲明映射
    var m1 map[string]int              // nil 映射
    fmt.Printf("nil 映射: %v (== nil: %t)\n", m1, m1 == nil)
    
    // 2. 使用 make 創建映射
    m2 := make(map[string]int)
    fmt.Printf("空映射: %v (== nil: %t)\n", m2, m2 == nil)
    
    // 3. 字面量初始化
    m3 := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }
    fmt.Printf("字面量映射: %v\n", m3)
    
    // 4. 部分初始化
    m4 := map[string]int{
        "one": 1,
        "two": 2,
        // 可以有尾隨逗號
    }
    fmt.Printf("部分初始化: %v\n", m4)
    
    // 5. 空映射初始化
    m5 := map[string]int{}
    fmt.Printf("空映射字面量: %v (== nil: %t)\n", m5, m5 == nil)
}
```

### 可以作為鍵的類型

```go
func demonstrateKeyTypes() {
    fmt.Println("\n--- 映射鍵類型 ---")
    
    // 基本類型作為鍵
    intMap := map[int]string{1: "one", 2: "two"}
    fmt.Printf("int 鍵: %v\n", intMap)
    
    stringMap := map[string]int{"hello": 1, "world": 2}
    fmt.Printf("string 鍵: %v\n", stringMap)
    
    boolMap := map[bool]string{true: "yes", false: "no"}
    fmt.Printf("bool 鍵: %v\n", boolMap)
    
    // 數組作為鍵（可比較）
    arrayMap := map[[3]int]string{
        {1, 2, 3}: "first",
        {4, 5, 6}: "second",
    }
    fmt.Printf("數組鍵: %v\n", arrayMap)
    
    // 結構體作為鍵（所有字段都可比較）
    type Point struct {
        X, Y int
    }
    pointMap := map[Point]string{
        {0, 0}: "origin",
        {1, 1}: "diagonal",
    }
    fmt.Printf("結構體鍵: %v\n", pointMap)
    
    // 以下類型不能作為鍵：
    // sliceMap := map[[]int]string{}     // 錯誤：slice 不可比較
    // mapMap := map[map[string]int]string{} // 錯誤：map 不可比較
    // funcMap := map[func()]string{}     // 錯誤：function 不可比較
}
```

## 🔧 映射的基本操作

### 增刪改查

```go
func demonstrateMapOperations() {
    fmt.Println("\n--- 映射基本操作 ---")
    
    // 創建映射
    scores := make(map[string]int)
    
    // 1. 添加/修改元素
    scores["Alice"] = 95
    scores["Bob"] = 87
    scores["Charlie"] = 92
    fmt.Printf("添加元素後: %v\n", scores)
    
    // 2. 獲取元素
    aliceScore := scores["Alice"]
    fmt.Printf("Alice 的分數: %d\n", aliceScore)
    
    // 3. 檢查鍵是否存在
    score, exists := scores["David"]
    if exists {
        fmt.Printf("David 的分數: %d\n", score)
    } else {
        fmt.Printf("David 不存在，默認值: %d\n", score)
    }
    
    // 4. 修改元素
    scores["Alice"] = 98
    fmt.Printf("修改 Alice 分數後: %v\n", scores)
    
    // 5. 刪除元素
    delete(scores, "Bob")
    fmt.Printf("刪除 Bob 後: %v\n", scores)
    
    // 6. 刪除不存在的鍵（安全操作）
    delete(scores, "NonExistent")
    fmt.Printf("刪除不存在的鍵後: %v\n", scores)
    
    // 7. 獲取映射長度
    fmt.Printf("映射長度: %d\n", len(scores))
}
```

### 遍歷映射

```go
func demonstrateMapIteration() {
    fmt.Println("\n--- 映射遍歷 ---")
    
    fruits := map[string]int{
        "apple":  10,
        "banana": 5,
        "orange": 8,
        "grape":  12,
    }
    
    // 1. 遍歷鍵值對
    fmt.Println("遍歷鍵值對:")
    for fruit, count := range fruits {
        fmt.Printf("  %s: %d\n", fruit, count)
    }
    
    // 2. 只遍歷鍵
    fmt.Print("只遍歷鍵: ")
    for fruit := range fruits {
        fmt.Printf("%s ", fruit)
    }
    fmt.Println()
    
    // 3. 只遍歷值
    fmt.Print("只遍歷值: ")
    for _, count := range fruits {
        fmt.Printf("%d ", count)
    }
    fmt.Println()
    
    // 4. 注意：映射遍歷順序是隨機的
    fmt.Println("多次遍歷順序可能不同:")
    for i := 0; i < 3; i++ {
        fmt.Printf("  第 %d 次: ", i+1)
        for fruit := range fruits {
            fmt.Printf("%s ", fruit)
        }
        fmt.Println()
    }
}
```

## 🔍 映射的高級用法

### 映射的零值處理

```go
func demonstrateMapZeroValues() {
    fmt.Println("\n--- 映射零值處理 ---")
    
    // 不同值類型的零值
    intMap := make(map[string]int)
    stringMap := make(map[string]string)
    boolMap := make(map[string]bool)
    sliceMap := make(map[string][]int)
    
    // 訪問不存在的鍵會返回零值
    fmt.Printf("不存在的 int 鍵: %d\n", intMap["nonexistent"])
    fmt.Printf("不存在的 string 鍵: '%s'\n", stringMap["nonexistent"])
    fmt.Printf("不存在的 bool 鍵: %t\n", boolMap["nonexistent"])
    fmt.Printf("不存在的 slice 鍵: %v\n", sliceMap["nonexistent"])
    
    // 利用零值的特性
    counter := make(map[string]int)
    words := []string{"hello", "world", "hello", "go", "world", "hello"}
    
    for _, word := range words {
        counter[word]++ // 零值是 0，直接可以遞增
    }
    fmt.Printf("單詞計數: %v\n", counter)
}
```

### 映射作為集合

```go
func demonstrateMapAsSet() {
    fmt.Println("\n--- 映射作為集合 ---")
    
    // 使用 map[T]bool 模擬集合
    set := make(map[string]bool)
    
    // 添加元素
    items := []string{"apple", "banana", "apple", "orange", "banana"}
    for _, item := range items {
        set[item] = true
    }
    
    fmt.Printf("集合內容: %v\n", set)
    
    // 檢查元素是否存在
    fmt.Printf("apple 在集合中: %t\n", set["apple"])
    fmt.Printf("grape 在集合中: %t\n", set["grape"])
    
    // 獲取集合大小
    fmt.Printf("集合大小: %d\n", len(set))
    
    // 遍歷集合
    fmt.Print("集合元素: ")
    for item := range set {
        fmt.Printf("%s ", item)
    }
    fmt.Println()
    
    // 刪除元素
    delete(set, "banana")
    fmt.Printf("刪除 banana 後: %v\n", set)
    
    // 使用 map[T]struct{} 節省內存
    efficientSet := make(map[string]struct{})
    efficientSet["item1"] = struct{}{}
    efficientSet["item2"] = struct{}{}
    
    // 檢查存在性
    _, exists := efficientSet["item1"]
    fmt.Printf("item1 存在: %t\n", exists)
}
```

### 映射的映射（嵌套映射）

```go
func demonstrateNestedMaps() {
    fmt.Println("\n--- 嵌套映射 ---")
    
    // 二維映射：學生 -> 科目 -> 分數
    grades := map[string]map[string]int{
        "Alice": {
            "Math":    95,
            "English": 87,
            "Science": 92,
        },
        "Bob": {
            "Math":    78,
            "English": 85,
            "Science": 88,
        },
    }
    
    // 訪問嵌套值
    fmt.Printf("Alice 的數學成績: %d\n", grades["Alice"]["Math"])
    
    // 安全地訪問可能不存在的鍵
    if studentGrades, exists := grades["Charlie"]; exists {
        if mathGrade, exists := studentGrades["Math"]; exists {
            fmt.Printf("Charlie 的數學成績: %d\n", mathGrade)
        }
    } else {
        fmt.Println("Charlie 不存在")
    }
    
    // 添加新學生
    grades["Charlie"] = make(map[string]int)
    grades["Charlie"]["Math"] = 90
    grades["Charlie"]["English"] = 88
    
    // 遍歷嵌套映射
    fmt.Println("所有學生成績:")
    for student, subjects := range grades {
        fmt.Printf("  %s:\n", student)
        for subject, grade := range subjects {
            fmt.Printf("    %s: %d\n", subject, grade)
        }
    }
}
```

## 📊 映射與切片的結合

### 映射的切片

```go
func demonstrateMapsWithSlices() {
    fmt.Println("\n--- 映射與切片結合 ---")
    
    // 映射的切片
    people := []map[string]interface{}{
        {"name": "Alice", "age": 30, "city": "New York"},
        {"name": "Bob", "age": 25, "city": "San Francisco"},
        {"name": "Charlie", "age": 35, "city": "Chicago"},
    }
    
    fmt.Println("人員列表:")
    for i, person := range people {
        fmt.Printf("  %d: %v\n", i, person)
    }
    
    // 切片作為映射的值
    groups := map[string][]string{
        "frontend":  {"Alice", "Bob"},
        "backend":   {"Charlie", "David"},
        "devops":    {"Eve"},
        "fullstack": {"Frank", "Grace"},
    }
    
    fmt.Println("團隊分組:")
    for team, members := range groups {
        fmt.Printf("  %s: %v\n", team, members)
    }
    
    // 向團隊添加成員
    groups["frontend"] = append(groups["frontend"], "Helen")
    fmt.Printf("添加成員後的前端團隊: %v\n", groups["frontend"])
    
    // 統計每個團隊的人數
    fmt.Println("團隊人數統計:")
    for team, members := range groups {
        fmt.Printf("  %s: %d 人\n", team, len(members))
    }
}
```

## 🔒 映射的併發安全

### 併發問題

```go
import "sync"

func demonstrateMapConcurrency() {
    fmt.Println("\n--- 映射併發安全 ---")
    
    // 普通映射不是線程安全的
    unsafeMap := make(map[int]int)
    
    // 使用 sync.Map 實現線程安全
    var safeMap sync.Map
    
    // 存儲值
    safeMap.Store(1, "one")
    safeMap.Store(2, "two")
    safeMap.Store(3, "three")
    
    // 加載值
    if value, ok := safeMap.Load(1); ok {
        fmt.Printf("sync.Map 值: %v\n", value)
    }
    
    // 刪除值
    safeMap.Delete(2)
    
    // 遍歷 sync.Map
    fmt.Println("sync.Map 內容:")
    safeMap.Range(func(key, value interface{}) bool {
        fmt.Printf("  %v: %v\n", key, value)
        return true // 繼續遍歷
    })
    
    // 使用互斥鎖保護普通映射
    type SafeCounter struct {
        mu    sync.Mutex
        count map[string]int
    }
    
    counter := SafeCounter{count: make(map[string]int)}
    
    // 安全的增加計數
    increment := func(key string) {
        counter.mu.Lock()
        defer counter.mu.Unlock()
        counter.count[key]++
    }
    
    // 安全的獲取計數
    getValue := func(key string) int {
        counter.mu.Lock()
        defer counter.mu.Unlock()
        return counter.count[key]
    }
    
    increment("clicks")
    increment("clicks")
    fmt.Printf("安全計數器 clicks: %d\n", getValue("clicks"))
}
```

## 💡 映射的最佳實踐

### 1. 初始化檢查

```go
func demonstrateMapBestPractices() {
    fmt.Println("\n--- 映射最佳實踐 ---")
    
    // 好的實踐：檢查映射是否為 nil
    var m map[string]int
    
    if m == nil {
        m = make(map[string]int)
    }
    m["key"] = 1
    
    // 更好的實踐：使用短聲明
    m2 := make(map[string]int)
    m2["key"] = 1
    
    // 最佳實踐：如果知道大致大小，預分配容量
    largeMap := make(map[string]int, 1000) // 預分配容量
    _ = largeMap
    
    fmt.Println("映射初始化最佳實踐演示完成")
}
```

### 2. 安全的鍵訪問

```go
func demonstrateSafeKeyAccess() {
    fmt.Println("\n--- 安全的鍵訪問 ---")
    
    m := map[string]int{
        "existing": 42,
    }
    
    // 不安全的訪問
    value := m["nonexistent"] // 返回零值，可能誤導
    fmt.Printf("不安全訪問: %d\n", value)
    
    // 安全的訪問
    if value, ok := m["existing"]; ok {
        fmt.Printf("安全訪問存在的鍵: %d\n", value)
    }
    
    if value, ok := m["nonexistent"]; ok {
        fmt.Printf("不會執行: %d\n", value)
    } else {
        fmt.Println("安全訪問：鍵不存在")
    }
    
    // 使用輔助函數
    getValue := func(m map[string]int, key string, defaultValue int) int {
        if value, ok := m[key]; ok {
            return value
        }
        return defaultValue
    }
    
    fmt.Printf("使用默認值: %d\n", getValue(m, "nonexistent", -1))
}
```

### 3. 映射的複製

```go
func demonstrateMapCopy() {
    fmt.Println("\n--- 映射複製 ---")
    
    original := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    // 淺複製
    shallow := make(map[string]int)
    for k, v := range original {
        shallow[k] = v
    }
    
    // 修改副本不影響原映射
    shallow["d"] = 4
    delete(shallow, "a")
    
    fmt.Printf("原映射: %v\n", original)
    fmt.Printf("淺複製: %v\n", shallow)
    
    // 對於嵌套映射的深複製
    nested := map[string]map[string]int{
        "group1": {"a": 1, "b": 2},
        "group2": {"c": 3, "d": 4},
    }
    
    deepCopy := make(map[string]map[string]int)
    for k, v := range nested {
        deepCopy[k] = make(map[string]int)
        for k2, v2 := range v {
            deepCopy[k][k2] = v2
        }
    }
    
    // 修改深複製不影響原映射
    deepCopy["group1"]["e"] = 5
    delete(deepCopy["group2"], "c")
    
    fmt.Printf("原嵌套映射: %v\n", nested)
    fmt.Printf("深複製: %v\n", deepCopy)
}
```

## 🎯 映射的實際應用

### 緩存實現

```go
type Cache struct {
    data map[string]interface{}
    mu   sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.data, key)
}

func (c *Cache) Size() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return len(c.data)
}
```

### 配置管理

```go
type Config struct {
    settings map[string]interface{}
    mu       sync.RWMutex
}

func NewConfig() *Config {
    return &Config{
        settings: make(map[string]interface{}),
    }
}

func (c *Config) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.settings[key] = value
}

func (c *Config) GetString(key string, defaultValue string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    if value, ok := c.settings[key]; ok {
        if str, ok := value.(string); ok {
            return str
        }
    }
    return defaultValue
}

func (c *Config) GetInt(key string, defaultValue int) int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    if value, ok := c.settings[key]; ok {
        if num, ok := value.(int); ok {
            return num
        }
    }
    return defaultValue
}
```

## 🎯 本章練習

1. 實現單詞頻率統計器
2. 創建學生成績管理系統
3. 實現 LRU 緩存
4. 創建配置文件解析器

---

**下一章：[接口](../11-interfaces/)**