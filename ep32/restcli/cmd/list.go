/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [-p | --page] [int] ",
	Short: "Fetches a list of users",
	Long:  `Fetches a page of users with ot without specifying the page number.`,
	Run: func(cmd *cobra.Command, args []string) {
		page, err := cmd.Flags().GetInt("page")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		users, err := users.List(page)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		PrintUserList(users)
	},
}

func init() {
	fetchCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("page", "p", 1, "Page number")
}
