package main

import "fmt"

// Person 結構體的方法定義

// 值接收者方法
func (p Person) GetFullInfo() string {
	return fmt.Sprintf("%s, %d歲, 住在%s", p.Name, p.Age, p.City)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) IsElder() bool {
	return p.Age >= 65
}

// 值接收者方法不能修改結構體
func (p Person) TryToAge() {
	p.Age++ // 這不會影響原始結構體
	fmt.Printf("在方法內部 Age 變成: %d\n", p.Age)
}

// 指針接收者方法可以修改結構體
func (p *Person) SetAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("年齡必須在 0-150 之間")
	}
	p.Age = age
	return nil
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("🎉 生日快樂！%s 現在 %d 歲了\n", p.Name, p.Age)
}

func (p *Person) MoveTo(city string) {
	oldCity := p.City
	p.City = city
	fmt.Printf("📍 %s 從 %s 搬到了 %s\n", p.Name, oldCity, city)
}

func (p *Person) UpdateInfo(name string, age int, city string) error {
	if name == "" {
		return fmt.Errorf("姓名不能為空")
	}
	if age < 0 || age > 150 {
		return fmt.Errorf("年齡必須在 0-150 之間")
	}
	
	p.Name = name
	p.Age = age
	p.City = city
	return nil
}

// 演示方法的使用
func demonstrateMethods() {
	fmt.Println("\n--- 結構體方法演示 ---")
	
	person := Person{
		Name: "Frank",
		Age:  25,
		City: "新竹",
	}
	
	// 調用值接收者方法
	fmt.Println("📝 個人信息:", person.GetFullInfo())
	fmt.Printf("🔞 是否成年: %t\n", person.IsAdult())
	fmt.Printf("👴 是否長者: %t\n", person.IsElder())
	
	// 值接收者方法不會修改原結構體
	fmt.Printf("🎂 嘗試增加年齡前: %d\n", person.Age)
	person.TryToAge()
	fmt.Printf("🎂 嘗試增加年齡後: %d (原結構體未改變)\n", person.Age)
	
	// 調用指針接收者方法
	err := person.SetAge(26)
	if err != nil {
		fmt.Printf("❌ 設置年齡失敗: %v\n", err)
	} else {
		fmt.Printf("✅ 設置年齡成功: %d\n", person.Age)
	}
	
	person.HaveBirthday()
	person.MoveTo("桃園")
	
	// 更新完整信息
	err = person.UpdateInfo("Frank Chen", 28, "台中")
	if err != nil {
		fmt.Printf("❌ 更新信息失敗: %v\n", err)
	} else {
		fmt.Printf("✅ 更新後信息: %s\n", person.GetFullInfo())
	}
}