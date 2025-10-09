package main

import (
	"fmt"
	"time"
)

// Person 結構體用於演示
type Person struct {
	Name string
	Age  int
	City string
}

// Node 結構體用於鏈表演示
type Node struct {
	Value int
	Next  *Node
}

// LinkedList 結構體
type LinkedList struct {
	Head *Node
	Size int
}

// LinkedList 方法
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

func main() {
	fmt.Println("=== Go 指針示例 ===")
	
	// 1. 指針基礎操作
	demonstratePointerBasics()
	
	// 2. 指針與函數
	demonstratePointerFunctions()
	
	// 3. 函數返回指針
	demonstrateReturnPointers()
	
	// 4. 指針與結構體
	demonstrateStructPointers()
	
	// 5. 結構體中的指針字段
	demonstrateStructWithPointers()
	
	// 6. 指針與數組
	demonstrateArrayPointers()
	
	// 7. 指針與切片
	demonstrateSlicePointers()
	
	// 8. 指針安全使用
	demonstratePointerSafety()
	
	// 9. 指針性能考量
	demonstratePointerPerformance()
}

func demonstratePointerBasics() {
	fmt.Println("\n--- 指針基礎操作 ---")
	
	// 聲明普通變數
	var num int = 42
	var name string = "Go語言"
	
	// 聲明指針變數
	var ptr *int
	var strPtr *string
	
	// 獲取變數的地址
	ptr = &num
	strPtr = &name
	
	fmt.Printf("📍 num 的值: %d\n", num)
	fmt.Printf("📍 num 的地址: %p\n", &num)
	fmt.Printf("📍 ptr 的值（即 num 的地址）: %p\n", ptr)
	fmt.Printf("📍 ptr 指向的值: %d\n", *ptr)
	
	fmt.Printf("📍 name 的值: %s\n", name)
	fmt.Printf("📍 strPtr 指向的值: %s\n", *strPtr)
	
	// 檢查指針是否為 nil
	var nilPtr *int
	fmt.Printf("📍 nilPtr 是否為 nil: %t\n", nilPtr == nil)
	
	// 指針的基本操作
	x := 100
	fmt.Printf("\n🔧 原始值 x: %d\n", x)
	
	ptrX := &x
	fmt.Printf("🔧 指針地址: %p\n", ptrX)
	fmt.Printf("🔧 指針指向的值: %d\n", *ptrX)
	
	// 通過指針修改值
	*ptrX = 200
	fmt.Printf("🔧 通過指針修改後 x: %d\n", x)
	
	// 指針的指針
	ptrPtr := &ptrX
	fmt.Printf("🔧 指針的指針地址: %p\n", ptrPtr)
	fmt.Printf("🔧 指針的指針最終指向的值: %d\n", **ptrPtr)
}

// 值傳遞：函數接收變數的副本
func doubleValue(x int) int {
	x = x * 2
	return x
}

// 指針傳遞：函數接收變數的地址
func doublePointer(x *int) {
	*x = *x * 2
}

// 交換兩個值（值傳遞版本 - 無效）
func swapValues(a, b int) {
	a, b = b, a
}

// 交換兩個值（指針版本 - 有效）
func swapPointers(a, b *int) {
	*a, *b = *b, *a
}

func demonstratePointerFunctions() {
	fmt.Println("\n--- 指針與函數 ---")
	
	// 值傳遞示例
	num1 := 10
	doubled := doubleValue(num1)
	fmt.Printf("🔄 值傳遞 - 原始值: %d, 加倍後: %d\n", num1, doubled)
	
	// 指針傳遞示例
	num2 := 10
	fmt.Printf("🔄 指針傳遞前: %d\n", num2)
	doublePointer(&num2)
	fmt.Printf("🔄 指針傳遞後: %d\n", num2)
	
	// 交換值示例
	a, b := 100, 200
	fmt.Printf("🔄 交換前: a=%d, b=%d\n", a, b)
	
	swapValues(a, b)
	fmt.Printf("🔄 值交換後: a=%d, b=%d (無變化)\n", a, b)
	
	swapPointers(&a, &b)
	fmt.Printf("🔄 指針交換後: a=%d, b=%d (已交換)\n", a, b)
}

// 返回局部變數的指針（Go 中是安全的）
func createInt(value int) *int {
	x := value
	return &x
}

// 工廠函數
func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: "未指定",
	}
}

func demonstrateReturnPointers() {
	fmt.Println("\n--- 函數返回指針 ---")
	
	// 獲取指向新建整數的指針
	intPtr := createInt(42)
	fmt.Printf("🏭 創建的整數: %d, 地址: %p\n", *intPtr, intPtr)
	
	// 創建 Person 實例
	person := newPerson("Alice", 25)
	fmt.Printf("🏭 創建的人員: %+v\n", *person)
	
	// 修改通過指針創建的對象
	person.City = "台北"
	person.Age = 26
	fmt.Printf("🏭 修改後的人員: %+v\n", *person)
}

func demonstrateStructPointers() {
	fmt.Println("\n--- 指針與結構體 ---")
	
	// 創建結構體實例
	person1 := Person{Name: "Bob", Age: 30, City: "台北"}
	fmt.Printf("🏠 person1: %+v\n", person1)
	
	// 創建指向結構體的指針
	personPtr := &person1
	fmt.Printf("🏠 指針地址: %p\n", personPtr)
	
	// Go 語言的語法糖：自動解引用
	fmt.Printf("🏠 姓名: %s (自動解引用)\n", personPtr.Name)
	fmt.Printf("🏠 年齡: %d (自動解引用)\n", personPtr.Age)
	
	// 通過指針修改結構體
	personPtr.Age = 31
	personPtr.City = "高雄"
	fmt.Printf("🏠 修改後: %+v\n", person1)
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
	fmt.Print("🔗 鏈表: ")
	current := list.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Printf(" (大小: %d)\n", list.Size)
}

func demonstrateArrayPointers() {
	fmt.Println("\n--- 指針與數組 ---")
	
	// 數組
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("📊 原始數組: %v\n", arr)
	
	// 數組指針：指向整個數組的指針
	arrPtr := &arr
	fmt.Printf("📊 數組指針指向的數組: %v\n", *arrPtr)
	
	// 通過數組指針修改數組
	(*arrPtr)[0] = 100
	fmt.Printf("📊 修改後的數組: %v\n", arr)
	
	// 指針數組：存儲指針的數組
	var ptrArray [3]*int
	a, b, c := 10, 20, 30
	ptrArray[0] = &a
	ptrArray[1] = &b
	ptrArray[2] = &c
	
	fmt.Printf("📊 指針數組指向的值: [%d, %d, %d]\n", *ptrArray[0], *ptrArray[1], *ptrArray[2])
}

func demonstrateSlicePointers() {
	fmt.Println("\n--- 指針與切片 ---")
	
	// 切片本身就是引用類型
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("🍕 原始切片: %v\n", slice)
	
	// 切片指針
	slicePtr := &slice
	fmt.Printf("🍕 切片指針指向的切片: %v\n", *slicePtr)
	
	// 通過切片指針修改
	(*slicePtr)[0] = 100
	fmt.Printf("🍕 修改後的切片: %v\n", slice)
	
	// 切片元素的指針
	elementPtr := &slice[1]
	*elementPtr = 200
	fmt.Printf("🍕 修改元素後的切片: %v\n", slice)
}

func demonstratePointerSafety() {
	fmt.Println("\n--- 指針安全使用 ---")
	
	// 1. 空指針檢查
	var ptr *int
	if ptr != nil {
		fmt.Printf("🛡️ 指針值: %d\n", *ptr)
	} else {
		fmt.Println("🛡️ 指針為 nil，不能解引用")
	}
	
	// 2. 正確的指針初始化
	num := 42
	ptr = &num
	if ptr != nil {
		fmt.Printf("🛡️ 安全的指針值: %d\n", *ptr)
	}
	
	// 3. 避免懸空指針（Go 的 GC 會處理）
	ptrFromFunc := createInt(100)
	fmt.Printf("🛡️ 函數返回的指針值: %d\n", *ptrFromFunc)
}

// 大結構體用於性能測試
type LargeStruct struct {
	Data [1000]int
	Name string
	Age  int
}

func processLargeStructByValue(ls LargeStruct) {
	_ = ls.Name
}

func processLargeStructByPointer(ls *LargeStruct) {
	_ = ls.Name
}

func demonstratePointerPerformance() {
	fmt.Println("\n--- 指針性能考量 ---")
	
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
	
	fmt.Printf("⚡ 值傳遞耗時: %v\n", valueDuration)
	fmt.Printf("⚡ 指針傳遞耗時: %v\n", pointerDuration)
	if pointerDuration > 0 {
		ratio := float64(valueDuration) / float64(pointerDuration)
		fmt.Printf("⚡ 性能提升: %.2fx\n", ratio)
	}
}