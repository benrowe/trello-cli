package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cardDetailsCmd represents the cardDetails command
var cardDetailsCmd = &cobra.Command{
	Use:   "card-details",
	Short: "Show details about a specified card",
	Long:  `Show details about a specified card`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cardDetails called")
	},
}

func init() {
	rootCmd.AddCommand(cardDetailsCmd)
}
