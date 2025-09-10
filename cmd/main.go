package main

import (
	"authorizate/internal/auth"
	"net/http"
)

func main() {
	router := auth.NewHandler()
	http.ListenAndServe(":8081", router)
}
