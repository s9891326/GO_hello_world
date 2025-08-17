package hello_world

import "fmt"

type Person struct {
	Questions string
}

// SetQuestion 如果方法需要修改 Person，用 *Person。
// 如果方法只是讀取 Person，用 Person。
// 指標接收者，能修改原始結構體
func (CIS *Person) SetQuestion(q string) {
	CIS.Questions = q
}

// GetQuestion 值接收者，不影響原始結構體
func (CIS Person) GetQuestion() string {
	return CIS.Questions
}

func Day11() {
	fmt.Println(callHaveReturnFunc())

	p1 := Person{Questions: "Initial Question"}
	fmt.Println(p1)
	p1.SetQuestion("New Question")
	fmt.Println(p1)
	fmt.Println(p1.GetQuestion())
}

func callHaveReturnFunc() (e string) {
	e = "hello"
	// 等校於 return e 因為已經在回傳值定義好回傳e這個變數了
	return
}
