package note

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 5.3 協程 Goroutine

var (
	c    int
	lock sync.Mutex
)

func primeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Println(n)
	lock.Lock() // 上互斥鎖
	c++
	lock.Unlock()
}

func Goroutine() {
	for i := 2; i < 10; i++ {
		go primeNum(i)
	}
	var key int
	fmt.Scanln(&key)
	fmt.Println("有幾個質數:", c)
}

func pushNum(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func pushPrimeNum(n int, c chan int)  {
	for i := 2; i < n; i++ {
		if n%i == 0{
			return
		}
	}
	c <- n
}


func Channel() {
	/**
	channel 主要用在 goroutine 中資料的傳遞，可以避免race condition等問題
	塞資料進channel c <- 1
	從channel取資料 print(<- c)
	**/

	// var c1 chan int = make(chan int)
	// go pushNum(c1)

	// for v:= range c1 {
	// 	fmt.Println(v)
	// }

	// for {
	// 	v, ok := <-c1
	// 	if ok {
	// 		fmt.Printf("%v\t", v)
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println(c1)

	c2 := make(chan int, 100)
	for i := 2; i < 100001; i++ {
		go pushPrimeNum(i, c2)
	}

	Print:
	for {
		select {
		case v:= <- c2:
			fmt.Printf("%v\t", v)
		default:
			fmt.Printf("done")
			break Print
		}
	}
}

func worker(c chan string) {
	for i := 0; i < 5; i++ {
		message := "hello" + strconv.Itoa(i)
		fmt.Println("Sent:", message)
		c <- message
		time.Sleep(3000)
	}
	close(c)
}


func ChannelWithGoroutine() {
	c := make(chan string)
	go worker(c)

	for v := range c {
		fmt.Println("receive:", v)
	}
}


