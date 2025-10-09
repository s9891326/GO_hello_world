// 練習 1 解答：基本 JSON 操作 - 學生成績管理系統
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

// Student 學生結構體
type Student struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Email  string  `json:"email"`
	Scores []Score `json:"scores"`
}

// Score 成績結構體
type Score struct {
	Subject string  `json:"subject"`
	Points  float64 `json:"points"`
	Date    string  `json:"date"`
}

// StudentManager 學生管理器
type StudentManager struct {
	students []Student
	nextID   int
}

// NewStudentManager 創建新的學生管理器
func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make([]Student, 0),
		nextID:   1,
	}
}

// AddStudent 添加學生
func (sm *StudentManager) AddStudent(student Student) error {
	// 驗證學生數據
	if student.Name == "" {
		return fmt.Errorf("學生姓名不能為空")
	}
	
	if student.Age <= 0 || student.Age > 100 {
		return fmt.Errorf("學生年齡無效: %d", student.Age)
	}
	
	if student.Email == "" {
		return fmt.Errorf("學生郵箱不能為空")
	}
	
	// 檢查 ID 是否已存在
	for _, existing := range sm.students {
		if existing.ID == student.ID {
			return fmt.Errorf("學生 ID %d 已存在", student.ID)
		}
	}
	
	// 如果沒有提供 ID，自動分配
	if student.ID == 0 {
		student.ID = sm.nextID
		sm.nextID++
	} else {
		// 更新 nextID 以避免衝突
		if student.ID >= sm.nextID {
			sm.nextID = student.ID + 1
		}
	}
	
	sm.students = append(sm.students, student)
	return nil
}

// GetStudentByID 根據 ID 獲取學生
func (sm *StudentManager) GetStudentByID(id int) (*Student, error) {
	for i, student := range sm.students {
		if student.ID == id {
			return &sm.students[i], nil
		}
	}
	return nil, fmt.Errorf("找不到 ID 為 %d 的學生", id)
}

// GetAllStudents 獲取所有學生
func (sm *StudentManager) GetAllStudents() []Student {
	return sm.students
}

// SaveToJSON 保存到 JSON 文件
func (sm *StudentManager) SaveToJSON(filename string) error {
	data, err := json.MarshalIndent(sm.students, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化學生數據錯誤: %w", err)
	}
	
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("寫入文件錯誤: %w", err)
	}
	
	return nil
}

// LoadFromJSON 從 JSON 文件加載
func (sm *StudentManager) LoadFromJSON(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("讀取文件錯誤: %w", err)
	}
	
	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		return fmt.Errorf("解析 JSON 錯誤: %w", err)
	}
	
	sm.students = students
	
	// 更新 nextID
	maxID := 0
	for _, student := range students {
		if student.ID > maxID {
			maxID = student.ID
		}
	}
	sm.nextID = maxID + 1
	
	return nil
}

// CalculateAverage 計算學生平均分
func (s *Student) CalculateAverage() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	
	total := 0.0
	for _, score := range s.Scores {
		total += score.Points
	}
	
	return total / float64(len(s.Scores))
}

// AddScore 為學生添加成績
func (sm *StudentManager) AddScore(studentID int, score Score) error {
	student, err := sm.GetStudentByID(studentID)
	if err != nil {
		return err
	}
	
	// 驗證成績數據
	if score.Subject == "" {
		return fmt.Errorf("科目名稱不能為空")
	}
	
	if score.Points < 0 || score.Points > 100 {
		return fmt.Errorf("成績必須在 0-100 之間")
	}
	
	if score.Date == "" {
		score.Date = time.Now().Format("2006-01-02")
	}
	
	student.Scores = append(student.Scores, score)
	return nil
}

// GetTopStudents 獲取成績最好的學生
func (sm *StudentManager) GetTopStudents(limit int) []Student {
	// 創建學生副本用於排序
	students := make([]Student, len(sm.students))
	copy(students, sm.students)
	
	// 按平均分排序
	sort.Slice(students, func(i, j int) bool {
		avgI := students[i].CalculateAverage()
		avgJ := students[j].CalculateAverage()
		return avgI > avgJ
	})
	
	// 返回前 limit 個學生
	if limit > len(students) {
		limit = len(students)
	}
	
	return students[:limit]
}

// GetStudentsBySubject 獲取指定科目的學生成績
func (sm *StudentManager) GetStudentsBySubject(subject string) []map[string]interface{} {
	var results []map[string]interface{}
	
	for _, student := range sm.students {
		for _, score := range student.Scores {
			if score.Subject == subject {
				results = append(results, map[string]interface{}{
					"student_id":   student.ID,
					"student_name": student.Name,
					"subject":      score.Subject,
					"points":       score.Points,
					"date":         score.Date,
				})
			}
		}
	}
	
	return results
}

// UpdateStudent 更新學生信息
func (sm *StudentManager) UpdateStudent(id int, updatedStudent Student) error {
	for i, student := range sm.students {
		if student.ID == id {
			// 保持原有的 ID 和成績
			updatedStudent.ID = id
			if len(updatedStudent.Scores) == 0 {
				updatedStudent.Scores = student.Scores
			}
			
			sm.students[i] = updatedStudent
			return nil
		}
	}
	return fmt.Errorf("找不到 ID 為 %d 的學生", id)
}

// DeleteStudent 刪除學生
func (sm *StudentManager) DeleteStudent(id int) error {
	for i, student := range sm.students {
		if student.ID == id {
			// 移除學生
			sm.students = append(sm.students[:i], sm.students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("找不到 ID 為 %d 的學生", id)
}

// 演示函數
func demonstrateStudentManager() {
	fmt.Println("=== 學生成績管理系統演示 ===")
	
	// 創建管理器
	sm := NewStudentManager()
	
	// 添加學生
	fmt.Println("\n--- 添加學生 ---")
	students := []Student{
		{
			Name:  "Alice Johnson",
			Age:   20,
			Email: "alice@university.edu",
			Scores: []Score{
				{Subject: "Math", Points: 95.5, Date: "2023-10-15"},
				{Subject: "Physics", Points: 88.0, Date: "2023-10-20"},
				{Subject: "Chemistry", Points: 92.5, Date: "2023-10-25"},
			},
		},
		{
			Name:  "Bob Smith",
			Age:   21,
			Email: "bob@university.edu",
			Scores: []Score{
				{Subject: "Math", Points: 87.0, Date: "2023-10-15"},
				{Subject: "Physics", Points: 94.5, Date: "2023-10-20"},
				{Subject: "Chemistry", Points: 89.0, Date: "2023-10-25"},
			},
		},
		{
			Name:  "Charlie Brown",
			Age:   19,
			Email: "charlie@university.edu",
			Scores: []Score{
				{Subject: "Math", Points: 92.0, Date: "2023-10-15"},
				{Subject: "Physics", Points: 85.5, Date: "2023-10-20"},
			},
		},
	}
	
	for _, student := range students {
		err := sm.AddStudent(student)
		if err != nil {
			fmt.Printf("添加學生錯誤: %v\n", err)
		} else {
			fmt.Printf("成功添加學生: %s (ID: %d)\n", student.Name, student.ID)
		}
	}
	
	// 顯示所有學生
	fmt.Println("\n--- 所有學生信息 ---")
	for _, student := range sm.GetAllStudents() {
		avg := student.CalculateAverage()
		fmt.Printf("ID: %d, 姓名: %s, 年齡: %d, 郵箱: %s, 平均分: %.2f\n",
			student.ID, student.Name, student.Age, student.Email, avg)
	}
	
	// 獲取成績最好的學生
	fmt.Println("\n--- 成績前 2 名學生 ---")
	topStudents := sm.GetTopStudents(2)
	for i, student := range topStudents {
		avg := student.CalculateAverage()
		fmt.Printf("%d. %s - 平均分: %.2f\n", i+1, student.Name, avg)
	}
	
	// 按科目查詢
	fmt.Println("\n--- Math 科目成績 ---")
	mathScores := sm.GetStudentsBySubject("Math")
	for _, score := range mathScores {
		fmt.Printf("%s: %.1f 分 (日期: %s)\n", 
			score["student_name"], score["points"], score["date"])
	}
	
	// 保存到 JSON 文件
	fmt.Println("\n--- 保存到 JSON 文件 ---")
	err := sm.SaveToJSON("students.json")
	if err != nil {
		fmt.Printf("保存錯誤: %v\n", err)
	} else {
		fmt.Println("成功保存到 students.json")
	}
	
	// 測試從文件加載
	fmt.Println("\n--- 從 JSON 文件加載 ---")
	newSM := NewStudentManager()
	err = newSM.LoadFromJSON("students.json")
	if err != nil {
		fmt.Printf("加載錯誤: %v\n", err)
	} else {
		fmt.Printf("成功從文件加載 %d 個學生\n", len(newSM.GetAllStudents()))
	}
	
	// 添加新成績
	fmt.Println("\n--- 添加新成績 ---")
	err = sm.AddScore(1, Score{
		Subject: "English",
		Points:  96.0,
		Date:    "2023-11-01",
	})
	if err != nil {
		fmt.Printf("添加成績錯誤: %v\n", err)
	} else {
		fmt.Println("成功為學生 1 添加 English 成績")
		student, _ := sm.GetStudentByID(1)
		fmt.Printf("更新後平均分: %.2f\n", student.CalculateAverage())
	}
}

// 測試錯誤處理
func testErrorHandling() {
	fmt.Println("\n=== 錯誤處理測試 ===")
	
	sm := NewStudentManager()
	
	// 測試無效學生數據
	fmt.Println("\n--- 測試無效數據 ---")
	invalidStudents := []Student{
		{Name: "", Age: 20, Email: "test@example.com"},    // 空姓名
		{Name: "Test", Age: -5, Email: "test@example.com"}, // 無效年齡
		{Name: "Test", Age: 20, Email: ""},                 // 空郵箱
	}
	
	for i, student := range invalidStudents {
		err := sm.AddStudent(student)
		if err != nil {
			fmt.Printf("測試 %d: 預期錯誤 - %v\n", i+1, err)
		}
	}
	
	// 測試重複 ID
	fmt.Println("\n--- 測試重複 ID ---")
	student1 := Student{ID: 100, Name: "Student1", Age: 20, Email: "s1@example.com"}
	student2 := Student{ID: 100, Name: "Student2", Age: 21, Email: "s2@example.com"}
	
	sm.AddStudent(student1)
	err := sm.AddStudent(student2)
	if err != nil {
		fmt.Printf("重複 ID 錯誤: %v\n", err)
	}
	
	// 測試查找不存在的學生
	fmt.Println("\n--- 測試查找不存在的學生 ---")
	_, err = sm.GetStudentByID(999)
	if err != nil {
		fmt.Printf("查找不存在學生錯誤: %v\n", err)
	}
	
	// 測試無效文件
	fmt.Println("\n--- 測試無效文件 ---")
	err = sm.LoadFromJSON("nonexistent.json")
	if err != nil {
		fmt.Printf("文件不存在錯誤: %v\n", err)
	}
}

// 清理測試文件
func cleanup() {
	os.Remove("students.json")
	fmt.Println("清理完成")
}

func main() {
	demonstrateStudentManager()
	testErrorHandling()
	
	fmt.Println("\n按 Enter 鍵清理測試文件...")
	fmt.Scanln()
	
	cleanup()
	
	fmt.Println("===== 演示完成 =====")
}