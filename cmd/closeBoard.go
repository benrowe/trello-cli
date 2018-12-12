package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// closeBoardCmd represents the closeBoard command
var closeBoardCmd = &cobra.Command{
	Use:   "close-board",
	Short: "Closes those board(s) where the specified text occurs in their name",
	Long:  `Closes those board(s) where the specified text occurs in their name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("closeBoard called")
	},
}

func init() {
	rootCmd.AddCommand(closeBoardCmd)
}
