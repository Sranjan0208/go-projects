package cli

import (
	"Sranjan0208/go-todo/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:  "update",
	Short: "Update a task's title and description",
	Run: func(cmd *cobra.Command, args []string){
		if len(args) < 3 {
			fmt.Println("Usage: kanban update <id> <new title> <new description>")
			return
		}

		id, _ := strconv.Atoi(args[0])
		title := args[1]
		description := args[2]

		task, err := boardInstance.GetTaskByID(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		task.UpdateTask(title, description)
		utils.SaveTasks(boardInstance.Tasks)
		fmt.Println("✅ Task updated successfully!")
	},
}