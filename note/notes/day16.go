package note

import "fmt"

func deferUtil() func(int) int  {
	i := 0
	return func(n int) int {
		fmt.Println("輸入為: ", n)
		i ++
		fmt.Println("第幾次使用: ", i)
		return i
	}
}

func main16() {
	f := deferUtil()
	//f(1)
	//f(2)

	//defer f(1)
	//f(2)

	//defer f(f(3))
	//f(2)

	defer f(1)
	defer f(2)
	defer f(3)
	f(4)

}
