// 練習 2 解答：公司配置系統
package main

import "fmt"

// 公司基本信息常數
const (
	CompanyName    = "ABC 科技股份有限公司"
	FoundedYear    = 2010
	CompanyAddress = "台北市信義區信義路五段7號"
	CompanyPhone   = "02-2720-1234"
	CompanyEmail   = "info@abc-tech.com"
)

// 員工職級枚舉
const (
	LevelIntern = iota // 實習生
	LevelJunior        // 初級工程師
	LevelMid           // 中級工程師
	LevelSenior        // 高級工程師
	LevelLead          // 技術主管
)

// 薪資等級常數 (新台幣/月)
const (
	SalaryIntern = 25000
	SalaryJunior = 45000
	SalaryMid    = 70000
	SalarySenior = 100000
	SalaryLead   = 150000
)

// 文件大小
const (
	Byte = 1
	_    = iota
	KB   = 1 << (10 * iota)
	MB
	GB
	TB
)

// 員工結構
type Employee struct {
	ID         int
	Name       string
	Level      int
	Salary     int
	Department string
}

func exercise1() {
	fmt.Println("=== 個人資料管理系統 ===")

	// 方式 1: 標準聲明
	var name string
	name = "李小明"

	// 方式 2: 聲明並初始化
	var age int = 28

	// 方式 3: 類型推導
	var height = 179.5

	// 方式 4: 短變數聲明
	isMarried := true
	email := "ming@example.com"

	// 多變數聲明
	var (
		address    = "台北市信義區"
		phone      = "0912-345-678"
		occupation = "軟體工程師"
	)

	// 顯示個人資料
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("年齡：%d 歲\n", age)
	fmt.Printf("身高：%.1f cm\n", height)
	fmt.Printf("婚姻狀況：%s\n", getMaritalStatus(isMarried))
	fmt.Printf("電子郵件：%s\n", email)
	fmt.Printf("地址：%s\n", address)
	fmt.Printf("電話：%s\n", phone)
	fmt.Printf("職業：%s\n", occupation)

	// 計算 BMI
	weight := 80.0
	bmi := weight / ((height / 100) * (height / 100))
	fmt.Printf("體重：%.1f kg\n", weight)
	fmt.Printf("BMI：%.1f (%s)\n", bmi, getBMICategory(bmi))
}

func exercise2() {
	fmt.Printf("=== %s ===\n", CompanyName)
	fmt.Printf("成立年份：%d\n", FoundedYear)
	fmt.Printf("公司地址：%s\n", CompanyAddress)
	fmt.Printf("聯絡電話：%s\n", CompanyPhone)
	fmt.Printf("電子郵件：%s\n", CompanyEmail)

	// 顯示職級系統
	fmt.Println("\n--- 員工職級系統 ---")
	levels := map[int]string{
		LevelIntern: "實習生",
		LevelJunior: "初級工程師",
		LevelMid:    "中級工程師",
		LevelSenior: "高級工程師",
		LevelLead:   "技術主管",
	}

	salaries := map[int]int{
		LevelIntern: SalaryIntern,
		LevelJunior: SalaryJunior,
		LevelMid:    SalaryMid,
		LevelSenior: SalarySenior,
		LevelLead:   SalaryLead,
	}

	for level := LevelIntern; level <= LevelLead; level++ {
		fmt.Printf("%d: %s (薪資: NT$%d)\n", level, levels[level], salaries[level])
	}

	// 創建員工實例
	fmt.Println("\n--- 員工資料 ---")
	employees := []Employee{
		{ID: 1001, Name: "張小華", Level: LevelSenior, Salary: SalarySenior, Department: "後端開發"},
		{ID: 1002, Name: "李小明", Level: LevelMid, Salary: SalaryMid, Department: "前端開發"},
		{ID: 1003, Name: "王小美", Level: LevelLead, Salary: SalaryLead, Department: "技術管理"},
		{ID: 1004, Name: "陳小剛", Level: LevelJunior, Salary: SalaryJunior, Department: "測試"},
	}

	for _, emp := range employees {
		fmt.Printf("員工編號：%d\n", emp.ID)
		fmt.Printf("姓名：%s\n", emp.Name)
		fmt.Printf("職級：%s\n", levels[emp.Level])
		fmt.Printf("薪資：NT$%d\n", emp.Salary)
		fmt.Printf("部門：%s\n", emp.Department)
		fmt.Println("---")
	}
}

func getMaritalStatus(married bool) string {
	if married {
		return "已婚"
	}
	return "未婚"
}

func getBMICategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "體重過輕"
	case bmi < 24:
		return "正常體重"
	case bmi < 27:
		return "體重過重"
	default:
		return "肥胖"
	}
}

func exercise3() {
	/*
		創建一個文件系統管理程序：

		要求：
		- 使用 iota 定義文件大小單位（Byte, KB, MB, GB, TB）
		- 定義文件類型常數（文檔、圖片、視頻、音頻）
		- 計算並顯示不同大小文件的字節數
		- 創建文件信息變數並顯示
	*/
	fmt.Println("--- 文件系統管理 ---")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)

	var (
		txt = "文檔"
		img = "圖片"
		vid = "視頻"
		aud = "音頻"
	)
	fmt.Printf("文件類型：%s, %s, %s, %s\n", txt, img, vid, aud)

}

func main() {
	// 練習 1：個人資料管理
	fmt.Println("\n=== 練習 1：個人資料管理 ===")
	exercise1()

	// 練習 2：公司配置系統
	fmt.Println("\n=== 練習 2：公司配置系統 ===")
	exercise2()

	// 練習 3：文件系統管理
	fmt.Println("\n=== 練習 3：文件系統管理 ===")
	exercise3()
}
