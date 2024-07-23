package redis

import "github.com/redis/go-redis/v9"

func InitDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.197.133:6379",
		Password: "", 
		DB:		0,
	})
}