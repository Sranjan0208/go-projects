package utils

import (
	"Sranjan0208/go-todo/task"
	"encoding/json"
	"fmt"
	"os"
)

const storageFile = "tasks.json"

// Save Tasks saves tasks to a file
func SaveTasks(tasks []task.Task){
	file,err := os.Create(storageFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return 
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}


// LoadTasks loads tasks from the file
func LoadTasks() []task.Task {
	file, err := os.Open(storageFile)
	if err != nil {
		return []task.Task{}
	}
	defer file.Close()

	var tasks []task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		return []task.Task{}
	}

	return tasks
}