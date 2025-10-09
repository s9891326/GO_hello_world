package main

import "fmt"

// 結構體方法：指針接收者
func (p *Person) SetAge(age int) {
	if age >= 0 && age <= 150 {
		p.Age = age
	}
}

func (p *Person) MoveTo(city string) {
	p.City = city
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("🎉 %s 生日快樂！現在 %d 歲了\n", p.Name, p.Age)
}

// 結構體方法：值接收者
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s (%d歲) 住在 %s", p.Name, p.Age, p.City)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func demonstrateStructPointers() {
	fmt.Println("\n--- 指針與結構體 ---")
	
	// 創建結構體實例
	person1 := Person{Name: "Bob", Age: 30, City: "台北"}
	fmt.Printf("🏠 person1: %+v\n", person1)
	
	// 創建指向結構體的指針
	personPtr := &person1
	fmt.Printf("🏠 指針地址: %p\n", personPtr)
	fmt.Printf("🏠 通過指針訪問: %+v\n", *personPtr)
	
	// Go 語言的語法糖：自動解引用
	fmt.Printf("🏠 姓名: %s (自動解引用)\n", personPtr.Name)
	fmt.Printf("🏠 年齡: %d (自動解引用)\n", personPtr.Age)
	
	// 通過指針修改結構體
	personPtr.Age = 31
	personPtr.City = "高雄"
	fmt.Printf("🏠 修改後: %+v\n", person1)
	
	// 使用 new 創建結構體指針
	person2 := new(Person)
	person2.Name = "Charlie"
	person2.Age = 28
	person2.City = "台中"
	fmt.Printf("🏠 new 創建: %+v\n", *person2)
	
	// 調用方法
	fmt.Printf("🏠 修改前信息: %s\n", person2.GetInfo())
	person2.SetAge(29)
	person2.MoveTo("台南")
	person2.HaveBirthday()
	fmt.Printf("🏠 修改後信息: %s\n", person2.GetInfo())
	fmt.Printf("🏠 是否成年: %t\n", person2.IsAdult())
}

// LinkedList 的方法
func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}
	
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) Display() {
	fmt.Print("🔗 鏈表: ")
	if ll.Head == nil {
		fmt.Println("空鏈表")
		return
	}
	
	current := ll.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Printf(" (大小: %d)\n", ll.Size)
}

func (ll *LinkedList) Find(value int) *Node {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList) Remove(value int) bool {
	if ll.Head == nil {
		return false
	}
	
	// 如果要刪除的是第一個節點
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}
	
	// 查找要刪除的節點
	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}
	
	return false
}

func demonstrateStructWithPointers() {
	fmt.Println("\n--- 結構體中的指針字段 ---")
	
	// 創建鏈表
	list := &LinkedList{}
	
	fmt.Println("🔗 創建空鏈表")
	list.Display()
	
	// 添加元素
	fmt.Println("🔗 添加元素 1, 2, 3, 4")
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Display()
	
	// 在開頭添加元素
	fmt.Println("🔗 在開頭添加元素 0")
	list.Prepend(0)
	list.Display()
	
	// 查找元素
	fmt.Println("🔗 查找元素")
	if node := list.Find(3); node != nil {
		fmt.Printf("   找到元素 3，地址: %p\n", node)
	} else {
		fmt.Println("   未找到元素 3")
	}
	
	if node := list.Find(10); node != nil {
		fmt.Printf("   找到元素 10，地址: %p\n", node)
	} else {
		fmt.Println("   未找到元素 10")
	}
	
	// 刪除元素
	fmt.Println("🔗 刪除元素 2")
	if list.Remove(2) {
		fmt.Println("   刪除成功")
	} else {
		fmt.Println("   刪除失敗")
	}
	list.Display()
	
	// 遍歷鏈表節點
	fmt.Println("🔗 遍歷節點地址:")
	current := list.Head
	index := 0
	for current != nil {
		fmt.Printf("   節點 %d: 值=%d, 地址=%p", index, current.Value, current)
		if current.Next != nil {
			fmt.Printf(", 下一個=%p", current.Next)
		}
		fmt.Println()
		current = current.Next
		index++
	}
}