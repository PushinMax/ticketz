package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "my_redis_password",            
        DB:       0,              
    })

	ctx := context.Background()

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Set(ctx, "key1", 111, 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	// обработайте ошибку err

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	// обработайте ошибку err
	fmt.Println("key: ", val)


	v, err := client.Get(ctx, "key1").Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	// обработайте ошибку err
	fmt.Println("key1: ", v)
}

