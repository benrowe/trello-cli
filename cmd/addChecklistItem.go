package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addChecklistItemCmd represents the addChecklistItem command
var addChecklistItemCmd = &cobra.Command{
	Use:   "add-checklist-item",
	Short: "Add a checklist item to a checklist",
	Long:  `Add a checklist item to a checklist`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addChecklistItem called")
	},
}

func init() {
	rootCmd.AddCommand(addChecklistItemCmd)
}
