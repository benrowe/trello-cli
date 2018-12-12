package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCardsCmd represents the showCards command
var showCardsCmd = &cobra.Command{
	Use:   "show-cards",
	Short: "Show the cards on a list",
	Long:  `Show the cards on a list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showCards called")
	},
}

func init() {
	rootCmd.AddCommand(showCardsCmd)
}
