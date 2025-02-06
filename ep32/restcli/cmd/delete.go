/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i | --id] [id]",
	Short: "deletes a record by id",
	Long: `Deletes a user record by id and returns error if any
	Examples:
	./restcli delete -i 2`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		err = users.Delete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("User with id %d deleted\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("id", "i", 0, "User id")
	deleteCmd.MarkFlagRequired("id")
}
