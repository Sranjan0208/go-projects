package task

import (
	"time"
)

type Status string

const (
	ToDo	Status = "Todo"
	Doing	Status = "Doing"
	Done 	Status = "Done"
)

type Task struct {
	ID			int			`json:"id"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Status		Status		`json:"status"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

// New Task 
func NewTask(id int, title, description string) Task {
	return Task{
		ID:	 			id,
		Title: 			title,
		Description: 	description,
		Status: 		ToDo,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}
}

// Update Task updates task details
func (t *Task) UpdateTask(title, description string){
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
}

// MoveTask moves task to a different column
func (t *Task) MoveTask(status Status) {
	t.Status = status
	t.UpdatedAt = time.Now()
}