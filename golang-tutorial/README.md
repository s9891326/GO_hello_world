# Golang 完整語法教學專案

這是一個從零開始的 Golang 語法教學專案，涵蓋了從基礎到進階的所有重要概念，並提供實際的工作應用範例。

## 📚 目錄結構

```
golang-tutorial/
├── README.md              # 本文件
├── go.mod                 # Go 模組配置
├── basics/                # 基礎語法
│   ├── 01-hello-world/    # Hello World 和環境設置
│   ├── 02-variables/      # 變數和常數
│   ├── 03-data-types/     # 數據類型
│   ├── 04-operators/      # 運算符
│   └── 05-control-flow/   # 流程控制
├── intermediate/          # 中級概念
│   ├── 06-functions/      # 函數
│   ├── 07-structs/        # 結構體
│   ├── 08-pointers/       # 指針
│   ├── 09-arrays-slices/  # 數組和切片
│   └── 10-maps/           # 映射
├── advanced/              # 進階概念
│   ├── 11-interfaces/     # 接口
│   ├── 12-goroutines/     # 協程
│   ├── 13-channels/       # 通道
│   ├── 14-error-handling/ # 錯誤處理
│   └── 15-packages/       # 包管理
├── practical/             # 實際應用
│   ├── 16-file-io/        # 文件操作
│   ├── 17-json-handling/  # JSON 處理
│   ├── 18-http-client/    # HTTP 客戶端
│   ├── 19-web-server/     # Web 伺服器
│   └── 20-database/       # 數據庫操作
└── projects/              # 實戰專案
    ├── cli-tool/          # 命令行工具
    ├── web-api/           # Web API
    └── microservice/      # 微服務
```

## 🚀 快速開始

### 1. 環境要求
- Go 1.21+
- Git
- VS Code (推薦) 或其他編輯器

### 2. 安裝 Go
```bash
# macOS (使用 Homebrew)
brew install go

# Windows (下載安裝包)
# 從 https://golang.org/dl/ 下載

# Linux (Ubuntu/Debian)
sudo apt-get install golang-go
```

### 3. 驗證安裝
```bash
go version
```

### 4. 初始化專案
```bash
cd golang-tutorial
go mod init golang-tutorial
```

## 📖 學習路徑

### 第一階段：基礎語法 (1-2 週)
1. **Hello World** - 環境設置和第一個程序
2. **變數和常數** - 數據存儲基礎
3. **數據類型** - 內建類型和類型轉換
4. **運算符** - 算術、比較、邏輯運算
5. **流程控制** - if/else、for、switch

### 第二階段：中級概念 (2-3 週)
6. **函數** - 函數定義、參數、返回值
7. **結構體** - 自定義數據類型
8. **指針** - 內存地址和引用
9. **數組和切片** - 集合數據處理
10. **映射** - 鍵值對數據結構

### 第三階段：進階概念 (3-4 週)
11. **接口** - 多態和抽象
12. **協程** - 併發程式設計
13. **通道** - 協程間通信
14. **錯誤處理** - 異常和錯誤管理
15. **包管理** - 模組化開發

### 第四階段：實際應用 (2-3 週)
16. **文件操作** - 讀寫文件
17. **JSON 處理** - 數據序列化
18. **HTTP 客戶端** - API 調用
19. **Web 伺服器** - 建立 REST API
20. **數據庫操作** - CRUD 操作

### 第五階段：實戰專案 (3-4 週)
21. **命令行工具** - CLI 應用開發
22. **Web API** - 完整的 REST API
23. **微服務** - 分布式系統基礎

## 🎯 學習目標

完成本教學後，您將能夠：

- ✅ 熟練掌握 Go 語言語法和特性
- ✅ 理解 Go 的併發程式設計模型
- ✅ 開發實際的業務應用程序
- ✅ 建立和維護微服務架構
- ✅ 處理常見的工程問題
- ✅ 在團隊中有效協作開發

## 📝 每章結構

每個章節都包含：

1. **README.md** - 概念說明和學習目標
2. **main.go** - 基礎範例代碼
3. **examples/** - 多個實際應用範例
4. **exercises/** - 練習題目
5. **solutions/** - 練習解答

## 🛠️ 實用工具和資源

### 開發工具
- **VS Code** + Go 擴展
- **GoLand** (JetBrains IDE)
- **Vim** + vim-go

### 實用網站
- [Go 官方文檔](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)

### 推薦書籍
- "The Go Programming Language" - Alan Donovan & Brian Kernighan
- "Go in Action" - William Kennedy
- "Concurrency in Go" - Katherine Cox-Buday

## 🤝 如何使用本教學

1. **按順序學習** - 從基礎開始，逐步深入
2. **動手實作** - 每個範例都要親自執行
3. **完成練習** - 練習是學習的關鍵
4. **實作專案** - 將知識應用到實際專案中
5. **參與社群** - 加入 Go 開發者社群

## 📞 支援和反饋

如果您在學習過程中遇到問題或有建議，歡迎：
- 提交 Issue
- 發送 Pull Request
- 參與討論

## 📄 授權

本專案採用 MIT 授權條款。

---

**準備好開始您的 Go 語言學習之旅了嗎？讓我們從 [第一章：Hello World](./basics/01-hello-world/) 開始！**