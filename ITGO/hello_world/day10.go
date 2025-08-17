package hello_world

import "fmt"

type Res struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func Day10() {
	showMap()
	showStruct()
	showNestedStructure()
}

func showNestedStructure() {
	type Wallet struct {
		Blue1000 int
		Red100   int
		Card     string
	}

	type PencilBox struct {
		Pencil string
		Pen    string
	}

	type Bag struct {
		Wallet
		PencilBox
		Books string
	}

	type Person struct {
		Bag
		Name string
	}

	bag := Bag{
		Wallet{Card: "卡", Red100: 5},
		PencilBox{Pencil: "Pentel", Pen: "Cross"},
		"讚讚",
	}

	var Tommy = Person{}
	Tommy.Name = "Tommy"
	Tommy.Bag = bag
	fmt.Println(Tommy)
	fmt.Printf("%+v\n", Tommy)
}

func showStruct() {
	res1 := new(Res)

	var res2 = new(Res)
	var res3 *Res
	res4 := &Res{Status: "ok", Msg: "ok"}
	fmt.Println(res1, res2, res3, res4)
	fmt.Printf("%+v, %+v, %+v, %+v,", res1, res2, res3, res4)
}

func showMap() {
	var Male = map[bool]string{
		true:  "公",
		false: "母",
	}
	fmt.Println(Male)

	var number = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	number["four"] = 4
	fmt.Println(Male[true], number["four"])

	for k, v := range number {
		fmt.Println(k, v)
	}
}
