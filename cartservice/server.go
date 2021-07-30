package main

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	redis  redisCache
	router *mux.Router
}

type redisCache struct {
	redisClient *redis.Client
	redisCache  *cache.Cache
}

func (s *server) initializeRoutes() {
	s.router.HandleFunc("/add", s.handleAddCart).Methods("POST")
}

func (s *server) Initialize() {
	s.redis.redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	s.redis.redisCache = cache.New(&cache.Options{
		Redis: s.redis.redisClient,
	})

	s.router = mux.NewRouter()

	s.initializeRoutes()
}

func (s *server) Run() {
	log.Fatal(http.ListenAndServe(":5001", s.router))
}
