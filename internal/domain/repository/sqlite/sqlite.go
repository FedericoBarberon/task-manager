package sqlite

import (
	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
)

type SqliteTaskRepository struct {
	conn string
}

func NewSqliteTaskRepository(conn string) repository.TaskRepository {
	return SqliteTaskRepository{conn}
}

func (s SqliteTaskRepository) Save(task *entity.Task) error {
	return nil
}
func (s SqliteTaskRepository) Update(task *entity.Task) error {
	return nil
}
func (s SqliteTaskRepository) Delete(id string) error {
	return nil
}
func (s SqliteTaskRepository) GetById(id string) (*entity.Task, error) {
	return nil, nil
}
func (s SqliteTaskRepository) GetAll() ([]*entity.Task, error) {
	return nil, nil
}
