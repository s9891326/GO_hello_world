// Go Modules 使用示例
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 演示 Go Modules 的使用和管理
func main() {
	fmt.Println("=== Go Modules 演示 ===")
	
	// 展示如何創建和管理 Go 模組
	demonstrateModuleCreation()
	demonstrateDependencyManagement()
	demonstrateVersioning()
	demonstrateWorkspaces()
}

// 演示模組創建
func demonstrateModuleCreation() {
	fmt.Println("\n--- 模組創建演示 ---")
	
	fmt.Println("1. 初始化新模組:")
	fmt.Println("   go mod init github.com/username/myproject")
	
	fmt.Println("\n2. 這會創建 go.mod 文件:")
	sampleGoMod := `module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/spf13/viper v1.16.0
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... 其他間接依賴
)

replace github.com/old/package => github.com/new/package v1.0.0

exclude github.com/problematic/package v1.2.3`

	fmt.Println(sampleGoMod)
}

// 演示依賴管理
func demonstrateDependencyManagement() {
	fmt.Println("\n--- 依賴管理演示 ---")
	
	commands := []struct {
		command     string
		description string
	}{
		{"go get github.com/gin-gonic/gin", "添加新依賴"},
		{"go get github.com/gin-gonic/gin@v1.9.1", "添加特定版本"},
		{"go get github.com/gin-gonic/gin@latest", "獲取最新版本"},
		{"go get -u github.com/gin-gonic/gin", "升級到最新版本"},
		{"go get -u ./...", "升級所有依賴"},
		{"go mod tidy", "清理未使用的依賴"},
		{"go mod download", "下載依賴到本地緩存"},
		{"go mod verify", "驗證依賴完整性"},
		{"go mod graph", "顯示依賴圖"},
		{"go mod why github.com/gin-gonic/gin", "解釋為什麼需要這個依賴"},
	}
	
	fmt.Println("常用依賴管理命令:")
	for _, cmd := range commands {
		fmt.Printf("  %-35s # %s\n", cmd.command, cmd.description)
	}
	
	// 演示 go.sum 文件的作用
	fmt.Println("\n3. go.sum 文件示例:")
	sampleGoSum := `github.com/gin-gonic/gin v1.9.1 h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=
github.com/gin-gonic/gin v1.9.1/go.mod h1:hPrL7YrpYKXt5YId3A/Tnip5kqbEAP+KLuI3SUcPTeU=
github.com/spf13/viper v1.16.0 h1:rGGH0XDZhdUOryiDWjmIvUSWpbNqisK8Wk0Vyefw8hc=
github.com/spf13/viper v1.16.0/go.mod h1:yg78JgCJcbrQOvV9YLXgkLaZqUidkY9K+Dd1FofRzQg=`
	fmt.Println(sampleGoSum)
	fmt.Println("\ngo.sum 文件用於:")
	fmt.Println("  - 確保依賴的完整性")
	fmt.Println("  - 防止依賴被篡改")
	fmt.Println("  - 確保構建的可重現性")
}

// 演示版本管理
func demonstrateVersioning() {
	fmt.Println("\n--- 版本管理演示 ---")
	
	fmt.Println("語義化版本控制 (SemVer):")
	fmt.Println("  格式: v主版本.次版本.修復版本")
	fmt.Println("  示例: v1.2.3")
	fmt.Println("    - 主版本: 不兼容的 API 變更")
	fmt.Println("    - 次版本: 向後兼容的功能新增")
	fmt.Println("    - 修復版本: 向後兼容的錯誤修復")
	
	fmt.Println("\nGo 模組版本約束:")
	constraints := []struct {
		constraint  string
		description string
	}{
		{"v1.2.3", "精確版本"},
		{">= v1.2.0", "最小版本"},
		{"< v2.0.0", "最大版本（不包含）"},
		{"~v1.2.3", "相容版本（~1.2.3 允許 >=1.2.3 且 <1.3.0）"},
		{"^v1.2.3", "兼容版本（^1.2.3 允許 >=1.2.3 且 <2.0.0）"},
		{"latest", "最新穩定版本"},
		{"main", "主分支的最新提交"},
	}
	
	for _, c := range constraints {
		fmt.Printf("  %-12s # %s\n", c.constraint, c.description)
	}
	
	fmt.Println("\n發布版本的步驟:")
	steps := []string{
		"git tag v1.0.0",
		"git push origin v1.0.0",
		"go list -m github.com/username/myproject@v1.0.0",
	}
	
	for i, step := range steps {
		fmt.Printf("  %d. %s\n", i+1, step)
	}
}

// 演示工作區 (Workspaces)
func demonstrateWorkspaces() {
	fmt.Println("\n--- Go 工作區演示 ---")
	
	fmt.Println("工作區允許同時開發多個相關模組:")
	
	workspaceStructure := `myworkspace/
├── go.work              # 工作區配置文件
├── moduleA/
│   ├── go.mod
│   └── main.go
├── moduleB/
│   ├── go.mod
│   └── lib.go
└── shared/
    ├── go.mod
    └── utils.go`
	
	fmt.Println(workspaceStructure)
	
	fmt.Println("\ngo.work 文件示例:")
	sampleGoWork := `go 1.21

use (
    ./moduleA
    ./moduleB
    ./shared
)

replace github.com/username/shared => ./shared`
	
	fmt.Println(sampleGoWork)
	
	fmt.Println("\n工作區命令:")
	workspaceCommands := []struct {
		command     string
		description string
	}{
		{"go work init ./moduleA ./moduleB", "初始化工作區"},
		{"go work use ./shared", "添加模組到工作區"},
		{"go work edit -dropuse ./moduleA", "從工作區移除模組"},
		{"go work sync", "同步工作區依賴"},
	}
	
	for _, cmd := range workspaceCommands {
		fmt.Printf("  %-30s # %s\n", cmd.command, cmd.description)
	}
}

// 演示模組代理和私有倉庫
func demonstrateModuleProxy() {
	fmt.Println("\n--- 模組代理演示 ---")
	
	fmt.Println("Go 模組代理配置:")
	fmt.Println("  GOPROXY=https://proxy.golang.org,direct")
	fmt.Println("  GOSUMDB=sum.golang.org")
	fmt.Println("  GOPRIVATE=git.company.com")
	fmt.Println("  GONOPROXY=git.company.com")
	fmt.Println("  GONOSUMDB=git.company.com")
	
	fmt.Println("\n私有倉庫配置:")
	fmt.Println("  1. 設置 GOPRIVATE 環境變數")
	fmt.Println("  2. 配置 Git 認證")
	fmt.Println("  3. 使用 replace 指令進行本地開發")
}

// 演示常見問題和解決方案
func demonstrateCommonIssues() {
	fmt.Println("\n--- 常見問題和解決方案 ---")
	
	issues := []struct {
		problem  string
		solution string
	}{
		{
			"依賴版本衝突",
			"使用 go mod tidy 清理，或手動解決版本約束",
		},
		{
			"無法下載私有倉庫",
			"配置 GOPRIVATE 和 Git 認證",
		},
		{
			"構建緩慢",
			"使用模組代理，配置 GOPROXY",
		},
		{
			"依賴校驗失敗",
			"檢查 go.sum 文件，使用 go mod verify",
		},
		{
			"循環依賴",
			"重新設計包結構，提取公共接口",
		},
	}
	
	for _, issue := range issues {
		fmt.Printf("問題: %s\n", issue.problem)
		fmt.Printf("解決: %s\n\n", issue.solution)
	}
}

// 輔助函數：獲取當前目錄
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("獲取當前目錄失敗: %v", err)
		return "unknown"
	}
	return filepath.Base(dir)
}