// 練習 1 解答：創建工具包
package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode"
)

// ============ math 包 ============

// Calculator 提供基本數學運算
type Calculator struct{}

// NewCalculator 創建新的計算器實例
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add 加法運算
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract 減法運算
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 乘法運算
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide 除法運算
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除數不能為零")
	}
	return a / b, nil
}

// Power 冪運算
func (c *Calculator) Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Sqrt 平方根運算
func (c *Calculator) Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("不能計算負數的平方根")
	}
	return math.Sqrt(x), nil
}

// ============ string 包 ============

// Processor 提供字符串處理功能
type Processor struct{}

// NewProcessor 創建新的字符串處理器實例
func NewProcessor() *Processor {
	return &Processor{}
}

// Reverse 反轉字符串
func (p *Processor) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 檢查是否為回文
func (p *Processor) IsPalindrome(s string) bool {
	// 移除空格並轉換為小寫
	cleaned := strings.ToLower(p.RemoveSpaces(s))
	return cleaned == p.Reverse(cleaned)
}

// RemoveSpaces 移除所有空格
func (p *Processor) RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// CountWords 計算單詞數量
func (p *Processor) CountWords(s string) int {
	// 使用正則表達式分割單詞
	re := regexp.MustCompile(`\S+`)
	words := re.FindAllString(s, -1)
	return len(words)
}

// ToSlug 轉換為 URL 友好格式
func (p *Processor) ToSlug(s string) string {
	// 轉換為小寫
	s = strings.ToLower(s)
	
	// 移除非字母數字字符，替換為連字符
	var result strings.Builder
	var lastWasHyphen bool
	
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
			lastWasHyphen = false
		} else if !lastWasHyphen {
			result.WriteRune('-')
			lastWasHyphen = true
		}
	}
	
	// 移除開頭和結尾的連字符
	slug := strings.Trim(result.String(), "-")
	return slug
}

// ============ time 包 ============

import (
	"time"
)

// Helper 提供時間處理功能
type TimeHelper struct{}

// NewTimeHelper 創建新的時間助手實例
func NewTimeHelper() *TimeHelper {
	return &TimeHelper{}
}

// FormatDuration 格式化時間間隔
func (h *TimeHelper) FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.0f秒", d.Seconds())
	} else if d < time.Hour {
		return fmt.Sprintf("%.0f分鐘", d.Minutes())
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.1f小時", d.Hours())
	} else {
		return fmt.Sprintf("%.1f天", d.Hours()/24)
	}
}

// IsWeekend 檢查是否為週末
func (h *TimeHelper) IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// AddBusinessDays 添加工作日
func (h *TimeHelper) AddBusinessDays(t time.Time, days int) time.Time {
	result := t
	addedDays := 0
	
	for addedDays < days {
		result = result.AddDate(0, 0, 1)
		if !h.IsWeekend(result) {
			addedDays++
		}
	}
	
	return result
}

// StartOfDay 獲取當天開始時間
func (h *TimeHelper) StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 獲取當天結束時間
func (h *TimeHelper) EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// ============ validator 包 ============

import (
	"net/mail"
	"regexp"
)

// Rules 提供驗證規則
type Rules struct{}

// NewRules 創建新的驗證規則實例
func NewRules() *Rules {
	return &Rules{}
}

// IsEmail 驗證郵箱格式
func (r *Rules) IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsPhone 驗證手機號碼（簡單格式）
func (r *Rules) IsPhone(phone string) bool {
	// 匹配常見的手機號格式
	patterns := []string{
		`^\+?1?[0-9]{10,11}$`,           // 美國格式
		`^09[0-9]{8}$`,                  // 台灣格式
		`^1[3-9][0-9]{9}$`,              // 中國大陸格式
		`^\+[1-9][0-9]{1,14}$`,          // 國際格式
	}
	
	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, phone)
		if matched {
			return true
		}
	}
	return false
}

// IsURL 驗證 URL 格式
func (r *Rules) IsURL(url string) bool {
	pattern := `^https?://[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]+$`
	matched, _ := regexp.MatchString(pattern, url)
	return matched
}

// IsStrongPassword 驗證強密碼
func (r *Rules) IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	
	// 檢查是否包含大寫字母
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	// 檢查是否包含小寫字母
	hasLower, _ := regexp.MatchString(`[a-z]`, password)
	// 檢查是否包含數字
	hasDigit, _ := regexp.MatchString(`[0-9]`, password)
	// 檢查是否包含特殊字符
	hasSpecial, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password)
	
	return hasUpper && hasLower && hasDigit && hasSpecial
}

// IsInRange 檢查數值是否在範圍內
func (r *Rules) IsInRange(value, min, max float64) bool {
	return value >= min && value <= max
}

// IsNotEmpty 檢查字符串是否非空
func (r *Rules) IsNotEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}

// ============ 主程序測試 ============

func main() {
	fmt.Println("=== 工具包使用示例 ===")
	
	// 測試數學計算器
	fmt.Println("\n--- 數學計算器測試 ---")
	calc := NewCalculator()
	
	fmt.Printf("5 + 3 = %.2f\n", calc.Add(5, 3))
	fmt.Printf("10 - 4 = %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("6 * 7 = %.2f\n", calc.Multiply(6, 7))
	
	if result, err := calc.Divide(15, 3); err != nil {
		fmt.Printf("除法錯誤: %v\n", err)
	} else {
		fmt.Printf("15 ÷ 3 = %.2f\n", result)
	}
	
	fmt.Printf("2^3 = %.2f\n", calc.Power(2, 3))
	
	if result, err := calc.Sqrt(16); err != nil {
		fmt.Printf("平方根錯誤: %v\n", err)
	} else {
		fmt.Printf("√16 = %.2f\n", result)
	}
	
	// 測試字符串處理器
	fmt.Println("\n--- 字符串處理器測試 ---")
	processor := NewProcessor()
	
	testStr := "Hello World"
	fmt.Printf("原字符串: %s\n", testStr)
	fmt.Printf("反轉: %s\n", processor.Reverse(testStr))
	fmt.Printf("移除空格: %s\n", processor.RemoveSpaces(testStr))
	fmt.Printf("單詞數量: %d\n", processor.CountWords(testStr))
	fmt.Printf("URL 格式: %s\n", processor.ToSlug(testStr))
	
	palindrome := "A man a plan a canal Panama"
	fmt.Printf("'%s' 是回文: %t\n", palindrome, processor.IsPalindrome(palindrome))
	
	// 測試時間助手
	fmt.Println("\n--- 時間助手測試 ---")
	timeHelper := NewTimeHelper()
	
	now := time.Now()
	fmt.Printf("當前時間: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("是否週末: %t\n", timeHelper.IsWeekend(now))
	
	businessDay := timeHelper.AddBusinessDays(now, 5)
	fmt.Printf("5個工作日後: %s\n", businessDay.Format("2006-01-02"))
	
	startOfDay := timeHelper.StartOfDay(now)
	endOfDay := timeHelper.EndOfDay(now)
	fmt.Printf("今天開始: %s\n", startOfDay.Format("2006-01-02 15:04:05"))
	fmt.Printf("今天結束: %s\n", endOfDay.Format("2006-01-02 15:04:05"))
	
	// 測試驗證規則
	fmt.Println("\n--- 驗證規則測試 ---")
	validator := NewRules()
	
	testCases := []struct {
		description string
		value       string
		validator   func(string) bool
	}{
		{"郵箱驗證", "test@example.com", validator.IsEmail},
		{"無效郵箱", "invalid-email", validator.IsEmail},
		{"手機號碼", "0912345678", validator.IsPhone},
		{"URL 驗證", "https://www.example.com", validator.IsURL},
		{"強密碼", "MySecure123!", validator.IsStrongPassword},
		{"弱密碼", "123456", validator.IsStrongPassword},
		{"非空字符串", "Hello", validator.IsNotEmpty},
		{"空字符串", "   ", validator.IsNotEmpty},
	}
	
	for _, tc := range testCases {
		result := tc.validator(tc.value)
		fmt.Printf("%s '%s': %t\n", tc.description, tc.value, result)
	}
	
	// 測試數值範圍
	fmt.Printf("數值 75 在 0-100 範圍內: %t\n", validator.IsInRange(75, 0, 100))
	fmt.Printf("數值 150 在 0-100 範圍內: %t\n", validator.IsInRange(150, 0, 100))
	
	fmt.Println("\n=== 測試完成 ===")
}