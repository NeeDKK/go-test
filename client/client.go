package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.25.10.120:6379",
		Password: "Netvine123#@!", // no password set
		DB:       1,               // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		val := client.Get(context.Background(), "current_base_line_warning").Val()
		fmt.Println(val == "")
	}
}
