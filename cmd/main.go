package main

import (
	"authorizate/config"
	"authorizate/internal/auth"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	router := auth.NewHandler(cfg)
	http.ListenAndServe(":8081", router)
}
