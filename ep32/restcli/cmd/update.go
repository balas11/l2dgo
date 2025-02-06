/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [-i | --id] [id][-n | --name] [name] [-j | --job] [job]",
	Short: "update a user record with name and job",
	Long: `Update an existing user by id with a name and job.
	Examples:
	./restcli update -i 2 -n "John Doe" -j "Programmer"`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		job, err := cmd.Flags().GetString("job")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		user, err := users.Update(id, users.EdtUser{Name: name, Job: job})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		printUpdatedUser(user)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("name", "n", "", "Name of the user")
	updateCmd.Flags().StringP("job", "j", "", "Job of the user")
	updateCmd.Flags().IntP("id", "i", 0, "User id")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("job")
	updateCmd.MarkFlagRequired("id")
}

func printUpdatedUser(user users.UpdRespUser) {
	fmt.Println("Name : ", user.Name)
	fmt.Println("Job : ", user.Job)
	fmt.Println("Updated At : ", user.UpdatedAt)
}
