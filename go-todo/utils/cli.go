package utils

import (
	"Sranjan0208/go-todo/task"

	"github.com/fatih/color"
)



// PrintBoard prints all tasks by status with colors
func PrintBoard(tasks []task.Task){
	statusColors := map[task.Status]*color.Color{
		task.ToDo:  color.New(color.FgBlue),
		task.Doing: color.New(color.FgYellow),
		task.Done: color.New(color.FgGreen),
	}

	for _,t := range tasks {
		c := statusColors[t.Status]
		c.Printf("[%s] %d: %s - %s\n", t.Status, t.ID, t.Title, t.Description)
	}
}