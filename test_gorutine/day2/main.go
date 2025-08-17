package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

/*
Day 2：併發控制與同步

WaitGroup: 用來等待一組goroutine全部執行完成的同步工具
Add(n) → 告訴 WaitGroup 有 n 個任務要做。
Done() → 告訴 WaitGroup 有一個任務完成了（計數器 -1）。
Wait() → 卡在這裡，直到計數器變成 0，才繼續往下執行。

sync.Mutex（互斥鎖）
單一鎖，鎖住後其他 goroutine 不能進入臨界區
適合保護「短時間」的共享資源寫入
缺點：讀的時候也會阻塞其他讀

sync.RWMutex（讀寫鎖）
允許多個 goroutine 同時讀（RLock），但寫的時候（Lock）必須獨佔
適合「讀多寫少」的場景，例如快取查詢

sync.Once
保證某段程式碼只會執行一次（常用於單例模式、初始化）
適合初始化資源（例如建立資料庫連線池）

sync.Cond（條件變數）
用來在 goroutine 間發送「事件通知」
適合用於「等待某條件成立再繼續」的情境，例如生產者消費者模型

sync.Map（併發安全 Map）
內建併發安全的 map，省去加鎖麻煩
適合 key 不固定、讀寫都很頻繁的情境

atomic（原子操作）
在 sync/atomic 裡提供的原子加減、比較交換（CAS）
不需要鎖，效能高，但只能保護單一變數的操作

工具	適用情境
Mutex	共享資源保護（短時間操作）
RWMutex	讀多寫少
Once	初始化只跑一次
Cond	條件等待／事件通知
sync.Map	併發安全 Map
atomic	高效保護單一數值

*/

func worker(id int, wg *sync.WaitGroup, mu *sync.Mutex, counter *int) {
	defer wg.Done()
	mu.Lock()
	*counter++
	fmt.Println("worker", id, "counter", *counter)
	mu.Unlock()
	time.Sleep(1 * time.Millisecond)
}

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(i, &wg, &mu, &counter)
	}
	wg.Wait()
}
