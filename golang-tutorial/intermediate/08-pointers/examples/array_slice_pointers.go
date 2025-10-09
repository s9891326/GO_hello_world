package main

import "fmt"

func demonstrateArrayPointers() {
	fmt.Println("\n--- 指針與數組 ---")
	
	// 數組
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("📊 原始數組: %v\n", arr)
	
	// 數組指針：指向整個數組的指針
	arrPtr := &arr
	fmt.Printf("📊 數組指針指向的數組: %v\n", *arrPtr)
	fmt.Printf("📊 數組指針地址: %p\n", arrPtr)
	fmt.Printf("📊 數組地址: %p\n", &arr)
	
	// 通過數組指針修改數組
	(*arrPtr)[0] = 100
	fmt.Printf("📊 通過數組指針修改後: %v\n", arr)
	
	// Go 的語法糖：自動解引用
	arrPtr[1] = 200  // 等同於 (*arrPtr)[1] = 200
	fmt.Printf("📊 語法糖修改後: %v\n", arr)
	
	// 指針數組：存儲指針的數組
	fmt.Println("\n📊 指針數組示例:")
	var ptrArray [3]*int
	a, b, c := 10, 20, 30
	ptrArray[0] = &a
	ptrArray[1] = &b
	ptrArray[2] = &c
	
	fmt.Printf("📊 變數地址: a=%p, b=%p, c=%p\n", &a, &b, &c)
	fmt.Printf("📊 指針數組: [%p, %p, %p]\n", ptrArray[0], ptrArray[1], ptrArray[2])
	fmt.Printf("📊 指針數組指向的值: [%d, %d, %d]\n", *ptrArray[0], *ptrArray[1], *ptrArray[2])
	
	// 修改指針數組指向的值
	*ptrArray[0] = 100
	*ptrArray[1] = 200
	*ptrArray[2] = 300
	fmt.Printf("📊 修改後的值: a=%d, b=%d, c=%d\n", a, b, c)
	
	// 數組元素的指針
	fmt.Println("\n📊 數組元素指針:")
	for i := range arr {
		elementPtr := &arr[i]
		fmt.Printf("📊 arr[%d] 地址: %p, 值: %d\n", i, elementPtr, *elementPtr)
	}
}

func demonstrateSlicePointers() {
	fmt.Println("\n--- 指針與切片 ---")
	
	// 切片本身就是引用類型
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("🍕 原始切片: %v\n", slice)
	fmt.Printf("🍕 切片長度: %d, 容量: %d\n", len(slice), cap(slice))
	
	// 切片指針
	slicePtr := &slice
	fmt.Printf("🍕 切片指針指向的切片: %v\n", *slicePtr)
	
	// 通過切片指針修改
	(*slicePtr)[0] = 100
	fmt.Printf("🍕 通過切片指針修改後: %v\n", slice)
	
	// 切片指針添加元素
	*slicePtr = append(*slicePtr, 6, 7)
	fmt.Printf("🍕 通過切片指針添加元素後: %v\n", slice)
	
	// 切片元素的指針
	fmt.Println("\n🍕 切片元素指針:")
	elementPtr := &slice[1]
	fmt.Printf("🍕 第二個元素的地址: %p, 值: %d\n", elementPtr, *elementPtr)
	
	*elementPtr = 200
	fmt.Printf("🍕 修改元素後的切片: %v\n", slice)
	
	// 指針切片：存儲指針的切片
	fmt.Println("\n🍕 指針切片示例:")
	var ptrSlice []*int
	
	// 創建一些變數
	values := []int{10, 20, 30, 40, 50}
	for i := range values {
		ptrSlice = append(ptrSlice, &values[i])
	}
	
	fmt.Printf("🍕 指針切片長度: %d\n", len(ptrSlice))
	fmt.Print("🍕 指針切片指向的值: [")
	for i, ptr := range ptrSlice {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", *ptr)
	}
	fmt.Println("]")
	
	// 修改指針切片指向的值
	for i, ptr := range ptrSlice {
		*ptr = *ptr * 10
		fmt.Printf("🍕 修改 values[%d] = %d\n", i, values[i])
	}
	fmt.Printf("🍕 修改後的 values: %v\n", values)
	
	// 切片傳遞函數示例
	fmt.Println("\n🍕 切片函數傳遞:")
	originalSlice := []int{1, 2, 3}
	fmt.Printf("🍕 函數調用前: %v\n", originalSlice)
	
	modifySlice(originalSlice)
	fmt.Printf("🍕 修改元素後: %v\n", originalSlice)
	
	modifySlicePointer(&originalSlice)
	fmt.Printf("🍕 修改切片本身後: %v\n", originalSlice)
}

// 修改切片元素（切片本身是引用類型）
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100  // 會影響原切片
	}
}

// 修改切片本身（需要切片指針）
func modifySlicePointer(s *[]int) {
	*s = append(*s, 4, 5, 6)  // 會影響原切片
}