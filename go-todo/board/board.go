package board

import (
	"Sranjan0208/go-todo/task"
	"errors"
)

type Board struct {
	Tasks []task.Task
}

// AddTask adds a new task
func (b *Board) AddTask(t task.Task){
	b.Tasks = append(b.Tasks, t)
}

// ListTasks lists tasks by status
func (b *Board) ListTasks(status task.Status) []task.Task{
	var filteredTasks []task.Task
	for _,t := range b.Tasks {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

// GetTaskByID finds task by ID
func (b *Board) GetTaskByID(id int) (*task.Task, error){
	for i,t := range b.Tasks {
		if t.ID == id {
			return &b.Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

// DeleteTask removes a task by ID
func (b *Board) DeleteTask(id int) error {
	for i, t := range b.Tasks {
		if t.ID == id {
			b.Tasks = append(b.Tasks[:i], b.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}