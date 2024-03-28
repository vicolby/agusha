package repository

import (
	"context"
	"github.com/gimmefear/dswv3/internal/domain"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{db: db}
}

func (r *WorkspaceRepository) CreateWorkspace(ctx context.Context, workspace domain.Workspace) error {
	result := r.db.Create(&workspace)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
