package repo

import (
	"context"
	"test.com/project-project/internal/data/task"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int) ([]task.MsTaskStagesTemplate, error)
}
