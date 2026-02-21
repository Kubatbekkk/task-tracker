package task

import (
	"errors"
	"time"
)

func AddTask(description string) (Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return Task{}, err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now()

	newTask := Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	return newTask, err
}

func UpdateTask(id int, description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("Task not found")
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTasks(tasks)
		}
	}

	return errors.New("Task not found")
}

func ChangeStatus(id int, status Status) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("Task Not Found")
}

func ListTasks(filter Status) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	if filter == "" {
		return tasks, nil
	}

	var filtered []Task
	for _, t := range tasks {
		if t.Status == filter {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}
