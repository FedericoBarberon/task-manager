package entity_test

import (
	"testing"

	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTask_Succesfull(t *testing.T) {
	name := "prueba"
	task, err := entity.NewTask(name)

	assert.NoError(t, err)
	assert.Equal(t, name, task.Name)
	assert.False(t, task.Completed, "task should not be completed by default")
	assert.NotZero(t, task.CreatedAt, "task CreatedAt should be set to current time")
}

func TestNewTask_ErrorEmptyName(t *testing.T) {
	_, err := entity.NewTask("  ")

	if assert.Error(t, err) {
		assert.Equal(t, entity.ErrEmptyName, err)
	}
}

func TestMarkCompleted(t *testing.T) {
	task, _ := entity.NewTask("prueba")

	assert.False(t, task.Completed)
	task.MarkCompleted()
	assert.True(t, task.Completed, "task should be marked as completed")
}

func TestUpdateName_Succesfull(t *testing.T) {
	task, _ := entity.NewTask("prueba")

	newName := "nuevo"
	err := task.UpdateName(newName)

	assert.NoError(t, err)
	assert.Equal(t, newName, task.Name)
}

func TestUpdateName_ErrorEmptyName(t *testing.T) {
	task, _ := entity.NewTask("prueba")

	err := task.UpdateName("  ")

	if assert.Error(t, err) {
		assert.Equal(t, entity.ErrEmptyName, err)
	}
}
