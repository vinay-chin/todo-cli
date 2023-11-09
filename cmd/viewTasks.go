/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"github.com/spf13/cobra"
)

// viewTasksCmd represents the viewTasks command
var viewTasksCmd = &cobra.Command{
	Use:   "view-tasks",
	Short: "View all Tasks",
	Long: `Shows all the tasks in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		viewall()
	},
}

func init() {
	rootCmd.AddCommand(viewTasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
