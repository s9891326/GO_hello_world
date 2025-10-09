// 練習 1 解答：雙向鏈表實現
package main

import (
	"fmt"
	"sort"
)

// 雙向鏈表節點
type DoublyNode struct {
	Value int
	Prev  *DoublyNode
	Next  *DoublyNode
}

// 雙向鏈表
type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
	Size int
}

// 創建新的雙向鏈表
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// 在鏈表末尾添加元素
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &DoublyNode{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
	
	if dll.Head == nil {
		// 空鏈表
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		// 非空鏈表，添加到末尾
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
	
	dll.Size++
	fmt.Printf("✅ 添加元素 %d，鏈表大小: %d\n", value, dll.Size)
}

// 在鏈表開頭添加元素
func (dll *DoublyLinkedList) Prepend(value int) {
	newNode := &DoublyNode{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
	
	if dll.Head == nil {
		// 空鏈表
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		// 非空鏈表，添加到開頭
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	
	dll.Size++
	fmt.Printf("✅ 在開頭添加元素 %d，鏈表大小: %d\n", value, dll.Size)
}

// 在指定位置插入元素
func (dll *DoublyLinkedList) InsertAt(index int, value int) error {
	if index < 0 || index > dll.Size {
		return fmt.Errorf("索引 %d 超出範圍 [0, %d]", index, dll.Size)
	}
	
	if index == 0 {
		dll.Prepend(value)
		return nil
	}
	
	if index == dll.Size {
		dll.Append(value)
		return nil
	}
	
	newNode := &DoublyNode{Value: value}
	current := dll.getNodeAt(index)
	
	// 插入到 current 之前
	newNode.Prev = current.Prev
	newNode.Next = current
	current.Prev.Next = newNode
	current.Prev = newNode
	
	dll.Size++
	fmt.Printf("✅ 在位置 %d 插入元素 %d\n", index, value)
	return nil
}

// 刪除指定值的元素
func (dll *DoublyLinkedList) Remove(value int) bool {
	current := dll.Head
	
	for current != nil {
		if current.Value == value {
			dll.removeNode(current)
			fmt.Printf("✅ 刪除元素 %d，鏈表大小: %d\n", value, dll.Size)
			return true
		}
		current = current.Next
	}
	
	fmt.Printf("❌ 未找到元素 %d\n", value)
	return false
}

// 刪除指定位置的元素
func (dll *DoublyLinkedList) RemoveAt(index int) error {
	if index < 0 || index >= dll.Size {
		return fmt.Errorf("索引 %d 超出範圍 [0, %d)", index, dll.Size)
	}
	
	nodeToRemove := dll.getNodeAt(index)
	dll.removeNode(nodeToRemove)
	fmt.Printf("✅ 刪除位置 %d 的元素 %d\n", index, nodeToRemove.Value)
	return nil
}

// 內部方法：刪除指定節點
func (dll *DoublyLinkedList) removeNode(node *DoublyNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// 刪除的是第一個節點
		dll.Head = node.Next
	}
	
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// 刪除的是最後一個節點
		dll.Tail = node.Prev
	}
	
	dll.Size--
}

// 查找元素
func (dll *DoublyLinkedList) Find(value int) *DoublyNode {
	current := dll.Head
	index := 0
	
	for current != nil {
		if current.Value == value {
			fmt.Printf("🔍 找到元素 %d，位置: %d，地址: %p\n", value, index, current)
			return current
		}
		current = current.Next
		index++
	}
	
	fmt.Printf("🔍 未找到元素 %d\n", value)
	return nil
}

// 獲取指定位置的節點
func (dll *DoublyLinkedList) getNodeAt(index int) *DoublyNode {
	if index < dll.Size/2 {
		// 從頭部開始搜索
		current := dll.Head
		for i := 0; i < index; i++ {
			current = current.Next
		}
		return current
	} else {
		// 從尾部開始搜索
		current := dll.Tail
		for i := dll.Size - 1; i > index; i-- {
			current = current.Prev
		}
		return current
	}
}

// 正向遍歷
func (dll *DoublyLinkedList) DisplayForward() {
	fmt.Print("➡️  正向遍歷: ")
	if dll.Head == nil {
		fmt.Println("空鏈表")
		return
	}
	
	current := dll.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		current = current.Next
	}
	fmt.Printf(" (大小: %d)\n", dll.Size)
}

// 反向遍歷
func (dll *DoublyLinkedList) DisplayBackward() {
	fmt.Print("⬅️  反向遍歷: ")
	if dll.Tail == nil {
		fmt.Println("空鏈表")
		return
	}
	
	current := dll.Tail
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.Prev
	}
	fmt.Printf(" (大小: %d)\n", dll.Size)
}

// 轉換為切片
func (dll *DoublyLinkedList) ToSlice() []int {
	result := make([]int, 0, dll.Size)
	current := dll.Head
	
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	
	return result
}

// 排序鏈表
func (dll *DoublyLinkedList) Sort() {
	if dll.Size <= 1 {
		return
	}
	
	// 轉換為切片排序
	values := dll.ToSlice()
	sort.Ints(values)
	
	// 清空鏈表並重新插入排序後的值
	dll.Clear()
	for _, value := range values {
		dll.Append(value)
	}
	
	fmt.Println("🔄 鏈表已排序")
}

// 清空鏈表
func (dll *DoublyLinkedList) Clear() {
	dll.Head = nil
	dll.Tail = nil
	dll.Size = 0
}

// 檢查鏈表的完整性
func (dll *DoublyLinkedList) Validate() bool {
	if dll.Size == 0 {
		return dll.Head == nil && dll.Tail == nil
	}
	
	if dll.Size == 1 {
		return dll.Head == dll.Tail && dll.Head.Prev == nil && dll.Head.Next == nil
	}
	
	// 檢查頭節點
	if dll.Head.Prev != nil {
		fmt.Println("❌ 頭節點的 Prev 不為 nil")
		return false
	}
	
	// 檢查尾節點
	if dll.Tail.Next != nil {
		fmt.Println("❌ 尾節點的 Next 不為 nil")
		return false
	}
	
	// 正向檢查
	count := 0
	current := dll.Head
	for current != nil {
		count++
		if current.Next != nil && current.Next.Prev != current {
			fmt.Printf("❌ 節點 %d 的鏈接不一致\n", current.Value)
			return false
		}
		current = current.Next
	}
	
	if count != dll.Size {
		fmt.Printf("❌ 節點數量不匹配：期望 %d，實際 %d\n", dll.Size, count)
		return false
	}
	
	return true
}

func main() {
	fmt.Println("=== 雙向鏈表測試 ===")
	
	// 創建雙向鏈表
	dll := NewDoublyLinkedList()
	fmt.Println("📝 創建空鏈表")
	dll.DisplayForward()
	
	// 添加元素
	fmt.Println("\n📝 添加元素測試")
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.DisplayForward()
	dll.DisplayBackward()
	
	// 在開頭添加元素
	fmt.Println("\n📝 在開頭添加元素")
	dll.Prepend(0)
	dll.DisplayForward()
	
	// 在指定位置插入
	fmt.Println("\n📝 插入元素測試")
	dll.InsertAt(2, 15)
	dll.DisplayForward()
	
	// 查找元素
	fmt.Println("\n📝 查找元素測試")
	dll.Find(15)
	dll.Find(99)
	
	// 刪除元素
	fmt.Println("\n📝 刪除元素測試")
	dll.Remove(15)
	dll.DisplayForward()
	
	dll.RemoveAt(0)
	dll.DisplayForward()
	
	// 排序測試
	fmt.Println("\n📝 排序測試")
	dll.Append(10)
	dll.Append(5)
	dll.Append(8)
	fmt.Println("排序前:")
	dll.DisplayForward()
	
	dll.Sort()
	fmt.Println("排序後:")
	dll.DisplayForward()
	
	// 驗證鏈表完整性
	fmt.Println("\n📝 鏈表完整性檢查")
	if dll.Validate() {
		fmt.Println("✅ 鏈表結構正確")
	} else {
		fmt.Println("❌ 鏈表結構有誤")
	}
	
	// 詳細節點信息
	fmt.Println("\n📝 節點詳細信息")
	current := dll.Head
	index := 0
	for current != nil {
		fmt.Printf("節點 %d: 值=%d, 地址=%p", index, current.Value, current)
		if current.Prev != nil {
			fmt.Printf(", Prev=%p(值:%d)", current.Prev, current.Prev.Value)
		} else {
			fmt.Print(", Prev=nil")
		}
		if current.Next != nil {
			fmt.Printf(", Next=%p(值:%d)", current.Next, current.Next.Value)
		} else {
			fmt.Print(", Next=nil")
		}
		fmt.Println()
		current = current.Next
		index++
	}
}