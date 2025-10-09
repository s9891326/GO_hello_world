# 第四章：運算符

## 🎯 學習目標

- 掌握 Go 語言的各種運算符
- 理解運算符的優先級和結合性
- 學會算術、比較、邏輯運算
- 了解位運算的應用場景
- 掌握賦值運算符的使用
- 學會指針運算符的基本操作

## ⚡ 運算符概覽

Go 語言提供了豐富的運算符，可以分為以下幾類：

```
Go 運算符分類
├── 算術運算符 (+, -, *, /, %, ++, --)
├── 比較運算符 (==, !=, <, >, <=, >=)
├── 邏輯運算符 (&&, ||, !)
├── 位運算符 (&, |, ^, <<, >>, &^)
├── 賦值運算符 (=, +=, -=, *=, /=, %=, etc.)
├── 指針運算符 (&, *)
└── 其他運算符 (<-, ...)
```

## 🧮 算術運算符

### 基本算術運算

```go
package main

import "fmt"

func arithmeticOperators() {
    a, b := 10, 3
    
    fmt.Printf("%d + %d = %d\n", a, b, a+b)   // 加法: 13
    fmt.Printf("%d - %d = %d\n", a, b, a-b)   // 減法: 7
    fmt.Printf("%d * %d = %d\n", a, b, a*b)   // 乘法: 30
    fmt.Printf("%d / %d = %d\n", a, b, a/b)   // 除法: 3 (整數除法)
    fmt.Printf("%d %% %d = %d\n", a, b, a%b)  // 取餘: 1
    
    // 浮點數除法
    fa, fb := 10.0, 3.0
    fmt.Printf("%.2f / %.2f = %.2f\n", fa, fb, fa/fb)  // 3.33
}
```

### 自增自減運算符

```go
func incrementDecrement() {
    i := 5
    
    fmt.Printf("初始值: %d\n", i)
    
    i++  // 等同於 i = i + 1
    fmt.Printf("i++ 後: %d\n", i)  // 6
    
    i--  // 等同於 i = i - 1
    fmt.Printf("i-- 後: %d\n", i)  // 5
    
    // 注意：Go 中 ++ 和 -- 是語句，不是表達式
    // j := i++  // 錯誤！
    // ++i       // 錯誤！Go 沒有前置遞增
}
```

## ⚖️ 比較運算符

```go
func comparisonOperators() {
    a, b := 10, 20
    
    fmt.Printf("%d == %d: %t\n", a, b, a == b)  // false
    fmt.Printf("%d != %d: %t\n", a, b, a != b)  // true
    fmt.Printf("%d < %d: %t\n", a, b, a < b)    // true
    fmt.Printf("%d > %d: %t\n", a, b, a > b)    // false
    fmt.Printf("%d <= %d: %t\n", a, b, a <= b)  // true
    fmt.Printf("%d >= %d: %t\n", a, b, a >= b)  // false
    
    // 字符串比較
    str1, str2 := "apple", "banana"
    fmt.Printf("%s < %s: %t\n", str1, str2, str1 < str2)  // true (字典序)
}
```

## 🧠 邏輯運算符

```go
func logicalOperators() {
    p, q := true, false
    
    fmt.Printf("p && q: %t\n", p && q)  // AND: false
    fmt.Printf("p || q: %t\n", p || q)  // OR: true
    fmt.Printf("!p: %t\n", !p)          // NOT: false
    fmt.Printf("!q: %t\n", !q)          // NOT: true
    
    // 短路求值
    age := 25
    hasLicense := true
    
    canDrive := age >= 18 && hasLicense
    fmt.Printf("可以開車: %t\n", canDrive)
    
    // 複合條件
    score := 85
    isPassing := score >= 60 && score <= 100
    fmt.Printf("成績及格: %t\n", isPassing)
}
```

## 🔢 位運算符

位運算符操作數字的二進制位：

```go
func bitwiseOperators() {
    a, b := 12, 25  // 12 = 1100, 25 = 11001
    
    fmt.Printf("a = %d (%08b)\n", a, a)
    fmt.Printf("b = %d (%08b)\n", b, b)
    
    fmt.Printf("a & b = %d (%08b)\n", a&b, a&b)   // AND: 8 (01000)
    fmt.Printf("a | b = %d (%08b)\n", a|b, a|b)   // OR: 29 (11101)
    fmt.Printf("a ^ b = %d (%08b)\n", a^b, a^b)   // XOR: 21 (10101)
    fmt.Printf("^a = %d (%08b)\n", ^a, ^a)        // NOT: -13
    
    // 位移運算
    fmt.Printf("a << 2 = %d (%08b)\n", a<<2, a<<2)  // 左移: 48
    fmt.Printf("a >> 1 = %d (%08b)\n", a>>1, a>>1)  // 右移: 6
    
    // 位清除 (AND NOT)
    fmt.Printf("a &^ b = %d (%08b)\n", a&^b, a&^b)  // 4 (00100)
}
```

### 位運算的實際應用

```go
// 權限管理系統
const (
    PermissionRead   = 1 << iota  // 1 (001)
    PermissionWrite               // 2 (010)
    PermissionExecute             // 4 (100)
)

func permissionExample() {
    // 設置權限
    userPermissions := PermissionRead | PermissionWrite  // 3 (011)
    
    // 檢查權限
    hasRead := (userPermissions & PermissionRead) != 0
    hasWrite := (userPermissions & PermissionWrite) != 0
    hasExecute := (userPermissions & PermissionExecute) != 0
    
    fmt.Printf("讀取權限: %t\n", hasRead)      // true
    fmt.Printf("寫入權限: %t\n", hasWrite)     // true
    fmt.Printf("執行權限: %t\n", hasExecute)   // false
    
    // 添加權限
    userPermissions |= PermissionExecute  // 7 (111)
    
    // 移除權限
    userPermissions &^= PermissionWrite   // 5 (101)
}
```

## ✏️ 賦值運算符

```go
func assignmentOperators() {
    var x int = 10
    
    fmt.Printf("初始值 x = %d\n", x)
    
    x += 5   // x = x + 5
    fmt.Printf("x += 5: %d\n", x)  // 15
    
    x -= 3   // x = x - 3
    fmt.Printf("x -= 3: %d\n", x)  // 12
    
    x *= 2   // x = x * 2
    fmt.Printf("x *= 2: %d\n", x)  // 24
    
    x /= 4   // x = x / 4
    fmt.Printf("x /= 4: %d\n", x)  // 6
    
    x %= 4   // x = x % 4
    fmt.Printf("x %%= 4: %d\n", x) // 2
    
    // 位運算賦值
    x <<= 2  // x = x << 2
    fmt.Printf("x <<= 2: %d\n", x) // 8
    
    x &= 12  // x = x & 12
    fmt.Printf("x &= 12: %d\n", x) // 8
}
```

## 👉 指針運算符

```go
func pointerOperators() {
    var x int = 42
    
    // & 取地址運算符
    ptr := &x
    fmt.Printf("x 的值: %d\n", x)
    fmt.Printf("x 的地址: %p\n", ptr)
    
    // * 解引用運算符
    fmt.Printf("ptr 指向的值: %d\n", *ptr)
    
    // 通過指針修改值
    *ptr = 100
    fmt.Printf("修改後 x 的值: %d\n", x)  // 100
}
```

## 📏 運算符優先級

Go 語言運算符優先級（從高到低）：

```go
// 優先級 5 (最高)
* / % << >> & &^

// 優先級 4
+ - | ^

// 優先級 3
== != < <= > >=

// 優先級 2
&&

// 優先級 1 (最低)
||
```

### 優先級示例

```go
func operatorPrecedence() {
    // 不使用括號
    result1 := 2 + 3 * 4        // 14 (不是 20)
    result2 := 10 > 5 && 3 < 7  // true
    result3 := 1 << 2 + 1       // 8 (不是 6)
    
    fmt.Printf("2 + 3 * 4 = %d\n", result1)
    fmt.Printf("10 > 5 && 3 < 7 = %t\n", result2)
    fmt.Printf("1 << 2 + 1 = %d\n", result3)
    
    // 使用括號明確優先級
    result4 := (2 + 3) * 4      // 20
    result5 := 1 << (2 + 1)     // 8
    
    fmt.Printf("(2 + 3) * 4 = %d\n", result4)
    fmt.Printf("1 << (2 + 1) = %d\n", result5)
}
```

## ⚠️ 常見陷阱和注意事項

### 1. 整數除法

```go
// 整數除法會截斷小數部分
a, b := 5, 2
result := a / b  // 2，不是 2.5

// 如果需要浮點結果
floatResult := float64(a) / float64(b)  // 2.5
```

### 2. 自增自減的限制

```go
i := 5
// j := i++     // 錯誤：++ 是語句，不是表達式
// k := ++i     // 錯誤：Go 沒有前置遞增

// 正確的方式
i++
j := i  // j = 6
```

### 3. 位移運算的注意事項

```go
// 右操作數必須是無符號整數或可以轉換為無符號整數的值
x := 10
n := 2
result := x << n  // 正確

// 小心大的位移值
large := x << 100  // 可能導致未定義行為
```

## 🎯 本章練習

1. 創建計算器程序，實現基本四則運算
2. 實現位運算的權限管理系統
3. 練習運算符優先級的使用
4. 創建一個表達式求值程序

---

**下一章：[流程控制](../05-control-flow/)**