package hello_world

import "fmt"

func Day5() {
	type Name struct {
		A string
		B bool
		C int
	}

	fmt.Printf("%v	\n", Name{})  // { false 0}
	fmt.Printf("%+v	\n", Name{}) // {A: B:false C:0}
	fmt.Printf("%#v	\n", Name{}) // main.Name{A:"", B:false, C:0}
}
