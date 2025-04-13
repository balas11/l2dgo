package main

import (
	"authtok/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("POST /auth", handlers.Login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
