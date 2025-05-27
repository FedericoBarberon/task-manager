package entity

import (
	"errors"
	"strings"
	"time"
)

var ErrEmptyName = errors.New("name cannot be empty")

type TaskID int64

func (id TaskID) IsValid() bool {
	return id > 0
}

type Task struct {
	Id        TaskID
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
