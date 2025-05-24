package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var ErrEmptyName = errors.New("name cannot be empty")

type Task struct {
	Id        string
	Name      string
	Completed bool
	CreatedAt time.Time
}

func NewTask(name string) (*Task, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, ErrEmptyName
	}

	return &Task{
		Id:        uuid.New().String(),
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
	}, nil
}

func (t *Task) MarkCompleted() {
	t.Completed = true
}

func (t *Task) UpdateName(newName string) error {
	newName = strings.TrimSpace(newName)

	if newName == "" {
		return ErrEmptyName
	}

	t.Name = newName

	return nil
}
