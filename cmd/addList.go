package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addListCmd represents the addList command
var addListCmd = &cobra.Command{
	Use:   "add-list",
	Short: "Add a new list to the spcified board with the specified name",
	Long:  `Add a new list to the spcified board with the specified name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addList called")
	},
}

func init() {
	rootCmd.AddCommand(addListCmd)
}
