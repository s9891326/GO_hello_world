package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go 通道示例 ===")
	
	// 1. 通道基礎
	demonstrateChannelBasics()
	
	// 2. 通道基本操作
	demonstrateChannelOperations()
	
	// 3. 無緩衝通道
	demonstrateUnbufferedChannel()
	
	// 4. 握手通信
	demonstrateHandshake()
	
	// 5. 有緩衝通道
	demonstrateBufferedChannel()
	
	// 6. 生產者消費者（緩衝）
	demonstrateProducerConsumerWithBuffer()
	
	// 7. 通道關閉
	demonstrateChannelClosure()
	
	// 8. 通道方向性
	demonstrateChannelDirections()
	
	// 9. 基本 Select
	demonstrateBasicSelect()
	
	// 10. Select 默認分支
	demonstrateSelectWithDefault()
	
	// 11. 高級 Select
	demonstrateAdvancedSelect()
	
	// 12. 扇入模式
	demonstrateFanIn()
	
	// 13. 管道模式
	demonstratePipeline()
	
	// 14. 死鎖預防
	demonstrateDeadlock()
	
	// 15. 優雅關閉
	demonstrateGracefulShutdown()
}

func demonstrateChannelBasics() {
	fmt.Println("\n--- 通道基礎演示 ---")
	
	// 創建無緩衝通道
	ch1 := make(chan int)
	ch2 := make(chan string)
	
	// 創建有緩衝通道
	ch3 := make(chan int, 5)
	ch4 := make(chan string, 3)
	
	fmt.Printf("📡 無緩衝整數通道: %T\n", ch1)
	fmt.Printf("📡 無緩衝字符串通道: %T\n", ch2)
	fmt.Printf("📦 有緩衝整數通道容量: %d\n", cap(ch3))
	fmt.Printf("📦 有緩衝字符串通道容量: %d\n", cap(ch4))
	
	// 測試通道狀態
	fmt.Printf("📊 ch3 長度/容量: %d/%d\n", len(ch3), cap(ch3))
	ch3 <- 1
	ch3 <- 2
	fmt.Printf("📊 添加2個元素後 ch3 長度/容量: %d/%d\n", len(ch3), cap(ch3))
	
	// 關閉通道（避免未使用變數錯誤）
	close(ch1)
	close(ch2)
	close(ch3)
	close(ch4)
}

func demonstrateChannelOperations() {
	fmt.Println("\n--- 通道基本操作 ---")
	
	// 無緩衝通道 - 同步通信
	ch := make(chan string)
	
	fmt.Println("🔄 無緩衝通道同步通信:")
	// 啟動接收協程
	go func() {
		message := <-ch // 接收數據
		fmt.Printf("   ✅ 接收到消息: %s\n", message)
	}()
	
	fmt.Println("   📤 準備發送消息...")
	ch <- "Hello, Channel!" // 發送數據（會阻塞直到有接收者）
	fmt.Println("   📤 消息發送完成")
	
	time.Sleep(100 * time.Millisecond)
	
	// 有緩衝通道 - 異步通信
	fmt.Println("\n🔄 有緩衝通道異步通信:")
	bufferedCh := make(chan int, 3)
	
	// 發送數據（不會阻塞，因為有緩衝）
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	fmt.Printf("   📦 緩衝通道狀態: %d/%d\n", len(bufferedCh), cap(bufferedCh))
	
	// 接收數據
	for i := 0; i < 3; i++ {
		value := <-bufferedCh
		fmt.Printf("   📥 從緩衝通道接收: %d (剩餘: %d)\n", value, len(bufferedCh))
	}
}

func demonstrateUnbufferedChannel() {
	fmt.Println("\n--- 無緩衝通道演示 ---")
	
	ch := make(chan string)
	
	fmt.Println("🤝 無緩衝通道的同步特性:")
	// 無緩衝通道的發送和接收是同步的
	go func() {
		fmt.Println("   🔄 協程準備發送數據...")
		ch <- "同步消息"
		fmt.Println("   ✅ 協程發送完成")
	}()
	
	time.Sleep(100 * time.Millisecond) // 模擬主協程在做其他事
	
	fmt.Println("   📥 主協程準備接收數據...")
	message := <-ch
	fmt.Printf("   ✅ 主協程接收到: %s\n", message)
}

func demonstrateHandshake() {
	fmt.Println("\n--- 握手通信演示 ---")
	
	done := make(chan bool)
	
	fmt.Println("🤝 Worker-Main 握手通信:")
	go func() {
		fmt.Println("   👷 Worker: 開始工作...")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("   👷 Worker: 工作完成")
		done <- true // 發送完成信號
	}()
	
	fmt.Println("   📋 Main: 等待 Worker 完成...")
	<-done // 等待完成信號
	fmt.Println("   ✅ Main: 收到完成信號，繼續執行")
}

func demonstrateBufferedChannel() {
	fmt.Println("\n--- 有緩衝通道演示 ---")
	
	// 創建容量為 3 的緩衝通道
	ch := make(chan int, 3)
	
	fmt.Println("📦 緩衝通道發送過程:")
	// 發送數據（不會阻塞）
	fmt.Println("   📤 發送數據到緩衝通道...")
	
	ch <- 1
	fmt.Printf("   📤 發送 1，通道狀態: %d/%d\n", len(ch), cap(ch))
	
	ch <- 2
	fmt.Printf("   📤 發送 2，通道狀態: %d/%d\n", len(ch), cap(ch))
	
	ch <- 3
	fmt.Printf("   📤 發送 3，通道狀態: %d/%d (已滿)\n", len(ch), cap(ch))
	
	// 現在通道已滿，再發送會阻塞
	
	// 開始接收數據
	fmt.Println("\n   📥 從緩衝通道接收數據...")
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("   📥 接收 %d，通道狀態: %d/%d\n", value, len(ch), cap(ch))
		time.Sleep(50 * time.Millisecond)
	}
}

func demonstrateProducerConsumerWithBuffer() {
	fmt.Println("\n--- 緩衝通道生產者消費者 ---")
	
	ch := make(chan int, 5) // 緩衝大小為5
	
	fmt.Println("🏭 啟動生產者-消費者模式:")
	
	// 生產者
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("   🏭 生產者生產: %d (緩衝: %d/%d)\n", i, len(ch), cap(ch))
			time.Sleep(50 * time.Millisecond)
		}
		close(ch)
		fmt.Println("   🏁 生產者完成")
	}()
	
	// 消費者（延遲啟動）
	time.Sleep(200 * time.Millisecond) // 讓生產者先生產一些
	fmt.Println("   🍽️ 消費者開始消費:")
	
	for value := range ch {
		fmt.Printf("   🍽️ 消費者消費: %d (緩衝: %d/%d)\n", value, len(ch), cap(ch))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("   🏁 消費者完成")
}

func demonstrateChannelClosure() {
	fmt.Println("\n--- 通道關閉演示 ---")
	
	ch := make(chan int, 3)
	
	fmt.Println("🔒 通道關閉和接收:")
	// 發送一些數據
	ch <- 1
	ch <- 2
	ch <- 3
	
	// 關閉通道
	close(ch)
	fmt.Println("   🔒 通道已關閉")
	
	// 關閉後仍可以接收數據
	fmt.Println("   📥 關閉通道後接收數據:")
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("   ❌ 通道已關閉且無數據")
			break
		}
		fmt.Printf("   ✅ 接收到: %d\n", value)
	}
	
	// 使用 range 遍歷已關閉的通道
	ch2 := make(chan string, 2)
	ch2 <- "Hello"
	ch2 <- "World"
	close(ch2)
	
	fmt.Println("\n   🔄 使用 range 遍歷:")
	for msg := range ch2 {
		fmt.Printf("   📨 接收到: %s\n", msg)
	}
}

// 只能發送的通道
func sender(ch chan<- string) {
	ch <- "Hello from sender"
	fmt.Println("   📤 Sender 發送完成")
}

// 只能接收的通道
func receiver(ch <-chan string) {
	message := <-ch
	fmt.Printf("   📥 Receiver 收到: %s\n", message)
}

func demonstrateChannelDirections() {
	fmt.Println("\n--- 通道方向性演示 ---")
	
	ch := make(chan string, 1)
	
	fmt.Println("🎯 單向通道演示:")
	// 將雙向通道轉為單向通道
	go sender(ch)   // 傳遞為只寫通道
	
	time.Sleep(50 * time.Millisecond)
	
	go receiver(ch) // 傳遞為只讀通道
	
	time.Sleep(100 * time.Millisecond)
}

func demonstrateBasicSelect() {
	fmt.Println("\n--- 基本 Select 演示 ---")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	fmt.Println("🔀 Select 多通道監聽:")
	
	// 啟動兩個發送協程
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "來自通道1的消息"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "來自通道2的消息"
	}()
	
	// 使用 select 同時監聽多個通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   📨 收到通道1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   📨 收到通道2: %s\n", msg2)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("   ⏰ 接收超時")
		}
	}
}

func demonstrateSelectWithDefault() {
	fmt.Println("\n--- Select 默認分支演示 ---")
	
	ch := make(chan int, 1)
	
	fmt.Println("🔀 非阻塞操作:")
	
	// 非阻塞發送
	select {
	case ch <- 42:
		fmt.Println("   ✅ 成功發送到通道")
	default:
		fmt.Println("   ❌ 通道已滿，無法發送")
	}
	
	// 非阻塞接收
	select {
	case value := <-ch:
		fmt.Printf("   📥 接收到值: %d\n", value)
	default:
		fmt.Println("   ❌ 通道為空，無法接收")
	}
	
	// 再次嘗試接收
	select {
	case value := <-ch:
		fmt.Printf("   📥 接收到值: %d\n", value)
	default:
		fmt.Println("   🔄 通道為空，使用默認分支")
	}
}

func demonstrateAdvancedSelect() {
	fmt.Println("\n--- 高級 Select 演示 ---")
	
	fmt.Println("⏰ 超時控制演示:")
	
	// 超時控制
	result := make(chan string, 1)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		result <- "操作完成"
	}()
	
	select {
	case res := <-result:
		fmt.Printf("   ✅ 操作結果: %s\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("   ⏰ 操作超時")
	}
	
	// 心跳檢測
	demonstrateHeartbeat()
}

func demonstrateHeartbeat() {
	fmt.Println("\n💓 心跳檢測演示:")
	
	heartbeat := time.Tick(100 * time.Millisecond)
	work := make(chan string)
	
	go func() {
		time.Sleep(250 * time.Millisecond)
		work <- "工作完成"
	}()
	
	for {
		select {
		case <-heartbeat:
			fmt.Println("   💓 心跳")
		case result := <-work:
			fmt.Printf("   📋 %s\n", result)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("   ⏰ 整體超時")
			return
		}
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	output := make(chan string)
	
	go func() {
		defer close(output)
		for {
			select {
			case s, ok := <-input1:
				if !ok {
					input1 = nil
				} else {
					output <- s
				}
			case s, ok := <-input2:
				if !ok {
					input2 = nil
				} else {
					output <- s
				}
			}
			
			if input1 == nil && input2 == nil {
				return
			}
		}
	}()
	
	return output
}

func demonstrateFanIn() {
	fmt.Println("\n--- 扇入模式演示 ---")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	fmt.Println("🌪️ 扇入模式 - 合併多個通道:")
	
	// 啟動兩個生產者
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("生產者1-消息%d", i+1)
			ch1 <- msg
			fmt.Printf("   📤 %s\n", msg)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("   🏁 生產者1完成")
	}()
	
	go func() {
		defer close(ch2)
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("生產者2-消息%d", i+1)
			ch2 <- msg
			fmt.Printf("   📤 %s\n", msg)
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Println("   🏁 生產者2完成")
	}()
	
	// 合並兩個通道
	merged := fanIn(ch1, ch2)
	
	// 接收合併的消息
	fmt.Println("   📥 扇入接收:")
	for msg := range merged {
		fmt.Printf("   🔄 扇入接收: %s\n", msg)
	}
	fmt.Println("🎯 扇入演示完成")
}

func demonstratePipeline() {
	fmt.Println("\n--- 管道模式演示 ---")
	
	fmt.Println("🚰 數據處理管道:")
	
	// 階段1：生成數字
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("   🔢 生成: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// 階段2：計算平方
	squares := make(chan int)
	go func() {
		defer close(squares)
		for num := range numbers {
			square := num * num
			squares <- square
			fmt.Printf("   📐 平方: %d -> %d\n", num, square)
		}
	}()
	
	// 階段3：過濾偶數
	evens := make(chan int)
	go func() {
		defer close(evens)
		for square := range squares {
			if square%2 == 0 {
				evens <- square
				fmt.Printf("   🔍 偶數: %d\n", square)
			} else {
				fmt.Printf("   ❌ 跳過奇數: %d\n", square)
			}
		}
	}()
	
	// 最終消費
	fmt.Println("   📋 最終結果:")
	for even := range evens {
		fmt.Printf("   ✅ 管道輸出: %d\n", even)
	}
}

func demonstrateDeadlock() {
	fmt.Println("\n--- 死鎖預防演示 ---")
	
	fmt.Println("🚫 死鎖預防技巧:")
	
	// 正確示例1：使用協程
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	value := <-ch1
	fmt.Printf("   ✅ 協程方式接收: %d\n", value)
	
	// 正確示例2：使用緩衝通道
	ch2 := make(chan int, 1)
	ch2 <- 2 // 不會阻塞
	value2 := <-ch2
	fmt.Printf("   ✅ 緩衝通道接收: %d\n", value2)
	
	// 正確示例3：使用 select 避免阻塞
	ch3 := make(chan int)
	select {
	case ch3 <- 3:
		fmt.Println("   ✅ 發送成功")
	default:
		fmt.Println("   🔄 發送失敗，但程序不會阻塞")
	}
}

func demonstrateGracefulShutdown() {
	fmt.Println("\n--- 優雅關閉演示 ---")
	
	jobs := make(chan int, 10)
	done := make(chan bool)
	
	fmt.Println("🛑 優雅關閉模式:")
	
	// Worker
	go func() {
		for {
			select {
			case job := <-jobs:
				fmt.Printf("   ⚙️ 處理任務: %d\n", job)
				time.Sleep(50 * time.Millisecond)
			case <-done:
				fmt.Println("   🛑 收到關閉信號，Worker 退出")
				return
			}
		}
	}()
	
	// 發送一些任務
	fmt.Println("   📤 發送任務:")
	for i := 1; i <= 5; i++ {
		jobs <- i
		fmt.Printf("   📋 任務 %d 已發送\n", i)
	}
	
	// 等待一段時間後優雅關閉
	time.Sleep(200 * time.Millisecond)
	fmt.Println("   🛑 發送關閉信號...")
	done <- true
	time.Sleep(50 * time.Millisecond)
	fmt.Println("✅ 優雅關閉完成")
}