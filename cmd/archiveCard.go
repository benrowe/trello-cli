package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// archiveCardCmd represents the archiveCard command
var archiveCardCmd = &cobra.Command{
	Use:   "archive-card",
	Short: "Archive a card from a board",
	Long:  `Archive a card from a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("archiveCard called")
	},
}

func init() {
	rootCmd.AddCommand(archiveCardCmd)
}
