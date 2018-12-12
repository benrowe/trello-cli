package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// assignedToMeCmd represents the assignedToMe command
var assignedToMeCmd = &cobra.Command{
	Use:   "assigned-to-me",
	Short: "Show cards that are currently assigned to yourself, or any member specified",
	Long:  `Show cards that are currently assigned to yourself, or any member specified`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("assignedToMe called")
	},
}

func init() {
	rootCmd.AddCommand(assignedToMeCmd)
}
