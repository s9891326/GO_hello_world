package hello_world

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

func Day16() {
	/*
		【sync.WaitGroup】
		可以透過內建的sync WaitGroup來等待線程結束，
		就像一群新兵在準備集合，等到每個人都到為止。
	*/
	//showSyncWaitGroup()

	/*
		在不使用 sync.Mutex的時候會發生race condition，多個CPU在搶count這個變數
		所以這個時候就需要上互斥鎖(Lock)
	*/
	showLock()
}

var count = 0

func showLock() {
	for i := 0; i < 10000; i++ {
		go race()
	}
	time.Sleep(time.Second)
	fmt.Println("count:", count)
}

func race() {
	lock.Lock()
	count++
	lock.Unlock()
}

func showSyncWaitGroup() {
	fmt.Println("你各位啊，現在開始休息，三秒鐘後記得回來。")
	wg := sync.WaitGroup{}
	wg.Add(2)

	go rest(&wg)
	go rest(&wg)

	fmt.Println("===你各位再慢慢來沒關係啊===")
	wg.Wait()
	fmt.Println("===集合完畢===")
}

func rest(s *sync.WaitGroup) {
	r := rand.Intn(10)
	fmt.Printf("r:%d, T:%T, time duration: %d, T:%T\n", r, r, time.Duration(r), time.Duration(r))
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println("新兵休息完畢。")
	s.Done()
}
