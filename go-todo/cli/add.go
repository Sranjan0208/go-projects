package cli

import (
	"Sranjan0208/go-todo/task"
	"Sranjan0208/go-todo/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:  "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string){
		if len(args) < 2 {
			fmt.Println("Usage: kanban add <title> <description>")
			return
		}
		title := args[0]
		description := args[1]
		t := task.NewTask(len(boardInstance.Tasks)+1, title, description)
		boardInstance.AddTask(t)
		utils.SaveTasks(boardInstance.Tasks)
		fmt.Println("✅ Task added successfully!")
	},
}