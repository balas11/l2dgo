package main

import (
	"net/http"
)

type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))
}
//Step 1
// func main() {
// 	http.ListenAndServe("127.0.0.1:8080", &HelloHandler{})
// }

// Step2
// func main() {
// 	server := &http.Server{
// 		Addr:    "127.0.0.1:8080",
// 		Handler: &HelloHandler{},
// 	}
// 	server.ListenAndServe()
// }

// Step 3
func main() {
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// Step 4
// func main() {
// 	server := &http.Server{
// 		Addr: "127.0.0.1:8080",
// 	}
// 	server.ListenAndServe()
// }
