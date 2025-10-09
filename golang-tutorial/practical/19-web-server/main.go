package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 用戶數據結構
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// 模擬用戶數據庫
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
	{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
}

var nextUserID = 4

// 響應結構
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 首頁處理器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	htmlContent := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web Server Demo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .nav { background: #f4f4f4; padding: 10px; margin-bottom: 20px; }
        .nav a { margin-right: 15px; text-decoration: none; color: #333; }
        .nav a:hover { color: #007bff; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Go Web Server 演示</h1>
        <div class="nav">
            <a href="/">首頁</a>
            <a href="/users">用戶列表</a>
            <a href="/api/users">API 用戶</a>
            <a href="/form">表單演示</a>
            <a href="/static/test.txt">靜態文件</a>
        </div>
        <h2>歡迎來到 Go Web Server!</h2>
        <p>當前時間: %s</p>
        <p>這是一個使用 Go 語言 net/http 包構建的 Web 服務器演示。</p>
        
        <h3>功能演示：</h3>
        <ul>
            <li><a href="/users">HTML 用戶列表</a> - 演示模板渲染</li>
            <li><a href="/api/users">JSON API</a> - 演示 REST API</li>
            <li><a href="/form">表單處理</a> - 演示表單提交</li>
            <li><a href="/upload">文件上傳</a> - 演示文件處理</li>
        </ul>
    </div>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, htmlContent, time.Now().Format("2006-01-02 15:04:05"))
}

// 用戶列表頁面（HTML模板）
func usersPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>用戶列表</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        table { border-collapse: collapse; width: 100%; }
        th, td { border: 1px solid #ddd; padding: 12px; text-align: left; }
        th { background-color: #f2f2f2; }
        .nav { margin-bottom: 20px; }
        .nav a { margin-right: 15px; text-decoration: none; color: #007bff; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">← 返回首頁</a>
    </div>
    <h1>用戶列表</h1>
    <table>
        <tr>
            <th>ID</th>
            <th>姓名</th>
            <th>郵箱</th>
            <th>年齡</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.Name}}</td>
            <td>{{.Email}}</td>
            <td>{{.Age}}</td>
        </tr>
        {{end}}
    </table>
    <p>總計: {{len .}} 個用戶</p>
</body>
</html>`

	t, err := template.New("users").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, users)
}

// REST API - 用戶列表
func apiUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case "GET":
		response := APIResponse{
			Success: true,
			Message: "獲取用戶列表成功",
			Data:    users,
		}
		json.NewEncoder(w).Encode(response)
		
	case "POST":
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := APIResponse{
				Success: false,
				Message: "JSON 解析錯誤: " + err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		// 驗證數據
		if newUser.Name == "" || newUser.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			response := APIResponse{
				Success: false,
				Message: "姓名和郵箱不能為空",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		// 添加新用戶
		newUser.ID = nextUserID
		nextUserID++
		users = append(users, newUser)
		
		w.WriteHeader(http.StatusCreated)
		response := APIResponse{
			Success: true,
			Message: "用戶創建成功",
			Data:    newUser,
		}
		json.NewEncoder(w).Encode(response)
		
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := APIResponse{
			Success: false,
			Message: "方法不允許",
		}
		json.NewEncoder(w).Encode(response)
	}
}

// REST API - 單個用戶操作
func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// 從 URL 路徑中提取用戶 ID
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		response := APIResponse{
			Success: false,
			Message: "無效的 URL 路徑",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	userID, err := strconv.Atoi(pathParts[3])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := APIResponse{
			Success: false,
			Message: "無效的用戶 ID",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// 查找用戶
	userIndex := -1
	var user User
	for i, u := range users {
		if u.ID == userID {
			userIndex = i
			user = u
			break
		}
	}
	
	switch r.Method {
	case "GET":
		if userIndex == -1 {
			w.WriteHeader(http.StatusNotFound)
			response := APIResponse{
				Success: false,
				Message: "用戶不存在",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		response := APIResponse{
			Success: true,
			Message: "獲取用戶成功",
			Data:    user,
		}
		json.NewEncoder(w).Encode(response)
		
	case "PUT":
		if userIndex == -1 {
			w.WriteHeader(http.StatusNotFound)
			response := APIResponse{
				Success: false,
				Message: "用戶不存在",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		var updatedUser User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := APIResponse{
				Success: false,
				Message: "JSON 解析錯誤: " + err.Error(),
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		// 保持原有 ID
		updatedUser.ID = userID
		users[userIndex] = updatedUser
		
		response := APIResponse{
			Success: true,
			Message: "用戶更新成功",
			Data:    updatedUser,
		}
		json.NewEncoder(w).Encode(response)
		
	case "DELETE":
		if userIndex == -1 {
			w.WriteHeader(http.StatusNotFound)
			response := APIResponse{
				Success: false,
				Message: "用戶不存在",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		
		// 刪除用戶
		users = append(users[:userIndex], users[userIndex+1:]...)
		
		response := APIResponse{
			Success: true,
			Message: "用戶刪除成功",
		}
		json.NewEncoder(w).Encode(response)
		
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := APIResponse{
			Success: false,
			Message: "方法不允許",
		}
		json.NewEncoder(w).Encode(response)
	}
}

// 表單處理演示
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 顯示表單
		formHTML := `
<!DOCTYPE html>
<html>
<head>
    <title>表單演示</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .form-group { margin-bottom: 15px; }
        label { display: block; margin-bottom: 5px; }
        input, textarea { width: 100%; padding: 8px; box-sizing: border-box; }
        button { background: #007bff; color: white; padding: 10px 20px; border: none; cursor: pointer; }
        button:hover { background: #0056b3; }
        .nav { margin-bottom: 20px; }
        .nav a { text-decoration: none; color: #007bff; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">← 返回首頁</a>
    </div>
    <h1>表單演示</h1>
    <form method="POST">
        <div class="form-group">
            <label for="name">姓名:</label>
            <input type="text" id="name" name="name" required>
        </div>
        <div class="form-group">
            <label for="email">郵箱:</label>
            <input type="email" id="email" name="email" required>
        </div>
        <div class="form-group">
            <label for="age">年齡:</label>
            <input type="number" id="age" name="age" required>
        </div>
        <div class="form-group">
            <label for="message">留言:</label>
            <textarea id="message" name="message" rows="4"></textarea>
        </div>
        <button type="submit">提交</button>
    </form>
</body>
</html>`
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, formHTML)
		
	} else if r.Method == "POST" {
		// 處理表單提交
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "解析表單錯誤", http.StatusBadRequest)
			return
		}
		
		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age")
		message := r.FormValue("message")
		
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "年齡格式無效", http.StatusBadRequest)
			return
		}
		
		// 顯示提交結果
		resultHTML := `
<!DOCTYPE html>
<html>
<head>
    <title>表單提交結果</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .result { background: #d4edda; border: 1px solid #c3e6cb; padding: 15px; border-radius: 5px; }
        .nav { margin-bottom: 20px; }
        .nav a { text-decoration: none; color: #007bff; margin-right: 15px; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">← 返回首頁</a>
        <a href="/form">← 返回表單</a>
    </div>
    <h1>表單提交成功！</h1>
    <div class="result">
        <h3>提交的信息：</h3>
        <p><strong>姓名:</strong> %s</p>
        <p><strong>郵箱:</strong> %s</p>
        <p><strong>年齡:</strong> %d</p>
        <p><strong>留言:</strong> %s</p>
    </div>
</body>
</html>`
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, resultHTML, name, email, age, message)
	}
}

// 文件上傳處理
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		uploadHTML := `
<!DOCTYPE html>
<html>
<head>
    <title>文件上傳</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .form-group { margin-bottom: 15px; }
        label { display: block; margin-bottom: 5px; }
        input { padding: 8px; }
        button { background: #007bff; color: white; padding: 10px 20px; border: none; cursor: pointer; }
        .nav a { text-decoration: none; color: #007bff; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">← 返回首頁</a>
    </div>
    <h1>文件上傳演示</h1>
    <form method="POST" enctype="multipart/form-data">
        <div class="form-group">
            <label for="file">選擇文件:</label>
            <input type="file" id="file" name="file" required>
        </div>
        <div class="form-group">
            <label for="description">文件描述:</label>
            <input type="text" id="description" name="description">
        </div>
        <button type="submit">上傳</button>
    </form>
</body>
</html>`
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, uploadHTML)
		
	} else if r.Method == "POST" {
		// 解析 multipart 表單
		err := r.ParseMultipartForm(10 << 20) // 10 MB 限制
		if err != nil {
			http.Error(w, "解析表單錯誤", http.StatusBadRequest)
			return
		}
		
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "獲取文件錯誤", http.StatusBadRequest)
			return
		}
		defer file.Close()
		
		description := r.FormValue("description")
		
		// 讀取文件內容（實際應用中應該保存到磁盤）
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "讀取文件錯誤", http.StatusInternalServerError)
			return
		}
		
		resultHTML := `
<!DOCTYPE html>
<html>
<head>
    <title>上傳成功</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .result { background: #d4edda; border: 1px solid #c3e6cb; padding: 15px; border-radius: 5px; }
        .nav a { text-decoration: none; color: #007bff; margin-right: 15px; }
    </style>
</head>
<body>
    <div class="nav">
        <a href="/">← 返回首頁</a>
        <a href="/upload">← 重新上傳</a>
    </div>
    <h1>文件上傳成功！</h1>
    <div class="result">
        <h3>文件信息：</h3>
        <p><strong>文件名:</strong> %s</p>
        <p><strong>文件大小:</strong> %d 字節</p>
        <p><strong>描述:</strong> %s</p>
        <p><strong>Content-Type:</strong> %s</p>
    </div>
</body>
</html>`
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, resultHTML, handler.Filename, len(fileBytes), description, handler.Header.Get("Content-Type"))
	}
}

// 日誌中間件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// 記錄請求
		log.Printf("開始 %s %s", r.Method, r.URL.Path)
		
		// 執行下一個處理器
		next.ServeHTTP(w, r)
		
		// 記錄完成
		duration := time.Since(start)
		log.Printf("完成 %s %s - 耗時: %v", r.Method, r.URL.Path, duration)
	})
}

// CORS 中間件
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 創建路由
	mux := http.NewServeMux()
	
	// 靜態文件服務
	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// 路由註冊
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", usersPageHandler)
	mux.HandleFunc("/api/users", apiUsersHandler)
	mux.HandleFunc("/api/users/", apiUserHandler)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/upload", uploadHandler)
	
	// 應用中間件
	handler := corsMiddleware(loggingMiddleware(mux))
	
	// 創建伺服器
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// 創建靜態文件目錄（如果不存在）
	err := createStaticDir()
	if err != nil {
		log.Printf("創建靜態目錄錯誤: %v", err)
	}
	
	fmt.Println("Web 伺服器啟動在 http://localhost:8080")
	fmt.Println("按 Ctrl+C 停止服務器")
	
	log.Fatal(server.ListenAndServe())
}

// 創建靜態文件目錄和測試文件
func createStaticDir() error {
	// 創建目錄
	err := os.MkdirAll("static", 0755)
	if err != nil {
		return err
	}
	
	// 創建測試文件
	testContent := "這是一個靜態文件測試內容。\n創建時間: " + time.Now().Format("2006-01-02 15:04:05")
	return os.WriteFile("static/test.txt", []byte(testContent), 0644)
}