/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restcli/users"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-n | --name] [name] [-j | --job] [job]",
	Short: "Creates a new user",
	Long: ` Create a new user with a name and job.
	Examples:
	./restcli create -n "John Doe" -j "Programmer"
	`,
	Run: func(cmd *cobra.Command, args []string) {
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
		user, err := users.Create(users.EdtUser{Name: name, Job: job})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		printCreatedUser(user)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the user")
	createCmd.Flags().StringP("job", "j", "", "Job of the user")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("job")
}

func printCreatedUser(user users.CrtRespUser) {
	fmt.Println("Id : ", user.Id)
	fmt.Println("Name : ", user.Name)
	fmt.Println("Job : ", user.Job)
	fmt.Println("Created At : ", user.CreatedAt)
}
