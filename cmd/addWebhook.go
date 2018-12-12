package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addWebhookCmd represents the addWebhook command
var addWebhookCmd = &cobra.Command{
	Use:   "add-webhook",
	Short: "Add a webhook to a board",
	Long:  `Add a webhook to a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addWebhook called")
	},
}

func init() {
	rootCmd.AddCommand(addWebhookCmd)
}
