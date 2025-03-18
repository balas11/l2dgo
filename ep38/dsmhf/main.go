package main

import "net/http"

func listCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Categories"))
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Products"))
}

func showCatalog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Catalog"))
}

func main() {

	http.HandleFunc("/", showCatalog)
	http.HandleFunc("/categories", listCategories)
	http.HandleFunc("/products", listProducts)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
