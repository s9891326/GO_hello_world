package main

import (
	"fmt"
	"sync"
)

func demonstrateMapZeroValues() {
	fmt.Println("\n--- 映射零值處理 ---")
	
	// 不同值類型的零值
	intMap := make(map[string]int)
	stringMap := make(map[string]string)
	boolMap := make(map[string]bool)
	sliceMap := make(map[string][]int)
	
	// 訪問不存在的鍵會返回零值
	fmt.Printf("💫 不存在的 int 鍵: %d\n", intMap["nonexistent"])
	fmt.Printf("💫 不存在的 string 鍵: '%s'\n", stringMap["nonexistent"])
	fmt.Printf("💫 不存在的 bool 鍵: %t\n", boolMap["nonexistent"])
	fmt.Printf("💫 不存在的 slice 鍵: %v (== nil: %t)\n", 
		sliceMap["nonexistent"], sliceMap["nonexistent"] == nil)
	
	// 利用零值的特性 - 計數器
	fmt.Println("💫 利用零值實現計數器:")
	counter := make(map[string]int)
	words := []string{"hello", "world", "hello", "go", "world", "hello", "programming"}
	
	for _, word := range words {
		counter[word]++ // 零值是 0，直接可以遞增
	}
	fmt.Printf("   單詞計數: %v\n", counter)
	
	// 利用零值特性 - 分組
	fmt.Println("💫 利用零值實現分組:")
	groups := make(map[int][]string)
	people := []struct {
		name string
		age  int
	}{
		{"Alice", 25}, {"Bob", 30}, {"Charlie", 25}, {"Diana", 30}, {"Eve", 25},
	}
	
	for _, person := range people {
		groups[person.age] = append(groups[person.age], person.name)
	}
	fmt.Println("   按年齡分組:")
	for age, names := range groups {
		fmt.Printf("     %d 歲: %v\n", age, names)
	}
}

func demonstrateMapAsSet() {
	fmt.Println("\n--- 映射作為集合 ---")
	
	// 使用 map[T]bool 模擬集合
	set := make(map[string]bool)
	
	// 添加元素
	items := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
	fmt.Printf("🔢 原始列表: %v\n", items)
	
	for _, item := range items {
		set[item] = true
	}
	fmt.Printf("🔢 去重後集合: %v\n", set)
	
	// 檢查元素是否存在
	fmt.Printf("🔢 apple 在集合中: %t\n", set["apple"])
	fmt.Printf("🔢 grape 在集合中: %t\n", set["grape"])
	fmt.Printf("🔢 kiwi 在集合中: %t\n", set["kiwi"])
	
	// 獲取集合大小
	fmt.Printf("🔢 集合大小: %d\n", len(set))
	
	// 遍歷集合
	fmt.Print("🔢 集合元素: ")
	for item := range set {
		fmt.Printf("%s ", item)
	}
	fmt.Println()
	
	// 集合操作
	set2 := map[string]bool{
		"banana": true,
		"grape":  true,
		"kiwi":   true,
		"mango":  true,
	}
	
	// 並集
	union := make(map[string]bool)
	for item := range set {
		union[item] = true
	}
	for item := range set2 {
		union[item] = true
	}
	fmt.Printf("🔢 並集: %v\n", getKeys(union))
	
	// 交集
	intersection := make(map[string]bool)
	for item := range set {
		if set2[item] {
			intersection[item] = true
		}
	}
	fmt.Printf("🔢 交集: %v\n", getKeys(intersection))
	
	// 使用 map[T]struct{} 節省內存
	fmt.Println("🔢 內存優化的集合:")
	efficientSet := make(map[string]struct{})
	efficientSet["item1"] = struct{}{}
	efficientSet["item2"] = struct{}{}
	efficientSet["item3"] = struct{}{}
	
	// 檢查存在性
	if _, exists := efficientSet["item1"]; exists {
		fmt.Printf("🔢 item1 存在於高效集合中\n")
	}
	
	fmt.Printf("🔢 高效集合大小: %d\n", len(efficientSet))
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func demonstrateNestedMaps() {
	fmt.Println("\n--- 嵌套映射 ---")
	
	// 二維映射：學生 -> 科目 -> 分數
	grades := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English": 87,
			"Science": 92,
			"History": 88,
		},
		"Bob": {
			"Math":    78,
			"English": 85,
			"Science": 88,
			"History": 82,
		},
		"Charlie": {
			"Math":    92,
			"English": 90,
			"Science": 85,
			"History": 91,
		},
	}
	
	// 訪問嵌套值
	fmt.Printf("📚 Alice 的數學成績: %d\n", grades["Alice"]["Math"])
	
	// 安全地訪問可能不存在的鍵
	if studentGrades, exists := grades["Diana"]; exists {
		if mathGrade, exists := studentGrades["Math"]; exists {
			fmt.Printf("📚 Diana 的數學成績: %d\n", mathGrade)
		}
	} else {
		fmt.Println("📚 Diana 不存在")
	}
	
	// 添加新學生
	grades["Diana"] = make(map[string]int)
	grades["Diana"]["Math"] = 90
	grades["Diana"]["English"] = 88
	grades["Diana"]["Science"] = 93
	grades["Diana"]["History"] = 86
	
	// 計算每個學生的平均分
	fmt.Println("📚 學生平均分:")
	for student, subjects := range grades {
		total := 0
		count := 0
		for _, grade := range subjects {
			total += grade
			count++
		}
		average := float64(total) / float64(count)
		fmt.Printf("   %s: %.1f\n", student, average)
	}
	
	// 計算每科的平均分
	fmt.Println("📚 科目平均分:")
	subjectTotals := make(map[string]int)
	subjectCounts := make(map[string]int)
	
	for _, subjects := range grades {
		for subject, grade := range subjects {
			subjectTotals[subject] += grade
			subjectCounts[subject]++
		}
	}
	
	for subject := range subjectTotals {
		average := float64(subjectTotals[subject]) / float64(subjectCounts[subject])
		fmt.Printf("   %s: %.1f\n", subject, average)
	}
	
	// 查找最高分
	maxGrade := 0
	maxStudent := ""
	maxSubject := ""
	
	for student, subjects := range grades {
		for subject, grade := range subjects {
			if grade > maxGrade {
				maxGrade = grade
				maxStudent = student
				maxSubject = subject
			}
		}
	}
	fmt.Printf("📚 最高分: %s 的 %s: %d 分\n", maxStudent, maxSubject, maxGrade)
}

func demonstrateMapsWithSlices() {
	fmt.Println("\n--- 映射與切片結合 ---")
	
	// 映射的切片
	people := []map[string]interface{}{
		{"name": "Alice", "age": 30, "city": "New York", "salary": 75000},
		{"name": "Bob", "age": 25, "city": "San Francisco", "salary": 80000},
		{"name": "Charlie", "age": 35, "city": "Chicago", "salary": 70000},
		{"name": "Diana", "age": 28, "city": "Seattle", "salary": 85000},
	}
	
	fmt.Println("👥 人員列表:")
	for i, person := range people {
		fmt.Printf("   %d: 姓名=%v, 年齡=%v, 城市=%v, 薪資=%v\n", 
			i+1, person["name"], person["age"], person["city"], person["salary"])
	}
	
	// 按城市分組
	cityGroups := make(map[string][]string)
	for _, person := range people {
		city := person["city"].(string)
		name := person["name"].(string)
		cityGroups[city] = append(cityGroups[city], name)
	}
	
	fmt.Println("👥 按城市分組:")
	for city, names := range cityGroups {
		fmt.Printf("   %s: %v\n", city, names)
	}
	
	// 切片作為映射的值
	teams := map[string][]string{
		"frontend":  {"Alice", "Bob"},
		"backend":   {"Charlie", "Diana"},
		"devops":    {"Eve"},
		"fullstack": {"Frank", "Grace"},
	}
	
	fmt.Println("👥 團隊分組:")
	for team, members := range teams {
		fmt.Printf("   %s (%d人): %v\n", team, len(members), members)
	}
	
	// 向團隊添加成員
	teams["frontend"] = append(teams["frontend"], "Helen")
	teams["backend"] = append(teams["backend"], "Ivan")
	
	fmt.Println("👥 添加成員後:")
	for team, members := range teams {
		fmt.Printf("   %s (%d人): %v\n", team, len(members), members)
	}
	
	// 統計總人數
	totalMembers := 0
	for _, members := range teams {
		totalMembers += len(members)
	}
	fmt.Printf("👥 總人數: %d\n", totalMembers)
	
	// 查找最大的團隊
	maxTeamSize := 0
	maxTeam := ""
	for team, members := range teams {
		if len(members) > maxTeamSize {
			maxTeamSize = len(members)
			maxTeam = team
		}
	}
	fmt.Printf("👥 最大團隊: %s (%d人)\n", maxTeam, maxTeamSize)
}

// 併發安全的映射示例
func demonstrateMapConcurrency() {
	fmt.Println("\n--- 映射併發安全 ---")
	
	// 使用 sync.Map 實現線程安全
	var safeMap sync.Map
	
	// 存儲值
	safeMap.Store("key1", "value1")
	safeMap.Store("key2", "value2")
	safeMap.Store("key3", "value3")
	
	// 加載值
	if value, ok := safeMap.Load("key1"); ok {
		fmt.Printf("🔒 sync.Map 值: %v\n", value)
	}
	
	// LoadOrStore：如果存在則加載，否則存儲
	actual, loaded := safeMap.LoadOrStore("key4", "value4")
	fmt.Printf("🔒 LoadOrStore - 值: %v, 是否已存在: %t\n", actual, loaded)
	
	// 刪除值
	safeMap.Delete("key2")
	
	// 遍歷 sync.Map
	fmt.Println("🔒 sync.Map 內容:")
	safeMap.Range(func(key, value interface{}) bool {
		fmt.Printf("   %v: %v\n", key, value)
		return true // 繼續遍歷
	})
	
	// 使用互斥鎖保護普通映射
	type SafeCounter struct {
		mu    sync.RWMutex
		count map[string]int
	}
	
	counter := SafeCounter{count: make(map[string]int)}
	
	// 安全的操作方法
	increment := func(key string) {
		counter.mu.Lock()
		defer counter.mu.Unlock()
		counter.count[key]++
	}
	
	getValue := func(key string) int {
		counter.mu.RLock()
		defer counter.mu.RUnlock()
		return counter.count[key]
	}
	
	getAll := func() map[string]int {
		counter.mu.RLock()
		defer counter.mu.RUnlock()
		result := make(map[string]int)
		for k, v := range counter.count {
			result[k] = v
		}
		return result
	}
	
	// 使用安全計數器
	increment("clicks")
	increment("views")
	increment("clicks")
	increment("downloads")
	increment("views")
	
	fmt.Printf("🔒 安全計數器 clicks: %d\n", getValue("clicks"))
	fmt.Printf("🔒 安全計數器 views: %d\n", getValue("views"))
	fmt.Printf("🔒 所有計數: %v\n", getAll())
}