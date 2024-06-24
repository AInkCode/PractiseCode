package repo

import (
	"context"
	"test.com/project-project/internal/data/menu"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*menu.ProjectMenu, error)
}
