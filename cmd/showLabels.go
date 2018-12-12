package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showLabelsCmd represents the showLabels command
var showLabelsCmd = &cobra.Command{
	Use:   "show-labels",
	Short: "Show labels defined on a board",
	Long:  `Show labels defined on a board`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showLabels called")
	},
}

func init() {
	rootCmd.AddCommand(showLabelsCmd)
}
