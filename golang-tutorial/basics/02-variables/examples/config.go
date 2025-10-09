// 配置管理示例
package main

import (
	"fmt"
	"time"
)

// 應用配置常數
const (
	// 應用基本信息
	AppName        = "MyWebApp"
	AppVersion     = "v1.2.3"
	AppDescription = "一個用 Go 編寫的 Web 應用程序"
)

// 服務器配置
const (
	DefaultHost = "localhost"
	DefaultPort = 8080
	
	// 超時配置
	ReadTimeout  = 30 * time.Second
	WriteTimeout = 30 * time.Second
	IdleTimeout  = 120 * time.Second
)

// 數據庫配置
const (
	MaxConnections    = 100
	MaxIdleConnections = 10
	ConnMaxLifetime   = 5 * time.Minute
)

// 日誌級別枚舉
const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

// HTTP 狀態碼
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

// 緩存配置
const (
	CacheExpiration = 1 * time.Hour
	CacheCleanup    = 10 * time.Minute
)

// 文件大小限制
const (
	_  = iota
	KB = 1 << (10 * iota) // 1024
	MB                    // 1024 * 1024
	GB                    // 1024 * 1024 * 1024
)

const (
	MaxFileSize    = 10 * MB  // 10MB
	MaxRequestSize = 32 * MB  // 32MB
)

// 配置結構
type Config struct {
	// 服務器配置
	Host string
	Port int
	
	// 數據庫配置
	DatabaseURL string
	
	// 功能開關
	EnableCache   bool
	EnableLogging bool
	EnableMetrics bool
	
	// 限制配置
	MaxUsers       int
	MaxConnections int
}

func main() {
	fmt.Println("=== 應用配置管理示例 ===")
	
	// 創建配置實例
	config := createConfig()
	
	// 顯示應用信息
	showAppInfo()
	
	// 顯示配置信息
	showConfig(config)
	
	// 顯示枚舉使用
	showLogLevels()
	
	// 顯示文件大小計算
	showFileSizes()
}

func createConfig() Config {
	// 使用常數初始化配置
	return Config{
		Host:           DefaultHost,
		Port:           DefaultPort,
		DatabaseURL:    "postgres://localhost:5432/myapp",
		EnableCache:    true,
		EnableLogging:  true,
		EnableMetrics:  false,
		MaxUsers:       1000,
		MaxConnections: MaxConnections,
	}
}

func showAppInfo() {
	fmt.Println("\n--- 應用信息 ---")
	fmt.Printf("應用名稱: %s\n", AppName)
	fmt.Printf("版本: %s\n", AppVersion)
	fmt.Printf("描述: %s\n", AppDescription)
}

func showConfig(config Config) {
	fmt.Println("\n--- 服務器配置 ---")
	fmt.Printf("主機: %s\n", config.Host)
	fmt.Printf("端口: %d\n", config.Port)
	fmt.Printf("讀取超時: %v\n", ReadTimeout)
	fmt.Printf("寫入超時: %v\n", WriteTimeout)
	fmt.Printf("空閒超時: %v\n", IdleTimeout)
	
	fmt.Println("\n--- 數據庫配置 ---")
	fmt.Printf("數據庫 URL: %s\n", config.DatabaseURL)
	fmt.Printf("最大連接數: %d\n", config.MaxConnections)
	fmt.Printf("最大空閒連接數: %d\n", MaxIdleConnections)
	fmt.Printf("連接最大生命週期: %v\n", ConnMaxLifetime)
	
	fmt.Println("\n--- 功能開關 ---")
	fmt.Printf("啟用緩存: %t\n", config.EnableCache)
	fmt.Printf("啟用日誌: %t\n", config.EnableLogging)
	fmt.Printf("啟用指標: %t\n", config.EnableMetrics)
	
	fmt.Println("\n--- 限制配置 ---")
	fmt.Printf("最大用戶數: %d\n", config.MaxUsers)
	fmt.Printf("最大文件大小: %d MB\n", MaxFileSize/MB)
	fmt.Printf("最大請求大小: %d MB\n", MaxRequestSize/MB)
}

func showLogLevels() {
	fmt.Println("\n--- 日誌級別 ---")
	logLevels := map[int]string{
		LogLevelDebug:   "DEBUG",
		LogLevelInfo:    "INFO",
		LogLevelWarning: "WARNING",
		LogLevelError:   "ERROR",
		LogLevelFatal:   "FATAL",
	}
	
	for level, name := range logLevels {
		fmt.Printf("級別 %d: %s\n", level, name)
	}
}

func showFileSizes() {
	fmt.Println("\n--- 文件大小計算 ---")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	
	// 計算實際文件大小
	fileSize := 2.5 * float64(MB)
	fmt.Printf("2.5 MB = %.0f bytes\n", fileSize)
}