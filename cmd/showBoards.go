package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showBoardsCmd represents the showBoards command
var showBoardsCmd = &cobra.Command{
	Use:   "show-boards",
	Short: "Show the list of boards",
	Long: `Show the list of cached boards.
	
In order to get the most up to date boards, run the refresh command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showBoards called")
	},
}

func init() {
	rootCmd.AddCommand(showBoardsCmd)
}
