package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func BenchmarkRedisPut(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()

	for i := 0; i < b.N; i++ {
		rdb.Set(context.TODO(), "key", "value", time.Second*20)
	}
}

func BenchmarkRedisPut2(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()

	for i := 0; i < b.N; i++ {
		rdb.Set(context.TODO(), "key"+fmt.Sprint(i), "value"+fmt.Sprint(i), time.Second*20)
	}
}

func BenchmarkRedisHset(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()

	for i := 0; i < b.N; i++ {
		rdb.HSet(context.TODO(), "hset"+fmt.Sprint(i), "k"+fmt.Sprint(i), "v"+fmt.Sprint(i), "k"+fmt.Sprint(i+1), "v"+fmt.Sprint(i+1), time.Second*20)
	}
}
