package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// 自定義類型
type UserID int
type Temperature float64
type Status int

// 枚舉常數
const (
	StatusInactive Status = iota
	StatusActive
	StatusSuspended
	StatusDeleted
)

func main() {
	fmt.Println("=== Go 數據類型示例 ===")
	
	// 1. 數值類型演示
	demonstrateNumericTypes()
	
	// 2. 字符串類型演示
	demonstrateStringTypes()
	
	// 3. 布爾類型演示
	demonstrateBooleanTypes()
	
	// 4. 類型轉換演示
	demonstrateTypeConversion()
	
	// 5. 自定義類型演示
	demonstrateCustomTypes()
	
	// 6. 類型信息演示
	demonstrateTypeInfo()
}

func demonstrateNumericTypes() {
	fmt.Println("\n--- 數值類型演示 ---")
	
	// 整數類型
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	
	fmt.Printf("int8 最大值: %d\n", i8)
	fmt.Printf("int16 最大值: %d\n", i16)
	fmt.Printf("int32 最大值: %d\n", i32)
	fmt.Printf("int64 最大值: %d\n", i64)
	fmt.Printf("uint8 最大值: %d\n", ui8)
	fmt.Printf("uint16 最大值: %d\n", ui16)
	
	// 浮點類型
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	
	fmt.Printf("float32: %.5f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)
	
	// 複數類型
	var c64 complex64 = 3 + 4i
	var c128 complex128 = 5 + 12i
	
	fmt.Printf("complex64: %v (實部: %.1f, 虛部: %.1f)\n", c64, real(c64), imag(c64))
	fmt.Printf("complex128: %v (模長: %.2f)\n", c128, math.Abs(real(c128)*real(c128)+imag(c128)*imag(c128)))
	
	// 特殊類型
	var b byte = 255        // uint8 別名
	var r rune = '中'       // int32 別名，Unicode 字符
	
	fmt.Printf("byte: %d (字符: %c)\n", b, b)
	fmt.Printf("rune: %d (字符: %c, Unicode: %U)\n", r, r, r)
}

func demonstrateStringTypes() {
	fmt.Println("\n--- 字符串類型演示 ---")
	
	// 字符串聲明
	s1 := "Hello, Go!"
	s2 := "你好，世界！"
	s3 := `這是一個多行字符串
可以包含換行符
和特殊字符 "quotes"`
	
	fmt.Printf("基本字符串: %s\n", s1)
	fmt.Printf("中文字符串: %s\n", s2)
	fmt.Printf("多行字符串:\n%s\n", s3)
	
	// 字符串操作
	fmt.Printf("s1 長度（字節）: %d\n", len(s1))
	fmt.Printf("s2 長度（字節）: %d\n", len(s2))
	fmt.Printf("s2 長度（字符）: %d\n", len([]rune(s2)))
	
	// 字符串索引和切片
	fmt.Printf("第一個字節: %c\n", s1[0])
	fmt.Printf("前5個字節: %s\n", s1[0:5])
	
	// 字符串拼接
	greeting := s1 + " " + s2
	fmt.Printf("拼接結果: %s\n", greeting)
	
	// 使用 strings 包
	words := []string{"apple", "banana", "orange"}
	joined := strings.Join(words, ", ")
	fmt.Printf("連接數組: %s\n", joined)
	
	split := strings.Split(joined, ", ")
	fmt.Printf("分割字符串: %v\n", split)
	
	// 字符串遍歷
	fmt.Println("字符遍歷:")
	for i, r := range s2 {
		fmt.Printf("  位置 %d: %c (Unicode: %U)\n", i, r, r)
	}
}

func demonstrateBooleanTypes() {
	fmt.Println("\n--- 布爾類型演示 ---")
	
	var isTrue bool = true
	var isFalse bool = false
	var isZero bool // 零值為 false
	
	fmt.Printf("isTrue: %t\n", isTrue)
	fmt.Printf("isFalse: %t\n", isFalse)
	fmt.Printf("isZero: %t\n", isZero)
	
	// 布爾運算
	fmt.Println("布爾運算:")
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("true || false = %t\n", true || false)
	fmt.Printf("!true = %t\n", !true)
	fmt.Printf("!false = %t\n", !false)
	
	// 比較運算產生布爾值
	a, b := 10, 20
	fmt.Printf("%d > %d = %t\n", a, b, a > b)
	fmt.Printf("%d == %d = %t\n", a, b, a == b)
	fmt.Printf("%d != %d = %t\n", a, b, a != b)
}

func demonstrateTypeConversion() {
	fmt.Println("\n--- 類型轉換演示 ---")
	
	// 數值類型轉換
	var i int = 42
	var f float64 = float64(i)
	var ui uint = uint(i)
	
	fmt.Printf("int %d -> float64 %f\n", i, f)
	fmt.Printf("int %d -> uint %d\n", i, ui)
	
	// 注意：可能會丟失精度或溢出
	var bigInt int64 = 1234567890123456789
	var smallInt int32 = int32(bigInt) // 可能溢出
	fmt.Printf("int64 %d -> int32 %d (可能溢出)\n", bigInt, smallInt)
	
	// 字符串轉換
	num := 12345
	str := strconv.Itoa(num)
	fmt.Printf("int %d -> string \"%s\"\n", num, str)
	
	parsed, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("string \"%s\" -> int %d\n", str, parsed)
	}
	
	// 浮點數轉換
	pi := 3.14159
	piStr := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Printf("float64 %f -> string \"%s\"\n", pi, piStr)
	
	parsedFloat, err := strconv.ParseFloat(piStr, 64)
	if err == nil {
		fmt.Printf("string \"%s\" -> float64 %f\n", piStr, parsedFloat)
	}
}

func demonstrateCustomTypes() {
	fmt.Println("\n--- 自定義類型演示 ---")
	
	// 使用自定義類型
	var uid UserID = 12345
	var temp Temperature = 36.5
	var status Status = StatusActive
	
	fmt.Printf("用戶ID: %d\n", uid)
	fmt.Printf("溫度: %s\n", temp.String())
	fmt.Printf("華氏溫度: %.1f°F\n", temp.Fahrenheit())
	fmt.Printf("狀態: %s\n", status.String())
	
	// 類型安全 - 以下會編譯錯誤
	// var pid ProductID = 456
	// if uid == pid { } // 錯誤：不同的自定義類型無法直接比較
}

func demonstrateTypeInfo() {
	fmt.Println("\n--- 類型信息演示 ---")
	
	var i int = 42
	var f float64 = 3.14
	var s string = "hello"
	var b bool = true
	
	// 使用 %T 格式符
	fmt.Printf("變數 i 的類型: %T, 值: %v\n", i, i)
	fmt.Printf("變數 f 的類型: %T, 值: %v\n", f, f)
	fmt.Printf("變數 s 的類型: %T, 值: %v\n", s, s)
	fmt.Printf("變數 b 的類型: %T, 值: %v\n", b, b)
	
	// 使用 reflect 包
	fmt.Println("\n使用 reflect 包:")
	fmt.Printf("i 的類型: %s, 種類: %s\n", reflect.TypeOf(i), reflect.TypeOf(i).Kind())
	fmt.Printf("f 的類型: %s, 種類: %s\n", reflect.TypeOf(f), reflect.TypeOf(f).Kind())
	
	// 類型大小
	fmt.Println("\n類型大小 (字節):")
	fmt.Printf("bool: %d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int: %d\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int8: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("float32: %d\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64: %d\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("string: %d\n", unsafe.Sizeof(string("")))
}

// Temperature 類型的方法
func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", t)
}

// Status 類型的方法
func (s Status) String() string {
	switch s {
	case StatusInactive:
		return "非活躍"
	case StatusActive:
		return "活躍"
	case StatusSuspended:
		return "暫停"
	case StatusDeleted:
		return "已刪除"
	default:
		return "未知狀態"
	}
}