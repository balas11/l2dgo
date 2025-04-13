package handlers

import (
	"authtok/data"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userName, password, err := decodeLogin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := data.LoginUser(userName, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func decodeLogin(r *http.Request) (string, string, error) {
	type loginUser struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var user loginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	return user.UserName, user.Password, err
}
