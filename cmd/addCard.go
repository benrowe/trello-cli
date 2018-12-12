package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCardCmd represents the addCard command
var addCardCmd = &cobra.Command{
	Use:   "add-card",
	Short: "Add a card to a board",
	Long:  `Add a card to a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addCard called")
	},
}

func init() {
	rootCmd.AddCommand(addCardCmd)
}
