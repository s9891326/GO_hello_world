// 演示如何創建自定義包的完整示例
package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

// ============ 數學工具包 ============

// MathUtils 提供數學計算功能
type MathUtils struct{}

// NewMathUtils 創建數學工具實例
func NewMathUtils() *MathUtils {
	return &MathUtils{}
}

// Max 返回兩個數中的較大者
func (m *MathUtils) Max(a, b float64) float64 {
	return math.Max(a, b)
}

// Min 返回兩個數中的較小者
func (m *MathUtils) Min(a, b float64) float64 {
	return math.Min(a, b)
}

// Average 計算數組的平均值
func (m *MathUtils) Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Median 計算數組的中位數
func (m *MathUtils) Median(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	// 創建副本並排序
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)
	
	length := len(sorted)
	if length%2 == 0 {
		// 偶數個元素，返回中間兩個數的平均值
		return (sorted[length/2-1] + sorted[length/2]) / 2
	} else {
		// 奇數個元素，返回中間的數
		return sorted[length/2]
	}
}

// StandardDeviation 計算標準差
func (m *MathUtils) StandardDeviation(numbers []float64) float64 {
	if len(numbers) <= 1 {
		return 0
	}
	
	avg := m.Average(numbers)
	sumSquares := 0.0
	
	for _, num := range numbers {
		diff := num - avg
		sumSquares += diff * diff
	}
	
	variance := sumSquares / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

// ============ 字符串工具包 ============

// StringUtils 提供字符串處理功能
type StringUtils struct{}

// NewStringUtils 創建字符串工具實例
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// Reverse 反轉字符串
func (s *StringUtils) Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 檢查是否為回文
func (s *StringUtils) IsPalindrome(str string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(str, " ", ""))
	return cleaned == s.Reverse(cleaned)
}

// WordCount 計算單詞數量
func (s *StringUtils) WordCount(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	wordCount := make(map[string]int)
	
	for _, word := range words {
		// 移除標點符號
		word = strings.Trim(word, ".,!?;:")
		if word != "" {
			wordCount[word]++
		}
	}
	
	return wordCount
}

// Capitalize 將每個單詞的首字母大寫
func (s *StringUtils) Capitalize(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// Truncate 截斷字符串到指定長度
func (s *StringUtils) Truncate(text string, maxLength int, suffix string) string {
	if len(text) <= maxLength {
		return text
	}
	
	if len(suffix) >= maxLength {
		return suffix[:maxLength]
	}
	
	return text[:maxLength-len(suffix)] + suffix
}

// ============ 時間工具包 ============

// TimeUtils 提供時間處理功能
type TimeUtils struct{}

// NewTimeUtils 創建時間工具實例
func NewTimeUtils() *TimeUtils {
	return &TimeUtils{}
}

// FormatDuration 格式化時間間隔
func (t *TimeUtils) FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.0f 秒", d.Seconds())
	} else if d < time.Hour {
		return fmt.Sprintf("%.0f 分鐘", d.Minutes())
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.1f 小時", d.Hours())
	} else {
		return fmt.Sprintf("%.1f 天", d.Hours()/24)
	}
}

// IsWeekend 檢查是否為週末
func (t *TimeUtils) IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// NextWorkday 獲取下一個工作日
func (t *TimeUtils) NextWorkday(date time.Time) time.Time {
	next := date.AddDate(0, 0, 1)
	for t.IsWeekend(next) {
		next = next.AddDate(0, 0, 1)
	}
	return next
}

// Age 計算年齡
func (t *TimeUtils) Age(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	
	// 如果還沒到生日，年齡減一
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	
	return age
}

// ============ 日誌工具包 ============

// LogLevel 日誌級別
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String 返回日誌級別的字符串表示
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger 自定義日誌記錄器
type Logger struct {
	level  LogLevel
	prefix string
}

// NewLogger 創建新的日誌記錄器
func NewLogger(level LogLevel, prefix string) *Logger {
	return &Logger{
		level:  level,
		prefix: prefix,
	}
}

// log 內部日誌方法
func (l *Logger) log(level LogLevel, message string) {
	if level >= l.level {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] %s [%s] %s\n", timestamp, level.String(), l.prefix, message)
	}
}

// Debug 記錄調試信息
func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

// Info 記錄一般信息
func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

// Warn 記錄警告信息
func (l *Logger) Warn(message string) {
	l.log(WARN, message)
}

// Error 記錄錯誤信息
func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

// Fatal 記錄致命錯誤並退出
func (l *Logger) Fatal(message string) {
	l.log(FATAL, message)
	log.Fatal("程序終止")
}

// ============ 配置管理包 ============

// Config 應用配置
type Config struct {
	AppName     string        `json:"app_name"`
	Version     string        `json:"version"`
	Port        int           `json:"port"`
	Debug       bool          `json:"debug"`
	Timeout     time.Duration `json:"timeout"`
	DatabaseURL string        `json:"database_url"`
}

// ConfigManager 配置管理器
type ConfigManager struct {
	config *Config
}

// NewConfigManager 創建配置管理器
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config: &Config{
			AppName:     "MyApp",
			Version:     "1.0.0",
			Port:        8080,
			Debug:       false,
			Timeout:     30 * time.Second,
			DatabaseURL: "localhost:5432",
		},
	}
}

// GetConfig 獲取配置
func (c *ConfigManager) GetConfig() *Config {
	return c.config
}

// SetPort 設置端口
func (c *ConfigManager) SetPort(port int) {
	c.config.Port = port
}

// SetDebug 設置調試模式
func (c *ConfigManager) SetDebug(debug bool) {
	c.config.Debug = debug
}

// Validate 驗證配置
func (c *ConfigManager) Validate() error {
	if c.config.Port <= 0 || c.config.Port > 65535 {
		return fmt.Errorf("無效的端口號: %d", c.config.Port)
	}
	
	if c.config.AppName == "" {
		return fmt.Errorf("應用名稱不能為空")
	}
	
	return nil
}

// ============ 主程序 ============

func main() {
	fmt.Println("=== 自定義包使用示例 ===")
	
	// 數學工具示例
	fmt.Println("\n--- 數學工具 ---")
	mathUtils := NewMathUtils()
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	fmt.Printf("數據: %v\n", numbers)
	fmt.Printf("平均值: %.2f\n", mathUtils.Average(numbers))
	fmt.Printf("中位數: %.2f\n", mathUtils.Median(numbers))
	fmt.Printf("標準差: %.2f\n", mathUtils.StandardDeviation(numbers))
	fmt.Printf("最大值: %.2f\n", mathUtils.Max(10, 5))
	fmt.Printf("最小值: %.2f\n", mathUtils.Min(10, 5))
	
	// 字符串工具示例
	fmt.Println("\n--- 字符串工具 ---")
	stringUtils := NewStringUtils()
	text := "Hello World! This is a test text."
	
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("反轉: %s\n", stringUtils.Reverse(text))
	fmt.Printf("首字母大寫: %s\n", stringUtils.Capitalize(text))
	fmt.Printf("截斷 (15字符): %s\n", stringUtils.Truncate(text, 15, "..."))
	fmt.Printf("是否回文 'madam': %t\n", stringUtils.IsPalindrome("madam"))
	
	wordCount := stringUtils.WordCount(text)
	fmt.Println("單詞統計:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}
	
	// 時間工具示例
	fmt.Println("\n--- 時間工具 ---")
	timeUtils := NewTimeUtils()
	now := time.Now()
	
	fmt.Printf("當前時間: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("是否週末: %t\n", timeUtils.IsWeekend(now))
	fmt.Printf("下一個工作日: %s\n", timeUtils.NextWorkday(now).Format("2006-01-02"))
	
	// 生日示例
	birthday := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	fmt.Printf("年齡 (生日 1990-05-15): %d 歲\n", timeUtils.Age(birthday))
	
	// 時間間隔格式化
	durations := []time.Duration{
		30 * time.Second,
		5 * time.Minute,
		2 * time.Hour,
		3 * 24 * time.Hour,
	}
	
	fmt.Println("時間間隔格式化:")
	for _, d := range durations {
		fmt.Printf("  %v = %s\n", d, timeUtils.FormatDuration(d))
	}
	
	// 日誌工具示例
	fmt.Println("\n--- 日誌工具 ---")
	logger := NewLogger(INFO, "APP")
	
	logger.Debug("這是調試信息 (不會顯示)")
	logger.Info("應用程序啟動")
	logger.Warn("這是警告信息")
	logger.Error("這是錯誤信息")
	
	// 配置管理示例
	fmt.Println("\n--- 配置管理 ---")
	configManager := NewConfigManager()
	config := configManager.GetConfig()
	
	fmt.Printf("應用配置: %+v\n", config)
	
	// 修改配置
	configManager.SetPort(9090)
	configManager.SetDebug(true)
	
	if err := configManager.Validate(); err != nil {
		logger.Error(fmt.Sprintf("配置驗證失敗: %v", err))
	} else {
		logger.Info("配置驗證成功")
		fmt.Printf("更新後配置: %+v\n", configManager.GetConfig())
	}
	
	fmt.Println("\n=== 示例完成 ===")
}