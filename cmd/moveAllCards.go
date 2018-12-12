package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// moveAllCardsCmd represents the moveAllCards command
var moveAllCardsCmd = &cobra.Command{
	Use:   "move-all-cards",
	Short: "Move all cards from one list to another",
	Long:  `Move all cards from one list to another`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("moveAllCards called")
	},
}

func init() {
	rootCmd.AddCommand(moveAllCardsCmd)
}
