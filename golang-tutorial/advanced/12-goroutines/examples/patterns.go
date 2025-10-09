package main

import (
	"fmt"
	"sync"
	"time"
)

// 扇出/扇入模式演示
func demonstrateFanOutFanIn() {
	fmt.Println("\n--- 扇出/扇入演示 ---")
	
	// 輸入通道
	input := make(chan int)
	
	// 扇出：一個輸入到多個處理器
	processor1 := processNumbers(input, "處理器1")
	processor2 := processNumbers(input, "處理器2")
	processor3 := processNumbers(input, "處理器3")
	
	// 扇入：多個處理器到一個輸出
	output := fanIn(processor1, processor2, processor3)
	
	// 發送數據
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			fmt.Printf("📤 發送數據: %d\n", i)
			input <- i
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("📤 數據發送完成")
	}()
	
	// 接收結果
	fmt.Println("📥 接收處理結果:")
	resultCount := 0
	for result := range output {
		resultCount++
		fmt.Printf("   結果 %d: %d\n", resultCount, result)
	}
	
	fmt.Printf("🎯 總共處理了 %d 個結果\n", resultCount)
}

func processNumbers(input <-chan int, name string) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			// 計算平方
			result := num * num
			fmt.Printf("⚙️ %s 處理 %d -> %d\n", name, num, result)
			time.Sleep(100 * time.Millisecond)
			output <- result
		}
		fmt.Printf("✅ %s 處理完成\n", name)
	}()
	return output
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup
	
	for i, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int, id int) {
			defer wg.Done()
			for value := range ch {
				fmt.Printf("🔄 扇入 %d: %d\n", id+1, value)
				output <- value
			}
		}(input, i)
	}
	
	go func() {
		wg.Wait()
		close(output)
		fmt.Println("🏁 扇入完成")
	}()
	
	return output
}

// Pipeline 模式演示
func demonstratePipeline() {
	fmt.Println("\n--- Pipeline 演示 ---")
	
	// 創建管道：數字生成 -> 平方 -> 過濾偶數 -> 求和
	numbers := generateNumbers(1, 10)
	squares := square(numbers)
	evens := filterEven(squares)
	sum := sumNumbers(evens)
	
	// 獲取最終結果
	result := <-sum
	fmt.Printf("🎯 Pipeline 最終結果: %d\n", result)
}

func generateNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("🔢 開始生成數字")
		for i := start; i <= end; i++ {
			fmt.Printf("   生成: %d\n", i)
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("🔢 數字生成完成")
	}()
	return ch
}

func square(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		fmt.Println("📐 開始計算平方")
		for num := range input {
			result := num * num
			fmt.Printf("   %d² = %d\n", num, result)
			output <- result
		}
		fmt.Println("📐 平方計算完成")
	}()
	return output
}

func filterEven(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		fmt.Println("🔍 開始過濾偶數")
		for num := range input {
			if num%2 == 0 {
				fmt.Printf("   保留偶數: %d\n", num)
				output <- num
			} else {
				fmt.Printf("   跳過奇數: %d\n", num)
			}
		}
		fmt.Println("🔍 過濾完成")
	}()
	return output
}

func sumNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		fmt.Println("➕ 開始求和")
		sum := 0
		for num := range input {
			sum += num
			fmt.Printf("   累加: %d (總和: %d)\n", num, sum)
		}
		fmt.Printf("➕ 求和完成，結果: %d\n", sum)
		output <- sum
	}()
	return output
}

// 工作竊取模式
func demonstrateWorkStealing() {
	fmt.Println("\n--- 工作竊取演示 ---")
	
	const numWorkers = 3
	const numTasks = 15
	
	// 為每個 worker 創建獨立的任務隊列
	queues := make([]chan int, numWorkers)
	for i := range queues {
		queues[i] = make(chan int, 10)
	}
	
	var wg sync.WaitGroup
	
	// 啟動 workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go workStealingWorker(i, queues, &wg)
	}
	
	// 分發任務到不同隊列
	for task := 1; task <= numTasks; task++ {
		queueIndex := (task - 1) % numWorkers
		fmt.Printf("📋 任務 %d 分配給隊列 %d\n", task, queueIndex)
		queues[queueIndex] <- task
	}
	
	// 關閉所有隊列
	for i := range queues {
		close(queues[i])
	}
	
	wg.Wait()
	fmt.Println("🎯 工作竊取演示完成")
}

func workStealingWorker(id int, queues []chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	myQueue := queues[id]
	processed := 0
	
	for {
		select {
		case task, ok := <-myQueue:
			if !ok {
				// 自己的隊列已關閉，嘗試從其他隊列竊取工作
				if stolenTask := stealWork(id, queues); stolenTask != -1 {
					processTask(id, stolenTask, "竊取")
					processed++
				} else {
					// 沒有更多工作
					fmt.Printf("🏁 Worker %d 完成，處理了 %d 個任務\n", id, processed)
					return
				}
			} else {
				processTask(id, task, "本地")
				processed++
			}
		default:
			// 自己隊列為空，嘗試竊取
			if stolenTask := stealWork(id, queues); stolenTask != -1 {
				processTask(id, stolenTask, "竊取")
				processed++
			} else {
				time.Sleep(10 * time.Millisecond) // 短暫休息
			}
		}
	}
}

func stealWork(workerID int, queues []chan int) int {
	for i, queue := range queues {
		if i == workerID {
			continue // 跳過自己的隊列
		}
		
		select {
		case task, ok := <-queue:
			if ok {
				fmt.Printf("🔄 Worker %d 從隊列 %d 竊取任務 %d\n", workerID, i, task)
				return task
			}
		default:
			// 隊列為空，繼續嘗試下一個
		}
	}
	return -1 // 沒有找到可竊取的任務
}

func processTask(workerID, task int, source string) {
	fmt.Printf("⚙️ Worker %d 處理任務 %d (%s)\n", workerID, task, source)
	// 模擬不同的工作負載
	workTime := time.Duration(50+task*10) * time.Millisecond
	time.Sleep(workTime)
	fmt.Printf("✅ Worker %d 完成任務 %d\n", workerID, task)
}

// 信號量模式
func demonstrateSemaphore() {
	fmt.Println("\n--- 信號量演示 ---")
	
	// 創建信號量，限制同時執行的任務數
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)
	
	var wg sync.WaitGroup
	
	// 啟動多個任務
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			
			// 獲取信號量
			sem <- struct{}{}
			fmt.Printf("🔐 任務 %d 獲得執行許可 (活躍: %d/%d)\n", taskID, len(sem), maxConcurrent)
			
			// 執行任務
			time.Sleep(time.Duration(200+taskID*50) * time.Millisecond)
			fmt.Printf("✅ 任務 %d 完成\n", taskID)
			
			// 釋放信號量
			<-sem
		}(i)
	}
	
	wg.Wait()
	fmt.Println("🎯 信號量演示完成")
}

func main() {
	demonstrateFanOutFanIn()
	demonstratePipeline()
	demonstrateWorkStealing()
	demonstrateSemaphore()
}