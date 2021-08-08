package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type server struct {
	router      *mux.Router
	redisClient *redis.Client
}

func (s *server) initializeRoutes() {
	s.router.HandleFunc("/api/v1/add-cart", s.handleAddCart).Methods("POST")
	s.router.HandleFunc("/api/v1/get-cart", s.handleGetCart).Methods("GET")
	s.router.HandleFunc("/api/v1/empty-cart", s.handleEmptyCart).Methods("DELETE")
}

func (s *server) Initialize() {
	s.redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	s.router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *server) Run() {
	srv := &http.Server{
		Addr:         "cartservice:5001",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      s.router,
	}

	log.Fatal(srv.ListenAndServe())
}
