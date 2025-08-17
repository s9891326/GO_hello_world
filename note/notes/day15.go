package note

import "fmt"

//func getRes(n1, n2 int) (sum, diff int) {
//	sum = n1 + n2
//	diff = n1 - n2
//	return
//}

func main15() {
	var getRes = func(n1, n2 int) (sum, diff int) {
		sum = n1 + n2
		diff = n1 - n2
		return
	}
	res1, res2 := getRes(5, 3)
	fmt.Println("res1: ", res1, " res2: ", res2)

	res3, res4 := func(n1, n2 int) (sum, diff int) {
		sum = n1 + n2
		diff = n1 - n2
		return
	}(5, 3)
	fmt.Println("res3: ", res3, " res4: ", res4)
}
