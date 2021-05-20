package main

import "github.com/go-redis/redis/v8"

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr : "localhost:6379",
		Password: "",
		DB: 0,
	}
	err :=
}