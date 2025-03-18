package main

import "net/http"

type CatalogHandler struct{}
type CategoriesHandler struct{}
type ProductsHandler struct{}

func (h CatalogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("catalog"))
}
func (h CategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("categories"))
}
func (h ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("products"))
}
func main() {
	http.Handle("/", CatalogHandler{})
	http.Handle("/categories", CategoriesHandler{})
	http.Handle("/products", ProductsHandler{})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
