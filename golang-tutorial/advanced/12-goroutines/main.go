package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Go 協程示例 ===")
	
	// 1. 基本協程演示
	demonstrateBasicGoroutine()
	
	// 2. 匿名函數協程
	demonstrateAnonymousGoroutine()
	
	// 3. WaitGroup 同步
	demonstrateWaitGroup()
	
	// 4. 通道同步
	demonstrateChannelSync()
	
	// 5. 工作池模式
	demonstrateWorkerPool()
	
	// 6. 限制協程數量
	demonstrateLimitedGoroutines()
	
	// 7. 生產者-消費者模式
	demonstrateProducerConsumer()
	
	// 8. 數據競爭問題
	demonstrateDataRace()
	
	// 9. 互斥鎖解決方案
	demonstrateMutex()
	
	// 10. 原子操作
	demonstrateAtomic()
	
	// 11. Context 使用
	demonstrateContext()
	
	// 12. 協程監控
	demonstrateGoroutineMonitoring()
}

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("👋 Hello, %s! (%d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func demonstrateBasicGoroutine() {
	fmt.Println("\n--- 基本協程演示 ---")
	
	// 普通函數調用（順序執行）
	fmt.Println("🔄 順序執行:")
	start := time.Now()
	sayHello("Alice")
	sayHello("Bob")
	sequential := time.Since(start)
	
	fmt.Println("\n🚀 併發執行:")
	start = time.Now()
	// 啟動協程（併發執行）
	go sayHello("Charlie")
	go sayHello("Diana")
	
	// 主協程等待一段時間
	time.Sleep(600 * time.Millisecond)
	concurrent := time.Since(start)
	
	fmt.Printf("📊 性能比較 - 順序: %v, 併發: %v\n", sequential, concurrent)
	fmt.Println("✅ 主函數結束")
}

func demonstrateAnonymousGoroutine() {
	fmt.Println("\n--- 匿名函數協程 ---")
	
	var wg sync.WaitGroup
	
	// 匿名函數協程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Printf("🔹 匿名協程 1: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// 帶參數的匿名函數協程
	wg.Add(1)
	go func(msg string, count int) {
		defer wg.Done()
		for i := 0; i < count; i++ {
			fmt.Printf("🔸 %s: %d\n", msg, i)
			time.Sleep(150 * time.Millisecond)
		}
	}("匿名協程 2", 3)
	
	// 閉包協程
	message := "閉包協程"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("🔺 %s 可以訪問外部變數\n", message)
	}()
	
	wg.Wait()
	fmt.Println("✅ 所有匿名協程完成")
}

func demonstrateWaitGroup() {
	fmt.Println("\n--- WaitGroup 同步演示 ---")
	
	var wg sync.WaitGroup
	
	// 啟動多個協程
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待計數
		
		go func(id int) {
			defer wg.Done() // 完成時減少計數
			
			fmt.Printf("👷 Worker %d 開始工作\n", id)
			// 模擬不同的工作時間
			workTime := time.Duration(id*100) * time.Millisecond
			time.Sleep(workTime)
			fmt.Printf("✅ Worker %d 完成工作 (耗時: %v)\n", id, workTime)
		}(i)
	}
	
	fmt.Println("⏳ 等待所有 Worker 完成...")
	start := time.Now()
	wg.Wait() // 等待所有協程完成
	elapsed := time.Since(start)
	fmt.Printf("🎉 所有 Worker 完成！總耗時: %v\n", elapsed)
}

func demonstrateChannelSync() {
	fmt.Println("\n--- 通道同步演示 ---")
	
	done := make(chan bool)
	result := make(chan string)
	
	go func() {
		fmt.Println("🔄 協程開始執行...")
		time.Sleep(300 * time.Millisecond)
		
		// 模擬一些工作
		workResult := "重要數據處理完成"
		result <- workResult
		
		fmt.Println("✅ 協程執行完成")
		done <- true // 發送完成信號
	}()
	
	fmt.Println("⏳ 等待協程完成...")
	
	// 同時等待結果和完成信號
	select {
	case data := <-result:
		fmt.Printf("📊 收到結果: %s\n", data)
		<-done // 等待完成信號
	case <-time.After(1 * time.Second):
		fmt.Println("⏰ 超時")
	}
	
	fmt.Println("🎯 主函數繼續執行")
}

func demonstrateWorkerPool() {
	fmt.Println("\n--- 工作池演示 ---")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// 啟動 worker 協程
	fmt.Printf("🏭 啟動 %d 個 worker\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// 發送任務
	fmt.Printf("📤 發送 %d 個任務\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 收集結果
	fmt.Println("📥 收集結果:")
	totalResult := 0
	for a := 1; a <= numJobs; a++ {
		result := <-results
		totalResult += result
		fmt.Printf("   任務結果: %d\n", result)
	}
	
	fmt.Printf("📊 所有任務完成，總結果: %d\n", totalResult)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("👷 Worker %d 開始處理任務 %d\n", id, j)
		
		// 模擬工作時間
		time.Sleep(time.Duration(100+j*10) * time.Millisecond)
		
		// 簡單的處理：計算平方
		result := j * j
		results <- result
		
		fmt.Printf("✅ Worker %d 完成任務 %d (結果: %d)\n", id, j, result)
	}
}

func demonstrateLimitedGoroutines() {
	fmt.Println("\n--- 限制協程數量演示 ---")
	
	const maxGoroutines = 3
	guard := make(chan struct{}, maxGoroutines)
	
	var wg sync.WaitGroup
	
	// 模擬10個任務，但同時只能有3個協程執行
	fmt.Printf("🚧 最多同時執行 %d 個協程\n", maxGoroutines)
	
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		guard <- struct{}{} // 獲取許可（如果滿了會阻塞）
		
		go func(taskID int) {
			defer wg.Done()
			defer func() { <-guard }() // 釋放許可
			
			fmt.Printf("🔄 Task %d 開始執行 (當前活躍協程: %d)\n", taskID, len(guard)+1)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("✅ Task %d 執行完成\n", taskID)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("🎉 所有任務完成")
}

func demonstrateProducerConsumer() {
	fmt.Println("\n--- 生產者-消費者演示 ---")
	
	ch := make(chan int, 5) // 緩衝通道
	var wg sync.WaitGroup
	
	// 生產者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		
		for i := 1; i <= 10; i++ {
			fmt.Printf("🏭 生產者生產: %d (緩衝區: %d/%d)\n", i, len(ch), cap(ch))
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("🏁 生產者完成")
	}()
	
	// 消費者1
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for {
			select {
			case value, ok := <-ch:
				if !ok {
					fmt.Println("🏁 消費者1完成")
					return
				}
				fmt.Printf("🍽️ 消費者1消費: %d\n", value)
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
	
	// 消費者2
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for {
			select {
			case value, ok := <-ch:
				if !ok {
					fmt.Println("🏁 消費者2完成")
					return
				}
				fmt.Printf("🥤 消費者2消費: %d\n", value)
				time.Sleep(120 * time.Millisecond)
			}
		}
	}()
	
	wg.Wait()
	fmt.Println("🎯 生產者-消費者演示完成")
}

// 全局變數用於演示數據競爭
var unsafeCounter int

func demonstrateDataRace() {
	fmt.Println("\n--- 數據競爭演示 ---")
	
	var wg sync.WaitGroup
	unsafeCounter = 0 // 重置計數器
	
	fmt.Println("⚠️ 警告：以下操作存在數據競爭")
	
	// 啟動多個協程同時修改共享變數
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				unsafeCounter++ // 不安全的操作
			}
			fmt.Printf("📊 協程 %d 完成\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("❌ 最終計數: %d (期望: 10000)\n", unsafeCounter)
	fmt.Println("💡 結果可能不是10000，說明存在數據競爭")
}

// 使用互斥鎖保護的計數器
var (
	safeCounter int
	mutex       sync.Mutex
)

func demonstrateMutex() {
	fmt.Println("\n--- 互斥鎖演示 ---")
	
	var wg sync.WaitGroup
	safeCounter = 0 // 重置計數器
	
	fmt.Println("🔒 使用互斥鎖保護共享變數")
	
	// 啟動多個協程，使用互斥鎖保護共享變數
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				safeCounter++
				mutex.Unlock()
			}
			fmt.Printf("🔐 協程 %d 安全完成\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("✅ 安全計數: %d\n", safeCounter)
}

// 原子操作計數器
var atomicCounter int64

func demonstrateAtomic() {
	fmt.Println("\n--- 原子操作演示 ---")
	
	var wg sync.WaitGroup
	atomic.StoreInt64(&atomicCounter, 0) // 重置計數器
	
	fmt.Println("⚛️ 使用原子操作")
	
	// 啟動多個協程，使用原子操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
			fmt.Printf("⚛️ 協程 %d 原子操作完成\n", id)
		}(i)
	}
	
	wg.Wait()
	result := atomic.LoadInt64(&atomicCounter)
	fmt.Printf("✅ 原子計數: %d\n", result)
}

func demonstrateContext() {
	fmt.Println("\n--- Context 演示 ---")
	
	// 演示取消操作
	fmt.Println("🚫 Context 取消演示:")
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("📡 發送取消信號")
		cancel()
	}()
	
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("⏰ 超時")
	case <-ctx.Done():
		fmt.Printf("✋ 收到取消信號: %v\n", ctx.Err())
	}
	
	// 演示超時
	fmt.Println("\n⏰ Context 超時演示:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel2()
	
	go longRunningTask(ctx2)
	time.Sleep(500 * time.Millisecond)
}

func longRunningTask(ctx context.Context) {
	taskID := time.Now().UnixNano() % 1000
	fmt.Printf("🔄 任務 %d 開始執行\n", taskID)
	
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 任務 %d 被取消: %v\n", taskID, ctx.Err())
			return
		default:
			fmt.Printf("💼 任務 %d 執行中... (%d/10)\n", taskID, i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	fmt.Printf("✅ 任務 %d 完成\n", taskID)
}

func demonstrateGoroutineMonitoring() {
	fmt.Println("\n--- 協程監控演示 ---")
	
	initialCount := runtime.NumGoroutine()
	fmt.Printf("📊 初始協程數: %d\n", initialCount)
	
	var wg sync.WaitGroup
	
	// 啟動一些協程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			currentCount := runtime.NumGoroutine()
			fmt.Printf("🔢 協程 %d 執行中，當前協程數: %d\n", id, currentCount)
			
			// 模擬工作
			time.Sleep(time.Duration(100+id*50) * time.Millisecond)
			
			fmt.Printf("✅ 協程 %d 完成\n", id)
		}(i)
	}
	
	// 在協程執行期間監控
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			count := runtime.NumGoroutine()
			fmt.Printf("📈 監控: 當前協程數 %d\n", count)
		}
	}()
	
	wg.Wait()
	
	// 強制垃圾回收
	runtime.GC()
	time.Sleep(10 * time.Millisecond)
	
	finalCount := runtime.NumGoroutine()
	fmt.Printf("📊 結束時協程數: %d\n", finalCount)
	
	// 顯示內存統計
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("💾 內存使用: %d KB\n", m.Alloc/1024)
}