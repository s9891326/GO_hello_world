package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Go 映射示例 ===")
	
	// 1. 映射基礎
	demonstrateMapBasics()
	
	// 2. 映射鍵類型
	demonstrateKeyTypes()
	
	// 3. 映射基本操作
	demonstrateMapOperations()
	
	// 4. 映射遍歷
	demonstrateMapIteration()
	
	// 5. 映射零值處理
	demonstrateMapZeroValues()
	
	// 6. 映射作為集合
	demonstrateMapAsSet()
	
	// 7. 嵌套映射
	demonstrateNestedMaps()
	
	// 8. 映射與切片結合
	demonstrateMapsWithSlices()
	
	// 9. 映射最佳實踐
	demonstrateMapBestPractices()
	
	// 10. 實際應用示例
	demonstrateRealWorldExamples()
}

func demonstrateMapBasics() {
	fmt.Println("\n--- 映射基礎 ---")
	
	// 1. 聲明映射
	var m1 map[string]int
	fmt.Printf("🗺️ nil 映射: %v (== nil: %t)\n", m1, m1 == nil)
	
	// 2. 使用 make 創建映射
	m2 := make(map[string]int)
	fmt.Printf("🗺️ 空映射: %v (== nil: %t)\n", m2, m2 == nil)
	
	// 3. 字面量初始化
	m3 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("🗺️ 字面量映射: %v\n", m3)
	
	// 4. 部分初始化
	m4 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("🗺️ 部分初始化: %v\n", m4)
	
	// 5. 空映射初始化
	m5 := map[string]int{}
	fmt.Printf("🗺️ 空映射字面量: %v (== nil: %t)\n", m5, m5 == nil)
	
	// 6. 映射不能比較（除了與 nil 比較）
	// fmt.Println(m2 == m3)  // 編譯錯誤
	fmt.Printf("🗺️ 映射長度: len(m3) = %d\n", len(m3))
}

func demonstrateKeyTypes() {
	fmt.Println("\n--- 映射鍵類型 ---")
	
	// 基本類型作為鍵
	intMap := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Printf("🔑 int 鍵: %v\n", intMap)
	
	stringMap := map[string]int{"hello": 1, "world": 2, "go": 3}
	fmt.Printf("🔑 string 鍵: %v\n", stringMap)
	
	boolMap := map[bool]string{true: "yes", false: "no"}
	fmt.Printf("🔑 bool 鍵: %v\n", boolMap)
	
	// 數組作為鍵（可比較）
	arrayMap := map[[3]int]string{
		{1, 2, 3}: "first",
		{4, 5, 6}: "second",
		{7, 8, 9}: "third",
	}
	fmt.Printf("🔑 數組鍵: %v\n", arrayMap)
	
	// 結構體作為鍵（所有字段都可比較）
	type Point struct {
		X, Y int
	}
	pointMap := map[Point]string{
		{0, 0}: "origin",
		{1, 1}: "diagonal",
		{5, 3}: "point",
	}
	fmt.Printf("🔑 結構體鍵: %v\n", pointMap)
	
	// 測試結構體鍵的使用
	p := Point{1, 1}
	if value, exists := pointMap[p]; exists {
		fmt.Printf("🔑 點 %v 對應的值: %s\n", p, value)
	}
	
	fmt.Println("🔑 注意：slice、map、function 不能作為鍵")
}

func demonstrateMapOperations() {
	fmt.Println("\n--- 映射基本操作 ---")
	
	// 創建映射
	scores := make(map[string]int)
	
	// 1. 添加/修改元素
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	scores["Diana"] = 88
	fmt.Printf("⚙️ 添加元素後: %v\n", scores)
	
	// 2. 獲取元素
	aliceScore := scores["Alice"]
	fmt.Printf("⚙️ Alice 的分數: %d\n", aliceScore)
	
	// 3. 檢查鍵是否存在
	score, exists := scores["David"]
	if exists {
		fmt.Printf("⚙️ David 的分數: %d\n", score)
	} else {
		fmt.Printf("⚙️ David 不存在，默認值: %d\n", score)
	}
	
	// 4. 安全獲取（推薦方式）
	if score, ok := scores["Alice"]; ok {
		fmt.Printf("⚙️ Alice 存在，分數: %d\n", score)
	}
	
	// 5. 修改元素
	scores["Alice"] = 98
	fmt.Printf("⚙️ 修改 Alice 分數後: %v\n", scores)
	
	// 6. 刪除元素
	delete(scores, "Bob")
	fmt.Printf("⚙️ 刪除 Bob 後: %v\n", scores)
	
	// 7. 刪除不存在的鍵（安全操作）
	delete(scores, "NonExistent")
	fmt.Printf("⚙️ 刪除不存在的鍵後: %v\n", scores)
	
	// 8. 獲取映射長度
	fmt.Printf("⚙️ 映射長度: %d\n", len(scores))
	
	// 9. 清空映射的方法
	for key := range scores {
		delete(scores, key)
	}
	fmt.Printf("⚙️ 清空後: %v (長度: %d)\n", scores, len(scores))
}

func demonstrateMapIteration() {
	fmt.Println("\n--- 映射遍歷 ---")
	
	fruits := map[string]int{
		"apple":      10,
		"banana":     5,
		"orange":     8,
		"grape":      12,
		"strawberry": 15,
	}
	
	// 1. 遍歷鍵值對
	fmt.Println("🔄 遍歷鍵值對:")
	for fruit, count := range fruits {
		fmt.Printf("   %s: %d\n", fruit, count)
	}
	
	// 2. 只遍歷鍵
	fmt.Print("🔄 只遍歷鍵: ")
	for fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()
	
	// 3. 只遍歷值
	fmt.Print("🔄 只遍歷值: ")
	for _, count := range fruits {
		fmt.Printf("%d ", count)
	}
	fmt.Println()
	
	// 4. 計算總和
	total := 0
	for _, count := range fruits {
		total += count
	}
	fmt.Printf("🔄 水果總數: %d\n", total)
	
	// 5. 查找最大值
	maxCount := 0
	maxFruit := ""
	for fruit, count := range fruits {
		if count > maxCount {
			maxCount = count
			maxFruit = fruit
		}
	}
	fmt.Printf("🔄 數量最多的水果: %s (%d)\n", maxFruit, maxCount)
	
	// 6. 注意：映射遍歷順序是隨機的
	fmt.Println("🔄 多次遍歷順序演示:")
	for i := 0; i < 3; i++ {
		fmt.Printf("   第 %d 次: ", i+1)
		count := 0
		for fruit := range fruits {
			fmt.Printf("%s ", fruit)
			count++
			if count >= 3 { // 只顯示前3個
				break
			}
		}
		fmt.Println("...")
	}
}