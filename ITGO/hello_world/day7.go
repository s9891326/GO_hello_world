package hello_world

import (
	"fmt"
	"log"
	"time"
)

func scanIf() {
	var a int
	fmt.Printf("input a:")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return
	}

	if a < 50 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func scanSwitch() {
	var i int
	fmt.Printf("input i:")
	_, err := fmt.Scanln(&i)
	if err != nil {
		return
	}
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2, 3")
	}
}

func forRange() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, num := range nums {
		fmt.Println(index, num)
	}

	for index := range nums {
		fmt.Println(index)
	}

	// map 是無序的，除非先把key抓出來排序後再抓出才能固定順序，不然就使用struct的方式來創建
	fruits := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range fruits {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range fruits {
		fmt.Printf("%s -> %s\n", k, fruits[k])
	}
}

func showLog() {
	/*
		const (
			Ldate         = 1 << iota     // local time zone: 2009/01/23
			Ltime               //2       // local time zone: 01:23:23
			Lmicroseconds       //4       // microsecond : 01:23:23.123123.
			Llongfile           //8       // full filename, line: /a/b/c/d.go:23
			Lshortfile          //16      // filename element, line: d.go:23
			LUTC                //32      // use UTC
			Lmsgprefix          //64      // move the "prefix" from the beginning of the line to before the message
			LstdFlags     = Ldate | Ltime    //3 預設Flag // initial standard logger
		)
	*/
	log.SetPrefix("安安，我是log: ")
	log.SetFlags(log.Lshortfile + log.LstdFlags)
	for i := 0; i < 3; i++ {
		log.Println("i:", i)
		fmt.Println("i:", i)
		time.Sleep(time.Second * 1)
		log.Println("hi")
		fmt.Println("hi")
		if i == 2 {
			// 用於發生錯誤時印出log並退出
			log.Fatalln("發生錯誤")
		}
	}
}

func Day7() {
	//scanIf()
	//scanSwitch()
	forRange()
	showLog()
}
