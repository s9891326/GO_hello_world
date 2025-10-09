# 第六章：函數

## 🎯 學習目標

- 理解函數的概念和作用
- 掌握函數的定義和調用
- 學會處理函數參數和返回值
- 了解函數的高級特性
- 掌握錯誤處理的函數模式
- 學會函數的最佳實踐

## 📝 函數基礎

函數是組織代碼的基本單位，用於封裝可重用的功能。Go 語言的函數設計簡潔而強大。

### 函數的基本語法

```go
func functionName(parameters) returnType {
    // 函數體
    return value
}
```

### 最簡單的函數

```go
package main

import "fmt"

// 無參數無返回值的函數
func sayHello() {
    fmt.Println("Hello, World!")
}

// 有參數無返回值的函數
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// 有參數有返回值的函數
func add(a, b int) int {
    return a + b
}

func main() {
    sayHello()                    // 調用無參數函數
    greet("Alice")                // 調用有參數函數
    result := add(3, 5)           // 調用有返回值函數
    fmt.Printf("3 + 5 = %d\n", result)
}
```

## 📥 函數參數

### 基本參數傳遞

```go
// 值傳遞（默認方式）
func doubleValue(x int) int {
    x = x * 2
    return x
}

// 指針傳遞
func doublePointer(x *int) {
    *x = *x * 2
}

func demonstrateParameters() {
    num := 10
    
    // 值傳遞不會改變原始值
    doubled := doubleValue(num)
    fmt.Printf("原始值: %d, 加倍後: %d\n", num, doubled)  // num 仍然是 10
    
    // 指針傳遞會改變原始值
    doublePointer(&num)
    fmt.Printf("通過指針修改後: %d\n", num)  // num 變成 20
}
```

### 多個參數

```go
// 相同類型的參數可以合併聲明
func calculateRectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // 命名返回值可以直接返回
}

// 不同類型的參數
func displayUserInfo(name string, age int, isActive bool) {
    fmt.Printf("用戶: %s, 年齡: %d, 活躍: %t\n", name, age, isActive)
}
```

### 可變參數（Variadic Parameters）

```go
// 可變參數函數
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 混合參數
func formatMessage(template string, args ...interface{}) string {
    return fmt.Sprintf(template, args...)
}

func demonstrateVariadic() {
    // 調用可變參數函數
    fmt.Printf("sum() = %d\n", sum())                    // 0
    fmt.Printf("sum(1) = %d\n", sum(1))                  // 1
    fmt.Printf("sum(1, 2, 3) = %d\n", sum(1, 2, 3))     // 6
    
    // 傳遞切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("sum(slice...) = %d\n", sum(numbers...)) // 15
    
    // 格式化消息
    msg := formatMessage("Hello %s, you are %d years old", "Alice", 25)
    fmt.Println(msg)
}
```

## 📤 函數返回值

### 單個返回值

```go
func square(x int) int {
    return x * x
}

func isEven(x int) bool {
    return x%2 == 0
}
```

### 多個返回值

```go
// 多個返回值
func divideAndRemainder(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

// 命名返回值
func calculateCircle(radius float64) (area, circumference float64) {
    area = 3.14159 * radius * radius
    circumference = 2 * 3.14159 * radius
    return  // 自動返回命名的變數
}

func demonstrateMultipleReturns() {
    q, r := divideAndRemainder(17, 5)
    fmt.Printf("17 ÷ 5 = %d 餘 %d\n", q, r)
    
    area, circ := calculateCircle(5.0)
    fmt.Printf("半徑 5 的圓：面積 %.2f，周長 %.2f\n", area, circ)
    
    // 忽略部分返回值
    _, remainder := divideAndRemainder(20, 3)
    fmt.Printf("20 除以 3 的餘數：%d\n", remainder)
}
```

## ⚠️ 錯誤處理

Go 語言使用顯式錯誤處理，通常函數的最後一個返回值是 error。

### 基本錯誤處理

```go
import "errors"

// 返回錯誤的函數
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除數不能為零")
    }
    return a / b, nil
}

// 自定義錯誤類型
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("驗證錯誤 - %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "年齡不能為負數",
        }
    }
    if age > 150 {
        return ValidationError{
            Field:   "age", 
            Message: "年齡不能超過 150",
        }
    }
    return nil
}

func demonstrateErrorHandling() {
    // 處理除法錯誤
    if result, err := divide(10, 0); err != nil {
        fmt.Printf("錯誤: %v\n", err)
    } else {
        fmt.Printf("結果: %.2f\n", result)
    }
    
    // 處理驗證錯誤
    if err := validateAge(-5); err != nil {
        fmt.Printf("驗證失敗: %v\n", err)
    }
}
```

## 🔧 函數作為值

在 Go 中，函數是一等公民，可以作為值傳遞。

### 函數變數

```go
// 定義函數類型
type MathOperation func(int, int) int

// 函數實現
func add(a, b int) int { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }

func demonstrateFunctionValues() {
    // 函數變數
    var operation MathOperation
    
    operation = add
    fmt.Printf("5 + 3 = %d\n", operation(5, 3))
    
    operation = multiply
    fmt.Printf("5 × 3 = %d\n", operation(5, 3))
    
    // 函數切片
    operations := []MathOperation{add, subtract, multiply}
    symbols := []string{"+", "-", "×"}
    
    for i, op := range operations {
        result := op(8, 3)
        fmt.Printf("8 %s 3 = %d\n", symbols[i], result)
    }
}
```

### 高階函數

```go
// 接受函數作為參數的函數
func applyOperation(a, b int, op MathOperation) int {
    return op(a, b)
}

// 返回函數的函數
func getCalculator(operation string) MathOperation {
    switch operation {
    case "add":
        return add
    case "subtract":
        return subtract
    case "multiply":
        return multiply
    default:
        return func(a, b int) int { return 0 }
    }
}

func demonstrateHigherOrderFunctions() {
    // 函數作為參數
    result := applyOperation(10, 5, subtract)
    fmt.Printf("使用函數參數: %d\n", result)
    
    // 函數作為返回值
    calculator := getCalculator("multiply")
    result = calculator(4, 7)
    fmt.Printf("使用返回的函數: %d\n", result)
}
```

## 🔒 匿名函數和閉包

### 匿名函數

```go
func demonstrateAnonymousFunctions() {
    // 立即執行的匿名函數
    result := func(x, y int) int {
        return x * x + y * y
    }(3, 4)
    fmt.Printf("3² + 4² = %d\n", result)
    
    // 賦值給變數的匿名函數
    square := func(x int) int {
        return x * x
    }
    fmt.Printf("5² = %d\n", square(5))
}
```

### 閉包

```go
// 返回閉包的函數
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 創建累加器
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func demonstrateClosures() {
    // 計數器閉包
    counter1 := makeCounter()
    counter2 := makeCounter()
    
    fmt.Printf("counter1: %d\n", counter1())  // 1
    fmt.Printf("counter1: %d\n", counter1())  // 2
    fmt.Printf("counter2: %d\n", counter2())  // 1 (獨立的計數器)
    fmt.Printf("counter1: %d\n", counter1())  // 3
    
    // 累加器閉包
    add10 := makeAdder(10)
    add100 := makeAdder(100)
    
    fmt.Printf("add10(5) = %d\n", add10(5))    // 15
    fmt.Printf("add100(5) = %d\n", add100(5))  // 105
}
```

## 🔄 遞歸函數

```go
// 階乘計算（遞歸）
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// 斐波那契數列（遞歸）
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// 優化的斐波那契（使用記憶化）
func fibonacciMemo(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
    return memo[n]
}

func demonstrateRecursion() {
    // 階乘
    for i := 0; i <= 5; i++ {
        fmt.Printf("%d! = %d\n", i, factorial(i))
    }
    
    // 斐波那契
    fmt.Println("斐波那契數列:")
    for i := 0; i < 10; i++ {
        fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
    }
    
    // 優化的斐波那契
    memo := make(map[int]int)
    fmt.Printf("F(40) = %d\n", fibonacciMemo(40, memo))
}
```

## 💡 函數最佳實踐

### 1. 函數命名

```go
// 好的函數命名
func calculateTotalPrice(items []Item) float64 { /* ... */ }
func isValidEmail(email string) bool { /* ... */ }
func getUserByID(id int) (*User, error) { /* ... */ }

// 避免的命名
func calc(items []Item) float64 { /* ... */ }        // 太簡短
func doSomething(data interface{}) { /* ... */ }     // 不明確
func getUserByIDAndReturnError(id int) (*User, error) { /* ... */ } // 太冗長
```

### 2. 函數大小

```go
// 保持函數簡短和專注
func processOrder(order Order) error {
    if err := validateOrder(order); err != nil {
        return err
    }
    
    if err := calculatePricing(order); err != nil {
        return err
    }
    
    if err := saveOrder(order); err != nil {
        return err
    }
    
    return sendConfirmation(order)
}
```

### 3. 錯誤處理

```go
// 一致的錯誤處理模式
func processUser(userID int) (*User, error) {
    user, err := getUserByID(userID)
    if err != nil {
        return nil, fmt.Errorf("獲取用戶失敗: %w", err)
    }
    
    if err := validateUser(user); err != nil {
        return nil, fmt.Errorf("用戶驗證失敗: %w", err)
    }
    
    return user, nil
}
```

## 🎯 本章練習

1. 創建數學計算函數庫
2. 實現字符串處理工具函數
3. 編寫遞歸解決問題的函數
4. 創建高階函數和閉包應用

---

**下一章：[結構體](../07-structs/)**