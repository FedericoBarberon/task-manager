package memory

import (
	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
)

type InMemoryTaskRepository struct {
	tasks  map[entity.TaskID]*entity.Task
	nextId entity.TaskID
}

func NewInMemoryTaskRepository() repository.TaskRepository {
	return &InMemoryTaskRepository{nextId: 1, tasks: make(map[entity.TaskID]*entity.Task)}
}

func (m *InMemoryTaskRepository) Save(task *entity.Task) error {
	task.Id = m.nextId
	m.tasks[m.nextId] = task
	m.nextId++
	return nil
}
func (m *InMemoryTaskRepository) Update(task *entity.Task) error {
	if _, ok := m.tasks[task.Id]; !ok {
		return repository.ErrTaskNotFound
	}
	m.tasks[task.Id] = task
	return nil
}
func (m *InMemoryTaskRepository) Delete(id entity.TaskID) error {
	if _, ok := m.tasks[id]; !ok {
		return repository.ErrTaskNotFound
	}

	delete(m.tasks, id)
	return nil
}
func (m *InMemoryTaskRepository) GetById(id entity.TaskID) (*entity.Task, error) {
	if task, ok := m.tasks[id]; !ok {
		return nil, repository.ErrTaskNotFound
	} else {
		return task, nil
	}
}
func (m *InMemoryTaskRepository) GetAll() ([]*entity.Task, error) {
	var tasks []*entity.Task

	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
