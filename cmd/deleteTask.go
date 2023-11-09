/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteTaskCmd represents the deleteTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "delete-task",
	Short: "Delete task by ID",
	Long: `Give your task ID to Delete your Task`,
	Run: func(cmd *cobra.Command, args []string) {
		var id int;
		fmt.Print("Enter your ID:")
		fmt.Scanf("%d", &id)
		deleteT(id)
		fmt.Println("Task Deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
