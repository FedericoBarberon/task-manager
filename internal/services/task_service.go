package services

import (
	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
	"github.com/FedericoBarberon/task-manager/internal/dto"
	"github.com/FedericoBarberon/task-manager/internal/mappers"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo}
}

func (ts *TaskService) GetTasks() ([]dto.TaskDTO, error) {
	tasks, err := ts.repo.GetAll()

	if err != nil {
		return nil, err
	}

	tasksDTOs := make([]dto.TaskDTO, len(tasks))

	for i := range tasks {
		tasksDTOs[i] = mappers.TaskToDTO(*tasks[i])
	}

	return tasksDTOs, nil
}

func (ts *TaskService) GetTaskById(id entity.TaskID) (dto.TaskDTO, error) {
	task, err := ts.repo.GetById(id)

	if err != nil {
		return dto.TaskDTO{}, err
	}

	return mappers.TaskToDTO(*task), nil
}

func (ts *TaskService) CreateTask(name string) (entity.TaskID, error) {
	task, err := entity.NewTask(name)

	if err != nil {
		return 0, err
	}

	err = ts.repo.Save(task)

	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func (ts *TaskService) MarkAsCompleted(id entity.TaskID) error {
	task, err := ts.repo.GetById(id)

	if err != nil {
		return err
	}

	task.MarkCompleted()
	err = ts.repo.Update(task)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) DeleteTask(id entity.TaskID) error {
	err := ts.repo.Delete(id)
	return err
}
