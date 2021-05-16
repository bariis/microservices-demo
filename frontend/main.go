package main

import "github.com/gorilla/mux"

func main() {
	server := FrontendServer{
		Router: mux.NewRouter(),
	}
	server.InitializeRoutes()
	server.Run(":8080")
}
