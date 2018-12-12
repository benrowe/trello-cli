package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addChecklistCmd represents the addChecklist command
var addChecklistCmd = &cobra.Command{
	Use:   "add-checklist",
	Short: "Add a new checklist to the specified card",
	Long:  `Add a new checklist to the specified card`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addChecklist called")
	},
}

func init() {
	rootCmd.AddCommand(addChecklistCmd)
}
