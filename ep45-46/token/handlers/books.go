package handlers

import (
	"authtok/data"
	"encoding/json"
	"net/http"
	"strconv"
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
	books := data.Data.GetBooks()
	writeResponse(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book := data.Data.GetBook(id)
	writeResponse(w, http.StatusOK, book)
}

func SaveBook(w http.ResponseWriter, r *http.Request) {
	book := data.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	book = data.Data.AddBook(book.Name, book.Author)
	writeResponse(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book := data.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	data.Data.UpdateBook(id, book.Name, book.Author)
	book = data.Data.GetBook(id)
	writeResponse(w, http.StatusOK, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Data.DeleteBook(id)
	writeResponse(w, http.StatusNoContent, nil)
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
