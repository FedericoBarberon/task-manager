package repository_test

import (
	"testing"

	"github.com/FedericoBarberon/task-manager/internal/domain/entity"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository/memory"
	"github.com/FedericoBarberon/task-manager/internal/domain/repository/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepository(t *testing.T) {
	factories := []repoFactory{
		{
			name: "in memory",
			init: memoryRepo,
		},
		// {
		// 	name: "sqlite",
		// 	init: sqliteRepo,
		// },
	}

	for _, factory := range factories {
		t.Run(factory.name+"/SaveAndGet", func(t *testing.T) {
			repo := factory.init()
			testTaskRepoSaveAndGet(t, repo)
			testTaskRepoGetNonExistingTask(t, repo)
		})
		t.Run(factory.name+"/Update", func(t *testing.T) {
			repo := factory.init()
			testTaskRepoUpdate(t, repo)
			testTaskRepoUpdateNonExistingTask(t, repo)
		})
		t.Run(factory.name+"/GetAll", func(t *testing.T) {
			repo := factory.init()
			testTaskRepoGetAll(t, repo)
		})
		t.Run(factory.name+"/Delete", func(t *testing.T) {
			repo := factory.init()
			testTaskRepoDelete(t, repo)
			testTaskRepoDeleteNonExistingTask(t, repo)
		})
	}
}

func testTaskRepoSaveAndGet(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	task, _ := entity.NewTask("prueba")
	err := taskRepo.Save(task)

	assert.NoError(t, err)

	got, err := taskRepo.GetById(task.Id)

	assert.NoError(t, err)

	assert.Equal(t, task, got)
}

func testTaskRepoGetNonExistingTask(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	_, err := taskRepo.GetById("non-existing-task-id")

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrTaskNotFound, err)
	}
}

func testTaskRepoUpdate(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	task, _ := entity.NewTask("prueba")
	err := taskRepo.Save(task)

	assert.NoError(t, err)

	newName := "new name"
	task.UpdateName(newName)

	err = taskRepo.Update(task)

	assert.NoError(t, err)

	got, err := taskRepo.GetById(task.Id)

	assert.NoError(t, err)
	assert.Equal(t, newName, got.Name)
}

func testTaskRepoUpdateNonExistingTask(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	newTask, _ := entity.NewTask("non-existing-task")
	err := taskRepo.Update(newTask)

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrTaskNotFound, err)
	}
}

func testTaskRepoGetAll(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()
	task1, _ := entity.NewTask("1")
	task2, _ := entity.NewTask("1")
	task3, _ := entity.NewTask("1")
	tasks := []*entity.Task{task1, task2, task3}

	taskRepo.Save(task1)
	taskRepo.Save(task2)
	taskRepo.Save(task3)

	got, err := taskRepo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, tasks, got)
}

func testTaskRepoDelete(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	task, _ := entity.NewTask("prueba")
	taskRepo.Save(task)

	err := taskRepo.Delete(task.Id)

	assert.NoError(t, err)

	_, err = taskRepo.GetById(task.Id)

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrTaskNotFound, err)
	}
}

func testTaskRepoDeleteNonExistingTask(t testing.TB, taskRepo repository.TaskRepository) {
	t.Helper()

	err := taskRepo.Delete("non-existing-task-id")

	if assert.Error(t, err) {
		assert.Equal(t, repository.ErrTaskNotFound, err)
	}
}

type repoFactory struct {
	name string
	init func() repository.TaskRepository
}

func memoryRepo() repository.TaskRepository {
	return memory.NewInMemoryTaskRepository()
}

func sqliteRepo() repository.TaskRepository {
	return sqlite.NewSqliteTaskRepository("in memory db connection")
}
