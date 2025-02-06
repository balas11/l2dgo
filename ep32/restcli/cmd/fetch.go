/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch [list | one] [flags]",
	Short: "Fetches one or a list of users",
	Long: `Fetches one user with an id or list of users by specifying a page number
	
		Examples:
		./restcli fetch list
		./restcli fetch list -p 2
		./restcli fetch one -i 2`,
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func PrintUserList(users users.RetUserList) {
	fmt.Println("Page : ", users.Page)
	fmt.Println("Per Page : ", users.PerPage)
	fmt.Println("Total : ", users.Total)
	fmt.Println("Total Pages : ", users.TotalPages)
	fmt.Println("Data : [")
	for _, data := range users.Data {
		fmt.Println("{")
		fmt.Println("Id : ", data.Id)
		fmt.Println("Email : ", data.Email)
		fmt.Println("First Name : ", data.FirstName)
		fmt.Println("Last Name : ", data.LastName)
		fmt.Println("Avatar : ", data.Avatar)
		fmt.Println("}")
	}
	fmt.Println("]")
	fmt.Println("Support : ", users.Support.Url, "-", users.Support.Text)
}

func PrintUser(user users.RetUserOne) {
	fmt.Println("Id : ", user.Data.Id)
	fmt.Println("Email : ", user.Data.Email)
	fmt.Println("First Name : ", user.Data.FirstName)
	fmt.Println("Last Name : ", user.Data.LastName)
	fmt.Println("Avatar : ", user.Data.Avatar)
	fmt.Println("Support : ", user.Support.Url, "-", user.Support.Text)
}
