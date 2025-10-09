// 練習 1 解答：學生管理系統
package main

import (
	"fmt"
	"time"
)

// 學生結構體
type Student struct {
	ID      string
	Name    string
	Age     int
	Major   string
	Email   string
	Grades  map[string]float64 // 課程ID -> 成績
	Credits map[string]int     // 課程ID -> 學分
}

// 課程結構體
type Course struct {
	ID       string
	Name     string
	Credits  int
	Teacher  string
	Semester string
}

// 成績記錄結構體
type Grade struct {
	StudentID string
	CourseID  string
	Score     float64
	Semester  string
	CreatedAt time.Time
}

// 學生方法
func (s *Student) AddCourse(course Course, score float64) {
	if s.Grades == nil {
		s.Grades = make(map[string]float64)
	}
	if s.Credits == nil {
		s.Credits = make(map[string]int)
	}
	
	s.Grades[course.ID] = score
	s.Credits[course.ID] = course.Credits
	fmt.Printf("✅ %s 成功註冊課程: %s (%d學分), 成績: %.1f\n", 
		s.Name, course.Name, course.Credits, score)
}

func (s Student) CalculateGPA() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	
	var totalPoints float64
	var totalCredits int
	
	for courseID, score := range s.Grades {
		credits := s.Credits[courseID]
		gp := scoreToGradePoint(score)
		totalPoints += gp * float64(credits)
		totalCredits += credits
	}
	
	if totalCredits == 0 {
		return 0.0
	}
	
	return totalPoints / float64(totalCredits)
}

func (s Student) GetTotalCredits() int {
	total := 0
	for _, credits := range s.Credits {
		total += credits
	}
	return total
}

func (s Student) GetAverageScore() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	
	total := 0.0
	for _, score := range s.Grades {
		total += score
	}
	return total / float64(len(s.Grades))
}

func (s Student) GetGradeLevel() string {
	gpa := s.CalculateGPA()
	switch {
	case gpa >= 3.7:
		return "優秀"
	case gpa >= 3.0:
		return "良好"
	case gpa >= 2.0:
		return "及格"
	default:
		return "需要改進"
	}
}

// 輔助函數
func scoreToGradePoint(score float64) float64 {
	switch {
	case score >= 90:
		return 4.0
	case score >= 80:
		return 3.0
	case score >= 70:
		return 2.0
	case score >= 60:
		return 1.0
	default:
		return 0.0
	}
}

// 構造函數
func NewStudent(id, name string, age int, major, email string) *Student {
	return &Student{
		ID:      id,
		Name:    name,
		Age:     age,
		Major:   major,
		Email:   email,
		Grades:  make(map[string]float64),
		Credits: make(map[string]int),
	}
}

func NewCourse(id, name string, credits int, teacher, semester string) Course {
	return Course{
		ID:       id,
		Name:     name,
		Credits:  credits,
		Teacher:  teacher,
		Semester: semester,
	}
}

func main() {
	fmt.Println("=== 學生管理系統 ===")
	
	// 創建學生
	student := NewStudent("S2024001", "張小明", 20, "計算機科學", "ming@university.edu")
	fmt.Printf("🎓 學生創建成功: %s (%s)\n", student.Name, student.ID)
	fmt.Printf("   專業: %s, 年齡: %d\n", student.Major, student.Age)
	
	// 創建課程
	courses := []Course{
		NewCourse("CS101", "資料結構", 3, "李教授", "2024春"),
		NewCourse("CS102", "演算法", 3, "王教授", "2024春"),
		NewCourse("CS103", "操作系統", 4, "陳教授", "2024春"),
		NewCourse("MATH201", "離散數學", 3, "林教授", "2024春"),
	}
	
	// 學生註冊課程並錄入成績
	fmt.Println("\n📚 課程註冊和成績錄入:")
	student.AddCourse(courses[0], 85.0)
	student.AddCourse(courses[1], 92.0)
	student.AddCourse(courses[2], 78.0)
	student.AddCourse(courses[3], 88.0)
	
	// 顯示學期總結
	fmt.Println("\n📊 學期總結:")
	fmt.Printf("   學生: %s (%s)\n", student.Name, student.Major)
	fmt.Printf("   總學分: %d\n", student.GetTotalCredits())
	fmt.Printf("   平均分: %.1f\n", student.GetAverageScore())
	fmt.Printf("   GPA: %.2f\n", student.CalculateGPA())
	fmt.Printf("   等級: %s\n", student.GetGradeLevel())
	
	// 詳細成績單
	fmt.Println("\n📋 詳細成績單:")
	courseMap := make(map[string]Course)
	for _, course := range courses {
		courseMap[course.ID] = course
	}
	
	for courseID, score := range student.Grades {
		course := courseMap[courseID]
		gp := scoreToGradePoint(score)
		fmt.Printf("   %s (%s): %.1f分 (%.1f GPA, %d學分)\n", 
			course.Name, course.ID, score, gp, course.Credits)
	}
}