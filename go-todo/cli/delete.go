package cli

import (
	"Sranjan0208/go-todo/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: kanban delete <id>")
			return
		}
		id, _ := strconv.Atoi(args[0])

		err := boardInstance.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		utils.SaveTasks(boardInstance.Tasks)
		fmt.Println("🗑️ Task deleted successfully!")
	},
}
