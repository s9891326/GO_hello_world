package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 模擬 SQL 查詢（假裝要 500ms）
func slowSQLQuery(userID string) string {
	time.Sleep(500 * time.Millisecond) // 模擬慢查詢
	return fmt.Sprintf("UserData{id=%s, name=Eddy}", userID)
}

func getUserHandler(c *gin.Context) {
	userID := c.Param("id")
	cacheKey := "user:" + userID

	// 1. 先查 Redis
	val, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"from": "cache", "data": val})
		return
	}

	// 2. 模擬 DB 查詢
	data := slowSQLQuery(userID)

	// 3. 存入 Redis
	err = rdb.Set(ctx, cacheKey, data, 10*time.Second).Err()
	if err != nil {
		log.Println("Redis SET error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"from": "db", "data": data})
}

func slowQuery(id string) (*User, error) {
	// 模擬慢 SQL (100ms)
	time.Sleep(100 * time.Millisecond)
	return &User{ID: 1, Name: "Eddy"}, nil
}

func wrkUserHandler(c *gin.Context) {
	userID := c.Param("id")
	cacheKey := "user:" + userID

	if val, err := rdb.Get(ctx, cacheKey).Result(); err == nil {
		var user User
		_ = json.Unmarshal([]byte(val), &user)
		c.JSON(http.StatusOK, gin.H{"from": "cache", "user": user})
		return
	}

	user, _ := slowQuery(userID)

	// 寫入快取
	data, _ := json.Marshal(user)
	_ = rdb.Set(ctx, cacheKey, data, 30*time.Second).Err()

	c.JSON(http.StatusOK, gin.H{"from": "db", "user": user})
}

//docker run --rm williamyeh/wrk -t4 -c100 -d10s http://192.168.213.43:8080/user2/1123

/*
wrk: 壓力測試，測QPS看看吞吐量如何
$ docker run --rm williamyeh/wrk -t4 -c100 -d10s http://host.docker.internal:8080/user2/1
Running 10s test @ http://host.docker.internal:8080/user2/1
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    16.03ms   27.67ms 316.58ms   96.62%
    Req/Sec     2.11k   533.49     3.48k    75.90%
  82347 requests in 10.04s, 12.96MB read
Requests/sec:   8204.04
Transfer/sec:      1.29MB

pprof: 效能測試，檢測
go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=30
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/profile?seconds=30
Saved profile in C:\Users\eddy\pprof\pprof.___go_build_main_go__4_.exe.samples.cpu.001.pb.gz
File: ___go_build_main_go__4_.exe
Build ID: C:\Users\eddy\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___go_build_main_go__4_.exe2025-08-16 20:15:58.6209976 +0800 CST
Type: cpu
Time: 2025-08-16 20:16:02 CST
Duration: 30.10s, Total samples = 10ms (0.033%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10ms, 100% of 10ms total
      flat  flat%   sum%        cum   cum%
      10ms   100%   100%       10ms   100%  runtime/pprof.StopCPUProfile
         0     0%   100%       10ms   100%  net/http.(*ServeMux).ServeHTTP
         0     0%   100%       10ms   100%  net/http.(*conn).serve
         0     0%   100%       10ms   100%  net/http.HandlerFunc.ServeHTTP
         0     0%   100%       10ms   100%  net/http.serverHandler.ServeHTTP
         0     0%   100%       10ms   100%  net/http/pprof.Profile
*/

func main() {
	// 連接 Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "fnq{8q*(DHd121}2d)",
		DB:       0,
	})

	// 啟用 pprof 在 6060 port
	// 效能測試
	go func() {
		log.Println("pprof listening on :6060")
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	// Gin Router
	r := gin.Default()
	r.GET("/user/:id", getUserHandler)
	r.GET("/user2/:id", wrkUserHandler)

	log.Println("API running on :8080")
	r.Run(":8080")

}
