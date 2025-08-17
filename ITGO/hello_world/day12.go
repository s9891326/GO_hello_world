package hello_world

import (
	"fmt"
	"reflect"
)

type Animal interface {
	eat()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (c Cat) eat() {
	fmt.Println(c.name, "吃飯")
}

func (d Dog) eat() {
	fmt.Print(d.name, " 吃吃")
}

func callEat(a Animal) {
	a.eat()
}

func Day12() {
	//showInterface()
	showInterfaceWithAnimal()

	//var hello interface{} = "hello"
	//helloStr, ok := hello.(string)
	//fmt.Println(helloStr, ok)
	//helloStr := hello.(int)
	//helloStr, ok := hello.(int)
	//fmt.Println(helloStr, ok)
}

func showInterfaceWithAnimal() {
	/*
		封裝（Encapsulation） 👉 用 struct + method
		多型（Polymorphism） 👉 用 interface
		組合（Composition） 👉 用 struct 內嵌其他 struct
	*/
	var c1 Animal = Cat{name: "肥貓"}
	//c1.eat()
	callEat(c1)

	var c2 Animal = Cat{name: "醜貓"}
	//c2.eat()
	callEat(c2)

	// 就可以快速換成狗了
	var dog1 Animal = Dog{name: "開心狗一號"}
	//dog1.eat()
	callEat(dog1)

	d := Dog{"eddy"}
	d.eat()
}

func showInterface() {
	var a interface{}
	fmt.Println(a, reflect.TypeOf(a))
	a = 123
	fmt.Println(a, reflect.TypeOf(a))
	a = "asdf"
	fmt.Println(a, reflect.TypeOf(a))
	a = true
	fmt.Println(a, reflect.TypeOf(a))
}
