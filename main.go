package main

import (
	"hello_world/note"
	// "hello_world/note/notes"
	// "hello_world/note/util"
)

// var Day17 = util.F("main.Day17")

// func init() {``
// 	util.F("main.init1")
// }

// func init() {
// 	util.F("main.init2")
// }

func main() {
	//fmt.Println(hello.Test)
	//fmt.Println(test_import.Test)

	// init函數: 每個package都可以有自己的init函數。
	// 執行順序(取決於包的依賴關係):
	// 被依賴包的全局變量 -> 被依賴包的init -> ... -> main包全局變量 -> main包的init
	// day17.HelloDay17()
	// note.Channel()
	note.PackageTime()

}
