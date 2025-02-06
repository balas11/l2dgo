/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// oneCmd represents the one command
var oneCmd = &cobra.Command{
	Use:   "one [-i | --id] [int] ",
	Short: "Fetch one user by id",
	Long:  `Fetch one user by id`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		user, err := users.Get(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		PrintUser(user)
	},
}

func init() {
	fetchCmd.AddCommand(oneCmd)

	oneCmd.Flags().IntP("id", "i", 0, "User id")
	oneCmd.MarkFlagRequired("id")
}
