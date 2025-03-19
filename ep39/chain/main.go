package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func LogHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, w.Header().Get("Content-Type"))
		userInfo := (r.Context().Value("userInfo")).(map[string]string)
		log.Printf("user : %s, role: %s", userInfo["username"], userInfo["role"])
		h.ServeHTTP(w, r)
	})
}

func TextHeaderHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		h.ServeHTTP(w, r)
	})
}

func UserInfoHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set user info in the request context
		userInfo := map[string]string{
			"username": "bala",
			"role":     "admin",
		}
		modReq := r.Clone(context.WithValue(r.Context(), "userInfo", userInfo))
		h.ServeHTTP(w, modReq)
	})
}
func ListProducts(w http.ResponseWriter, r *http.Request) {
	userInfo := (r.Context().Value("userInfo")).(map[string]string)
	username := userInfo["username"]
	w.Write([]byte(fmt.Sprintf("\nHello %s!\n", username)))
	w.Write([]byte("Product List here..."))
}

type HandlerPipeStage func(http.HandlerFunc) http.HandlerFunc

func Pipeline(h http.HandlerFunc, stages ...HandlerPipeStage) http.HandlerFunc {
	if len(stages) < 1 {
		return h
	}

	staged := h

	for i := len(stages) - 1; i >= 0; i-- {
		staged = stages[i](staged)
	}
	return staged
}

func main() {
	textAPIStages := []HandlerPipeStage{
		UserInfoHandler,
		TextHeaderHandler,
		LogHandler,
	}
	// http.HandleFunc("/products", UserInfoHandler(TextHeaderHandler(LogHandler(ListProducts))))
	http.HandleFunc("/products", Pipeline(ListProducts, textAPIStages...))
	log.Fatalln(http.ListenAndServe("127.0.0.1:8080", nil))
}
