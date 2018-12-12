package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteWebhookCmd represents the deleteWebhook command
var deleteWebhookCmd = &cobra.Command{
	Use:   "delete-webhook",
	Short: "Remove a card from a board",
	Long:  `Remove a card from a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteWebhook called")
	},
}

func init() {
	rootCmd.AddCommand(deleteWebhookCmd)
}
