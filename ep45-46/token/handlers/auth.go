package handlers

import (
	"authtok/data"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("The quick brown fox")

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

	token, err := createToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	retMap := map[string]string{"token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retMap)
}

func CheckAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		token = token[len("Bearer "):]
		user, err := verifyToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		modReq := r.Clone(context.WithValue(r.Context(), "user", user))
		h.ServeHTTP(w, modReq)
	})
}

func CheckAdmin(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(data.User)
		if !ok || user.Role != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
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

func createToken(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"role":     user.Role,
	})

	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) (data.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return data.User{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return data.User{UserName: claims["username"].(string), Role: claims["role"].(string)}, nil
	}
	return data.User{}, errors.New("Invalid token")
}
