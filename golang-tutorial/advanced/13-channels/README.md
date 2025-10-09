# 第十三章：通道（Channels）

## 🎯 學習目標

- 理解通道的概念和作用
- 掌握通道的創建和基本操作
- 學會有緩衝和無緩衝通道的使用
- 了解通道的方向性和關閉機制
- 掌握 select 語句的使用
- 學會通道的高級應用模式

## 📡 通道基礎

通道（Channel）是 Go 語言中協程間通信的主要機制。它體現了 Go 的設計哲學："不要通過共享內存來通信，而要通過通信來共享內存"。

### 通道的特點

```
Channel 的關鍵特性：
┌─────────────────────────────────────┐
│ • 類型安全的通信機制                  │
│ • 可以是有緩衝或無緩衝的              │
│ • 支援方向性（只讀、只寫、讀寫）      │
│ • 可以被關閉                        │
│ • 支援非阻塞操作（select）           │
│ • 遵循先進先出（FIFO）原則           │
└─────────────────────────────────────┘
```

### 通道的創建

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateChannelBasics() {
    fmt.Println("--- 通道基礎演示 ---")
    
    // 創建無緩衝通道
    ch1 := make(chan int)
    ch2 := make(chan string)
    
    // 創建有緩衝通道
    ch3 := make(chan int, 5)     // 容量為5的整數通道
    ch4 := make(chan string, 3)  // 容量為3的字符串通道
    
    fmt.Printf("無緩衝整數通道: %T\n", ch1)
    fmt.Printf("無緩衝字符串通道: %T\n", ch2)
    fmt.Printf("有緩衝整數通道容量: %d\n", cap(ch3))
    fmt.Printf("有緩衝字符串通道容量: %d\n", cap(ch4))
    
    // 關閉通道（避免未使用變數錯誤）
    close(ch1)
    close(ch2)
    close(ch3)
    close(ch4)
}
```

### 基本操作

```go
func demonstrateChannelOperations() {
    fmt.Println("\n--- 通道基本操作 ---")
    
    // 無緩衝通道 - 同步通信
    ch := make(chan string)
    
    // 啟動接收協程
    go func() {
        message := <-ch // 接收數據
        fmt.Printf("接收到消息: %s\n", message)
    }()
    
    // 發送數據
    ch <- "Hello, Channel!" // 發送數據
    
    time.Sleep(100 * time.Millisecond)
    
    // 有緩衝通道 - 異步通信
    bufferedCh := make(chan int, 3)
    
    // 發送數據（不會阻塞，因為有緩衝）
    bufferedCh <- 1
    bufferedCh <- 2
    bufferedCh <- 3
    
    fmt.Printf("緩衝通道長度: %d/%d\n", len(bufferedCh), cap(bufferedCh))
    
    // 接收數據
    for i := 0; i < 3; i++ {
        value := <-bufferedCh
        fmt.Printf("從緩衝通道接收: %d\n", value)
    }
}
```

## 🔄 無緩衝 vs 有緩衝通道

### 無緩衝通道（同步通道）

```go
func demonstrateUnbufferedChannel() {
    fmt.Println("\n--- 無緩衝通道演示 ---")
    
    ch := make(chan string)
    
    // 無緩衝通道的發送和接收是同步的
    go func() {
        fmt.Println("協程準備發送數據...")
        ch <- "同步消息"
        fmt.Println("協程發送完成")
    }()
    
    time.Sleep(100 * time.Millisecond) // 模擬主協程在做其他事
    
    fmt.Println("主協程準備接收數據...")
    message := <-ch
    fmt.Printf("主協程接收到: %s\n", message)
}

func demonstrateHandshake() {
    fmt.Println("\n--- 握手通信演示 ---")
    
    done := make(chan bool)
    
    go func() {
        fmt.Println("Worker: 開始工作...")
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Worker: 工作完成")
        done <- true // 發送完成信號
    }()
    
    fmt.Println("Main: 等待 Worker 完成...")
    <-done // 等待完成信號
    fmt.Println("Main: 收到完成信號，繼續執行")
}
```

### 有緩衝通道（異步通道）

```go
func demonstrateBufferedChannel() {
    fmt.Println("\n--- 有緩衝通道演示 ---")
    
    // 創建容量為 3 的緩衝通道
    ch := make(chan int, 3)
    
    // 發送數據（不會阻塞）
    fmt.Println("發送數據到緩衝通道...")
    ch <- 1
    fmt.Printf("發送 1，通道長度: %d/%d\n", len(ch), cap(ch))
    
    ch <- 2
    fmt.Printf("發送 2，通道長度: %d/%d\n", len(ch), cap(ch))
    
    ch <- 3
    fmt.Printf("發送 3，通道長度: %d/%d\n", len(ch), cap(ch))
    
    // 現在通道已滿，再發送會阻塞
    
    // 開始接收數據
    fmt.Println("\n從緩衝通道接收數據...")
    for i := 0; i < 3; i++ {
        value := <-ch
        fmt.Printf("接收 %d，通道長度: %d/%d\n", value, len(ch), cap(ch))
    }
}

func demonstrateProducerConsumerWithBuffer() {
    fmt.Println("\n--- 緩衝通道生產者消費者 ---")
    
    ch := make(chan int, 5) // 緩衝大小為5
    
    // 生產者
    go func() {
        for i := 1; i <= 10; i++ {
            ch <- i
            fmt.Printf("生產者生產: %d (緩衝: %d/%d)\n", i, len(ch), cap(ch))
            time.Sleep(50 * time.Millisecond)
        }
        close(ch)
        fmt.Println("生產者完成")
    }()
    
    // 消費者
    time.Sleep(200 * time.Millisecond) // 讓生產者先生產一些
    
    for value := range ch {
        fmt.Printf("消費者消費: %d (緩衝: %d/%d)\n", value, len(ch), cap(ch))
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println("消費者完成")
}
```

## 🔒 通道的關閉

### 關閉通道的規則

```go
func demonstrateChannelClosure() {
    fmt.Println("\n--- 通道關閉演示 ---")
    
    ch := make(chan int, 3)
    
    // 發送一些數據
    ch <- 1
    ch <- 2
    ch <- 3
    
    // 關閉通道
    close(ch)
    
    // 關閉後仍可以接收數據
    fmt.Println("關閉通道後接收數據:")
    for {
        value, ok := <-ch
        if !ok {
            fmt.Println("通道已關閉且無數據")
            break
        }
        fmt.Printf("接收到: %d\n", value)
    }
    
    // 使用 range 遍歷已關閉的通道
    ch2 := make(chan string, 2)
    ch2 <- "Hello"
    ch2 <- "World"
    close(ch2)
    
    fmt.Println("\n使用 range 遍歷:")
    for msg := range ch2 {
        fmt.Printf("接收到: %s\n", msg)
    }
}

func demonstrateClosePatterns() {
    fmt.Println("\n--- 通道關閉模式 ---")
    
    // 模式1: 發送者關閉通道
    numbers := make(chan int)
    
    go func() {
        defer close(numbers) // 確保關閉通道
        for i := 1; i <= 5; i++ {
            numbers <- i
            time.Sleep(50 * time.Millisecond)
        }
        fmt.Println("發送者完成並關閉通道")
    }()
    
    // 接收者檢查通道狀態
    for {
        select {
        case num, ok := <-numbers:
            if !ok {
                fmt.Println("通道已關閉")
                return
            }
            fmt.Printf("接收到數字: %d\n", num)
        case <-time.After(200 * time.Millisecond):
            fmt.Println("接收超時")
            return
        }
    }
}
```

## 🎛️ 通道方向性

### 只讀和只寫通道

```go
// 只能發送的通道
func sender(ch chan<- string) {
    ch <- "Hello from sender"
    // value := <-ch // 編譯錯誤：不能從只寫通道接收
}

// 只能接收的通道
func receiver(ch <-chan string) {
    message := <-ch
    fmt.Printf("Receiver got: %s\n", message)
    // ch <- "response" // 編譯錯誤：不能向只讀通道發送
}

func demonstrateChannelDirections() {
    fmt.Println("\n--- 通道方向性演示 ---")
    
    ch := make(chan string, 1)
    
    // 將雙向通道轉為單向通道
    go sender(ch)   // 傳遞為只寫通道
    go receiver(ch) // 傳遞為只讀通道
    
    time.Sleep(100 * time.Millisecond)
}

// 實際應用：創建管道
func createPipeline() (<-chan int, chan<- bool) {
    numbers := make(chan int)
    done := make(chan bool)
    
    go func() {
        defer close(numbers)
        for i := 1; i <= 10; i++ {
            select {
            case numbers <- i:
                fmt.Printf("生成數字: %d\n", i)
            case <-done:
                fmt.Println("收到停止信號")
                return
            }
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    return numbers, done
}

func demonstratePipeline() {
    fmt.Println("\n--- 管道方向性演示 ---")
    
    numbers, done := createPipeline()
    
    go func() {
        time.Sleep(300 * time.Millisecond)
        done <- true // 發送停止信號
    }()
    
    for num := range numbers {
        fmt.Printf("處理數字: %d\n", num)
    }
}
```

## 🔀 Select 語句

### 基本 Select 使用

```go
func demonstrateBasicSelect() {
    fmt.Println("\n--- 基本 Select 演示 ---")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
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
            fmt.Printf("收到通道1: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("收到通道2: %s\n", msg2)
        case <-time.After(200 * time.Millisecond):
            fmt.Println("接收超時")
        }
    }
}

func demonstrateSelectWithDefault() {
    fmt.Println("\n--- Select 默認分支演示 ---")
    
    ch := make(chan int, 1)
    
    // 非阻塞發送
    select {
    case ch <- 42:
        fmt.Println("成功發送到通道")
    default:
        fmt.Println("通道已滿，無法發送")
    }
    
    // 非阻塞接收
    select {
    case value := <-ch:
        fmt.Printf("接收到值: %d\n", value)
    default:
        fmt.Println("通道為空，無法接收")
    }
    
    // 再次嘗試接收
    select {
    case value := <-ch:
        fmt.Printf("接收到值: %d\n", value)
    default:
        fmt.Println("通道為空，使用默認分支")
    }
}
```

### Select 的高級使用

```go
func demonstrateAdvancedSelect() {
    fmt.Println("\n--- 高級 Select 演示 ---")
    
    // 超時控制
    result := make(chan string, 1)
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        result <- "操作完成"
    }()
    
    select {
    case res := <-result:
        fmt.Printf("操作結果: %s\n", res)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("操作超時")
    }
    
    // 心跳檢測
    demonstrateHeartbeat()
}

func demonstrateHeartbeat() {
    fmt.Println("\n心跳檢測演示:")
    
    heartbeat := time.Tick(100 * time.Millisecond)
    work := make(chan string)
    
    go func() {
        time.Sleep(250 * time.Millisecond)
        work <- "工作完成"
    }()
    
    for {
        select {
        case <-heartbeat:
            fmt.Println("💓 心跳")
        case result := <-work:
            fmt.Printf("📋 %s\n", result)
            return
        case <-time.After(500 * time.Millisecond):
            fmt.Println("⏰ 整體超時")
            return
        }
    }
}
```

## 🔄 通道模式

### 扇入模式（Fan-in）

```go
func fanIn(input1, input2 <-chan string) <-chan string {
    output := make(chan string)
    
    go func() {
        for {
            select {
            case s := <-input1:
                output <- s
            case s := <-input2:
                output <- s
            }
        }
    }()
    
    return output
}

func demonstrateFanIn() {
    fmt.Println("\n--- 扇入模式演示 ---")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // 啟動兩個生產者
    go func() {
        for i := 0; i < 3; i++ {
            ch1 <- fmt.Sprintf("生產者1-消息%d", i+1)
            time.Sleep(100 * time.Millisecond)
        }
        close(ch1)
    }()
    
    go func() {
        for i := 0; i < 3; i++ {
            ch2 <- fmt.Sprintf("生產者2-消息%d", i+1)
            time.Sleep(150 * time.Millisecond)
        }
        close(ch2)
    }()
    
    // 合並兩個通道
    merged := fanIn(ch1, ch2)
    
    // 接收合併的消息
    timeout := time.After(1 * time.Second)
    for {
        select {
        case msg := <-merged:
            fmt.Printf("扇入接收: %s\n", msg)
        case <-timeout:
            fmt.Println("扇入演示完成")
            return
        }
    }
}
```

### 扇出模式（Fan-out）

```go
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        
        go func(out chan<- int) {
            defer close(out)
            for data := range input {
                out <- data * data // 計算平方
                time.Sleep(50 * time.Millisecond)
            }
        }(output)
    }
    
    return outputs
}

func demonstrateFanOut() {
    fmt.Println("\n--- 扇出模式演示 ---")
    
    input := make(chan int)
    
    // 生產數據
    go func() {
        defer close(input)
        for i := 1; i <= 6; i++ {
            input <- i
            fmt.Printf("輸入數據: %d\n", i)
        }
    }()
    
    // 扇出到3個worker
    outputs := fanOut(input, 3)
    
    // 收集所有結果
    var results []int
    for _, output := range outputs {
        for result := range output {
            results = append(results, result)
            fmt.Printf("工作結果: %d\n", result)
        }
    }
    
    fmt.Printf("總共收到 %d 個結果\n", len(results))
}
```

### 管道模式（Pipeline）

```go
func pipeline() {
    fmt.Println("\n--- 管道模式演示 ---")
    
    // 階段1：生成數字
    numbers := make(chan int)
    go func() {
        defer close(numbers)
        for i := 1; i <= 5; i++ {
            numbers <- i
            fmt.Printf("生成: %d\n", i)
        }
    }()
    
    // 階段2：計算平方
    squares := make(chan int)
    go func() {
        defer close(squares)
        for num := range numbers {
            square := num * num
            squares <- square
            fmt.Printf("平方: %d -> %d\n", num, square)
        }
    }()
    
    // 階段3：過濾偶數
    evens := make(chan int)
    go func() {
        defer close(evens)
        for square := range squares {
            if square%2 == 0 {
                evens <- square
                fmt.Printf("偶數: %d\n", square)
            }
        }
    }()
    
    // 最終消費
    for even := range evens {
        fmt.Printf("最終結果: %d\n", even)
    }
}
```

## 🚫 通道的常見陷阱

### 死鎖

```go
func demonstrateDeadlock() {
    fmt.Println("\n--- 死鎖預防演示 ---")
    
    // 錯誤示例（註釋避免死鎖）
    /*
    ch := make(chan int)
    ch <- 1  // 死鎖：沒有接收者
    */
    
    // 正確示例1：使用協程
    ch1 := make(chan int)
    go func() {
        ch1 <- 1
    }()
    value := <-ch1
    fmt.Printf("協程方式接收: %d\n", value)
    
    // 正確示例2：使用緩衝通道
    ch2 := make(chan int, 1)
    ch2 <- 2 // 不會阻塞
    value2 := <-ch2
    fmt.Printf("緩衝通道接收: %d\n", value2)
    
    // 正確示例3：使用 select 避免阻塞
    ch3 := make(chan int)
    select {
    case ch3 <- 3:
        fmt.Println("發送成功")
    default:
        fmt.Println("發送失敗，但程序不會阻塞")
    }
}
```

### 通道洩漏

```go
func demonstrateChannelLeak() {
    fmt.Println("\n--- 通道洩漏預防 ---")
    
    // 錯誤示例：協程可能永遠阻塞
    // 正確做法：使用超時或取消機制
    
    timeout := time.After(200 * time.Millisecond)
    result := make(chan string, 1)
    
    go func() {
        // 模擬長時間運行的任務
        time.Sleep(300 * time.Millisecond)
        result <- "任務完成"
    }()
    
    select {
    case res := <-result:
        fmt.Printf("任務結果: %s\n", res)
    case <-timeout:
        fmt.Println("任務超時，避免協程洩漏")
    }
}
```

## 💡 通道最佳實踐

### 1. 通道所有權

```go
// 好的實踐：誰創建誰關閉
func goodChannelOwnership() <-chan int {
    ch := make(chan int)
    
    go func() {
        defer close(ch) // 創建者負責關閉
        for i := 1; i <= 3; i++ {
            ch <- i
        }
    }()
    
    return ch
}

func demonstrateOwnership() {
    fmt.Println("\n--- 通道所有權演示 ---")
    
    numbers := goodChannelOwnership()
    
    for num := range numbers {
        fmt.Printf("接收數字: %d\n", num)
    }
    fmt.Println("通道已正確關閉")
}
```

### 2. 優雅關閉

```go
func demonstrateGracefulShutdown() {
    fmt.Println("\n--- 優雅關閉演示 ---")
    
    jobs := make(chan int, 10)
    done := make(chan bool)
    
    // Worker
    go func() {
        for {
            select {
            case job := <-jobs:
                fmt.Printf("處理任務: %d\n", job)
                time.Sleep(50 * time.Millisecond)
            case <-done:
                fmt.Println("收到關閉信號，Worker 退出")
                return
            }
        }
    }()
    
    // 發送一些任務
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    
    // 等待一段時間後優雅關閉
    time.Sleep(200 * time.Millisecond)
    done <- true
    time.Sleep(50 * time.Millisecond)
}
```

## 🎯 本章練習

1. 實現生產者-消費者隊列
2. 創建工作調度器
3. 實現請求/響應模式
4. 創建事件總線系統

---

**下一章：[錯誤處理](../14-error-handling/)**