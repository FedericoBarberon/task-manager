package dto

import "github.com/FedericoBarberon/task-manager/internal/domain/entity"

type TaskDTO struct {
	Id        entity.TaskID
	Name      string
	Completed bool
}
