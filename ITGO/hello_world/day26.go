package hello_world

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func Day26() {
	dsn := "postgres://root:Root&123@localhost:5432/go?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("成功連接 PostgreSQL")
	createTable()
}

func createTable() {
	result, err := db.Exec(`CREATE TABLE IF NOT EXISTS "student" ("name" VARCHAR(10))`)
	if err != nil {
		log.Fatal("Create table failed:", err)
		return
	}
	fmt.Println(result)
}
