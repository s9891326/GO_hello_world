# 第二十章：數據庫操作

## 🎯 學習目標

- 掌握 Go 語言數據庫操作基礎
- 理解 database/sql 包的使用
- 學會使用不同數據庫驅動
- 掌握 CRUD 操作和事務處理
- 了解連接池和性能優化
- 學會使用 ORM 框架
- 掌握數據庫遷移和版本控制

## 🗄️ 數據庫操作概述

Go 語言通過 `database/sql` 包提供了統一的數據庫接口。

### 核心組件

```
數據庫操作架構：
┌─────────────────────────────────────┐
│ database/sql                         │
├─────────────────────────────────────┤
│ • DB (數據庫連接池)                  │
│ • Tx (事務)                         │
│ • Stmt (預編譯語句)                  │
│ • Rows (查詢結果)                    │
└─────────────────────────────────────┘
         ↓
┌─────────────────────────────────────┐
│ Driver (驅動程序)                    │
├─────────────────────────────────────┤
│ • MySQL: github.com/go-sql-driver/mysql │
│ • PostgreSQL: github.com/lib/pq     │
│ • SQLite: github.com/mattn/go-sqlite3   │
│ • SQL Server: github.com/denisenkom/go-mssqldb │
└─────────────────────────────────────┘
```

## 🔌 數據庫連接

### 1. MySQL 連接

```go
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func connectMySQL() (*sql.DB, error) {
    // 數據源名稱 (DSN)
    dsn := "username:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("打開數據庫錯誤: %w", err)
    }
    
    // 驗證連接
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("連接數據庫錯誤: %w", err)
    }
    
    // 設置連接池參數
    db.SetMaxOpenConns(25)                 // 最大打開連接數
    db.SetMaxIdleConns(10)                 // 最大空閒連接數
    db.SetConnMaxLifetime(5 * time.Minute) // 連接最大生存時間
    
    return db, nil
}
```

### 2. PostgreSQL 連接

```go
import (
    _ "github.com/lib/pq"
)

func connectPostgreSQL() (*sql.DB, error) {
    dsn := "host=localhost port=5432 user=username password=password dbname=database_name sslmode=disable"
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    return db, db.Ping()
}
```

### 3. SQLite 連接

```go
import (
    _ "github.com/mattn/go-sqlite3"
)

func connectSQLite() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./database.db")
    if err != nil {
        return nil, err
    }
    
    return db, db.Ping()
}
```

## 📊 基本 CRUD 操作

### 1. 創建表結構

```go
func createTables(db *sql.DB) error {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        age INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`
    
    _, err := db.Exec(createUsersTable)
    if err != nil {
        return fmt.Errorf("創建用戶表錯誤: %w", err)
    }
    
    createPostsTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT,
        title VARCHAR(200) NOT NULL,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    )`
    
    _, err = db.Exec(createPostsTable)
    if err != nil {
        return fmt.Errorf("創建文章表錯誤: %w", err)
    }
    
    return nil
}
```

### 2. 插入數據 (Create)

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// 插入單個用戶
func createUser(db *sql.DB, user User) (int64, error) {
    query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`
    
    result, err := db.Exec(query, user.Name, user.Email, user.Age)
    if err != nil {
        return 0, fmt.Errorf("插入用戶錯誤: %w", err)
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("獲取插入 ID 錯誤: %w", err)
    }
    
    return id, nil
}

// 批量插入用戶
func createUsersBatch(db *sql.DB, users []User) error {
    // 準備批量插入語句
    query := `INSERT INTO users (name, email, age) VALUES `
    values := []interface{}{}
    
    for i, user := range users {
        if i > 0 {
            query += ", "
        }
        query += "(?, ?, ?)"
        values = append(values, user.Name, user.Email, user.Age)
    }
    
    _, err := db.Exec(query, values...)
    if err != nil {
        return fmt.Errorf("批量插入用戶錯誤: %w", err)
    }
    
    return nil
}

// 使用預編譯語句插入
func createUserWithStmt(db *sql.DB, user User) (int64, error) {
    stmt, err := db.Prepare(`INSERT INTO users (name, email, age) VALUES (?, ?, ?)`)
    if err != nil {
        return 0, fmt.Errorf("準備語句錯誤: %w", err)
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(user.Name, user.Email, user.Age)
    if err != nil {
        return 0, fmt.Errorf("執行語句錯誤: %w", err)
    }
    
    return result.LastInsertId()
}
```

### 3. 查詢數據 (Read)

```go
// 查詢單個用戶
func getUserByID(db *sql.DB, id int) (*User, error) {
    query := `SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?`
    
    var user User
    err := db.QueryRow(query, id).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.Age,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("用戶不存在")
        }
        return nil, fmt.Errorf("查詢用戶錯誤: %w", err)
    }
    
    return &user, nil
}

// 查詢多個用戶
func getUsers(db *sql.DB, limit, offset int) ([]User, error) {
    query := `SELECT id, name, email, age, created_at, updated_at 
              FROM users 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`
    
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("查詢用戶列表錯誤: %w", err)
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
            &user.Age,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("掃描用戶數據錯誤: %w", err)
        }
        users = append(users, user)
    }
    
    // 檢查迭代錯誤
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("遍歷結果錯誤: %w", err)
    }
    
    return users, nil
}

// 條件查詢
func searchUsers(db *sql.DB, name string, minAge int) ([]User, error) {
    query := `SELECT id, name, email, age, created_at, updated_at 
              FROM users 
              WHERE name LIKE ? AND age >= ?
              ORDER BY name`
    
    rows, err := db.Query(query, "%"+name+"%", minAge)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
        users = append(users, user)
    }
    
    return users, rows.Err()
}

// 聚合查詢
func getUserStats(db *sql.DB) (map[string]interface{}, error) {
    query := `SELECT 
                COUNT(*) as total_users,
                AVG(age) as average_age,
                MIN(age) as min_age,
                MAX(age) as max_age
              FROM users`
    
    var stats map[string]interface{} = make(map[string]interface{})
    
    err := db.QueryRow(query).Scan(
        &stats["total_users"],
        &stats["average_age"],
        &stats["min_age"],
        &stats["max_age"],
    )
    
    if err != nil {
        return nil, fmt.Errorf("查詢統計錯誤: %w", err)
    }
    
    return stats, nil
}
```

### 4. 更新數據 (Update)

```go
// 更新用戶信息
func updateUser(db *sql.DB, id int, user User) error {
    query := `UPDATE users SET name = ?, email = ?, age = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
    
    result, err := db.Exec(query, user.Name, user.Email, user.Age, id)
    if err != nil {
        return fmt.Errorf("更新用戶錯誤: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("獲取影響行數錯誤: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("用戶不存在或沒有變化")
    }
    
    return nil
}

// 部分更新
func updateUserPartial(db *sql.DB, id int, updates map[string]interface{}) error {
    if len(updates) == 0 {
        return fmt.Errorf("沒有要更新的字段")
    }
    
    // 動態構建 SQL 語句
    setParts := []string{}
    values := []interface{}{}
    
    for field, value := range updates {
        setParts = append(setParts, field+" = ?")
        values = append(values, value)
    }
    
    query := fmt.Sprintf("UPDATE users SET %s, updated_at = CURRENT_TIMESTAMP WHERE id = ?", 
        strings.Join(setParts, ", "))
    values = append(values, id)
    
    result, err := db.Exec(query, values...)
    if err != nil {
        return fmt.Errorf("部分更新用戶錯誤: %w", err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("用戶不存在")
    }
    
    return nil
}
```

### 5. 刪除數據 (Delete)

```go
// 刪除用戶
func deleteUser(db *sql.DB, id int) error {
    query := `DELETE FROM users WHERE id = ?`
    
    result, err := db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("刪除用戶錯誤: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("獲取影響行數錯誤: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("用戶不存在")
    }
    
    return nil
}

// 批量刪除
func deleteUsersByAge(db *sql.DB, maxAge int) (int64, error) {
    query := `DELETE FROM users WHERE age > ?`
    
    result, err := db.Exec(query, maxAge)
    if err != nil {
        return 0, fmt.Errorf("批量刪除用戶錯誤: %w", err)
    }
    
    return result.RowsAffected()
}

// 軟刪除（添加刪除標記）
func softDeleteUser(db *sql.DB, id int) error {
    query := `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`
    
    result, err := db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("軟刪除用戶錯誤: %w", err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("用戶不存在或已被刪除")
    }
    
    return nil
}
```

## 🔄 事務處理

### 1. 基本事務

```go
func transferMoney(db *sql.DB, fromUserID, toUserID int, amount float64) error {
    // 開始事務
    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("開始事務錯誤: %w", err)
    }
    
    // 使用 defer 確保事務被回滾或提交
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()
    
    // 檢查發送方餘額
    var balance float64
    err = tx.QueryRow("SELECT balance FROM accounts WHERE user_id = ?", fromUserID).Scan(&balance)
    if err != nil {
        return fmt.Errorf("查詢發送方餘額錯誤: %w", err)
    }
    
    if balance < amount {
        return fmt.Errorf("餘額不足")
    }
    
    // 扣除發送方餘額
    _, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE user_id = ?", amount, fromUserID)
    if err != nil {
        return fmt.Errorf("扣除發送方餘額錯誤: %w", err)
    }
    
    // 增加接收方餘額
    _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE user_id = ?", amount, toUserID)
    if err != nil {
        return fmt.Errorf("增加接收方餘額錯誤: %w", err)
    }
    
    // 記錄轉帳歷史
    _, err = tx.Exec("INSERT INTO transfers (from_user_id, to_user_id, amount) VALUES (?, ?, ?)", 
        fromUserID, toUserID, amount)
    if err != nil {
        return fmt.Errorf("記錄轉帳歷史錯誤: %w", err)
    }
    
    // 提交事務
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("提交事務錯誤: %w", err)
    }
    
    return nil
}
```

### 2. 事務選項

```go
import "context"

func transactionWithOptions(db *sql.DB) error {
    // 設置事務選項
    txOptions := &sql.TxOptions{
        Isolation: sql.LevelReadCommitted, // 隔離級別
        ReadOnly:  false,                  // 是否只讀
    }
    
    ctx := context.Background()
    tx, err := db.BeginTx(ctx, txOptions)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // 執行事務操作
    _, err = tx.ExecContext(ctx, "INSERT INTO users (name, email, age) VALUES (?, ?, ?)", 
        "Alice", "alice@example.com", 25)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

## 🏗️ 連接池管理

### 連接池配置

```go
func configureConnectionPool(db *sql.DB) {
    // 設置最大打開連接數
    db.SetMaxOpenConns(25)
    
    // 設置最大空閒連接數
    db.SetMaxIdleConns(10)
    
    // 設置連接的最大生存時間
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // 設置連接的最大空閒時間
    db.SetConnMaxIdleTime(30 * time.Second)
}

// 監控連接池狀態
func monitorDBStats(db *sql.DB) {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        stats := db.Stats()
        fmt.Printf("DB Stats - Open: %d, InUse: %d, Idle: %d\n", 
            stats.OpenConnections, stats.InUse, stats.Idle)
    }
}
```

## 🔍 高級查詢

### 1. 聯接查詢

```go
type UserWithPosts struct {
    User
    Posts []Post `json:"posts"`
}

type Post struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

// 獲取用戶及其文章
func getUserWithPosts(db *sql.DB, userID int) (*UserWithPosts, error) {
    // 查詢用戶信息
    user, err := getUserByID(db, userID)
    if err != nil {
        return nil, err
    }
    
    // 查詢用戶的文章
    query := `SELECT p.id, p.user_id, p.title, p.content, p.created_at
              FROM posts p
              WHERE p.user_id = ?
              ORDER BY p.created_at DESC`
    
    rows, err := db.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("查詢用戶文章錯誤: %w", err)
    }
    defer rows.Close()
    
    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return &UserWithPosts{
        User:  *user,
        Posts: posts,
    }, nil
}

// 複雜聯接查詢
func getUsersWithPostCount(db *sql.DB) ([]map[string]interface{}, error) {
    query := `SELECT u.id, u.name, u.email, COUNT(p.id) as post_count
              FROM users u
              LEFT JOIN posts p ON u.id = p.user_id
              GROUP BY u.id, u.name, u.email
              ORDER BY post_count DESC`
    
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var results []map[string]interface{}
    for rows.Next() {
        var id, postCount int
        var name, email string
        
        err := rows.Scan(&id, &name, &email, &postCount)
        if err != nil {
            return nil, err
        }
        
        results = append(results, map[string]interface{}{
            "id":         id,
            "name":       name,
            "email":      email,
            "post_count": postCount,
        })
    }
    
    return results, nil
}
```

### 2. 分頁查詢

```go
type PaginationResult struct {
    Data       []User `json:"data"`
    Page       int    `json:"page"`
    PerPage    int    `json:"per_page"`
    Total      int    `json:"total"`
    TotalPages int    `json:"total_pages"`
}

func getUsersPaginated(db *sql.DB, page, perPage int) (*PaginationResult, error) {
    // 計算偏移量
    offset := (page - 1) * perPage
    
    // 查詢總數
    var total int
    err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
    if err != nil {
        return nil, err
    }
    
    // 查詢分頁數據
    users, err := getUsers(db, perPage, offset)
    if err != nil {
        return nil, err
    }
    
    // 計算總頁數
    totalPages := (total + perPage - 1) / perPage
    
    return &PaginationResult{
        Data:       users,
        Page:       page,
        PerPage:    perPage,
        Total:      total,
        TotalPages: totalPages,
    }, nil
}
```

## 🎯 最佳實踐

### 1. 錯誤處理
- 適當處理 `sql.ErrNoRows`
- 區分業務錯誤和系統錯誤
- 提供有意義的錯誤信息

### 2. 安全考慮
- 使用參數化查詢防止 SQL 注入
- 驗證和清理用戶輸入
- 最小權限原則

### 3. 性能優化
- 合理使用索引
- 避免 N+1 查詢問題
- 使用連接池
- 預編譯常用語句

### 4. 資源管理
- 及時關閉 Rows、Stmt
- 正確處理事務
- 監控連接池狀態

---

這四章 (17-20) 涵蓋了 Go 語言實際應用的核心主題，為構建完整的 Web 應用程序奠定了基礎。