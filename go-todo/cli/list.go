package cli

import (
	"Sranjan0208/go-todo/utils"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:  "list",
	Short: "List tasks by status",
	Run: func(cmd *cobra.Command, args []string){
		utils.PrintBoard(boardInstance.Tasks)
	},
}