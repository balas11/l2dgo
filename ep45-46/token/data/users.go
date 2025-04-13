package data

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users []User

func NewUser(username, password, role string) User {
	hPassword := hashPassword(password)
	return User{UserName: username, Password: hPassword, Role: role}
}

func FindUser(username string) (User, error) {

	for _, user := range Users {
		if user.UserName == username {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

func GetRole(username string) string {
	user, _ := FindUser(username)
	return user.Role
}

func init() {
	Users = append(Users, NewUser("admin", "admin", "admin"))
	Users = append(Users, NewUser("user", "user", "user"))
	for _, user := range Users {
		fmt.Println(user)
	}
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginUser(username, password string) (User, error) {
	user, _ := FindUser(username)
	if CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return User{}, errors.New("Invalid Login")
}
