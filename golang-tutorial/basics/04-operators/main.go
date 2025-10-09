package main

import "fmt"

func main() {
	fmt.Println("=== Go 運算符示例 ===")
	
	// 1. 算術運算符
	demonstrateArithmetic()
	
	// 2. 比較運算符
	demonstrateComparison()
	
	// 3. 邏輯運算符
	demonstrateLogical()
	
	// 4. 位運算符
	demonstrateBitwise()
	
	// 5. 賦值運算符
	demonstrateAssignment()
	
	// 6. 指針運算符
	demonstratePointer()
	
	// 7. 運算符優先級
	demonstratePrecedence()
	
	// 8. 實際應用示例
	demonstrateRealWorldExamples()
}

func demonstrateArithmetic() {
	fmt.Println("\n--- 算術運算符 ---")
	
	a, b := 15, 4
	fmt.Printf("a = %d, b = %d\n", a, b)
	
	fmt.Printf("加法: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("減法: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("乘法: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("除法: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("取餘: %d %% %d = %d\n", a, b, a%b)
	
	// 浮點數運算
	fa, fb := 15.0, 4.0
	fmt.Printf("浮點除法: %.2f / %.2f = %.2f\n", fa, fb, fa/fb)
	
	// 自增自減
	x := 10
	fmt.Printf("x 初始值: %d\n", x)
	x++
	fmt.Printf("x++ 後: %d\n", x)
	x--
	fmt.Printf("x-- 後: %d\n", x)
}

func demonstrateComparison() {
	fmt.Println("\n--- 比較運算符 ---")
	
	a, b := 10, 20
	fmt.Printf("a = %d, b = %d\n", a, b)
	
	fmt.Printf("a == b: %t\n", a == b)
	fmt.Printf("a != b: %t\n", a != b)
	fmt.Printf("a < b: %t\n", a < b)
	fmt.Printf("a > b: %t\n", a > b)
	fmt.Printf("a <= b: %t\n", a <= b)
	fmt.Printf("a >= b: %t\n", a >= b)
	
	// 字符串比較
	str1, str2 := "apple", "banana"
	fmt.Printf("\n字符串比較:\n")
	fmt.Printf("\"%s\" < \"%s\": %t\n", str1, str2, str1 < str2)
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str1, str1 == str1)
}

func demonstrateLogical() {
	fmt.Println("\n--- 邏輯運算符 ---")
	
	p, q := true, false
	fmt.Printf("p = %t, q = %t\n", p, q)
	
	fmt.Printf("p && q (AND): %t\n", p && q)
	fmt.Printf("p || q (OR): %t\n", p || q)
	fmt.Printf("!p (NOT): %t\n", !p)
	fmt.Printf("!q (NOT): %t\n", !q)
	
	// 實際應用示例
	age := 25
	hasLicense := true
	hasInsurance := false
	
	fmt.Printf("\n實際應用:\n")
	fmt.Printf("年齡: %d, 有駕照: %t, 有保險: %t\n", age, hasLicense, hasInsurance)
	
	canDrive := age >= 18 && hasLicense
	canDriveCommercial := canDrive && hasInsurance
	
	fmt.Printf("可以開車: %t\n", canDrive)
	fmt.Printf("可以開商用車: %t\n", canDriveCommercial)
}

func demonstrateBitwise() {
	fmt.Println("\n--- 位運算符 ---")
	
	a, b := 12, 25  // 12 = 1100, 25 = 11001
	fmt.Printf("a = %d (%08b)\n", a, a)
	fmt.Printf("b = %d (%08b)\n", b, b)
	
	fmt.Printf("a & b  = %d (%08b)  // AND\n", a&b, a&b)
	fmt.Printf("a | b  = %d (%08b)  // OR\n", a|b, a|b)
	fmt.Printf("a ^ b  = %d (%08b)  // XOR\n", a^b, a^b)
	fmt.Printf("^a     = %d (%08b)  // NOT\n", ^a, uint8(^a))
	
	// 位移運算
	fmt.Printf("\n位移運算:\n")
	fmt.Printf("a << 2 = %d (%08b)  // 左移\n", a<<2, a<<2)
	fmt.Printf("a >> 1 = %d (%08b)  // 右移\n", a>>1, a>>1)
	fmt.Printf("a &^ b = %d (%08b)  // AND NOT\n", a&^b, a&^b)
	
	// 權限管理示例
	demonstratePermissions()
}

// 權限管理系統
const (
	PermissionRead    = 1 << iota // 1 (001)
	PermissionWrite               // 2 (010)
	PermissionExecute             // 4 (100)
)

func demonstratePermissions() {
	fmt.Println("\n權限管理系統:")
	fmt.Printf("讀取權限: %d (%03b)\n", PermissionRead, PermissionRead)
	fmt.Printf("寫入權限: %d (%03b)\n", PermissionWrite, PermissionWrite)
	fmt.Printf("執行權限: %d (%03b)\n", PermissionExecute, PermissionExecute)
	
	// 設置用戶權限
	userPermissions := PermissionRead | PermissionWrite
	fmt.Printf("\n用戶權限: %d (%03b)\n", userPermissions, userPermissions)
	
	// 檢查權限
	hasRead := (userPermissions & PermissionRead) != 0
	hasWrite := (userPermissions & PermissionWrite) != 0
	hasExecute := (userPermissions & PermissionExecute) != 0
	
	fmt.Printf("有讀取權限: %t\n", hasRead)
	fmt.Printf("有寫入權限: %t\n", hasWrite)
	fmt.Printf("有執行權限: %t\n", hasExecute)
	
	// 添加執行權限
	userPermissions |= PermissionExecute
	fmt.Printf("添加執行權限後: %d (%03b)\n", userPermissions, userPermissions)
	
	// 移除寫入權限
	userPermissions &^= PermissionWrite
	fmt.Printf("移除寫入權限後: %d (%03b)\n", userPermissions, userPermissions)
}

func demonstrateAssignment() {
	fmt.Println("\n--- 賦值運算符 ---")
	
	x := 10
	fmt.Printf("初始值 x = %d\n", x)
	
	x += 5
	fmt.Printf("x += 5: x = %d\n", x)
	
	x -= 3
	fmt.Printf("x -= 3: x = %d\n", x)
	
	x *= 2
	fmt.Printf("x *= 2: x = %d\n", x)
	
	x /= 4
	fmt.Printf("x /= 4: x = %d\n", x)
	
	x %= 3
	fmt.Printf("x %%= 3: x = %d\n", x)
	
	// 位運算賦值
	y := 8
	fmt.Printf("\n位運算賦值 (y = %d):\n", y)
	
	y <<= 1
	fmt.Printf("y <<= 1: y = %d\n", y)
	
	y >>= 2
	fmt.Printf("y >>= 2: y = %d\n", y)
	
	y &= 3
	fmt.Printf("y &= 3: y = %d\n", y)
}

func demonstratePointer() {
	fmt.Println("\n--- 指針運算符 ---")
	
	var x int = 42
	fmt.Printf("變數 x 的值: %d\n", x)
	
	// 取地址運算符 &
	ptr := &x
	fmt.Printf("x 的地址: %p\n", ptr)
	fmt.Printf("指針 ptr 的值: %p\n", ptr)
	
	// 解引用運算符 *
	fmt.Printf("ptr 指向的值: %d\n", *ptr)
	
	// 通過指針修改值
	*ptr = 100
	fmt.Printf("通過指針修改後，x = %d\n", x)
	
	// 指針的指針
	ptrPtr := &ptr
	fmt.Printf("指針的指針: %p\n", ptrPtr)
	fmt.Printf("**ptrPtr = %d\n", **ptrPtr)
}

func demonstratePrecedence() {
	fmt.Println("\n--- 運算符優先級 ---")
	
	// 算術運算符優先級
	result1 := 2 + 3 * 4
	fmt.Printf("2 + 3 * 4 = %d (乘法優先)\n", result1)
	
	result2 := (2 + 3) * 4
	fmt.Printf("(2 + 3) * 4 = %d (括號改變優先級)\n", result2)
	
	// 位移和加法
	result3 := 1 << 2 + 1
	fmt.Printf("1 << 2 + 1 = %d (加法優先於位移)\n", result3)
	
	result4 := 1 << (2 + 1)
	fmt.Printf("1 << (2 + 1) = %d (括號明確優先級)\n", result4)
	
	// 邏輯運算符優先級
	a, b, c := 5, 10, 15
	result5 := a < b && b < c
	fmt.Printf("%d < %d && %d < %d = %t\n", a, b, b, c, result5)
	
	result6 := a < b || b > c && c > a
	fmt.Printf("%d < %d || %d > %d && %d > %d = %t (注意 && 優先於 ||)\n", 
		a, b, b, c, c, a, result6)
}

func demonstrateRealWorldExamples() {
	fmt.Println("\n--- 實際應用示例 ---")
	
	// 1. 計算器
	fmt.Println("1. 簡單計算器:")
	calculateAndDisplay(10, 5, "+")
	calculateAndDisplay(10, 5, "-")
	calculateAndDisplay(10, 5, "*")
	calculateAndDisplay(10, 5, "/")
	calculateAndDisplay(10, 5, "%")
	
	// 2. 條件判斷
	fmt.Println("\n2. 成績評定:")
	gradeStudent(95)
	gradeStudent(75)
	gradeStudent(55)
	
	// 3. 狀態管理
	fmt.Println("\n3. 設備狀態管理:")
	deviceStatus := 0
	deviceStatus = setDeviceState(deviceStatus, "power", true)
	deviceStatus = setDeviceState(deviceStatus, "network", true)
	deviceStatus = setDeviceState(deviceStatus, "bluetooth", false)
	displayDeviceStatus(deviceStatus)
}

func calculateAndDisplay(a, b int, operator string) {
	var result int
	var valid bool = true
	
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Printf("錯誤：除數不能為零\n")
			valid = false
		}
	case "%":
		if b != 0 {
			result = a % b
		} else {
			fmt.Printf("錯誤：除數不能為零\n")
			valid = false
		}
	default:
		fmt.Printf("錯誤：不支援的運算符 %s\n", operator)
		valid = false
	}
	
	if valid {
		fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
	}
}

func gradeStudent(score int) {
	var grade string
	var passed bool
	
	switch {
	case score >= 90:
		grade = "A"
		passed = true
	case score >= 80:
		grade = "B"
		passed = true
	case score >= 70:
		grade = "C"
		passed = true
	case score >= 60:
		grade = "D"
		passed = true
	default:
		grade = "F"
		passed = false
	}
	
	status := "不及格"
	if passed {
		status = "及格"
	}
	
	fmt.Printf("分數 %d: 等級 %s (%s)\n", score, grade, status)
}

// 設備狀態位定義
const (
	StatePower     = 1 << iota // 1 (001)
	StateNetwork               // 2 (010)
	StateBluetooth             // 4 (100)
)

func setDeviceState(status int, device string, enabled bool) int {
	var flag int
	
	switch device {
	case "power":
		flag = StatePower
	case "network":
		flag = StateNetwork
	case "bluetooth":
		flag = StateBluetooth
	default:
		return status
	}
	
	if enabled {
		status |= flag  // 設置位
	} else {
		status &^= flag // 清除位
	}
	
	return status
}

func displayDeviceStatus(status int) {
	fmt.Printf("設備狀態 (%03b):\n", status)
	fmt.Printf("  電源: %t\n", (status&StatePower) != 0)
	fmt.Printf("  網路: %t\n", (status&StateNetwork) != 0)
	fmt.Printf("  藍牙: %t\n", (status&StateBluetooth) != 0)
}