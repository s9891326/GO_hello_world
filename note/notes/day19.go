package note

import (
	"fmt"
	"hello_world/note/util"
)

func main19() {
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4]
	fmt.Println(s1)

	s1[0] = 0
	fmt.Println("s1=", s1)
	fmt.Println("array=", array)

	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array=", array)

	var s3 []int
	fmt.Println("s3 is nil? ", s3 == nil)
	s3 = make([]int, 3, 5)
	fmt.Println("s1= ", s1)
	s3 = append(s1, 4, 5, 6, 7) // 底層創建了新的數組，不再使用原數組，array不會跟著被修改
	s3[1] = 222
	fmt.Println("s3= ", s3)
	fmt.Println("array=", array)

	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	s6 := make([]int, 3)
	copy(s6, s5)           // 容量能接收多少，就接收多少
	fmt.Println("s6=", s6) // s6=[0 0 4]
	copy(s5, s6)           // 容量能接收多少，就接收多少
	fmt.Println("s5=", s5) // s5=[0 0 4 0 4]

	fmt.Println("\n4.2.5 string 與 []byte")
	str := "hello world"
	fmt.Printf("[]byte(str):%v\n[]byte(str): %s\n", []byte(str), []byte(str))

	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}
	key := util.SelectByKey("aaa", "bbbb", "ccc")
	fmt.Println("key=", key)

}
