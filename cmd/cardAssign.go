package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cardAssignCmd represents the cardAssign command
var cardAssignCmd = &cobra.Command{
	Use:   "card-assign",
	Short: "Add or remove a member to a card",
	Long:  `Add or remove a member to a card`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cardAssign called")
	},
}

func init() {
	rootCmd.AddCommand(cardAssignCmd)
}
