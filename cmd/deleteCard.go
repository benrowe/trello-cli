package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCardCmd represents the deleteCard command
var deleteCardCmd = &cobra.Command{
	Use:   "delete-card",
	Short: "Remove a card from a board",
	Long:  `Remove a card from a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteCard called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCardCmd)
}
