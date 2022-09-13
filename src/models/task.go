package models

type Task struct {
	TaskId int32 `json:"id"`
	TaskName string `json:"task_name"`
	Description string `json:"description"`
	IsComplete bool `json:"is_complete"`
}
