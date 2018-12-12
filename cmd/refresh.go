package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh all your board/list names",
	Long:  `Refresh all your board/list names`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("refresh called")
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
