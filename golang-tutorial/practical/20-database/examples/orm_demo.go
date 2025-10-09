package main

import (
	"fmt"
	"log"
	"time"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用戶模型
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Age       int       `json:"age"`
	Birthday  time.Time `json:"birthday"`
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
	// 關聯關係
	Profile *Profile `json:"profile,omitempty"`
	Posts   []Post   `json:"posts,omitempty"`
	Roles   []Role   `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

// Profile 用戶檔案
type Profile struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id" gorm:"uniqueIndex"`
	Bio      string `json:"bio" gorm:"type:text"`
	Avatar   string `json:"avatar"`
	Website  string `json:"website"`
	Location string `json:"location"`
	User     User   `json:"user"`
}

// Post 文章模型
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"size:200;not null"`
	Content   string    `json:"content" gorm:"type:text"`
	Published bool      `json:"published" gorm:"default:false"`
	AuthorID  uint      `json:"author_id"`
	Author    User      `json:"author"`
	Tags      []Tag     `json:"tags" gorm:"many2many:post_tags;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tag 標籤模型
type Tag struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"size:50;uniqueIndex"`
	Posts []Post `json:"posts" gorm:"many2many:post_tags;"`
}

// Role 角色模型
type Role struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"size:50;uniqueIndex"`
	Users []User `json:"users" gorm:"many2many:user_roles;"`
}

// 初始化數據庫
func initDB() (*gorm.DB, error) {
	// 配置日誌
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	
	// 連接數據庫
	db, err := gorm.Open(sqlite.Open("gorm_demo.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("連接數據庫失敗: %w", err)
	}
	
	// 自動遷移
	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Tag{}, &Role{})
	if err != nil {
		return nil, fmt.Errorf("數據遷移失敗: %w", err)
	}
	
	return db, nil
}

// 演示基本 CRUD 操作
func demonstrateBasicCRUD(db *gorm.DB) {
	fmt.Println("=== 基本 CRUD 操作 ===")
	
	// Create - 創建用戶
	fmt.Println("\n--- 創建用戶 ---")
	users := []User{
		{
			Name:     "Alice Johnson",
			Email:    "alice@example.com",
			Age:      25,
			Birthday: time.Date(1998, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:     "Bob Smith",
			Email:    "bob@example.com",
			Age:      30,
			Birthday: time.Date(1993, 8, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:     "Charlie Brown",
			Email:    "charlie@example.com",
			Age:      35,
			Birthday: time.Date(1988, 12, 3, 0, 0, 0, 0, time.UTC),
		},
	}
	
	// 批量創建
	result := db.Create(&users)
	if result.Error != nil {
		log.Printf("創建用戶錯誤: %v", result.Error)
	} else {
		fmt.Printf("成功創建 %d 個用戶\n", result.RowsAffected)
	}
	
	// Read - 查詢用戶
	fmt.Println("\n--- 查詢用戶 ---")
	
	// 查詢第一個用戶
	var firstUser User
	db.First(&firstUser)
	fmt.Printf("第一個用戶: %+v\n", firstUser)
	
	// 根據 ID 查詢
	var userByID User
	db.First(&userByID, users[0].ID)
	fmt.Printf("根據 ID 查詢: %+v\n", userByID)
	
	// 根據條件查詢
	var userByEmail User
	db.Where("email = ?", "alice@example.com").First(&userByEmail)
	fmt.Printf("根據郵箱查詢: %+v\n", userByEmail)
	
	// 查詢多個用戶
	var allUsers []User
	db.Find(&allUsers)
	fmt.Printf("所有用戶數量: %d\n", len(allUsers))
	
	// Update - 更新用戶
	fmt.Println("\n--- 更新用戶 ---")
	
	// 更新單個字段
	db.Model(&firstUser).Update("age", 26)
	fmt.Printf("更新後的用戶年齡: %d\n", firstUser.Age)
	
	// 更新多個字段
	db.Model(&firstUser).Updates(User{Name: "Alice Updated", Age: 27})
	fmt.Printf("更新後的用戶姓名: %s, 年齡: %d\n", firstUser.Name, firstUser.Age)
	
	// 批量更新
	db.Model(&User{}).Where("age > ?", 30).Update("active", false)
	
	// Delete - 刪除用戶
	fmt.Println("\n--- 刪除用戶 ---")
	
	// 軟刪除（如果模型有 DeletedAt 字段）
	db.Delete(&User{}, users[2].ID)
	
	// 查詢未刪除的用戶
	var activeUsers []User
	db.Find(&activeUsers)
	fmt.Printf("活躍用戶數量: %d\n", len(activeUsers))
	
	// 包含軟刪除的查詢
	var allUsersIncludeDeleted []User
	db.Unscoped().Find(&allUsersIncludeDeleted)
	fmt.Printf("所有用戶（包括已刪除）: %d\n", len(allUsersIncludeDeleted))
}

// 演示關聯操作
func demonstrateAssociations(db *gorm.DB) {
	fmt.Println("\n=== 關聯操作演示 ===")
	
	// 創建用戶檔案
	fmt.Println("\n--- 一對一關聯 ---")
	var user User
	db.First(&user)
	
	profile := Profile{
		UserID:   user.ID,
		Bio:      "這是我的個人簡介",
		Avatar:   "avatar.jpg",
		Website:  "https://example.com",
		Location: "台北市",
	}
	
	db.Create(&profile)
	
	// 查詢用戶及其檔案
	var userWithProfile User
	db.Preload("Profile").First(&userWithProfile, user.ID)
	fmt.Printf("用戶檔案: %+v\n", userWithProfile.Profile)
	
	// 創建文章（一對多關聯）
	fmt.Println("\n--- 一對多關聯 ---")
	posts := []Post{
		{
			Title:     "第一篇文章",
			Content:   "這是第一篇文章的內容",
			Published: true,
			AuthorID:  user.ID,
		},
		{
			Title:     "第二篇文章",
			Content:   "這是第二篇文章的內容",
			Published: false,
			AuthorID:  user.ID,
		},
	}
	
	db.Create(&posts)
	
	// 查詢用戶及其文章
	var userWithPosts User
	db.Preload("Posts").First(&userWithPosts, user.ID)
	fmt.Printf("用戶文章數量: %d\n", len(userWithPosts.Posts))
	
	// 創建標籤和角色（多對多關聯）
	fmt.Println("\n--- 多對多關聯 ---")
	
	// 創建標籤
	tags := []Tag{
		{Name: "Go語言"},
		{Name: "Web開發"},
		{Name: "數據庫"},
	}
	db.Create(&tags)
	
	// 創建角色
	roles := []Role{
		{Name: "管理員"},
		{Name: "編輯者"},
		{Name: "普通用戶"},
	}
	db.Create(&roles)
	
	// 為文章添加標籤
	var post Post
	db.First(&post)
	db.Model(&post).Association("Tags").Append([]Tag{tags[0], tags[1]})
	
	// 為用戶添加角色
	db.Model(&user).Association("Roles").Append([]Role{roles[0], roles[2]})
	
	// 查詢文章及其標籤
	var postWithTags Post
	db.Preload("Tags").First(&postWithTags, post.ID)
	fmt.Printf("文章標籤數量: %d\n", len(postWithTags.Tags))
	
	// 查詢用戶及其角色
	var userWithRoles User
	db.Preload("Roles").First(&userWithRoles, user.ID)
	fmt.Printf("用戶角色數量: %d\n", len(userWithRoles.Roles))
}

// 演示高級查詢
func demonstrateAdvancedQueries(db *gorm.DB) {
	fmt.Println("\n=== 高級查詢演示 ===")
	
	// 條件查詢
	fmt.Println("\n--- 條件查詢 ---")
	var users []User
	
	// WHERE 條件
	db.Where("age > ?", 25).Find(&users)
	fmt.Printf("年齡大於25的用戶: %d個\n", len(users))
	
	// 多個條件
	db.Where("age > ? AND active = ?", 20, true).Find(&users)
	fmt.Printf("年齡大於20且活躍的用戶: %d個\n", len(users))
	
	// IN 查詢
	db.Where("name IN ?", []string{"Alice Updated", "Bob Smith"}).Find(&users)
	fmt.Printf("指定姓名的用戶: %d個\n", len(users))
	
	// LIKE 查詢
	db.Where("email LIKE ?", "%@example.com").Find(&users)
	fmt.Printf("郵箱包含example.com的用戶: %d個\n", len(users))
	
	// 排序和限制
	fmt.Println("\n--- 排序和分頁 ---")
	
	// 排序
	db.Order("age desc").Find(&users)
	fmt.Printf("按年齡降序排列的第一個用戶年齡: %d\n", users[0].Age)
	
	// 限制數量
	db.Limit(2).Find(&users)
	fmt.Printf("限制查詢結果: %d個\n", len(users))
	
	// 分頁
	db.Offset(1).Limit(2).Find(&users)
	fmt.Printf("分頁查詢結果: %d個\n", len(users))
	
	// 聚合查詢
	fmt.Println("\n--- 聚合查詢 ---")
	
	var count int64
	db.Model(&User{}).Count(&count)
	fmt.Printf("用戶總數: %d\n", count)
	
	var avgAge float64
	db.Model(&User{}).Select("AVG(age)").Row().Scan(&avgAge)
	fmt.Printf("平均年齡: %.2f\n", avgAge)
	
	// 分組查詢
	type AgeGroup struct {
		AgeRange string
		Count    int64
	}
	
	var ageGroups []AgeGroup
	db.Model(&User{}).
		Select("CASE WHEN age < 25 THEN '<25' WHEN age < 35 THEN '25-34' ELSE '35+' END as age_range, COUNT(*) as count").
		Group("age_range").
		Scan(&ageGroups)
	
	fmt.Println("年齡分布:")
	for _, group := range ageGroups {
		fmt.Printf("  %s: %d人\n", group.AgeRange, group.Count)
	}
}

// 演示事務操作
func demonstrateTransactions(db *gorm.DB) {
	fmt.Println("\n=== 事務操作演示 ===")
	
	// 自動事務
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事務中創建用戶
		user := User{
			Name:  "Transaction User",
			Email: "tx@example.com",
			Age:   28,
		}
		
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		
		// 在事務中創建檔案
		profile := Profile{
			UserID:   user.ID,
			Bio:      "事務創建的檔案",
			Location: "事務城市",
		}
		
		if err := tx.Create(&profile).Error; err != nil {
			return err
		}
		
		// 模擬條件錯誤
		if user.Age < 30 {
			// return fmt.Errorf("年齡太小，回滾事務")
		}
		
		return nil
	})
	
	if err != nil {
		fmt.Printf("事務失敗: %v\n", err)
	} else {
		fmt.Println("事務成功提交")
	}
	
	// 手動事務
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	user := User{
		Name:  "Manual Transaction User",
		Email: "manual@example.com",
		Age:   32,
	}
	
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Printf("手動事務失敗: %v\n", err)
		return
	}
	
	if err := tx.Commit().Error; err != nil {
		fmt.Printf("提交事務失敗: %v\n", err)
		return
	}
	
	fmt.Println("手動事務成功")
}

// 演示原生 SQL
func demonstrateRawSQL(db *gorm.DB) {
	fmt.Println("\n=== 原生 SQL 演示 ===")
	
	// 原生查詢
	var users []User
	db.Raw("SELECT * FROM users WHERE age > ?", 25).Scan(&users)
	fmt.Printf("原生 SQL 查詢結果: %d個用戶\n", len(users))
	
	// 原生執行
	result := db.Exec("UPDATE users SET active = ? WHERE age > ?", false, 35)
	fmt.Printf("原生 SQL 更新了 %d 行\n", result.RowsAffected)
	
	// 使用 Row
	var name string
	var email string
	row := db.Raw("SELECT name, email FROM users WHERE id = ?", 1).Row()
	row.Scan(&name, &email)
	fmt.Printf("使用 Row 查詢: %s (%s)\n", name, email)
	
	// 使用 Rows
	rows, err := db.Raw("SELECT name, age FROM users").Rows()
	if err != nil {
		fmt.Printf("查詢錯誤: %v\n", err)
		return
	}
	defer rows.Close()
	
	fmt.Println("使用 Rows 遍歷:")
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Printf("  %s: %d歲\n", name, age)
	}
}

func main() {
	fmt.Println("=== GORM ORM 演示 ===")
	
	// 初始化數據庫
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	
	// 獲取底層 sql.DB 進行配置
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	
	// 設置連接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	// 演示各種功能
	demonstrateBasicCRUD(db)
	demonstrateAssociations(db)
	demonstrateAdvancedQueries(db)
	demonstrateTransactions(db)
	demonstrateRawSQL(db)
	
	fmt.Println("\n=== 演示完成 ===")
}