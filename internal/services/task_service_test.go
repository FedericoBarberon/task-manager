package services_test

import (
	"testing"

	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository/memory"
	"github.com/FedericoBarberon/task-manager/internal/dto"
	"github.com/FedericoBarberon/task-manager/internal/mappers"
	"github.com/FedericoBarberon/task-manager/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	repo := memory.NewInMemoryTaskRepository()

	task1, _ := entity.NewTask("task 1")
	task2, _ := entity.NewTask("task 2")
	task3, _ := entity.NewTask("task 3")
	repo.Save(task1)
	repo.Save(task2)
	repo.Save(task3)

	taskService := services.NewTaskService(repo)

	tasks, err := taskService.GetTasks()

	assert.NoError(t, err)

	want := []dto.TaskDTO{
		{Id: 1, Name: "task 1", Completed: false},
		{Id: 2, Name: "task 2", Completed: false},
		{Id: 3, Name: "task 3", Completed: false},
	}

	assert.ElementsMatch(t, want, tasks)
}

func TestGetTaskById(t *testing.T) {
	repo := memory.NewInMemoryTaskRepository()
	task, _ := entity.NewTask("task")
	repo.Save(task)

	taskService := services.NewTaskService(repo)
	got, err := taskService.GetTaskById(task.Id)
	want := mappers.TaskToDTO(*task)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestCreateTask(t *testing.T) {
	taskService := services.NewTaskService(memory.NewInMemoryTaskRepository())
	id, err := taskService.CreateTask("task")

	assert.NoError(t, err)

	got, _ := taskService.GetTaskById(id)
	want := dto.TaskDTO{Id: id, Name: "task", Completed: false}

	assert.Equal(t, want, got)
}

func TestDeleteTask(t *testing.T) {
	taskService := services.NewTaskService(memory.NewInMemoryTaskRepository())
	id, _ := taskService.CreateTask("task")

	err := taskService.DeleteTask(id)

	assert.NoError(t, err)

	_, err = taskService.GetTaskById(id)

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrTaskNotFound, err)
	}
}
