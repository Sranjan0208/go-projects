package cli

import (
	"Sranjan0208/go-todo/utils"
	"Sranjan0208/go-todo/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

)

var MoveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a task to another column",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: kanban move <id> <todo|doing|done>")
			return
		}

		id, _ := strconv.Atoi(args[0])
		status := args[1]

		// Get the task from boardInstance
		taskInstance, err := boardInstance.GetTaskByID(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		var newStatus task.Status
		switch status {
		case "todo":
			newStatus = task.ToDo
		case "doing":
			newStatus = task.Doing
		case "done":
			newStatus = task.Done
		default:
			fmt.Println("❌ Invalid status! Use 'todo', 'doing', or 'done'.")
			return
		}

		// Move the task to the new status
		taskInstance.MoveTask(newStatus)
		utils.SaveTasks(boardInstance.Tasks)
		fmt.Println("🚀 Task moved successfully!")
	},
}
