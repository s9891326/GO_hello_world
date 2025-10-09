package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
)

// User 用戶結構體
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 初始化數據庫
func initDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./demo.db")
	if err != nil {
		return nil, fmt.Errorf("打開數據庫錯誤: %w", err)
	}
	
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("連接數據庫錯誤: %w", err)
	}
	
	// 設置連接池參數
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	return db, nil
}

// 創建表
func createTables(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	
	_, err := db.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("創建用戶表錯誤: %w", err)
	}
	
	fmt.Println("數據表創建成功")
	return nil
}

// 插入用戶
func createUser(db *sql.DB, user User) (int64, error) {
	query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`
	
	result, err := db.Exec(query, user.Name, user.Email, user.Age)
	if err != nil {
		return 0, fmt.Errorf("插入用戶錯誤: %w", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("獲取插入ID錯誤: %w", err)
	}
	
	return id, nil
}

// 查詢所有用戶
func getAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, age, created_at, updated_at FROM users ORDER BY created_at DESC`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查詢用戶錯誤: %w", err)
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("掃描用戶數據錯誤: %w", err)
		}
		users = append(users, user)
	}
	
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍歷結果錯誤: %w", err)
	}
	
	return users, nil
}

// 根據ID查詢用戶
func getUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?`
	
	var user User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("用戶不存在")
		}
		return nil, fmt.Errorf("查詢用戶錯誤: %w", err)
	}
	
	return &user, nil
}

// 更新用戶
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
		return fmt.Errorf("用戶不存在")
	}
	
	return nil
}

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

// 演示事務處理
func demonstrateTransaction(db *sql.DB) {
	fmt.Println("\n=== 事務處理演示 ===")
	
	// 開始事務
	tx, err := db.Begin()
	if err != nil {
		log.Printf("開始事務錯誤: %v", err)
		return
	}
	
	// 使用defer確保事務被處理
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("事務已回滾")
		}
	}()
	
	// 在事務中插入多個用戶
	users := []User{
		{Name: "事務用戶1", Email: "tx1@example.com", Age: 25},
		{Name: "事務用戶2", Email: "tx2@example.com", Age: 30},
		{Name: "事務用戶3", Email: "tx3@example.com", Age: 35},
	}
	
	for _, user := range users {
		_, err = tx.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", 
			user.Name, user.Email, user.Age)
		if err != nil {
			log.Printf("事務中插入用戶錯誤: %v", err)
			return
		}
	}
	
	// 提交事務
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事務錯誤: %v", err)
		return
	}
	
	fmt.Println("事務提交成功，插入了3個用戶")
}

// 演示預編譯語句
func demonstratePreparedStatement(db *sql.DB) {
	fmt.Println("\n=== 預編譯語句演示 ===")
	
	// 準備語句
	stmt, err := db.Prepare("INSERT INTO users (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("準備語句錯誤: %v", err)
		return
	}
	defer stmt.Close()
	
	// 使用預編譯語句插入多個用戶
	users := []User{
		{Name: "預編譯用戶1", Email: "prep1@example.com", Age: 22},
		{Name: "預編譯用戶2", Email: "prep2@example.com", Age: 27},
		{Name: "預編譯用戶3", Email: "prep3@example.com", Age: 32},
	}
	
	for _, user := range users {
		result, err := stmt.Exec(user.Name, user.Email, user.Age)
		if err != nil {
			log.Printf("執行預編譯語句錯誤: %v", err)
			continue
		}
		
		id, _ := result.LastInsertId()
		fmt.Printf("插入用戶: %s (ID: %d)\n", user.Name, id)
	}
}

// 演示統計查詢
func demonstrateAggregateQueries(db *sql.DB) {
	fmt.Println("\n=== 統計查詢演示 ===")
	
	// 用戶總數
	var totalUsers int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalUsers)
	if err != nil {
		log.Printf("查詢用戶總數錯誤: %v", err)
	} else {
		fmt.Printf("用戶總數: %d\n", totalUsers)
	}
	
	// 平均年齡
	var avgAge float64
	err = db.QueryRow("SELECT AVG(age) FROM users").Scan(&avgAge)
	if err != nil {
		log.Printf("查詢平均年齡錯誤: %v", err)
	} else {
		fmt.Printf("平均年齡: %.2f\n", avgAge)
	}
	
	// 年齡分布
	query := `SELECT 
		CASE 
			WHEN age < 25 THEN '18-24'
			WHEN age < 35 THEN '25-34'
			WHEN age < 45 THEN '35-44'
			ELSE '45+'
		END as age_group,
		COUNT(*) as count
		FROM users 
		GROUP BY age_group`
	
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("查詢年齡分布錯誤: %v", err)
		return
	}
	defer rows.Close()
	
	fmt.Println("年齡分布:")
	for rows.Next() {
		var ageGroup string
		var count int
		
		err := rows.Scan(&ageGroup, &count)
		if err != nil {
			log.Printf("掃描年齡分布錯誤: %v", err)
			continue
		}
		
		fmt.Printf("  %s: %d 人\n", ageGroup, count)
	}
}

func main() {
	fmt.Println("=== Go 數據庫操作演示 ===")
	
	// 初始化數據庫
	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// 創建表
	err = createTables(db)
	if err != nil {
		log.Fatal(err)
	}
	
	// 1. 插入用戶演示
	fmt.Println("\n--- 插入用戶 ---")
	users := []User{
		{Name: "Alice Johnson", Email: "alice@example.com", Age: 25},
		{Name: "Bob Smith", Email: "bob@example.com", Age: 30},
		{Name: "Charlie Brown", Email: "charlie@example.com", Age: 35},
	}
	
	for _, user := range users {
		id, err := createUser(db, user)
		if err != nil {
			log.Printf("插入用戶錯誤: %v", err)
		} else {
			fmt.Printf("插入用戶成功: %s (ID: %d)\n", user.Name, id)
		}
	}
	
	// 2. 查詢所有用戶
	fmt.Println("\n--- 查詢所有用戶 ---")
	allUsers, err := getAllUsers(db)
	if err != nil {
		log.Printf("查詢用戶錯誤: %v", err)
	} else {
		for _, user := range allUsers {
			fmt.Printf("ID: %d, 姓名: %s, 郵箱: %s, 年齡: %d\n", 
				user.ID, user.Name, user.Email, user.Age)
		}
	}
	
	// 3. 根據ID查詢用戶
	fmt.Println("\n--- 根據ID查詢用戶 ---")
	if len(allUsers) > 0 {
		user, err := getUserByID(db, allUsers[0].ID)
		if err != nil {
			log.Printf("查詢用戶錯誤: %v", err)
		} else {
			fmt.Printf("查詢到用戶: %+v\n", user)
		}
	}
	
	// 4. 更新用戶
	fmt.Println("\n--- 更新用戶 ---")
	if len(allUsers) > 0 {
		updateData := User{
			Name:  "Alice Updated",
			Email: "alice.updated@example.com",
			Age:   26,
		}
		
		err := updateUser(db, allUsers[0].ID, updateData)
		if err != nil {
			log.Printf("更新用戶錯誤: %v", err)
		} else {
			fmt.Printf("用戶更新成功: ID %d\n", allUsers[0].ID)
		}
	}
	
	// 5. 事務處理演示
	demonstrateTransaction(db)
	
	// 6. 預編譯語句演示
	demonstratePreparedStatement(db)
	
	// 7. 統計查詢演示
	demonstrateAggregateQueries(db)
	
	// 8. 刪除用戶演示
	fmt.Println("\n--- 刪除用戶 ---")
	// 刪除最後一個用戶
	currentUsers, _ := getAllUsers(db)
	if len(currentUsers) > 0 {
		lastUser := currentUsers[len(currentUsers)-1]
		err := deleteUser(db, lastUser.ID)
		if err != nil {
			log.Printf("刪除用戶錯誤: %v", err)
		} else {
			fmt.Printf("用戶刪除成功: %s (ID: %d)\n", lastUser.Name, lastUser.ID)
		}
	}
	
	// 最終用戶統計
	fmt.Println("\n--- 最終用戶統計 ---")
	finalUsers, _ := getAllUsers(db)
	fmt.Printf("數據庫中共有 %d 個用戶\n", len(finalUsers))
	
	fmt.Println("\n=== 演示完成 ===")
}