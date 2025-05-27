package repository

import (
	"errors"

	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
)

var (
	ErrTaskNotFound = errors.New("task doesnt exists")
)

type TaskRepository interface {
	Save(task *entity.Task) error
	Update(task *entity.Task) error
	Delete(id entity.TaskID) error
	GetById(id entity.TaskID) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
}
