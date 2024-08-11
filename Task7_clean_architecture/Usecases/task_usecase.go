package usecases

import (
	"context"
	"time"

	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) CreateTask(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.CreateTask(ctx, task)
}

func (tu *taskUsecase) Fetch(c context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Fetch(ctx)
}

func (tu *taskUsecase) GetByID(c context.Context, id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetByID(ctx, id)
}

func (tu *taskUsecase) UpdateTask(c context.Context, task *domain.TaskUpdate, id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.UpdateTask(ctx, task, id)
}

func (tu *taskUsecase) DeleteTask(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.DeleteTask(ctx, id)
}
