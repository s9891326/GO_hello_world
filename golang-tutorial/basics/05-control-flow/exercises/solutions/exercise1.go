// 練習 1 解答：猜數字遊戲
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== 猜數字遊戲 ===")
	
	for {
		playGuessGame()
		
		fmt.Print("再玩一局嗎？(y/n): ")
		var playAgain string
		fmt.Scanf("%s", &playAgain)
		
		if playAgain != "y" && playAgain != "Y" {
			fmt.Println("謝謝遊戲！再見！")
			break
		}
		fmt.Println()
	}
}

func playGuessGame() {
	target := rand.Intn(100) + 1
	maxAttempts := 7
	
	fmt.Printf("我想了一個 1-100 的數字，你能猜中嗎？\n")
	fmt.Printf("你有 %d 次機會！\n\n", maxAttempts)
	
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		var guess int
		fmt.Printf("第 %d 次猜測，請輸入數字：", attempt)
		fmt.Scanf("%d", &guess)
		
		// 驗證輸入範圍
		if guess < 1 || guess > 100 {
			fmt.Println("請輸入 1-100 之間的數字！")
			attempt-- // 不計入失敗次數
			continue
		}
		
		// 判斷猜測結果
		switch {
		case guess == target:
			fmt.Printf("🎉 恭喜！你猜中了！\n")
			evaluatePerformance(attempt)
			return
		case guess < target:
			distance := target - guess
			fmt.Printf("太小了！請猜大一點的數字")
			giveHint(distance)
		case guess > target:
			distance := guess - target
			fmt.Printf("太大了！請猜小一點的數字")
			giveHint(distance)
		}
		
		// 顯示剩餘機會
		remaining := maxAttempts - attempt
		if remaining > 0 {
			fmt.Printf("還有 %d 次機會\n", remaining)
		}
		fmt.Println()
	}
	
	fmt.Printf("💔 很遺憾，你沒有猜中！正確答案是 %d\n", target)
}

func giveHint(distance int) {
	switch {
	case distance <= 3:
		fmt.Printf("（非常接近了！）")
	case distance <= 8:
		fmt.Printf("（很接近了！）")
	case distance <= 15:
		fmt.Printf("（比較接近）")
	default:
		fmt.Printf("（還差得遠）")
	}
	fmt.Println()
}

func evaluatePerformance(attempts int) {
	switch {
	case attempts == 1:
		fmt.Println("太神了！一次就猜中，你是神算子！")
	case attempts <= 3:
		fmt.Printf("太厲害了！你只用了 %d 次就猜中了！\n", attempts)
	case attempts <= 5:
		fmt.Printf("不錯！你用了 %d 次猜中，表現很好！\n", attempts)
	default:
		fmt.Printf("雖然用了 %d 次，但最終還是猜中了！\n", attempts)
	}
}