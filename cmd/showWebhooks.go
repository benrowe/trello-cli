package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showWebhooksCmd represents the showWebhooks command
var showWebhooksCmd = &cobra.Command{
	Use:   "show-webhooks",
	Short: "Display webhooks for current user applications",
	Long:  `Display webhooks for current user applications`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showWebhooks called")
	},
}

func init() {
	rootCmd.AddCommand(showWebhooksCmd)
}
