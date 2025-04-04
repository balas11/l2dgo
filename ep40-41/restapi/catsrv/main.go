package main

import (
	"catsrv/handlers"
	"log"
	"net/http"

	"rfcrud/dal"
)

func main() {
	dbURI := "postgres://postgres:postgres@localhost:5432/store?sslmode=disable"

	err := dal.OpenDB(dbURI)
	if err != nil {
		panic(err)
	}
	defer dal.CloseDB()
	if err := dal.PingTest(); err != nil {
		log.Println(err)
		return
	}

	//Test handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	//Define handlers
	http.HandleFunc("GET /categories", handlers.ListCategories)
	http.HandleFunc("GET /categories/{id}", handlers.GetCategory)
	http.HandleFunc("DELETE /categories/{id}", handlers.DeleteCategory)

	http.HandleFunc("GET /categories/{id}/products", handlers.ListProductsByCategoryID)

	http.HandleFunc("GET /products", handlers.ListProducts)
	http.HandleFunc("GET /products/{id}", handlers.GetProduct)
	http.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Println(err)
		return
	}
}
