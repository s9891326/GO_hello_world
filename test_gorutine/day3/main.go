package main

import (
	"context"
	"fmt"
	"time"
)

/*
Day 3：context 與限流

context:
- 控制任務生命週期
可以讓 goroutine 在超時、取消、或父任務結束時自動停止
避免 goroutine 洩漏（長時間沒退出佔用資源）
- 在 goroutine 間傳遞資訊
比如傳遞請求 ID、用戶資訊等（用 context.WithValue）

限流:
- 固定頻率（Token Bucket / Leaky Bucket）
控制每秒最多處理多少請求
- 併發數限制
同時最多允許 N 個任務在跑


限流控制「同時多少請求進來」
context 控制「請求最多活多久」

兩個搭配：
先用限流保護系統不超載
再用 context 保證請求超時後自動結束，避免卡死

*/

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker %d running\n", id)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	//defer cancel()
	//
	//for i := 0; i < 3; i++ {
	//	go worker(ctx, i)
	//}
	//time.Sleep(time.Second * 3)

	limiter := time.Tick(200 * time.Millisecond) // 每 200ms 發一個 token
	for i := 0; i < 10; i++ {
		<-limiter
		fmt.Println("發送任務", i)
	}

	// 限制同時最多 3 個併發
	sem := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		sem <- struct{}{} // 佔用一個位置
		go func(id int) {
			defer func() {
				<-sem
			}() // 釋放位置

			fmt.Println("處理任務", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

}
