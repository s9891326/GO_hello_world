# 第二章：變數和常數

## 🎯 學習目標

- 理解變數的概念和用途
- 學會聲明和初始化變數
- 掌握不同的變數聲明方式
- 了解常數的特點和使用
- 理解作用域和生命週期
- 學會命名規範和最佳實踐

## 📖 變數基礎

### 什麼是變數？

變數是用來存儲數據的命名存儲位置。在 Go 中，每個變數都有特定的類型，用來決定變數的內存大小和布局。

### 變數聲明語法

Go 提供了多種聲明變數的方式：

#### 1. 標準聲明
```go
var variableName dataType
var name string
var age int
var isStudent bool
```

#### 2. 聲明並初始化
```go
var variableName dataType = value
var name string = "張三"
var age int = 25
var isStudent bool = true
```

#### 3. 類型推導
```go
var variableName = value
var name = "張三"        // string 類型
var age = 25            // int 類型
var isStudent = true    // bool 類型
```

#### 4. 短變數聲明（最常用）
```go
variableName := value
name := "張三"
age := 25
isStudent := true
```

### 多變數聲明

#### 同類型多變數
```go
var a, b, c int
var x, y, z = 1, 2, 3
a, b, c := 1, 2, 3
```

#### 不同類型多變數
```go
var (
    name     string = "張三"
    age      int    = 25
    isActive bool   = true
)
```

## 🔄 零值

在 Go 中，聲明但未初始化的變數會被賦予零值：

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
var slice []int  // nil
var m map[string]int // nil
```

## 📋 常數

### 常數聲明

常數是固定不變的值，在編譯時確定：

```go
const pi = 3.14159
const greeting = "Hello, World!"
const maxUsers = 100
```

### 常數組
```go
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusInternalServerError = 500
)
```

### iota 枚舉器

`iota` 是 Go 的常數生成器：

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### iota 的高級用法
1024 512 256 128 64 32 16 8 4 2 1
```go
const (
    _  = iota        // 跳過 0
    KB = 1 << (10 * iota) // 1024
    MB               // 1024 * 1024
    GB               // 1024 * 1024 * 1024
)
```

## 🌍 作用域

### 包級別變數
```go
package main

var globalVar = "我是全局變數"

func main() {
    fmt.Println(globalVar) // 可以訪問
}
```

### 函數級別變數
```go
func main() {
    var localVar = "我是局部變數"
    fmt.Println(localVar)
} // localVar 在這裡被銷毀
```

### 塊級別作用域
```go
func main() {
    if true {
        var blockVar = "我是塊級變數"
        fmt.Println(blockVar)
    }
    // fmt.Println(blockVar) // 錯誤：無法訪問
}
```

## 📝 命名規範

### 變數命名規則

1. **必須以字母或下劃線開頭**
2. **可以包含字母、數字、下劃線**
3. **區分大小寫**
4. **不能使用關鍵字**

### 命名風格

```go
// 好的命名
userName := "john"
userAge := 25
isActive := true
maxRetryCount := 3

// 不好的命名
u := "john"           // 太短，不清楚
user_name := "john"   // Go 不推薦下劃線
UserName := "john"    // 除非是導出變數，否則不要大寫開頭
```

### 導出和未導出

```go
// 大寫開頭 = 導出（public）
var ExportedVar = "其他包可以訪問"

// 小寫開頭 = 未導出（private）
var unexportedVar = "只有本包可以訪問"
```

## 💡 最佳實踐

### 1. 選擇合適的聲明方式

```go
// 零值初始化
var count int

// 明確類型
var timeout time.Duration = 30 * time.Second

// 類型推導
name := "張三"
```

### 2. 變數分組

```go
var (
    host     = "localhost"
    port     = 8080
    database = "myapp"
)
```

### 3. 常數使用

```go
const (
    DefaultTimeout = 30 * time.Second
    MaxRetries     = 3
    BufferSize     = 1024
)
```

## ⚠️ 常見錯誤

### 1. 未使用的變數
```go
func main() {
    name := "張三" // 錯誤：declared and not used
    age := 25
    fmt.Println(age)
}
```

### 2. 重複聲明
```go
func main() {
    var name string
    var name string // 錯誤：重複聲明
}
```

### 3. 短聲明的限制
```go
var name string
func main() {
    name := "張三" // 這是新的局部變數，不是賦值
    fmt.Println(name)
}
```

## 🔧 工具使用

### 1. 查看變數類型
```go
import "reflect"

name := "張三"
fmt.Printf("類型：%T，值：%v\n", name, name)
fmt.Printf("類型：%s\n", reflect.TypeOf(name))
```

### 2. 格式化輸出變數
```go
name := "張三"
age := 25

fmt.Printf("姓名：%s，年齡：%d\n", name, age)
fmt.Printf("變數：%+v\n", struct{Name string; Age int}{name, age})
```

## 🎯 本章練習

1. 聲明不同類型的變數並初始化
2. 使用 iota 創建枚舉常數
3. 練習變數作用域
4. 實現一個簡單的配置管理

---

**下一章：[數據類型](../03-data-types/)**