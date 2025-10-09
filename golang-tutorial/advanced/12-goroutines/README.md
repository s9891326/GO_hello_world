# 第十二章：協程（Goroutines）

## 🎯 學習目標

- 理解協程的概念和優勢
- 掌握協程的創建和管理
- 學會協程間的通信機制
- 了解協程調度器的工作原理
- 掌握併發編程的最佳實踐
- 學會處理併發中的常見問題

## 🚀 協程基礎

協程（Goroutine）是 Go 語言實現併發的核心機制。它是一種輕量級的線程，由 Go 運行時管理，比操作系統線程更高效。

### 協程的特點

```
Goroutine 的關鍵特性：
┌─────────────────────────────────────┐
│ • 輕量級（初始棧大小約 2KB）          │
│ • 高效的調度器（M:P:G 模型）          │
│ • 內建在語言中                        │
│ • 通過通道進行通信                    │
│ • 垃圾回收器感知                      │
│ • 可創建百萬級協程                    │
└─────────────────────────────────────┘
```

### 創建協程

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i+1)
        time.Sleep(100 * time.Millisecond)
    }
}

func demonstrateBasicGoroutine() {
    fmt.Println("--- 基本協程演示 ---")
    
    // 普通函數調用（順序執行）
    fmt.Println("順序執行:")
    sayHello("Alice")
    sayHello("Bob")
    
    fmt.Println("\n併發執行:")
    // 啟動協程（併發執行）
    go sayHello("Charlie")
    go sayHello("Diana")
    
    // 主協程等待一段時間
    time.Sleep(600 * time.Millisecond)
    
    fmt.Println("主函數結束")
}
```

### 匿名函數協程

```go
func demonstrateAnonymousGoroutine() {
    fmt.Println("\n--- 匿名函數協程 ---")
    
    // 匿名函數協程
    go func() {
        for i := 0; i < 3; i++ {
            fmt.Printf("匿名協程 1: %d\n", i)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    // 帶參數的匿名函數協程
    go func(msg string, count int) {
        for i := 0; i < count; i++ {
            fmt.Printf("%s: %d\n", msg, i)
            time.Sleep(150 * time.Millisecond)
        }
    }("匿名協程 2", 3)
    
    // 主協程等待
    time.Sleep(500 * time.Millisecond)
}
```

## ⏱️ 協程同步

### 使用 WaitGroup

```go
import "sync"

func demonstrateWaitGroup() {
    fmt.Println("\n--- WaitGroup 同步演示 ---")
    
    var wg sync.WaitGroup
    
    // 啟動多個協程
    for i := 1; i <= 3; i++ {
        wg.Add(1) // 增加等待計數
        
        go func(id int) {
            defer wg.Done() // 完成時減少計數
            
            fmt.Printf("Worker %d 開始工作\n", id)
            time.Sleep(time.Duration(id*100) * time.Millisecond)
            fmt.Printf("Worker %d 完成工作\n", id)
        }(i)
    }
    
    fmt.Println("等待所有 Worker 完成...")
    wg.Wait() // 等待所有協程完成
    fmt.Println("所有 Worker 完成！")
}
```

### 使用通道同步

```go
func demonstrateChannelSync() {
    fmt.Println("\n--- 通道同步演示 ---")
    
    done := make(chan bool)
    
    go func() {
        fmt.Println("協程開始執行...")
        time.Sleep(300 * time.Millisecond)
        fmt.Println("協程執行完成")
        done <- true // 發送完成信號
    }()
    
    fmt.Println("等待協程完成...")
    <-done // 等待完成信號
    fmt.Println("主函數繼續執行")
}
```

## 🏭 協程池模式

### Worker Pool

```go
func demonstrateWorkerPool() {
    fmt.Println("\n--- 工作池演示 ---")
    
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // 啟動 worker 協程
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }
    
    // 發送任務
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 收集結果
    for a := 1; a <= numJobs; a++ {
        result := <-results
        fmt.Printf("任務結果: %d\n", result)
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d 開始處理任務 %d\n", id, j)
        time.Sleep(100 * time.Millisecond)
        results <- j * 2 // 簡單的處理：乘以2
        fmt.Printf("Worker %d 完成任務 %d\n", id, j)
    }
}
```

### 限制協程數量

```go
func demonstrateLimitedGoroutines() {
    fmt.Println("\n--- 限制協程數量演示 ---")
    
    const maxGoroutines = 3
    guard := make(chan struct{}, maxGoroutines)
    
    var wg sync.WaitGroup
    
    // 模擬10個任務
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        guard <- struct{}{} // 獲取許可
        
        go func(taskID int) {
            defer wg.Done()
            defer func() { <-guard }() // 釋放許可
            
            fmt.Printf("Task %d 開始執行\n", taskID)
            time.Sleep(200 * time.Millisecond)
            fmt.Printf("Task %d 執行完成\n", taskID)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("所有任務完成")
}
```

## 🔄 協程通信模式

### 生產者-消費者模式

```go
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
            fmt.Printf("生產者生產: %d\n", i)
            ch <- i
            time.Sleep(50 * time.Millisecond)
        }
        fmt.Println("生產者完成")
    }()
    
    // 消費者
    wg.Add(1)
    go func() {
        defer wg.Done()
        
        for value := range ch {
            fmt.Printf("消費者消費: %d\n", value)
            time.Sleep(100 * time.Millisecond)
        }
        fmt.Println("消費者完成")
    }()
    
    wg.Wait()
}
```

### 扇出/扇入模式

```go
func demonstrateFanOutFanIn() {
    fmt.Println("\n--- 扇出/扇入演示 ---")
    
    // 輸入通道
    input := make(chan int)
    
    // 扇出：一個輸入到多個處理器
    processor1 := processNumbers(input)
    processor2 := processNumbers(input)
    processor3 := processNumbers(input)
    
    // 扇入：多個處理器到一個輸出
    output := fanIn(processor1, processor2, processor3)
    
    // 發送數據
    go func() {
        defer close(input)
        for i := 1; i <= 10; i++ {
            input <- i
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    // 接收結果
    for result := range output {
        fmt.Printf("處理結果: %d\n", result)
    }
}

func processNumbers(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for num := range input {
            result := num * num // 計算平方
            time.Sleep(100 * time.Millisecond)
            output <- result
        }
    }()
    return output
}

func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for value := range ch {
                output <- value
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}
```

## 🔒 併發安全

### 數據競爭問題

```go
var counter int

func demonstrateDataRace() {
    fmt.Println("\n--- 數據競爭演示 ---")
    
    var wg sync.WaitGroup
    
    // 啟動多個協程同時修改共享變數
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                counter++ // 不安全的操作
            }
        }()
    }
    
    wg.Wait()
    fmt.Printf("最終計數: %d (期望: 10000)\n", counter)
    fmt.Println("注意：結果可能不是10000，說明存在數據競爭")
}
```

### 使用互斥鎖

```go
var (
    safeCounter int
    mutex       sync.Mutex
)

func demonstrateMutex() {
    fmt.Println("\n--- 互斥鎖演示 ---")
    
    var wg sync.WaitGroup
    
    // 重置計數器
    safeCounter = 0
    
    // 啟動多個協程，使用互斥鎖保護共享變數
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                mutex.Lock()
                safeCounter++
                mutex.Unlock()
            }
        }()
    }
    
    wg.Wait()
    fmt.Printf("安全計數: %d\n", safeCounter)
}
```

### 使用原子操作

```go
import "sync/atomic"

var atomicCounter int64

func demonstrateAtomic() {
    fmt.Println("\n--- 原子操作演示 ---")
    
    var wg sync.WaitGroup
    
    // 重置計數器
    atomic.StoreInt64(&atomicCounter, 0)
    
    // 啟動多個協程，使用原子操作
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                atomic.AddInt64(&atomicCounter, 1)
            }
        }()
    }
    
    wg.Wait()
    result := atomic.LoadInt64(&atomicCounter)
    fmt.Printf("原子計數: %d\n", result)
}
```

## 🎯 高級併發模式

### Context 模式

```go
import "context"

func demonstrateContext() {
    fmt.Println("\n--- Context 演示 ---")
    
    // 創建可取消的 context
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        fmt.Println("取消信號發送")
        cancel() // 發送取消信號
    }()
    
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("超時")
    case <-ctx.Done():
        fmt.Printf("收到取消信號: %v\n", ctx.Err())
    }
}

func demonstrateContextWithTimeout() {
    fmt.Println("\n--- Context 超時演示 ---")
    
    // 創建帶超時的 context
    ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
    defer cancel()
    
    go longRunningTask(ctx)
    
    time.Sleep(500 * time.Millisecond)
}

func longRunningTask(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("任務被取消: %v\n", ctx.Err())
            return
        default:
            fmt.Println("任務執行中...")
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

### Pipeline 模式

```go
func demonstratePipeline() {
    fmt.Println("\n--- Pipeline 演示 ---")
    
    // 創建管道
    numbers := generateNumbers(1, 10)
    squares := square(numbers)
    output := filter(squares, func(n int) bool {
        return n%2 == 0 // 過濾偶數
    })
    
    // 處理結果
    for result := range output {
        fmt.Printf("管道結果: %d\n", result)
    }
}

func generateNumbers(start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := start; i <= end; i++ {
            ch <- i
            time.Sleep(50 * time.Millisecond)
        }
    }()
    return ch
}

func square(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for num := range input {
            output <- num * num
        }
    }()
    return output
}

func filter(input <-chan int, predicate func(int) bool) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for num := range input {
            if predicate(num) {
                output <- num
            }
        }
    }()
    return output
}
```

## 🐛 常見陷阱和解決方案

### 協程洩漏

```go
func demonstrateGoroutineLeak() {
    fmt.Println("\n--- 協程洩漏演示 ---")
    
    // 錯誤的方式：可能導致協程洩漏
    ch := make(chan string)
    
    go func() {
        // 這個協程會一直阻塞
        result := <-ch
        fmt.Println("收到:", result)
    }()
    
    // 如果不發送數據，上面的協程會一直等待
    fmt.Println("可能的協程洩漏（未發送數據）")
    
    // 正確的方式：使用帶緩衝的通道或確保發送
    time.Sleep(100 * time.Millisecond)
    ch <- "數據" // 避免洩漏
}

func demonstrateProperCleanup() {
    fmt.Println("\n--- 正確清理演示 ---")
    
    ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
    defer cancel()
    
    ch := make(chan string, 1) // 使用緩衝通道
    
    go func() {
        select {
        case result := <-ch:
            fmt.Println("收到:", result)
        case <-ctx.Done():
            fmt.Println("協程被取消，避免洩漏")
            return
        }
    }()
    
    // 模擬可能不發送數據的情況
    time.Sleep(300 * time.Millisecond)
}
```

### 死鎖問題

```go
func demonstrateDeadlock() {
    fmt.Println("\n--- 死鎖預防演示 ---")
    
    // 使用 select 避免死鎖
    ch := make(chan int)
    
    select {
    case ch <- 42:
        fmt.Println("數據已發送")
    case <-time.After(100 * time.Millisecond):
        fmt.Println("發送超時，避免死鎖")
    }
    
    select {
    case data := <-ch:
        fmt.Println("收到數據:", data)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("接收超時，避免死鎖")
    }
}
```

## 📊 性能監控

### 協程數量監控

```go
func demonstrateGoroutineMonitoring() {
    fmt.Println("\n--- 協程監控演示 ---")
    
    initialCount := runtime.NumGoroutine()
    fmt.Printf("初始協程數: %d\n", initialCount)
    
    var wg sync.WaitGroup
    
    // 啟動一些協程
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            currentCount := runtime.NumGoroutine()
            fmt.Printf("協程 %d 執行中，當前協程數: %d\n", id, currentCount)
            time.Sleep(100 * time.Millisecond)
        }(i)
    }
    
    wg.Wait()
    
    finalCount := runtime.NumGoroutine()
    fmt.Printf("結束時協程數: %d\n", finalCount)
}
```

## 💡 最佳實踐

### 1. 協程設計原則

```go
// 好的實踐：明確的開始和結束
func goodGoroutinePattern() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go func() {
        defer wg.Done()
        // 明確的工作邏輯
        fmt.Println("執行具體任務")
    }()
    
    wg.Wait() // 明確等待完成
}

// 避免：無限制的協程創建
func avoidUnlimitedGoroutines() {
    // 使用工作池限制併發數
    semaphore := make(chan struct{}, 10) // 最多10個並發
    
    for i := 0; i < 100; i++ {
        semaphore <- struct{}{}
        go func(id int) {
            defer func() { <-semaphore }()
            // 工作邏輯
            time.Sleep(100 * time.Millisecond)
        }(i)
    }
}
```

### 2. 錯誤處理

```go
func demonstrateErrorHandling() {
    fmt.Println("\n--- 協程錯誤處理 ---")
    
    errCh := make(chan error, 1)
    resultCh := make(chan string, 1)
    
    go func() {
        // 模擬可能失敗的操作
        time.Sleep(100 * time.Millisecond)
        
        if time.Now().UnixNano()%2 == 0 {
            errCh <- fmt.Errorf("模擬錯誤")
        } else {
            resultCh <- "成功結果"
        }
    }()
    
    select {
    case err := <-errCh:
        fmt.Printf("協程錯誤: %v\n", err)
    case result := <-resultCh:
        fmt.Printf("協程結果: %s\n", result)
    case <-time.After(200 * time.Millisecond):
        fmt.Println("協程超時")
    }
}
```

## 🎯 本章練習

1. 實現並發爬蟲
2. 創建任務調度器
3. 實現生產者-消費者隊列
4. 創建協程池管理器

---

**下一章：[通道](../13-channels/)**