package main

import (
	"fmt"
	"log"
	"restcli/users"
)

func main() {
	// user := users.EdtUser{
	// 	Name: "morpheus",
	// 	Job:  "leader",
	// }

	err := users.Delete(2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("User Deleted")
}

// func getUser(id int) {
// 	user, err := users.Get(id)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Printf("%+v\n", user)
// }

// func listUsers(page int) {
// 	users, err := users.List(page)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("Page : ", users.Page)
// 	fmt.Println("Per Page : ", users.PerPage)
// 	fmt.Println("Total : ", users.Total)
// 	fmt.Println("Total Pages : ", users.TotalPages)
// 	for _, data := range users.Data {
// 		fmt.Println("Id : ", data.Id)
// 		fmt.Println("Email : ", data.Email)
// 		fmt.Println("First Name : ", data.FirstName)
// 		fmt.Println("Last Name : ", data.LastName)
// 		fmt.Println("Avatar : ", data.Avatar)
// 	}
// 	fmt.Println("Support : ", users.Suppor.Url, users.Suppor.Text)
// }

// func createUser(user users.EdtUser) {
// 	resp, err := users.Create(user)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Printf("%+v\n", resp)
// }

// func updateUser(id int, user users.EdtUser) {
// 	resp, err := users.Update(id, user)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Printf("%+v\n", resp)
// }
