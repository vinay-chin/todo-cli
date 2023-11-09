/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

// updateTaskCmd represents the updateTask command
var updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Update task by ID",
	Long: `This takes an ID as input and update the specific task as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter Task ID: ")
		var id int 
		fmt.Scanf("%d", &id)
		update(id)
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
