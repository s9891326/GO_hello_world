package main

import "fmt"

func demonstrateSliceTraps() {
	fmt.Println("\n--- 切片陷阱 ---")
	
	// 陷阱 1：切片共享底層數組
	fmt.Println("⚠️ 陷阱 1: 切片共享底層數組")
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:3]  // [2, 3]
	slice2 := arr[2:4]  // [3, 4]
	
	fmt.Printf("   原數組: %v\n", arr)
	fmt.Printf("   slice1 (arr[1:3]): %v\n", slice1)
	fmt.Printf("   slice2 (arr[2:4]): %v\n", slice2)
	
	slice1[1] = 100  // 修改 slice1[1]，實際修改的是 arr[2]
	fmt.Printf("   修改 slice1[1] = 100 後:\n")
	fmt.Printf("   原數組: %v\n", arr)
	fmt.Printf("   slice1: %v\n", slice1)
	fmt.Printf("   slice2: %v (slice2[0] 也變了！)\n", slice2)
	
	// 陷阱 2：append 可能改變底層數組
	fmt.Println("\n⚠️ 陷阱 2: append 可能改變底層數組")
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice3 := arr2[1:3]  // [2, 3]，容量為 4
	
	fmt.Printf("   原數組: %v\n", arr2)
	fmt.Printf("   slice3 (arr2[1:3]): %v (容量: %d)\n", slice3, cap(slice3))
	
	slice3 = append(slice3, 99)  // 容量足夠，直接修改底層數組
	fmt.Printf("   append 99 後:\n")
	fmt.Printf("   原數組: %v (arr2[3] 變成了 99)\n", arr2)
	fmt.Printf("   slice3: %v\n", slice3)
	
	// 陷阱 3：切片擴容後脫離原數組
	fmt.Println("\n⚠️ 陷阱 3: 切片擴容後脫離原數組")
	arr3 := [3]int{1, 2, 3}
	slice4 := arr3[:]  // 容量為 3
	
	fmt.Printf("   原數組: %v\n", arr3)
	fmt.Printf("   slice4: %v (容量: %d)\n", slice4, cap(slice4))
	
	slice4 = append(slice4, 4, 5, 6)  // 超出容量，分配新數組
	fmt.Printf("   append 4,5,6 後:\n")
	fmt.Printf("   原數組: %v (沒有變化)\n", arr3)
	fmt.Printf("   slice4: %v (新數組，容量: %d)\n", slice4, cap(slice4))
	
	slice4[0] = 999
	fmt.Printf("   修改 slice4[0] = 999 後:\n")
	fmt.Printf("   原數組: %v (仍然沒有變化)\n", arr3)
	fmt.Printf("   slice4: %v\n", slice4)
	
	// 陷阱 4：nil 切片 vs 空切片
	fmt.Println("\n⚠️ 陷阱 4: nil 切片 vs 空切片")
	var nilSlice []int
	emptySlice := []int{}
	makeSlice := make([]int, 0)
	
	fmt.Printf("   nil 切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("   空切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("   make 空切片: %v (長度: %d, 容量: %d, == nil: %t)\n", 
		makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)
	
	// 在 JSON 編碼時的差異
	// nil 切片編碼為 null，空切片編碼為 []
}

func demonstrateSliceGrowth() {
	fmt.Println("\n--- 切片擴容機制 ---")
	
	slice := make([]int, 0, 1)
	fmt.Printf("📈 初始: 長度=%d, 容量=%d\n", len(slice), cap(slice))
	
	for i := 1; i <= 20; i++ {
		slice = append(slice, i)
		fmt.Printf("📈 append %2d: 長度=%2d, 容量=%2d", i, len(slice), cap(slice))
		
		// 檢查是否發生了擴容
		if i > 1 && cap(slice) > cap(slice[:len(slice)-1]) {
			oldCap := len(slice) - 1
			if oldCap == 0 {
				oldCap = 1
			}
			ratio := float64(cap(slice)) / float64(oldCap)
			fmt.Printf(" (擴容: %.1fx)", ratio)
		}
		fmt.Println()
	}
	
	fmt.Println("📈 擴容規律總結:")
	fmt.Println("   - 當容量 < 1024 時，新容量約為舊容量的 2 倍")
	fmt.Println("   - 當容量 >= 1024 時，新容量約為舊容量的 1.25 倍")
	fmt.Println("   - 實際容量會根據內存對齊進行調整")
}

func demonstrateSliceMemoryOptimization() {
	fmt.Println("\n--- 切片內存優化 ---")
	
	// 問題：大切片的小切片可能導致內存泄漏
	fmt.Println("💾 內存泄漏問題演示:")
	largeSlice := make([]int, 1000000)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	fmt.Printf("   大切片: 長度=%d, 容量=%d\n", len(largeSlice), cap(largeSlice))
	
	// 不好的做法：保留對大切片的引用
	smallSliceBad := largeSlice[0:5]
	fmt.Printf("   不好的小切片: 長度=%d, 容量=%d (仍引用大數組)\n", 
		len(smallSliceBad), cap(smallSliceBad))
	
	// 好的做法：複製需要的部分
	smallSliceGood := make([]int, 5)
	copy(smallSliceGood, largeSlice[0:5])
	fmt.Printf("   好的小切片: 長度=%d, 容量=%d (獨立數組)\n", 
		len(smallSliceGood), cap(smallSliceGood))
	
	fmt.Println("💾 內存優化建議:")
	fmt.Println("   - 如果只需要大切片的一小部分，使用 copy 創建獨立切片")
	fmt.Println("   - 避免長期持有大切片的小切片引用")
	fmt.Println("   - 適當預分配切片容量，減少擴容次數")
	
	// 預分配示例
	fmt.Println("\n💾 預分配容量示例:")
	
	// 不好的做法：頻繁擴容
	var badSlice []int
	fmt.Printf("   頻繁擴容前: 容量=%d\n", cap(badSlice))
	for i := 0; i < 1000; i++ {
		badSlice = append(badSlice, i)
	}
	fmt.Printf("   頻繁擴容後: 容量=%d\n", cap(badSlice))
	
	// 好的做法：預分配容量
	goodSlice := make([]int, 0, 1000)
	fmt.Printf("   預分配前: 容量=%d\n", cap(goodSlice))
	for i := 0; i < 1000; i++ {
		goodSlice = append(goodSlice, i)
	}
	fmt.Printf("   預分配後: 容量=%d (無擴容)\n", cap(goodSlice))
}

func demonstrateSliceAsParameter() {
	fmt.Println("\n--- 切片作為函數參數 ---")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("🔄 原始切片: %v (長度: %d, 容量: %d)\n", slice, len(slice), cap(slice))
	
	// 修改切片元素
	fmt.Println("🔄 修改切片元素:")
	modifySliceElements(slice)
	fmt.Printf("   修改元素後: %v\n", slice)
	
	// 嘗試修改切片本身（不會影響原切片）
	fmt.Println("🔄 嘗試修改切片本身:")
	tryModifySlice(slice)
	fmt.Printf("   嘗試修改切片後: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
	
	// 正確修改切片本身的方法
	fmt.Println("🔄 正確修改切片本身:")
	slice = correctModifySlice(slice)
	fmt.Printf("   正確修改切片後: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
	
	// 使用指針修改切片
	fmt.Println("🔄 使用指針修改切片:")
	modifySliceByPointer(&slice)
	fmt.Printf("   指針修改後: %v (長度: %d, 容量: %d)\n", 
		slice, len(slice), cap(slice))
}

func modifySliceElements(s []int) {
	fmt.Printf("   函數內修改前: %v\n", s)
	for i := range s {
		s[i] *= 2
	}
	fmt.Printf("   函數內修改後: %v\n", s)
}

func tryModifySlice(s []int) {
	fmt.Printf("   函數內修改前: %v (長度: %d, 容量: %d)\n", s, len(s), cap(s))
	s = append(s, 100)
	fmt.Printf("   函數內修改後: %v (長度: %d, 容量: %d)\n", s, len(s), cap(s))
	fmt.Println("   注意：這個修改不會影響原切片！")
}

func correctModifySlice(s []int) []int {
	fmt.Printf("   函數內修改前: %v\n", s)
	result := append(s, 200)
	fmt.Printf("   函數內修改後: %v\n", result)
	return result
}

func modifySliceByPointer(s *[]int) {
	fmt.Printf("   函數內修改前: %v\n", *s)
	*s = append(*s, 300)
	fmt.Printf("   函數內修改後: %v\n", *s)
}