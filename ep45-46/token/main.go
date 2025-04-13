package main

import (
	"authtok/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("POST /auth", handlers.Login)

	http.HandleFunc("GET /books",
		handlers.Pipeline(handlers.ListBooks, handlers.ProtectedStages...))
	http.HandleFunc("GET /books/{id}",
		handlers.Pipeline(handlers.GetBook, handlers.ProtectedStages...))
	http.HandleFunc("POST /books",
		handlers.Pipeline(handlers.SaveBook, handlers.AdminStages...))
	http.HandleFunc("PUT /books/{id}",
		handlers.Pipeline(handlers.UpdateBook, handlers.AdminStages...))
	http.HandleFunc("DELETE /books/{id}",
		handlers.Pipeline(handlers.DeleteBook, handlers.AdminStages...))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
