package hello_world

import (
	"fmt"
	"time"
)

func Day14() {
	/*
		併發(Concurrency) 不是併行(Parallelism)

		併發是共享時間運算，在一段時間內輪流享有時間資源
		併行是平行運算，一直都能享有時間資源

		併發是把時間切成很小很小段，在這小段的時間裡先後執行多項任務。
		併行是CPU有多個核心，可以同時處理多個任務。

		簡單來說，拿人來比喻的話：
		併發：一個人在一段時間內做兩件事
		併行：兩個人同時在做事
	*/
	printTwo()
	showAnonymousFunc()

	fmt.Println("\nstart")
	go showPanic()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

func showPanic() {
	fmt.Println("panic!!!")
	panic("空難@@@@")
}

func showAnonymousFunc() {
	fmt.Println()
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Print("Anonymous")
		}
	}()

	for i := 0; i < 1000; i++ {
		fmt.Print("main")
	}
	time.Sleep(time.Second)
}

func printTwo() {
	/*
		runtime.GOMAXPROCS(n) 這一參數限制程式執行時 CPU用到的最大核心數量。
		如果設置小於1，等於沒設，預設值是電腦核心數。
		但限制了一核心之後，為什麼還是可以把兩個print func都印到呢，
		怎麼不是只印出一個直到時間到？說好的單核心？

		原來是Go Routine會去排程，執行A線程一小段時間後會跳到線程B去，
		這才公平合理嘛！不然CPU資源都被其中一個線程給佔住，作業系統就卡死啦。
		所以這次看到的輸出會是 Ｏ很多 再來－很多，兩者都連續印很多的情況下交錯著。
	*/
	//runtime.GOMAXPROCS(2)
	go print1()
	go print2()

	time.Sleep(1 * time.Second)
}

func print1() {
	for i := 0; i < 1000; i++ {
		fmt.Print("Ｏ")
	}
}

func print2() {
	for i := 0; i < 1000; i++ {
		fmt.Print("-")
	}
}
