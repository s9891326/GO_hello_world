package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Go 數組和切片示例 ===")
	
	// 1. 數組基礎操作
	demonstrateArrayBasics()
	
	// 2. 數組基本操作
	demonstrateArrayOperations()
	
	// 3. 多維數組
	demonstrateMultiDimensionalArrays()
	
	// 4. 切片基礎
	demonstrateSliceBasics()
	
	// 5. 切片操作
	demonstrateSliceOperations()
	
	// 6. 切片陷阱
	demonstrateSliceTraps()
	
	// 7. 切片擴容機制
	demonstrateSliceGrowth()
	
	// 8. 切片內存優化
	demonstrateSliceMemoryOptimization()
	
	// 9. 切片作為函數參數
	demonstrateSliceAsParameter()
	
	// 10. 實際應用示例
	demonstrateRealWorldExamples()
}

func demonstrateArrayBasics() {
	fmt.Println("\n--- 數組基礎 ---")
	
	// 1. 聲明數組
	var arr1 [5]int
	var arr2 [3]string
	
	fmt.Printf("📊 零值數組 arr1: %v\n", arr1)
	fmt.Printf("📊 零值數組 arr2: %v\n", arr2)
	
	// 2. 聲明並初始化
	var arr3 [4]int = [4]int{1, 2, 3, 4}
	fmt.Printf("📊 初始化數組 arr3: %v\n", arr3)
	
	// 3. 簡化初始化
	arr4 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("📊 簡化初始化 arr4: %v\n", arr4)
	
	// 4. 部分初始化
	arr5 := [5]int{1, 2}
	fmt.Printf("📊 部分初始化 arr5: %v\n", arr5)
	
	// 5. 指定索引初始化
	arr6 := [5]int{0: 100, 2: 200, 4: 400}
	fmt.Printf("📊 指定索引初始化 arr6: %v\n", arr6)
	
	// 6. 自動推導長度
	arr7 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("📊 自動推導長度 arr7: %v (長度: %d)\n", arr7, len(arr7))
	
	// 7. 數組大小
	fmt.Printf("📊 數組類型大小: arr1=%d bytes, arr4=%d bytes\n", 
		unsafe.Sizeof(arr1), unsafe.Sizeof(arr4))
}

func demonstrateArrayOperations() {
	fmt.Println("\n--- 數組基本操作 ---")
	
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("🔧 原始數組: %v\n", arr)
	
	// 1. 訪問元素
	fmt.Printf("🔧 第一個元素: %d\n", arr[0])
	fmt.Printf("🔧 最後一個元素: %d\n", arr[len(arr)-1])
	
	// 2. 修改元素
	arr[0] = 100
	arr[4] = 500
	fmt.Printf("🔧 修改後: %v\n", arr)
	
	// 3. 數組長度
	fmt.Printf("🔧 數組長度: %d\n", len(arr))
	
	// 4. 遍歷數組
	fmt.Print("🔧 for-range 遍歷: ")
	for index, value := range arr {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println()
	
	fmt.Print("🔧 傳統 for 遍歷: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
	
	// 5. 只要值，忽略索引
	fmt.Print("🔧 只取值: ")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	
	// 6. 數組比較
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	
	fmt.Printf("🔧 arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("🔧 arr1 == arr3: %t\n", arr1 == arr3)
}

func demonstrateMultiDimensionalArrays() {
	fmt.Println("\n--- 多維數組 ---")
	
	// 二維數組
	var matrix [3][4]int
	fmt.Printf("🗃️ 零值二維數組:\n")
	printMatrix(matrix)
	
	// 初始化二維數組
	matrix2 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Printf("🗃️ 初始化二維數組:\n")
	printMatrix(matrix2)
	
	// 部分初始化
	matrix3 := [3][4]int{
		{1, 2},
		{5, 6, 7},
	}
	fmt.Printf("🗃️ 部分初始化二維數組:\n")
	printMatrix(matrix3)
	
	// 修改二維數組元素
	matrix3[0][2] = 33
	matrix3[2][1] = 99
	fmt.Printf("🗃️ 修改後:\n")
	printMatrix(matrix3)
	
	// 三維數組示例
	var cube [2][3][4]int
	cube[1][2][3] = 100
	fmt.Printf("🗃️ 三維數組 cube[1][2][3] = %d\n", cube[1][2][3])
}

func printMatrix(matrix [3][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("   ")
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func demonstrateSliceBasics() {
	fmt.Println("\n--- 切片基礎 ---")
	
	// 1. 聲明切片
	var slice1 []int
	fmt.Printf("🍕 nil 切片: %v (長度: %d, 容量: %d, 是否為 nil: %t)\n", 
		slice1, len(slice1), cap(slice1), slice1 == nil)
	
	// 2. 使用 make 創建切片
	slice2 := make([]int, 5)
	fmt.Printf("🍕 make 切片: %v (長度: %d, 容量: %d)\n", 
		slice2, len(slice2), cap(slice2))
	
	slice3 := make([]int, 3, 8)
	fmt.Printf("🍕 make 切片(指定容量): %v (長度: %d, 容量: %d)\n", 
		slice3, len(slice3), cap(slice3))
	
	// 3. 字面量初始化
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("🍕 字面量切片: %v (長度: %d, 容量: %d)\n", 
		slice4, len(slice4), cap(slice4))
	
	// 4. 從數組創建切片
	arr := [6]int{10, 20, 30, 40, 50, 60}
	slice5 := arr[1:4]
	fmt.Printf("🍕 從數組創建切片 arr[1:4]: %v (長度: %d, 容量: %d)\n", 
		slice5, len(slice5), cap(slice5))
	
	slice6 := arr[:]
	fmt.Printf("🍕 整個數組切片 arr[:]: %v (長度: %d, 容量: %d)\n", 
		slice6, len(slice6), cap(slice6))
	
	// 5. 從切片創建切片
	slice7 := slice4[1:3]
	fmt.Printf("🍕 從切片創建切片: %v (長度: %d, 容量: %d)\n", 
		slice7, len(slice7), cap(slice7))
	
	// 6. 空切片 vs nil 切片
	emptySlice := []int{}
	makeEmptySlice := make([]int, 0)
	
	fmt.Printf("🍕 空切片: %v (== nil: %t)\n", emptySlice, emptySlice == nil)
	fmt.Printf("🍕 make 空切片: %v (== nil: %t)\n", makeEmptySlice, makeEmptySlice == nil)
}

func demonstrateSliceOperations() {
	fmt.Println("\n--- 切片操作 ---")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("⚙️ 原始切片: %v\n", slice)
	
	// 1. 切片切分
	fmt.Printf("⚙️ slice[1:3]: %v\n", slice[1:3])
	fmt.Printf("⚙️ slice[:3]: %v\n", slice[:3])
	fmt.Printf("⚙️ slice[2:]: %v\n", slice[2:])
	fmt.Printf("⚙️ slice[:]: %v\n", slice[:])
	
	// 2. 三參數切片 slice[low:high:max]
	slice8 := slice[1:3:4]  // 限制容量
	fmt.Printf("⚙️ slice[1:3:4]: %v (長度: %d, 容量: %d)\n", 
		slice8, len(slice8), cap(slice8))
	
	// 3. 修改切片
	slice[0] = 100
	fmt.Printf("⚙️ 修改後: %v\n", slice)
	
	// 4. append 操作
	slice = append(slice, 6)
	fmt.Printf("⚙️ append 6: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
	
	slice = append(slice, 7, 8, 9)
	fmt.Printf("⚙️ append 多個: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
	
	// 5. append 另一個切片
	other := []int{10, 11, 12}
	slice = append(slice, other...)
	fmt.Printf("⚙️ append 切片: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
	
	// 6. copy 操作
	dest := make([]int, len(slice))
	n := copy(dest, slice)
	fmt.Printf("⚙️ copy 結果: %v (複製了 %d 個元素)\n", dest, n)
	
	// 7. 部分 copy
	partialDest := make([]int, 3)
	n2 := copy(partialDest, slice)
	fmt.Printf("⚙️ 部分 copy: %v (複製了 %d 個元素)\n", partialDest, n2)
	
	// 8. 刪除元素
	index := 2
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("⚙️ 刪除前: %v\n", originalSlice)
	result := append(originalSlice[:index], originalSlice[index+1:]...)
	fmt.Printf("⚙️ 刪除索引 %d 後: %v\n", index, result)
	
	// 9. 插入元素
	insertSlice := []int{1, 2, 4, 5}
	insertIndex := 2
	insertValue := 3
	fmt.Printf("⚙️ 插入前: %v\n", insertSlice)
	insertSlice = append(insertSlice[:insertIndex], append([]int{insertValue}, insertSlice[insertIndex:]...)...)
	fmt.Printf("⚙️ 在索引 %d 插入 %d 後: %v\n", insertIndex, insertValue, insertSlice)
}

func demonstrateSliceTraps() {
	fmt.Println("\n--- 切片陷阱 ---")
	
	// 陷阱 1：切片共享底層數組
	fmt.Println("⚠️ 陷阱 1: 切片共享底層數組")
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:3]
	slice2 := arr[2:4]
	
	fmt.Printf("   原數組: %v\n", arr)
	fmt.Printf("   slice1: %v, slice2: %v\n", slice1, slice2)
	
	slice1[1] = 100
	fmt.Printf("   修改 slice1[1] 後: slice2 也受影響 %v\n", slice2)
}

func demonstrateSliceGrowth() {
	fmt.Println("\n--- 切片擴容機制 ---")
	
	slice := make([]int, 0, 1)
	fmt.Printf("📈 初始: 長度=%d, 容量=%d\n", len(slice), cap(slice))
	
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Printf("📈 append %d: 長度=%d, 容量=%d\n", i, len(slice), cap(slice))
	}
}

func demonstrateSliceMemoryOptimization() {
	fmt.Println("\n--- 切片內存優化 ---")
	
	largeSlice := make([]int, 1000000)
	fmt.Printf("💾 大切片: 長度=%d\n", len(largeSlice))
	
	// 不好的做法
	smallSliceBad := largeSlice[0:5]
	fmt.Printf("💾 不好的小切片: 容量=%d (引用大數組)\n", cap(smallSliceBad))
	
	// 好的做法
	smallSliceGood := make([]int, 5)
	copy(smallSliceGood, largeSlice[0:5])
	fmt.Printf("💾 好的小切片: 容量=%d (獨立數組)\n", cap(smallSliceGood))
}

func demonstrateSliceAsParameter() {
	fmt.Println("\n--- 切片作為函數參數 ---")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("🔄 原始切片: %v\n", slice)
	
	modifySliceElements(slice)
	fmt.Printf("🔄 修改元素後: %v\n", slice)
	
	slice = correctModifySlice(slice)
	fmt.Printf("🔄 正確修改切片後: %v\n", slice)
}

func modifySliceElements(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func correctModifySlice(s []int) []int {
	return append(s, 100)
}

func demonstrateRealWorldExamples() {
	fmt.Println("\n--- 實際應用示例 ---")
	
	// 動態數組示例
	fmt.Println("🎯 動態數組:")
	da := NewDynamicArray()
	da.Add(1)
	da.Add(2)
	da.Add(3)
	fmt.Printf("   動態數組: %v\n", da.ToSlice())
	
	// 矩陣操作示例
	fmt.Println("🎯 矩陣操作:")
	matrix := NewMatrix(3, 3)
	matrix.Set(0, 0, 1)
	matrix.Set(1, 1, 2)
	matrix.Set(2, 2, 3)
	fmt.Println("   3x3 矩陣:")
	matrix.Display()
}

// 動態數組實現
type DynamicArray struct {
	data []int
	size int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data: make([]int, 0, 4),
		size: 0,
	}
}

func (da *DynamicArray) Add(value int) {
	da.data = append(da.data, value)
	da.size++
}

func (da *DynamicArray) ToSlice() []int {
	result := make([]int, da.size)
	copy(result, da.data)
	return result
}

// 矩陣實現
type Matrix [][]int

func NewMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func (m Matrix) Set(row, col, value int) {
	if row >= 0 && row < len(m) && col >= 0 && col < len(m[0]) {
		m[row][col] = value
	}
}

func (m Matrix) Display() {
	for _, row := range m {
		fmt.Print("      ")
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}