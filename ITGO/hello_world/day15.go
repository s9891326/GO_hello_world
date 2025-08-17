package hello_world

import (
	"fmt"
	"log"
	"time"
)

func Day15() {
	/*
		Buffer 是拿來緩衝用的，Unbuffered Channel則是0緩衝，就是沒有緩衝啦！
		Unbuffered 是需要有同時有 一頭寫入、另一頭讀出，才能動的。

		有Buffer(有儲存空間限制的通道) Buffered channel 跟 無Buffer(無限制的通道) Unbuffered channel
	*/

	// Unbuffered channel
	//showSample1()
	//showSample2()

	// Buffered channel
	//showSample3()
	showBlock()
}

func showBlock() {
	log.SetFlags(5)
	ch := make(chan int, 5)
	go func(ch chan int) {
		for {
			time.Sleep(time.Millisecond * 100)
			log.Println("got:", <-ch)
		}
	}(ch)

	for i := 0; i < 10; i++ {
		ch <- i
		log.Println("main set channel:", i)
	}
	time.Sleep(time.Second)
}

func showSample3() {
	// 因為通道有限制長度，所以只能有兩個值，不然就deadlock了
	fmt.Println("show sample3")
	ch := make(chan int, 2)
	go func(ch chan int) {
		fmt.Println(ch)
	}(ch)
	ch <- 1
	ch <- 2
	ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	time.Sleep(time.Second)
}

func showSample2() {
	fmt.Println("show sample2")
	ch := make(chan int)
	go func(ch chan int) {
		time.Sleep(time.Second)
		ch <- 999
	}(ch)
	fmt.Println("ch got:", <-ch)
}

func showSample1() {
	ch := make(chan int)
	go func(ch chan int) {
		fmt.Println("in goroutine")
		i := <-ch
		fmt.Println("i:", i)
	}(ch)
	fmt.Println("input 100 in channel", ch)
	ch <- 100
	time.Sleep(time.Second)
}

func func1(ch chan int) {
	i := <-ch
	fmt.Println("i:", i)
}

// 1. func中不能再有func了嗎? 除了Anonymous func外?
// A: 可以用Anonymous func 或是 closed func(閉包)
/* 閉包
innerFunc := func(msg string) {
	fmt.Println("Inner function says:", msg)
}
innerFunc("Hello!")
*/

// 2. 建立了channel之後，在goroutine的程式碼中，就會持續等待channel是否有資料? 否則下面的程式都不會執行了嗎?
// A: 對，當 goroutine 在 channel 讀取（<-channel）時，如果沒有資料，它會阻塞，直到有資料進來。
//這會導致 goroutine 卡住，直到其他 goroutine 將數據送進 channel，否則它不會往下執行。
