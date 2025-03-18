package main

import "net/http"

type CatalogServer struct{}

func (c *CatalogServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI

	switch uri {
	case "/":
		w.Write([]byte("Catalog"))
	case "/products":
		w.Write([]byte("Products"))
	case "/categories":
		w.Write([]byte("Categories"))
	}
}
func main() {
	http.ListenAndServe("127.0.0.1:8080", &CatalogServer{})
}
