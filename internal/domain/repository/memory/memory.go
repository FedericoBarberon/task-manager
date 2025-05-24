package memory

import (
	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
)

type InMemoryTaskRepository map[string]*entity.Task

func NewInMemoryTaskRepository() repository.TaskRepository {
	return InMemoryTaskRepository{}
}

func (m InMemoryTaskRepository) Save(task *entity.Task) error {
	m[task.Id] = task
	return nil
}
func (m InMemoryTaskRepository) Update(task *entity.Task) error {
	if _, ok := m[task.Id]; !ok {
		return repository.ErrTaskNotFound
	}
	m[task.Id] = task
	return nil
}
func (m InMemoryTaskRepository) Delete(id string) error {
	if _, ok := m[id]; !ok {
		return repository.ErrTaskNotFound
	}

	delete(m, id)
	return nil
}
func (m InMemoryTaskRepository) GetById(id string) (*entity.Task, error) {
	if task, ok := m[id]; !ok {
		return nil, repository.ErrTaskNotFound
	} else {
		return task, nil
	}
}
func (m InMemoryTaskRepository) GetAll() ([]*entity.Task, error) {
	var tasks []*entity.Task

	for _, task := range m {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
