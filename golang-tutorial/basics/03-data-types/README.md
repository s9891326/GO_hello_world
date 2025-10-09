# 第三章：數據類型

## 🎯 學習目標

- 掌握 Go 語言的基本數據類型
- 理解數值類型的特點和使用場景
- 學會字符串操作和處理
- 了解布爾類型的應用
- 掌握類型轉換和類型斷言
- 學會自定義類型的創建和使用

## 📊 Go 數據類型概覽

Go 是靜態類型語言，所有變數都有明確的類型。Go 的類型系統包括：

```
Go 數據類型
├── 基本類型 (Basic Types)
│   ├── 數值類型
│   │   ├── 整數類型 (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
│   │   └── 浮點類型 (float32, float64)
│   │   └── 複數類型 (complex64, complex128)
│   ├── 字符串類型 (string)
│   └── 布爾類型 (bool)
├── 復合類型 (Composite Types)
│   ├── 數組 (Array)
│   ├── 切片 (Slice)
│   ├── 映射 (Map)
│   ├── 結構體 (Struct)
│   ├── 通道 (Channel)
│   ├── 函數 (Function)
│   └── 接口 (Interface)
└── 指針類型 (Pointer)
```

## 🔢 數值類型

### 整數類型

#### 有符號整數
```go
var i8 int8 = -128     // -128 到 127
var i16 int16 = -32768 // -32768 到 32767
var i32 int32 = -2147483648
var i64 int64 = -9223372036854775808
var i int = -42        // 平台相關 (32位或64位)
```

#### 無符號整數
```go
var ui8 uint8 = 255     // 0 到 255 (等同於 byte)
var ui16 uint16 = 65535 // 0 到 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615
var ui uint = 42        // 平台相關 (32位或64位)
```

#### 特殊整數類型
```go
var b byte = 255        // uint8 的別名
var r rune = '中'       // int32 的別名，用於 Unicode 字符
var ptr uintptr         // 存儲指針的整數類型
```

### 浮點類型

```go
var f32 float32 = 3.14159
var f64 float64 = 3.141592653589793
```

### 複數類型

```go
var c64 complex64 = 3 + 4i
var c128 complex128 = 5 + 12i

// 複數操作
real := real(c128)    // 實部
imag := imag(c128)    // 虛部
```

## 📝 字符串類型

### 字符串基礎

```go
var s1 string = "Hello, World!"
s2 := "你好，世界！"
s3 := `這是一個
多行字符串`
```

### 字符串操作

```go
package main

import (
    "fmt"
    "strings"
    "strconv"
)

func stringOperations() {
    s := "Hello, Go!"
    
    // 長度
    fmt.Printf("長度: %d\n", len(s))
    
    // 索引訪問
    fmt.Printf("第一個字符: %c\n", s[0])
    
    // 子字符串
    fmt.Printf("子字符串: %s\n", s[0:5])
    
    // 字符串連接
    s1 := "Hello"
    s2 := "World"
    result := s1 + ", " + s2 + "!"
    
    // 使用 strings 包
    words := strings.Split("apple,banana,orange", ",")
    joined := strings.Join(words, " | ")
    
    // 字符串轉換
    num := 42
    str := strconv.Itoa(num)        // 整數轉字符串
    parsed, _ := strconv.Atoi(str)  // 字符串轉整數
}
```

### Unicode 和 Rune

```go
s := "Hello 世界"
fmt.Printf("字節長度: %d\n", len(s))
fmt.Printf("字符長度: %d\n", len([]rune(s)))

// 遍歷字符串
for i, r := range s {
    fmt.Printf("位置 %d: %c (Unicode: %U)\n", i, r, r)
}
```

## ✅ 布爾類型

```go
var isTrue bool = true
var isFalse bool = false
var isZero bool         // 零值是 false

// 布爾運算
result := true && false  // false
result = true || false   // true
result = !true          // false
```

## 🔄 類型轉換

### 顯式類型轉換

Go 不支援隱式類型轉換，必須顯式轉換：

```go
var i int = 42
var f float64 = float64(i)  // 正確
var u uint = uint(i)        // 正確

// var f2 float64 = i       // 錯誤：無法隱式轉換
```

### 字符串轉換

```go
import "strconv"

// 數字轉字符串
i := 42
s := strconv.Itoa(i)
f := 3.14
fs := strconv.FormatFloat(f, 'f', 2, 64)

// 字符串轉數字
s = "42"
i, err := strconv.Atoi(s)
fs = "3.14"
f, err = strconv.ParseFloat(fs, 64)
```

## 🏷️ 自定義類型

### Type 定義

```go
// 基於現有類型創建新類型
type UserID int
type UserName string
type Temperature float64

var uid UserID = 12345
var name UserName = "Alice"
var temp Temperature = 36.5
```

### 類型方法

```go
type Temperature float64

// 為自定義類型添加方法
func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", t)
}
```

## 📏 類型信息

### 獲取類型信息

```go
import (
    "fmt"
    "reflect"
)

func typeInfo() {
    var i int = 42
    var f float64 = 3.14
    var s string = "hello"
    
    // 使用 %T 格式符
    fmt.Printf("i 的類型: %T\n", i)
    fmt.Printf("f 的類型: %T\n", f)
    fmt.Printf("s 的類型: %T\n", s)
    
    // 使用 reflect 包
    fmt.Printf("i 的類型: %s\n", reflect.TypeOf(i))
    fmt.Printf("i 的值: %v\n", reflect.ValueOf(i))
}
```

## 📐 類型大小

```go
import "unsafe"

func typeSizes() {
    fmt.Printf("bool: %d bytes\n", unsafe.Sizeof(bool(true)))
    fmt.Printf("int8: %d bytes\n", unsafe.Sizeof(int8(0)))
    fmt.Printf("int16: %d bytes\n", unsafe.Sizeof(int16(0)))
    fmt.Printf("int32: %d bytes\n", unsafe.Sizeof(int32(0)))
    fmt.Printf("int64: %d bytes\n", unsafe.Sizeof(int64(0)))
    fmt.Printf("float32: %d bytes\n", unsafe.Sizeof(float32(0)))
    fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(float64(0)))
    fmt.Printf("string: %d bytes\n", unsafe.Sizeof(string("")))
}
```

## ⚠️ 常見陷阱

### 1. 整數溢出

```go
var i8 int8 = 127
i8++  // 溢出，變成 -128
```

### 2. 浮點精度問題

```go
var f float32 = 0.1 + 0.2
fmt.Printf("%.10f\n", f)  // 可能不等於 0.3
```

### 3. 字符串不可變

```go
s := "hello"
// s[0] = 'H'  // 錯誤：字符串不可修改

// 正確的方式
runes := []rune(s)
runes[0] = 'H'
s = string(runes)
```

## 💡 最佳實踐

### 1. 選擇合適的數值類型

```go
// 一般用途
var count int           // 推薦使用 int
var price float64       // 推薦使用 float64

// 特定用途
var age uint8           // 年齡 (0-255)
var fileSize int64      // 文件大小
var percentage float32  // 百分比
```

### 2. 字符串處理

```go
// 大量字符串拼接使用 strings.Builder
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("hello")
}
result := builder.String()
```

### 3. 類型安全

```go
type UserID int
type ProductID int

var uid UserID = 123
var pid ProductID = 456

// 編譯時捕獲錯誤
// if uid == pid { } // 錯誤：類型不匹配
```

## 🎯 本章練習

1. 創建不同數值類型的變數並進行計算
2. 實現字符串處理功能
3. 創建自定義類型並添加方法
4. 練習類型轉換和錯誤處理

---

**下一章：[運算符](../04-operators/)**