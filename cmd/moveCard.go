package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// moveCardCmd represents the moveCard command
var moveCardCmd = &cobra.Command{
	Use:   "move-card",
	Short: "Move a card on a board",
	Long:  `Move a card on a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("moveCard called")
	},
}

func init() {
	rootCmd.AddCommand(moveCardCmd)
}
