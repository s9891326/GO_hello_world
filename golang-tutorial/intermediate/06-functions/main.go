package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== Go 函數示例 ===")
	
	// 1. 基本函數
	demonstrateBasicFunctions()
	
	// 2. 參數傳遞
	demonstrateParameters()
	
	// 3. 多返回值和錯誤處理
	demonstrateReturnsAndErrors()
	
	// 4. 函數作為值
	demonstrateFunctionValues()
	
	// 5. 匿名函數和閉包
	demonstrateClosures()
	
	// 6. 遞歸函數
	demonstrateRecursion()
	
	// 7. 實際應用示例
	demonstrateRealWorldExamples()
}

// =========================
// 基本函數演示
// =========================

func demonstrateBasicFunctions() {
	fmt.Println("\n--- 基本函數演示 ---")
	
	// 無參數無返回值
	sayHello()
	
	// 有參數無返回值
	greetUser("Alice")
	
	// 有參數有返回值
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)
	
	// 多個參數
	area, perimeter := calculateRectangle(5.0, 3.0)
	fmt.Printf("矩形 (5×3): 面積=%.1f, 周長=%.1f\n", area, perimeter)
	
	// 可變參數
	total := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total)
}

func sayHello() {
	fmt.Println("Hello, World!")
}

func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func calculateRectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 命名返回值可以直接返回
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// =========================
// 參數傳遞演示
// =========================

func demonstrateParameters() {
	fmt.Println("\n--- 參數傳遞演示 ---")
	
	// 值傳遞
	num := 10
	fmt.Printf("原始值: %d\n", num)
	doubled := doubleValue(num)
	fmt.Printf("值傳遞後 - 原始值: %d, 返回值: %d\n", num, doubled)
	
	// 指針傳遞
	doubleByPointer(&num)
	fmt.Printf("指針傳遞後 - 修改後的值: %d\n", num)
	
	// 切片傳遞（引用類型）
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片: %v\n", numbers)
	doubleSlice(numbers)
	fmt.Printf("切片傳遞後: %v\n", numbers)
	
	// 結構體傳遞
	person := Person{Name: "Bob", Age: 25}
	fmt.Printf("原始結構體: %+v\n", person)
	
	// 值傳遞結構體
	aged1 := agePersonByValue(person)
	fmt.Printf("值傳遞後 - 原始: %+v, 返回: %+v\n", person, aged1)
	
	// 指針傳遞結構體
	agePersonByPointer(&person)
	fmt.Printf("指針傳遞後: %+v\n", person)
}

func doubleValue(x int) int {
	x = x * 2
	return x
}

func doubleByPointer(x *int) {
	*x = *x * 2
}

func doubleSlice(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}

type Person struct {
	Name string
	Age  int
}

func agePersonByValue(p Person) Person {
	p.Age++
	return p
}

func agePersonByPointer(p *Person) {
	p.Age++
}

// =========================
// 多返回值和錯誤處理
// =========================

func demonstrateReturnsAndErrors() {
	fmt.Println("\n--- 多返回值和錯誤處理 ---")
	
	// 多返回值
	quotient, remainder := divideWithRemainder(17, 5)
	fmt.Printf("17 ÷ 5 = %d 餘 %d\n", quotient, remainder)
	
	// 忽略部分返回值
	_, onlyRemainder := divideWithRemainder(20, 3)
	fmt.Printf("20 除以 3 的餘數: %d\n", onlyRemainder)
	
	// 錯誤處理
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}
	
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("錯誤: %v\n", err)
	}
	
	// 自定義錯誤處理
	if err := validateEmail("invalid-email"); err != nil {
		fmt.Printf("郵箱驗證失敗: %v\n", err)
	}
	
	if err := validateEmail("user@example.com"); err != nil {
		fmt.Printf("郵箱驗證失敗: %v\n", err)
	} else {
		fmt.Println("郵箱驗證成功")
	}
}

func divideWithRemainder(dividend, divisor int) (int, int) {
	return dividend / divisor, dividend % divisor
}

func safeDivide(a, b float64) (float64, error) {
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
	return fmt.Sprintf("驗證錯誤 [%s]: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return ValidationError{"email", "郵箱不能為空"}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{"email", "郵箱格式不正確，缺少 @ 符號"}
	}
	if !strings.Contains(email, ".") {
		return ValidationError{"email", "郵箱格式不正確，缺少域名"}
	}
	return nil
}

// =========================
// 函數作為值
// =========================

// 定義函數類型
type Operation func(int, int) int
type StringProcessor func(string) string

func demonstrateFunctionValues() {
	fmt.Println("\n--- 函數作為值演示 ---")
	
	// 函數變數
	var op Operation
	
	op = addOp
	fmt.Printf("使用加法: 5 + 3 = %d\n", op(5, 3))
	
	op = multiplyOp
	fmt.Printf("使用乘法: 5 × 3 = %d\n", op(5, 3))
	
	// 函數切片
	operations := []Operation{addOp, subtractOp, multiplyOp}
	symbols := []string{"+", "-", "×"}
	
	a, b := 8, 3
	for i, operation := range operations {
		result := operation(a, b)
		fmt.Printf("%d %s %d = %d\n", a, symbols[i], b, result)
	}
	
	// 高階函數
	result := applyOperation(10, 5, subtractOp)
	fmt.Printf("高階函數結果: %d\n", result)
	
	// 函數工廠
	calculator := createCalculator("multiply")
	fmt.Printf("函數工廠結果: %d\n", calculator(6, 7))
	
	// 字符串處理函數
	processors := []StringProcessor{
		strings.ToUpper,
		strings.ToLower,
		strings.TrimSpace,
		func(s string) string { return strings.ReplaceAll(s, " ", "_") },
	}
	
	text := "  Hello World  "
	fmt.Printf("原始文本: '%s'\n", text)
	for i, processor := range processors {
		processed := processor(text)
		fmt.Printf("處理器 %d: '%s'\n", i+1, processed)
	}
}

func addOp(a, b int) int      { return a + b }
func subtractOp(a, b int) int { return a - b }
func multiplyOp(a, b int) int { return a * b }

func applyOperation(a, b int, op Operation) int {
	return op(a, b)
}

func createCalculator(opType string) Operation {
	switch opType {
	case "add":
		return addOp
	case "subtract":
		return subtractOp
	case "multiply":
		return multiplyOp
	default:
		return func(a, b int) int { return 0 }
	}
}

// =========================
// 匿名函數和閉包
// =========================

func demonstrateClosures() {
	fmt.Println("\n--- 匿名函數和閉包演示 ---")
	
	// 立即執行的匿名函數
	result := func(x, y int) int {
		return x*x + y*y
	}(3, 4)
	fmt.Printf("立即執行匿名函數: 3² + 4² = %d\n", result)
	
	// 閉包 - 計數器
	counter1 := createCounter()
	counter2 := createCounter()
	
	fmt.Printf("counter1: %d\n", counter1()) // 1
	fmt.Printf("counter1: %d\n", counter1()) // 2
	fmt.Printf("counter2: %d\n", counter2()) // 1 (獨立的計數器)
	fmt.Printf("counter1: %d\n", counter1()) // 3
	
	// 閉包 - 累加器
	add10 := createAdder(10)
	add100 := createAdder(100)
	
	fmt.Printf("add10(5) = %d\n", add10(5))   // 15
	fmt.Printf("add100(5) = %d\n", add100(5)) // 105
	
	// 閉包 - 狀態維護
	bank := createBankAccount(1000)
	
	fmt.Printf("餘額: %.2f\n", bank("balance", 0))
	fmt.Printf("存款 200: %.2f\n", bank("deposit", 200))
	fmt.Printf("提款 150: %.2f\n", bank("withdraw", 150))
	fmt.Printf("提款 2000: %.2f\n", bank("withdraw", 2000)) // 餘額不足
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func createBankAccount(initialBalance float64) func(string, float64) float64 {
	balance := initialBalance
	
	return func(operation string, amount float64) float64 {
		switch operation {
		case "deposit":
			balance += amount
		case "withdraw":
			if amount <= balance {
				balance -= amount
			} else {
				fmt.Printf("餘額不足，無法提款 %.2f\n", amount)
			}
		case "balance":
			// 只查詢餘額
		}
		return balance
	}
}

// =========================
// 遞歸函數
// =========================

func demonstrateRecursion() {
	fmt.Println("\n--- 遞歸函數演示 ---")
	
	// 階乘
	fmt.Println("階乘計算:")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
	
	// 斐波那契數列
	fmt.Println("\n斐波那契數列:")
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}
	
	// 最大公約數
	fmt.Printf("\ngcd(48, 18) = %d\n", gcd(48, 18))
	fmt.Printf("gcd(100, 25) = %d\n", gcd(100, 25))
	
	// 樹狀遍歷
	fmt.Println("\n目錄結構:")
	printTree(0, "項目根目錄")
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func printTree(level int, name string) {
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s\n", indent, name)
	
	if level < 2 { // 限制遞歸深度
		subItems := []string{"src/", "tests/", "docs/"}
		for _, item := range subItems {
			printTree(level+1, item)
		}
	}
}

// =========================
// 實際應用示例
// =========================

func demonstrateRealWorldExamples() {
	fmt.Println("\n--- 實際應用示例 ---")
	
	// 1. 數據處理管道
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	even := filter(numbers, isEven)
	squared := mapFunc(even, square)
	sum := reduce(squared, 0, addOp)
	
	fmt.Printf("原始數據: %v\n", numbers)
	fmt.Printf("偶數: %v\n", even)
	fmt.Printf("平方: %v\n", squared)
	fmt.Printf("求和: %d\n", sum)
	
	// 2. 配置驗證器
	config := map[string]interface{}{
		"host":    "localhost",
		"port":    8080,
		"timeout": 30,
		"debug":   true,
	}
	
	validators := []func(map[string]interface{}) error{
		validateHost,
		validatePort,
		validateTimeout,
	}
	
	fmt.Println("\n配置驗證:")
	for _, validator := range validators {
		if err := validator(config); err != nil {
			fmt.Printf("驗證失敗: %v\n", err)
		}
	}
	fmt.Println("所有配置驗證通過")
}

// 函數式編程風格的輔助函數
func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func mapFunc(slice []int, mapper func(int) int) []int {
	result := make([]int, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

func reduce(slice []int, initial int, reducer func(int, int) int) int {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func square(n int) int {
	return n * n
}

// 配置驗證函數
func validateHost(config map[string]interface{}) error {
	host, ok := config["host"].(string)
	if !ok || len(host) == 0 {
		return errors.New("host 配置無效")
	}
	return nil
}

func validatePort(config map[string]interface{}) error {
	port, ok := config["port"].(int)
	if !ok || port <= 0 || port > 65535 {
		return errors.New("port 配置無效")
	}
	return nil
}

func validateTimeout(config map[string]interface{}) error {
	timeout, ok := config["timeout"].(int)
	if !ok || timeout <= 0 {
		return errors.New("timeout 配置無效")
	}
	return nil
}