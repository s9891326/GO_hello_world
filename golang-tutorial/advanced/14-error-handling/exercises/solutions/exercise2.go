// 練習 2 解答：自定義錯誤類型 - 用戶管理系統
package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 自定義錯誤類型
type UserError struct {
	Code    int
	Field   string
	Message string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("用戶錯誤 [%d] %s: %s", e.Code, e.Field, e.Message)
}

// 預定義錯誤
var (
	ErrUsernameEmpty    = &UserError{1001, "username", "用戶名不能為空"}
	ErrEmailInvalid     = &UserError{1002, "email", "郵箱格式無效"}
	ErrAgeInvalid       = &UserError{1003, "age", "年齡必須在 18-100 之間"}
	ErrUsernameExists   = &UserError{1004, "username", "用戶名已存在"}
	ErrUserNotFound     = &UserError{1005, "id", "用戶不存在"}
)

type User struct {
	ID       int
	Username string
	Email    string
	Age      int
}

type UserManager struct {
	users  map[int]User
	nextID int
}

func NewUserManager() *UserManager {
	return &UserManager{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (um *UserManager) validateUser(user User) error {
	// 檢查用戶名
	if strings.TrimSpace(user.Username) == "" {
		return ErrUsernameEmpty
	}
	
	// 檢查郵箱格式
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid
	}
	
	// 檢查年齡
	if user.Age < 18 || user.Age > 100 {
		return ErrAgeInvalid
	}
	
	return nil
}

func (um *UserManager) usernameExists(username string) bool {
	for _, user := range um.users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func (um *UserManager) CreateUser(user User) error {
	// 驗證用戶數據
	if err := um.validateUser(user); err != nil {
		return err
	}
	
	// 檢查用戶名是否已存在
	if um.usernameExists(user.Username) {
		return ErrUsernameExists
	}
	
	// 創建用戶
	user.ID = um.nextID
	um.users[um.nextID] = user
	um.nextID++
	
	fmt.Printf("用戶創建成功: ID=%d, Username=%s\n", user.ID, user.Username)
	return nil
}

func (um *UserManager) GetUser(id int) (User, error) {
	user, exists := um.users[id]
	if !exists {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (um *UserManager) UpdateUser(id int, updatedUser User) error {
	// 檢查用戶是否存在
	_, exists := um.users[id]
	if !exists {
		return ErrUserNotFound
	}
	
	// 驗證更新的用戶數據
	if err := um.validateUser(updatedUser); err != nil {
		return err
	}
	
	// 檢查用戶名是否與其他用戶衝突
	for userID, user := range um.users {
		if userID != id && user.Username == updatedUser.Username {
			return ErrUsernameExists
		}
	}
	
	// 更新用戶
	updatedUser.ID = id
	um.users[id] = updatedUser
	
	fmt.Printf("用戶更新成功: ID=%d, Username=%s\n", id, updatedUser.Username)
	return nil
}

func (um *UserManager) ListUsers() {
	fmt.Println("\n=== 用戶列表 ===")
	if len(um.users) == 0 {
		fmt.Println("沒有用戶")
		return
	}
	
	for _, user := range um.users {
		fmt.Printf("ID: %d, Username: %s, Email: %s, Age: %d\n", 
			user.ID, user.Username, user.Email, user.Age)
	}
}

func main() {
	um := NewUserManager()
	
	fmt.Println("=== 用戶管理系統測試 ===")
	
	// 測試創建用戶
	fmt.Println("\n--- 創建用戶測試 ---")
	testUsers := []User{
		{Username: "alice", Email: "alice@example.com", Age: 25},
		{Username: "", Email: "bob@example.com", Age: 30},        // 用戶名為空
		{Username: "charlie", Email: "invalid-email", Age: 35},   // 無效郵箱
		{Username: "david", Email: "david@example.com", Age: 15}, // 年齡無效
		{Username: "alice", Email: "alice2@example.com", Age: 28}, // 用戶名重複
		{Username: "eve", Email: "eve@example.com", Age: 22},
	}
	
	for i, user := range testUsers {
		fmt.Printf("\n測試用戶 %d: %+v\n", i+1, user)
		err := um.CreateUser(user)
		if err != nil {
			fmt.Printf("創建失敗: %v\n", err)
			
			// 檢查錯誤類型
			var userErr *UserError
			if errors.As(err, &userErr) {
				fmt.Printf("錯誤代碼: %d, 字段: %s\n", userErr.Code, userErr.Field)
			}
		}
	}
	
	// 列出所有用戶
	um.ListUsers()
	
	// 測試獲取用戶
	fmt.Println("\n--- 獲取用戶測試 ---")
	for _, id := range []int{1, 999} {
		user, err := um.GetUser(id)
		if err != nil {
			fmt.Printf("獲取用戶 ID %d 失敗: %v\n", id, err)
		} else {
			fmt.Printf("獲取用戶 ID %d 成功: %+v\n", id, user)
		}
	}
	
	// 測試更新用戶
	fmt.Println("\n--- 更新用戶測試 ---")
	updatedUser := User{Username: "alice_updated", Email: "alice_new@example.com", Age: 26}
	err := um.UpdateUser(1, updatedUser)
	if err != nil {
		fmt.Printf("更新失敗: %v\n", err)
	}
	
	// 測試更新不存在的用戶
	err = um.UpdateUser(999, updatedUser)
	if err != nil {
		fmt.Printf("更新不存在用戶失敗: %v\n", err)
	}
	
	// 最終用戶列表
	um.ListUsers()
}