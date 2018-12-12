package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addBoardCmd represents the addBoard command
var addBoardCmd = &cobra.Command{
	Use:   "add-board",
	Short: "Add a new board with the specified name",
	Long:  `Add a new board with the specified name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addBoard called")
	},
}

func init() {
	rootCmd.AddCommand(addBoardCmd)
}
