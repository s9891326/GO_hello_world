// 練習 1 解答：動態數組實現
package main

import (
	"fmt"
	"sort"
)

// DynamicArray 動態數組結構
type DynamicArray struct {
	data     []int
	size     int
	capacity int
}

// NewDynamicArray 創建新的動態數組
func NewDynamicArray() *DynamicArray {
	initialCapacity := 4
	return &DynamicArray{
		data:     make([]int, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// NewDynamicArrayWithCapacity 創建指定容量的動態數組
func NewDynamicArrayWithCapacity(capacity int) *DynamicArray {
	if capacity < 1 {
		capacity = 4
	}
	return &DynamicArray{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// Add 添加元素到數組末尾
func (da *DynamicArray) Add(value int) {
	da.ensureCapacity(da.size + 1)
	da.data[da.size] = value
	da.size++
}

// Insert 在指定位置插入元素
func (da *DynamicArray) Insert(index int, value int) error {
	if index < 0 || index > da.size {
		return fmt.Errorf("索引超出範圍: %d, 有效範圍: [0, %d]", index, da.size)
	}
	
	da.ensureCapacity(da.size + 1)
	
	// 移動元素為新元素騰出空間
	for i := da.size; i > index; i-- {
		da.data[i] = da.data[i-1]
	}
	
	da.data[index] = value
	da.size++
	return nil
}

// Remove 刪除指定索引的元素
func (da *DynamicArray) Remove(index int) error {
	if index < 0 || index >= da.size {
		return fmt.Errorf("索引超出範圍: %d, 有效範圍: [0, %d)", index, da.size)
	}
	
	// 向前移動元素填補空隙
	for i := index; i < da.size-1; i++ {
		da.data[i] = da.data[i+1]
	}
	
	da.size--
	da.shrinkIfNeeded()
	return nil
}

// RemoveValue 刪除第一個匹配的值
func (da *DynamicArray) RemoveValue(value int) bool {
	index := da.IndexOf(value)
	if index == -1 {
		return false
	}
	da.Remove(index)
	return true
}

// Get 獲取指定索引的元素
func (da *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= da.size {
		return 0, fmt.Errorf("索引超出範圍: %d, 有效範圍: [0, %d)", index, da.size)
	}
	return da.data[index], nil
}

// Set 設置指定索引的元素值
func (da *DynamicArray) Set(index int, value int) error {
	if index < 0 || index >= da.size {
		return fmt.Errorf("索引超出範圍: %d, 有效範圍: [0, %d)", index, da.size)
	}
	da.data[index] = value
	return nil
}

// IndexOf 查找元素的索引
func (da *DynamicArray) IndexOf(value int) int {
	for i := 0; i < da.size; i++ {
		if da.data[i] == value {
			return i
		}
	}
	return -1
}

// Contains 檢查是否包含指定元素
func (da *DynamicArray) Contains(value int) bool {
	return da.IndexOf(value) != -1
}

// Size 返回數組大小
func (da *DynamicArray) Size() int {
	return da.size
}

// Capacity 返回數組容量
func (da *DynamicArray) Capacity() int {
	return da.capacity
}

// IsEmpty 檢查數組是否為空
func (da *DynamicArray) IsEmpty() bool {
	return da.size == 0
}

// Clear 清空數組
func (da *DynamicArray) Clear() {
	da.size = 0
	da.shrinkIfNeeded()
}

// ToSlice 轉換為切片
func (da *DynamicArray) ToSlice() []int {
	result := make([]int, da.size)
	copy(result, da.data[:da.size])
	return result
}

// Sort 排序數組
func (da *DynamicArray) Sort() {
	if da.size <= 1 {
		return
	}
	sort.Ints(da.data[:da.size])
}

// Reverse 反轉數組
func (da *DynamicArray) Reverse() {
	for i, j := 0, da.size-1; i < j; i, j = i+1, j-1 {
		da.data[i], da.data[j] = da.data[j], da.data[i]
	}
}

// ensureCapacity 確保容量足夠
func (da *DynamicArray) ensureCapacity(minCapacity int) {
	if minCapacity > da.capacity {
		da.grow(minCapacity)
	}
}

// grow 擴容
func (da *DynamicArray) grow(minCapacity int) {
	oldCapacity := da.capacity
	newCapacity := oldCapacity * 2
	
	if newCapacity < minCapacity {
		newCapacity = minCapacity
	}
	
	newData := make([]int, newCapacity)
	copy(newData, da.data)
	da.data = newData
	da.capacity = newCapacity
	
	fmt.Printf("🔄 擴容: %d -> %d\n", oldCapacity, newCapacity)
}

// shrinkIfNeeded 根據需要縮容
func (da *DynamicArray) shrinkIfNeeded() {
	if da.capacity > 4 && da.size < da.capacity/4 {
		da.shrink()
	}
}

// shrink 縮容
func (da *DynamicArray) shrink() {
	oldCapacity := da.capacity
	newCapacity := da.capacity / 2
	
	if newCapacity < 4 {
		newCapacity = 4
	}
	
	newData := make([]int, newCapacity)
	copy(newData, da.data[:da.size])
	da.data = newData
	da.capacity = newCapacity
	
	fmt.Printf("🔄 縮容: %d -> %d\n", oldCapacity, newCapacity)
}

// Stats 返回統計信息
func (da *DynamicArray) Stats() (int, int, float64) {
	usageRate := 0.0
	if da.capacity > 0 {
		usageRate = float64(da.size) / float64(da.capacity) * 100
	}
	return da.size, da.capacity, usageRate
}

// String 實現 Stringer 接口
func (da *DynamicArray) String() string {
	if da.size == 0 {
		return "[]"
	}
	
	result := "["
	for i := 0; i < da.size; i++ {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d", da.data[i])
	}
	result += "]"
	return result
}

// Iterator 迭代器結構
type Iterator struct {
	array *DynamicArray
	index int
}

// NewIterator 創建迭代器
func (da *DynamicArray) NewIterator() *Iterator {
	return &Iterator{
		array: da,
		index: 0,
	}
}

// HasNext 檢查是否有下一個元素
func (it *Iterator) HasNext() bool {
	return it.index < it.array.size
}

// Next 獲取下一個元素
func (it *Iterator) Next() (int, error) {
	if !it.HasNext() {
		return 0, fmt.Errorf("沒有更多元素")
	}
	value := it.array.data[it.index]
	it.index++
	return value, nil
}

// Reset 重置迭代器
func (it *Iterator) Reset() {
	it.index = 0
}

func main() {
	fmt.Println("=== 動態數組測試 ===")
	
	// 創建動態數組
	da := NewDynamicArray()
	size, capacity, usage := da.Stats()
	fmt.Printf("📝 創建動態數組，初始狀態: 大小=%d, 容量=%d, 使用率=%.1f%%\n", 
		size, capacity, usage)
	
	// 添加元素
	fmt.Println("\n📝 添加元素測試:")
	for i := 1; i <= 10; i++ {
		da.Add(i)
		if i <= 5 || i == 10 {
			size, capacity, usage := da.Stats()
			fmt.Printf("   添加 %d: %s (大小=%d, 容量=%d, 使用率=%.1f%%)\n", 
				i, da.String(), size, capacity, usage)
		}
	}
	
	// 插入元素
	fmt.Println("\n📝 插入元素測試:")
	err := da.Insert(2, 99)
	if err != nil {
		fmt.Printf("   插入失敗: %v\n", err)
	} else {
		fmt.Printf("   在索引 2 插入 99: %s\n", da.String())
	}
	
	// 刪除元素
	fmt.Println("\n📝 刪除元素測試:")
	err = da.Remove(1)
	if err != nil {
		fmt.Printf("   刪除失敗: %v\n", err)
	} else {
		fmt.Printf("   刪除索引 1: %s\n", da.String())
	}
	
	// 查找元素
	fmt.Println("\n📝 查找元素測試:")
	index := da.IndexOf(99)
	if index != -1 {
		fmt.Printf("   找到元素 99，索引: %d\n", index)
	} else {
		fmt.Printf("   未找到元素 99\n", index)
	}
	
	index = da.IndexOf(100)
	if index != -1 {
		fmt.Printf("   找到元素 100，索引: %d\n", index)
	} else {
		fmt.Printf("   未找到元素 100\n")
	}
	
	// 排序
	fmt.Println("\n📝 排序測試:")
	fmt.Printf("   排序前: %s\n", da.String())
	da.Sort()
	fmt.Printf("   排序後: %s\n", da.String())
	
	// 迭代器測試
	fmt.Println("\n📝 迭代器測試:")
	fmt.Print("   迭代結果: ")
	iter := da.NewIterator()
	for iter.HasNext() {
		value, _ := iter.Next()
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	
	// 刪除一些元素觸發縮容
	fmt.Println("\n📝 縮容測試:")
	originalSize := da.Size()
	for i := 0; i < originalSize-2; i++ {
		da.Remove(0)
	}
	size, capacity, usage = da.Stats()
	fmt.Printf("   刪除大部分元素後: %s (大小=%d, 容量=%d, 使用率=%.1f%%)\n", 
		da.String(), size, capacity, usage)
}