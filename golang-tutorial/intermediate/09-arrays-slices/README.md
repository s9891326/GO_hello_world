# 第九章：數組和切片

## 🎯 學習目標

- 理解數組的基本概念和使用
- 掌握切片的原理和操作
- 學會數組和切片的區別
- 了解切片的內部實現
- 掌握切片的高級操作
- 學會數組和切片的最佳實踐

## 📊 數組基礎

數組是具有固定長度的相同類型元素的序列。在 Go 中，數組的長度是類型的一部分。

### 數組的聲明和初始化

```go
package main

import "fmt"

func demonstrateArrayBasics() {
    // 1. 聲明數組
    var arr1 [5]int                    // 零值初始化，所有元素為 0
    var arr2 [3]string                 // 零值初始化，所有元素為 ""
    
    fmt.Printf("零值數組 arr1: %v\n", arr1)
    fmt.Printf("零值數組 arr2: %v\n", arr2)
    
    // 2. 聲明並初始化
    var arr3 [4]int = [4]int{1, 2, 3, 4}
    fmt.Printf("初始化數組 arr3: %v\n", arr3)
    
    // 3. 簡化初始化
    arr4 := [5]int{10, 20, 30, 40, 50}
    fmt.Printf("簡化初始化 arr4: %v\n", arr4)
    
    // 4. 部分初始化
    arr5 := [5]int{1, 2}              // 其餘元素為零值
    fmt.Printf("部分初始化 arr5: %v\n", arr5)
    
    // 5. 指定索引初始化
    arr6 := [5]int{0: 100, 2: 200, 4: 400}
    fmt.Printf("指定索引初始化 arr6: %v\n", arr6)
    
    // 6. 自動推導長度
    arr7 := [...]int{1, 2, 3, 4, 5, 6}  // 編譯器計算長度
    fmt.Printf("自動推導長度 arr7: %v (長度: %d)\n", arr7, len(arr7))
}
```

### 數組的基本操作

```go
func demonstrateArrayOperations() {
    fmt.Println("\n--- 數組基本操作 ---")
    
    arr := [5]int{10, 20, 30, 40, 50}
    fmt.Printf("原始數組: %v\n", arr)
    
    // 1. 訪問元素
    fmt.Printf("第一個元素: %d\n", arr[0])
    fmt.Printf("最後一個元素: %d\n", arr[len(arr)-1])
    
    // 2. 修改元素
    arr[0] = 100
    arr[4] = 500
    fmt.Printf("修改後: %v\n", arr)
    
    // 3. 數組長度
    fmt.Printf("數組長度: %d\n", len(arr))
    
    // 4. 遍歷數組
    fmt.Print("for-range 遍歷: ")
    for index, value := range arr {
        fmt.Printf("[%d]=%d ", index, value)
    }
    fmt.Println()
    
    fmt.Print("傳統 for 遍歷: ")
    for i := 0; i < len(arr); i++ {
        fmt.Printf("[%d]=%d ", i, arr[i])
    }
    fmt.Println()
    
    // 5. 只要值，忽略索引
    fmt.Print("只取值: ")
    for _, value := range arr {
        fmt.Printf("%d ", value)
    }
    fmt.Println()
    
    // 6. 只要索引，忽略值
    fmt.Print("只取索引: ")
    for index := range arr {
        fmt.Printf("%d ", index)
    }
    fmt.Println()
}
```

### 多維數組

```go
func demonstrateMultiDimensionalArrays() {
    fmt.Println("\n--- 多維數組 ---")
    
    // 二維數組
    var matrix [3][4]int
    fmt.Printf("零值二維數組:\n")
    printMatrix(matrix)
    
    // 初始化二維數組
    matrix2 := [3][4]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    fmt.Printf("初始化二維數組:\n")
    printMatrix(matrix2)
    
    // 部分初始化
    matrix3 := [3][4]int{
        {1, 2},
        {5, 6, 7},
    }
    fmt.Printf("部分初始化二維數組:\n")
    printMatrix(matrix3)
    
    // 修改二維數組元素
    matrix3[0][2] = 33
    matrix3[2][1] = 99
    fmt.Printf("修改後:\n")
    printMatrix(matrix3)
}

func printMatrix(matrix [3][4]int) {
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%3d ", matrix[i][j])
        }
        fmt.Println()
    }
}
```

## 🍕 切片基礎

切片是對數組的抽象，提供了更靈活的接口。切片是引用類型，包含指向底層數組的指針、長度和容量。

### 切片的內部結構

```
切片的內部結構：
┌─────────────────────────────────────┐
│    slice header (24 bytes)         │
├─────────────────────────────────────┤
│ ptr   │ len   │ cap               │
│ 8字節  │ 8字節  │ 8字節              │
└─────────────────────────────────────┘
        │
        ▼
┌─────────────────────────────────────┐
│     underlying array              │
│ [0] [1] [2] [3] [4] [5] ...       │
└─────────────────────────────────────┘
```

### 切片的聲明和初始化

```go
func demonstrateSliceBasics() {
    fmt.Println("\n--- 切片基礎 ---")
    
    // 1. 聲明切片
    var slice1 []int              // nil 切片
    fmt.Printf("nil 切片: %v (長度: %d, 容量: %d, 是否為 nil: %t)\n", 
        slice1, len(slice1), cap(slice1), slice1 == nil)
    
    // 2. 使用 make 創建切片
    slice2 := make([]int, 5)      // 長度為 5，容量為 5
    fmt.Printf("make 切片: %v (長度: %d, 容量: %d)\n", 
        slice2, len(slice2), cap(slice2))
    
    slice3 := make([]int, 3, 8)   // 長度為 3，容量為 8
    fmt.Printf("make 切片(指定容量): %v (長度: %d, 容量: %d)\n", 
        slice3, len(slice3), cap(slice3))
    
    // 3. 字面量初始化
    slice4 := []int{1, 2, 3, 4, 5}
    fmt.Printf("字面量切片: %v (長度: %d, 容量: %d)\n", 
        slice4, len(slice4), cap(slice4))
    
    // 4. 從數組創建切片
    arr := [6]int{10, 20, 30, 40, 50, 60}
    slice5 := arr[1:4]            // 包含索引 1, 2, 3
    fmt.Printf("從數組創建切片: %v (長度: %d, 容量: %d)\n", 
        slice5, len(slice5), cap(slice5))
    
    slice6 := arr[:]              // 整個數組
    fmt.Printf("整個數組切片: %v (長度: %d, 容量: %d)\n", 
        slice6, len(slice6), cap(slice6))
}
```

### 切片操作

```go
func demonstrateSliceOperations() {
    fmt.Println("\n--- 切片操作 ---")
    
    slice := []int{1, 2, 3, 4, 5}
    fmt.Printf("原始切片: %v\n", slice)
    
    // 1. 切片切分
    fmt.Printf("slice[1:3]: %v\n", slice[1:3])     // [2, 3]
    fmt.Printf("slice[:3]: %v\n", slice[:3])       // [1, 2, 3]
    fmt.Printf("slice[2:]: %v\n", slice[2:])       // [3, 4, 5]
    fmt.Printf("slice[:]: %v\n", slice[:])         // [1, 2, 3, 4, 5]
    
    // 2. 修改切片
    slice[0] = 100
    fmt.Printf("修改後: %v\n", slice)
    
    // 3. append 操作
    slice = append(slice, 6)
    fmt.Printf("append 6: %v (長度: %d, 容量: %d)\n", 
        slice, len(slice), cap(slice))
    
    slice = append(slice, 7, 8, 9)
    fmt.Printf("append 多個: %v (長度: %d, 容量: %d)\n", 
        slice, len(slice), cap(slice))
    
    // 4. append 另一個切片
    other := []int{10, 11, 12}
    slice = append(slice, other...)
    fmt.Printf("append 切片: %v (長度: %d, 容量: %d)\n", 
        slice, len(slice), cap(slice))
    
    // 5. copy 操作
    dest := make([]int, len(slice))
    n := copy(dest, slice)
    fmt.Printf("copy 結果: %v (複製了 %d 個元素)\n", dest, n)
    
    // 6. 刪除元素
    index := 2
    slice = append(slice[:index], slice[index+1:]...)
    fmt.Printf("刪除索引 %d: %v\n", index, slice)
}
```

### 切片的陷阱

```go
func demonstrateSliceTraps() {
    fmt.Println("\n--- 切片陷阱 ---")
    
    // 陷阱 1：切片共享底層數組
    fmt.Println("陷阱 1: 切片共享底層數組")
    arr := [5]int{1, 2, 3, 4, 5}
    slice1 := arr[1:3]  // [2, 3]
    slice2 := arr[2:4]  // [3, 4]
    
    fmt.Printf("原數組: %v\n", arr)
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("slice2: %v\n", slice2)
    
    slice1[1] = 100  // 修改 slice1[1]，實際修改的是 arr[2]
    fmt.Printf("修改 slice1[1] = 100 後:\n")
    fmt.Printf("原數組: %v\n", arr)
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("slice2: %v\n", slice2)  // slice2[0] 也變了！
    
    // 陷阱 2：append 可能改變底層數組
    fmt.Println("\n陷阱 2: append 可能改變底層數組")
    arr2 := [5]int{1, 2, 3, 4, 5}
    slice3 := arr2[1:3]  // [2, 3]，容量為 4
    
    fmt.Printf("原數組: %v\n", arr2)
    fmt.Printf("slice3: %v (容量: %d)\n", slice3, cap(slice3))
    
    slice3 = append(slice3, 99)  // 容量足夠，直接修改底層數組
    fmt.Printf("append 99 後:\n")
    fmt.Printf("原數組: %v\n", arr2)  // arr2[3] 變成了 99
    fmt.Printf("slice3: %v\n", slice3)
    
    // 陷阱 3：nil 切片 vs 空切片
    fmt.Println("\n陷阱 3: nil 切片 vs 空切片")
    var nilSlice []int
    emptySlice := []int{}
    makeSlice := make([]int, 0)
    
    fmt.Printf("nil 切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
        nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
    fmt.Printf("空切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
        emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
    fmt.Printf("make 空切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
        makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)
}
```

## 🔧 高級切片操作

### 切片的擴容機制

```go
func demonstrateSliceGrowth() {
    fmt.Println("\n--- 切片擴容機制 ---")
    
    slice := make([]int, 0, 1)
    fmt.Printf("初始: 長度=%d, 容量=%d\n", len(slice), cap(slice))
    
    for i := 1; i <= 10; i++ {
        slice = append(slice, i)
        fmt.Printf("append %d: 長度=%d, 容量=%d\n", i, len(slice), cap(slice))
    }
    
    // 觀察擴容規律：
    // 當容量小於 1024 時，每次擴容翻倍
    // 當容量大於等於 1024 時，每次擴容 25%
}
```

### 切片的內存優化

```go
func demonstrateSliceMemoryOptimization() {
    fmt.Println("\n--- 切片內存優化 ---")
    
    // 問題：大切片的小切片可能導致內存泄漏
    largeSlice := make([]int, 1000000)
    for i := range largeSlice {
        largeSlice[i] = i
    }
    
    // 不好的做法：保留對大切片的引用
    smallSliceBad := largeSlice[0:5]
    fmt.Printf("不好的小切片: 長度=%d, 容量=%d\n", len(smallSliceBad), cap(smallSliceBad))
    
    // 好的做法：複製需要的部分
    smallSliceGood := make([]int, 5)
    copy(smallSliceGood, largeSlice[0:5])
    fmt.Printf("好的小切片: 長度=%d, 容量=%d\n", len(smallSliceGood), cap(smallSliceGood))
    
    // 現在 largeSlice 可以被垃圾回收了（如果沒有其他引用）
}
```

### 切片作為函數參數

```go
func demonstrateSliceAsParameter() {
    fmt.Println("\n--- 切片作為函數參數 ---")
    
    slice := []int{1, 2, 3, 4, 5}
    fmt.Printf("原始切片: %v\n", slice)
    
    // 修改切片元素
    modifySliceElements(slice)
    fmt.Printf("修改元素後: %v\n", slice)
    
    // 嘗試修改切片本身（不會影響原切片）
    tryModifySlice(slice)
    fmt.Printf("嘗試修改切片後: %v\n", slice)
    
    // 正確修改切片本身的方法
    slice = correctModifySlice(slice)
    fmt.Printf("正確修改切片後: %v\n", slice)
}

func modifySliceElements(s []int) {
    for i := range s {
        s[i] *= 2
    }
}

func tryModifySlice(s []int) {
    s = append(s, 100)  // 這不會影響原切片
}

func correctModifySlice(s []int) []int {
    return append(s, 100)  // 返回新切片
}
```

## 🎨 實際應用場景

### 動態數組實現

```go
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

func (da *DynamicArray) Get(index int) (int, error) {
    if index < 0 || index >= da.size {
        return 0, fmt.Errorf("索引超出範圍")
    }
    return da.data[index], nil
}

func (da *DynamicArray) Remove(index int) error {
    if index < 0 || index >= da.size {
        return fmt.Errorf("索引超出範圍")
    }
    da.data = append(da.data[:index], da.data[index+1:]...)
    da.size--
    return nil
}

func (da *DynamicArray) Size() int {
    return da.size
}

func (da *DynamicArray) ToSlice() []int {
    result := make([]int, da.size)
    copy(result, da.data)
    return result
}
```

### 矩陣操作

```go
type Matrix [][]int

func NewMatrix(rows, cols int) Matrix {
    matrix := make(Matrix, rows)
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }
    return matrix
}

func (m Matrix) Set(row, col, value int) error {
    if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
        return fmt.Errorf("索引超出範圍")
    }
    m[row][col] = value
    return nil
}

func (m Matrix) Get(row, col int) (int, error) {
    if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
        return 0, fmt.Errorf("索引超出範圍")
    }
    return m[row][col], nil
}

func (m Matrix) Display() {
    for _, row := range m {
        for _, val := range row {
            fmt.Printf("%4d ", val)
        }
        fmt.Println()
    }
}
```

## 💡 最佳實踐

### 1. 選擇數組還是切片

```go
// 使用數組的場景：
// - 長度固定且已知
// - 需要值語義（複製整個數組）
// - 作為哈希表的鍵
func useArray() {
    var buffer [1024]byte     // 固定大小緩衝區
    var rgb [3]uint8          // RGB 顏色值
}

// 使用切片的場景：
// - 長度動態變化
// - 需要引用語義
// - 大部分情況下
func useSlice() {
    var items []string        // 動態字符串列表
    var numbers []int         // 數字集合
}
```

### 2. 切片的安全操作

```go
// 安全的切片操作
func safeSliceOperations() {
    // 檢查切片是否為 nil
    var slice []int
    if slice != nil {
        fmt.Println("切片不為 nil")
    }
    
    // 檢查索引範圍
    if len(slice) > 0 {
        first := slice[0]
        fmt.Println("第一個元素:", first)
    }
    
    // 預分配容量
    slice = make([]int, 0, 100)  // 如果知道大概大小
    
    // 避免切片泄漏
    bigSlice := make([]int, 1000000)
    smallSlice := make([]int, 10)
    copy(smallSlice, bigSlice[:10])  // 而不是 bigSlice[:10]
}
```

## 🎯 本章練習

1. 實現動態數組類
2. 創建矩陣運算庫
3. 實現環形緩衝區
4. 創建排序算法集合

---

**下一章：[映射](../10-maps/)**