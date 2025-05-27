package mappers

import (
	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/dto"
)

func TaskToDTO(task entity.Task) dto.TaskDTO {
	return dto.TaskDTO{
		Id:        task.Id,
		Name:      task.Name,
		Completed: task.Completed,
	}
}
