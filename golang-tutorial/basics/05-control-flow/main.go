package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go 流程控制示例 ===")
	
	// 1. 條件語句演示
	demonstrateIf()
	
	// 2. 循環語句演示
	demonstrateFor()
	
	// 3. 選擇語句演示
	demonstrateSwitch()
	
	// 4. 跳轉語句演示
	demonstrateJump()
	
	// 5. 實際應用示例
	demonstrateRealExamples()
}

func demonstrateIf() {
	fmt.Println("\n--- 條件語句演示 ---")
	
	// 基本 if 語句
	age := 20
	if age >= 18 {
		fmt.Printf("年齡 %d：已成年\n", age)
	}
	
	// if-else 語句
	temperature := 25
	if temperature > 30 {
		fmt.Println("天氣很熱")
	} else {
		fmt.Println("天氣不錯")
	}
	
	// if-else if-else 語句
	score := 85
	if score >= 90 {
		fmt.Println("成績：優秀")
	} else if score >= 80 {
		fmt.Println("成績：良好")
	} else if score >= 70 {
		fmt.Println("成績：中等")
	} else if score >= 60 {
		fmt.Println("成績：及格")
	} else {
		fmt.Println("成績：不及格")
	}
	
	// 帶初始化的 if 語句
	if currentYear := 2024; currentYear-1990 >= 18 {
		fmt.Println("1990年出生的人已成年")
	}
	
	// 錯誤處理模式
	if email := "test@example.com"; validateEmail(email) {
		fmt.Printf("郵箱 %s 格式正確\n", email)
	} else {
		fmt.Printf("郵箱 %s 格式錯誤\n", email)
	}
}

func validateEmail(email string) bool {
	return len(email) > 0 && strings.Contains(email, "@") && strings.Contains(email, ".")
}

func demonstrateFor() {
	fmt.Println("\n--- 循環語句演示 ---")
	
	// 基本 for 循環
	fmt.Println("基本 for 循環：")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  第 %d 次循環\n", i)
	}
	
	// while 風格的 for 循環
	fmt.Println("while 風格循環：")
	count := 1
	for count <= 3 {
		fmt.Printf("  計數：%d\n", count)
		count++
	}
	
	// for-range 遍歷切片
	fmt.Println("遍歷數字切片：")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("  索引 %d：值 %d\n", index, value)
	}
	
	// for-range 遍歷字符串
	fmt.Println("遍歷字符串：")
	text := "Hello"
	for i, char := range text {
		fmt.Printf("  位置 %d：字符 %c\n", i, char)
	}
	
	// for-range 遍歷映射
	fmt.Println("遍歷映射：")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("  %s：%d 歲\n", name, age)
	}
}

func demonstrateSwitch() {
	fmt.Println("\n--- 選擇語句演示 ---")
	
	// 基本 switch 語句
	day := 3
	fmt.Printf("今天是星期 %d：", day)
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7:
		fmt.Println("週末")
	default:
		fmt.Println("無效日期")
	}
	
	// 無表達式的 switch
	score := 75
	fmt.Printf("分數 %d 的等級：", score)
	switch {
	case score >= 90:
		fmt.Println("A級")
	case score >= 80:
		fmt.Println("B級")
	case score >= 70:
		fmt.Println("C級")
	case score >= 60:
		fmt.Println("D級")
	default:
		fmt.Println("F級")
	}
	
	// 帶初始化的 switch
	switch hour := time.Now().Hour(); {
	case hour < 6:
		fmt.Println("凌晨時間")
	case hour < 12:
		fmt.Println("上午時間")
	case hour < 18:
		fmt.Println("下午時間")
	default:
		fmt.Println("晚上時間")
	}
}

func demonstrateJump() {
	fmt.Println("\n--- 跳轉語句演示 ---")
	
	// break 示例
	fmt.Println("break 示例：")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			fmt.Printf("  在 %d 處中斷\n", i)
			break
		}
		fmt.Printf("  i = %d\n", i)
	}
	
	// continue 示例
	fmt.Println("continue 示例（只顯示奇數）：")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳過偶數
		}
		fmt.Printf("  奇數：%d\n", i)
	}
	
	// 嵌套循環的標籤使用
	fmt.Println("嵌套循環標籤示例：")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("  在 i=%d, j=%d 處跳出外層循環\n", i, j)
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
}

func demonstrateRealExamples() {
	fmt.Println("\n--- 實際應用示例 ---")
	
	// 1. 用戶認證系統
	demonstrateUserAuth()
	
	// 2. 數據處理
	demonstrateDataProcessing()
	
	// 3. 遊戲邏輯
	demonstrateGameLogic()
}

func demonstrateUserAuth() {
	fmt.Println("\n用戶認證系統：")
	
	users := map[string]string{
		"admin": "admin123",
		"user1": "password1",
		"user2": "password2",
	}
	
	testCredentials := []struct {
		username, password string
	}{
		{"admin", "admin123"},
		{"user1", "wrongpass"},
		{"user2", "password2"},
		{"unknown", "test"},
	}
	
	for _, cred := range testCredentials {
		fmt.Printf("嘗試登錄 %s: ", cred.username)
		
		if password, exists := users[cred.username]; !exists {
			fmt.Println("用戶不存在")
			continue
		} else if password != cred.password {
			fmt.Println("密碼錯誤")
			continue
		} else {
			fmt.Println("登錄成功")
		}
	}
}

func demonstrateDataProcessing() {
	fmt.Println("\n數據處理示例：")
	
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// 計算統計信息
	var sum, evenCount, oddCount int
	var max, min int = numbers[0], numbers[0]
	
	for _, num := range numbers {
		sum += num
		
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	
	average := float64(sum) / float64(len(numbers))
	
	fmt.Printf("數據：%v\n", numbers)
	fmt.Printf("總和：%d\n", sum)
	fmt.Printf("平均值：%.2f\n", average)
	fmt.Printf("最大值：%d\n", max)
	fmt.Printf("最小值：%d\n", min)
	fmt.Printf("偶數個數：%d\n", evenCount)
	fmt.Printf("奇數個數：%d\n", oddCount)
}

func demonstrateGameLogic() {
	fmt.Println("\n簡單猜數字遊戲邏輯：")
	
	// 設置隨機數種子
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	
	guesses := []int{50, 75, 60, 65, 63, 64}
	
	fmt.Printf("目標數字：%d\n", target)
	fmt.Println("遊戲開始！")
	
	for attempt, guess := range guesses {
		fmt.Printf("第 %d 次猜測：%d ", attempt+1, guess)
		
		switch {
		case guess == target:
			fmt.Println("🎉 恭喜！猜中了！")
			return
		case guess < target:
			fmt.Println("太小了，再試試更大的數字")
		case guess > target:
			fmt.Println("太大了，再試試更小的數字")
		}
		
		// 給出距離提示
		distance := target - guess
		if distance < 0 {
			distance = -distance
		}
		
		switch {
		case distance <= 5:
			fmt.Println("  提示：非常接近了！")
		case distance <= 10:
			fmt.Println("  提示：很接近了！")
		case distance <= 20:
			fmt.Println("  提示：比較接近")
		default:
			fmt.Println("  提示：還差得遠")
		}
	}
	
	fmt.Printf("遊戲結束！正確答案是 %d\n", target)
}