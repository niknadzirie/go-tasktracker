package service

import (
	"tasktracker/internal/helper"
	"tasktracker/internal/model"
	"tasktracker/internal/repository"
	"time"
)

type taskService struct {
	repo  repository.TaskRepository
	tasks []model.Task
}

func NewTaskSerivce(repo repository.TaskRepository) *taskService {
	tasks, err := repo.ReadAllFromJsonFile()
	if err != nil {
		return nil
	}
	return &taskService{repo: repo, tasks: tasks}
}

func (s *taskService) AddTask(desc string, status model.Status) error {
	newTask := &model.Task{
		ID:          len(s.tasks) + 1,
		Description: desc,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks = append(s.tasks, *newTask)

	err := s.repo.WriteAllToJsonFile(s.tasks)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) TestDisplay() {
	helper.Display(s.tasks)
}
