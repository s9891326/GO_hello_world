package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/*
Day 4：高併發 I/O 與快取
*/

var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "fnq{8q*(DHd121}2d)",
	DB:       0,
})

func getData(ctx context.Context, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("Cache miss, query DB...")
		val = "DB Data for " + key
		rdb.Set(ctx, key, val, 0)
		return val
	}
	fmt.Println("Cache hit")
	return val
}

func main() {
	ctx := context.Background()
	fmt.Println(getData(ctx, "user:1"))
	fmt.Println(getData(ctx, "user:1"))
}
