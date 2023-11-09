/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"bufio"
	"os"
	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "add-task",
	Short: "Add a tittle and description for your task",
	Long: `Give your task tittle and description to add your Task into the list`,
	Run: func(cmd *cobra.Command, args []string) {
		var title, description string

		fmt.Print("Title: ")
		in := bufio.NewReader(os.Stdin)
		title, err := in.ReadString('\n')

		fmt.Print("Description: ")
		in = bufio.NewReader(os.Stdin)
		description, err = in.ReadString('\n')

		if err != nil {
			fmt.Println(("error at reading"))
		}

		fmt.Println(title + description)

		msg := add(title,description)

		fmt.Println(msg)

	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
