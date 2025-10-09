# 第五章：流程控制

## 🎯 學習目標

- 掌握條件語句的使用（if/else）
- 學會循環語句的各種形式（for）
- 理解選擇語句的應用（switch）
- 了解跳轉語句（break、continue、goto）
- 學會控制流程的最佳實踐
- 掌握錯誤處理的基本模式

## 🔀 流程控制概覽

程式的執行流程控制是程式設計的核心概念。Go 語言提供了簡潔而強大的流程控制語句：

```
Go 流程控制語句
├── 條件語句
│   ├── if 語句
│   ├── if-else 語句
│   └── if-else if-else 語句
├── 循環語句
│   ├── for 循環（唯一的循環語句）
│   ├── for-range 循環
│   └── 無限循環
├── 選擇語句
│   ├── switch 語句
│   ├── type switch
│   └── select 語句（通道相關）
└── 跳轉語句
    ├── break
    ├── continue
    ├── goto
    └── return
```

## ❓ 條件語句（if）

### 基本 if 語句

```go
package main

import "fmt"

func basicIf() {
    age := 18
    
    if age >= 18 {
        fmt.Println("已成年")
    }
    
    // 條件可以是任何布爾表達式
    score := 85
    if score >= 60 && score <= 100 {
        fmt.Println("成績及格")
    }
}
```

### if-else 語句

```go
func ifElse() {
    temperature := 25
    
    if temperature > 30 {
        fmt.Println("天氣很熱")
    } else {
        fmt.Println("天氣不錯")
    }
    
    // 三元運算符的替代
    var status string
    if temperature > 30 {
        status = "熱"
    } else {
        status = "涼"
    }
    fmt.Printf("天氣狀態: %s\n", status)
}
```

### if-else if-else 語句

```go
func ifElseIf() {
    score := 85
    
    if score >= 90 {
        fmt.Println("優秀")
    } else if score >= 80 {
        fmt.Println("良好")
    } else if score >= 70 {
        fmt.Println("中等")
    } else if score >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }
}
```

### 帶初始化的 if 語句

Go 的 if 語句可以包含一個初始化語句：

```go
func ifWithInit() {
    // 在 if 語句中初始化變數
    if age := calculateAge(1995); age >= 18 {
        fmt.Printf("年齡 %d，已成年\n", age)
    } else {
        fmt.Printf("年齡 %d，未成年\n", age)
    }
    // age 變數在此處不可訪問
    
    // 實際應用：錯誤處理
    if err := validateInput("test@example.com"); err != nil {
        fmt.Printf("驗證失敗: %v\n", err)
        return
    }
    fmt.Println("驗證成功")
}

func calculateAge(birthYear int) int {
    return 2024 - birthYear
}

func validateInput(email string) error {
    if len(email) == 0 {
        return fmt.Errorf("郵箱不能為空")
    }
    if !strings.Contains(email, "@") {
        return fmt.Errorf("郵箱格式不正確")
    }
    return nil
}
```

## 🔄 循環語句（for）

Go 只有 `for` 一種循環語句，但它非常靈活，可以實現其他語言中的各種循環。

### 基本 for 循環

```go
func basicFor() {
    // 標準的三部分 for 循環
    for i := 0; i < 5; i++ {
        fmt.Printf("i = %d\n", i)
    }
    
    // 可以省略初始化
    j := 0
    for ; j < 3; j++ {
        fmt.Printf("j = %d\n", j)
    }
    
    // 可以省略後置語句
    k := 0
    for k < 3 {
        fmt.Printf("k = %d\n", k)
        k++
    }
}
```

### while 風格的 for 循環

```go
func whileStyleFor() {
    i := 0
    for i < 5 {  // 等同於 while (i < 5)
        fmt.Printf("while 風格: i = %d\n", i)
        i++
    }
}
```

### 無限循環

```go
func infiniteLoop() {
    count := 0
    for {  // 無限循環
        fmt.Printf("無限循環: %d\n", count)
        count++
        
        if count >= 3 {
            break  // 跳出循環
        }
    }
}
```

### for-range 循環

`for-range` 用於遍歷數組、切片、映射、字符串等：

```go
func forRange() {
    // 遍歷數組/切片
    numbers := []int{10, 20, 30, 40, 50}
    
    // 獲取索引和值
    for index, value := range numbers {
        fmt.Printf("索引 %d: 值 %d\n", index, value)
    }
    
    // 只要值，忽略索引
    for _, value := range numbers {
        fmt.Printf("值: %d\n", value)
    }
    
    // 只要索引，忽略值
    for index := range numbers {
        fmt.Printf("索引: %d\n", index)
    }
    
    // 遍歷字符串
    text := "Hello, 世界"
    for i, char := range text {
        fmt.Printf("位置 %d: 字符 %c (Unicode: %U)\n", i, char, char)
    }
    
    // 遍歷映射
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Charlie": 35,
    }
    
    for name, age := range ages {
        fmt.Printf("%s 的年齡是 %d\n", name, age)
    }
}
```

## 🔀 選擇語句（switch）

### 基本 switch 語句

```go
func basicSwitch() {
    day := 3
    
    switch day {
    case 1:
        fmt.Println("星期一")
    case 2:
        fmt.Println("星期二")
    case 3:
        fmt.Println("星期三")
    case 4:
        fmt.Println("星期四")
    case 5:
        fmt.Println("星期五")
    case 6, 7:  // 多個值
        fmt.Println("週末")
    default:
        fmt.Println("無效的日期")
    }
}
```

### 帶表達式的 switch

```go
func expressionSwitch() {
    score := 85
    
    switch {  // 沒有表達式的 switch
    case score >= 90:
        fmt.Println("優秀")
    case score >= 80:
        fmt.Println("良好")
    case score >= 70:
        fmt.Println("中等")
    case score >= 60:
        fmt.Println("及格")
    default:
        fmt.Println("不及格")
    }
}
```

### 帶初始化的 switch

```go
func switchWithInit() {
    switch grade := calculateGrade(85); grade {
    case "A":
        fmt.Println("優秀成績！")
    case "B":
        fmt.Println("良好成績！")
    case "C":
        fmt.Println("中等成績")
    case "D":
        fmt.Println("及格成績")
    default:
        fmt.Println("需要努力")
    }
}

func calculateGrade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    } else if score >= 60 {
        return "D"
    }
    return "F"
}
```

### fallthrough 語句

```go
func fallthroughExample() {
    number := 2
    
    switch number {
    case 1:
        fmt.Println("一")
        fallthrough  // 繼續執行下一個 case
    case 2:
        fmt.Println("二")
        fallthrough
    case 3:
        fmt.Println("三")
    default:
        fmt.Println("其他")
    }
    // 輸出: 二、三
}
```

## 🏃 跳轉語句

### break 語句

```go
func breakExample() {
    // 在循環中使用 break
    for i := 0; i < 10; i++ {
        if i == 5 {
            break  // 跳出循環
        }
        fmt.Printf("i = %d\n", i)
    }
    
    // 在嵌套循環中使用標籤
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break OuterLoop  // 跳出外層循環
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
```

### continue 語句

```go
func continueExample() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue  // 跳過當前迭代
        }
        fmt.Printf("奇數: %d\n", i)
    }
    
    // 在嵌套循環中使用標籤
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue OuterLoop  // 跳到外層循環的下一次迭代
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
```

### goto 語句

```go
func gotoExample() {
    i := 0
    
Loop:
    if i < 5 {
        fmt.Printf("i = %d\n", i)
        i++
        goto Loop
    }
    
    fmt.Println("循環結束")
    
    // goto 用於錯誤處理（不推薦，僅作示例）
    if err := someOperation(); err != nil {
        goto ErrorHandler
    }
    
    fmt.Println("操作成功")
    return
    
ErrorHandler:
    fmt.Println("處理錯誤")
}

func someOperation() error {
    return nil  // 模擬操作
}
```

## 🎯 實際應用示例

### 1. 用戶輸入驗證

```go
func validateUserInput() {
    inputs := []string{"", "test", "test@", "test@example.com"}
    
    for _, input := range inputs {
        fmt.Printf("驗證輸入: '%s' -> ", input)
        
        if len(input) == 0 {
            fmt.Println("錯誤: 輸入不能為空")
            continue
        }
        
        if len(input) < 3 {
            fmt.Println("錯誤: 輸入太短")
            continue
        }
        
        if !strings.Contains(input, "@") {
            fmt.Println("錯誤: 缺少 @ 符號")
            continue
        }
        
        fmt.Println("驗證通過")
    }
}
```

### 2. 菜單驅動程序

```go
func menuDrivenProgram() {
    for {
        fmt.Println("\n=== 主菜單 ===")
        fmt.Println("1. 計算圓面積")
        fmt.Println("2. 計算矩形面積")
        fmt.Println("3. 計算三角形面積")
        fmt.Println("0. 退出")
        fmt.Print("請選擇: ")
        
        var choice int
        fmt.Scanf("%d", &choice)
        
        switch choice {
        case 1:
            calculateCircleArea()
        case 2:
            calculateRectangleArea()
        case 3:
            calculateTriangleArea()
        case 0:
            fmt.Println("再見！")
            return
        default:
            fmt.Println("無效選擇，請重試")
        }
    }
}

func calculateCircleArea() {
    var radius float64
    fmt.Print("請輸入半徑: ")
    fmt.Scanf("%f", &radius)
    
    if radius <= 0 {
        fmt.Println("半徑必須大於 0")
        return
    }
    
    area := 3.14159 * radius * radius
    fmt.Printf("圓的面積: %.2f\n", area)
}

func calculateRectangleArea() {
    var width, height float64
    fmt.Print("請輸入寬度和高度: ")
    fmt.Scanf("%f %f", &width, &height)
    
    if width <= 0 || height <= 0 {
        fmt.Println("寬度和高度必須大於 0")
        return
    }
    
    area := width * height
    fmt.Printf("矩形面積: %.2f\n", area)
}

func calculateTriangleArea() {
    var base, height float64
    fmt.Print("請輸入底邊和高: ")
    fmt.Scanf("%f %f", &base, &height)
    
    if base <= 0 || height <= 0 {
        fmt.Println("底邊和高必須大於 0")
        return
    }
    
    area := 0.5 * base * height
    fmt.Printf("三角形面積: %.2f\n", area)
}
```

## 💡 最佳實踐

### 1. 條件語句

```go
// 好的實踐
if user != nil && user.IsActive {
    // 處理邏輯
}

// 避免深度嵌套
if user == nil {
    return errors.New("用戶不能為空")
}
if !user.IsActive {
    return errors.New("用戶未激活")
}
// 主要邏輯
```

### 2. 循環語句

```go
// 使用有意義的變數名
for userIndex, userData := range users {
    // 而不是 i, v
}

// 避免無限循環
maxRetries := 3
for attempt := 0; attempt < maxRetries; attempt++ {
    if success := tryOperation(); success {
        break
    }
}
```

### 3. Switch 語句

```go
// 使用 switch 而不是長串 if-else
switch userRole {
case "admin":
    handleAdminRequest()
case "user":
    handleUserRequest()
case "guest":
    handleGuestRequest()
default:
    handleUnknownRole()
}
```

## 🎯 本章練習

1. 創建一個猜數字遊戲
2. 實現簡單的計算器菜單
3. 編寫九九乘法表
4. 創建學生成績統計程序

---

**下一章：[函數](../06-functions/)**