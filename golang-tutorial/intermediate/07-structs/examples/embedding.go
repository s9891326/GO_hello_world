package main

import "fmt"

// 演示結構體組合和嵌入

// Manager 的方法
func (m Manager) GetManagementInfo() string {
	return fmt.Sprintf("經理 %s 管理 %s 部門的 %d 個人", 
		m.Name, m.Department, m.TeamSize)
}

func (m *Manager) PromoteEmployee() {
	m.TeamSize++
	fmt.Printf("👔 %s 的團隊增加了一個成員，現在有 %d 人\n", m.Name, m.TeamSize)
}

func (m *Manager) SetDepartment(dept string) {
	oldDept := m.Department
	m.Department = dept
	fmt.Printf("🏢 %s 從 %s 調到了 %s\n", m.Name, oldDept, dept)
}

// Employee 的方法
func (e Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("員工ID: %d, %s, 職位: %s, 薪資: %.2f", 
		e.ID, e.Person.GetFullInfo(), e.Position, e.Salary)
}

func (e *Employee) Promote(newPosition string, salaryIncrease float64) {
	oldPosition := e.Position
	e.Position = newPosition
	e.Salary += salaryIncrease
	fmt.Printf("🎯 %s 從 %s 晉升為 %s，薪資增加 %.2f\n", 
		e.Person.Name, oldPosition, newPosition, salaryIncrease)
}

func demonstrateComposition() {
	fmt.Println("\n--- 結構體組合和嵌入演示 ---")
	
	// 創建組合結構體（Employee）
	employee := Employee{
		ID: 1001,
		Person: Person{
			Name: "Grace",
			Age:  28,
			City: "台北",
		},
		Address: Address{
			Street:  "信義路100號",
			City:    "台北市",
			ZipCode: "110",
			Country: "台灣",
		},
		Contact: Contact{
			Email: "grace@company.com",
			Phone: "02-1234-5678",
		},
		Salary:   60000,
		Position: "軟體工程師",
	}
	
	fmt.Println("👩‍💼 員工信息:")
	fmt.Println("  ", employee.GetEmployeeInfo())
	fmt.Printf("   聯繫方式: %s, %s\n", employee.Contact.Email, employee.Contact.Phone)
	fmt.Printf("   地址: %s, %s\n", employee.Address.Street, employee.Address.City)
	
	// 晉升員工
	employee.Promote("高級軟體工程師", 10000)
	
	// 創建嵌入結構體（Manager）
	manager := Manager{
		Person: Person{
			Name: "Henry",
			Age:  35,
			City: "台北",
		},
		Address: Address{
			Street:  "敦化南路200號",
			City:    "台北市",
			ZipCode: "106",
			Country: "台灣",
		},
		Contact: Contact{
			Email: "henry@company.com",
			Phone: "02-5678-9012",
		},
		EmployeeID: 2001,
		Department: "工程部",
		TeamSize:   5,
	}
	
	fmt.Println("\n👨‍💼 經理信息:")
	// 直接訪問嵌入結構體的字段
	fmt.Printf("   姓名: %s (來自 Person)\n", manager.Name)
	fmt.Printf("   年齡: %d (來自 Person)\n", manager.Age)
	fmt.Printf("   郵箱: %s (來自 Contact)\n", manager.Email)
	
	// 訪問嵌入結構體的方法
	fmt.Println("   個人信息:", manager.GetFullInfo()) // 來自 Person
	fmt.Println("   管理信息:", manager.GetManagementInfo())
	
	// 調用 Manager 的方法
	manager.PromoteEmployee()
	manager.SetDepartment("產品部")
	
	// 演示字段名衝突
	demonstrateFieldConflicts()
}

// 演示字段名衝突的處理
type Student struct {
	Person
	Address
	StudentID string
	Grade     int
	Major     string
}

func (s Student) GetStudentInfo() string {
	return fmt.Sprintf("學生 %s，%d 年級，主修 %s", s.Name, s.Grade, s.Major)
}

func demonstrateFieldConflicts() {
	fmt.Println("\n--- 字段名衝突處理 ---")
	
	student := Student{
		Person: Person{
			Name: "Isabella",
			Age:  20,
			City: "學校宿舍",
		},
		Address: Address{
			Street:  "大學路300號",
			City:    "台中市",
			ZipCode: "402",
			Country: "台灣",
		},
		StudentID: "S2024001",
		Grade:     2,
		Major:     "資訊工程",
	}
	
	// 處理 City 字段衝突
	fmt.Printf("👩‍🎓 學生居住地: %s (Person.City)\n", student.Person.City)
	fmt.Printf("🏫 學校地址: %s (Address.City)\n", student.Address.City)
	
	// 無衝突的字段可以直接訪問
	fmt.Printf("📚 學生信息: %s\n", student.GetStudentInfo())
	fmt.Printf("📧 學生郵箱: %s (如果 Contact 也嵌入會怎樣？)\n", "需要明確指定")
}