package repository

import (
	"context"
	"fmt"

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
	result := r.db.WithContext(ctx).Create(&workspace)

	if result.Error != nil {
		return fmt.Errorf("Workspace creation failed: %s", result.Error)
	}

	return nil
}

func (r *WorkspaceRepository) DeleteWorkspace(ctx context.Context, workspace domain.Workspace) error {
	result := r.db.WithContext(ctx).Delete(&workspace)

	if result.Error != nil {
		return fmt.Errorf("Workspace deletion failed: %s", result.Error)
	}

	return nil
}

func (r *WorkspaceRepository) GetWorkspaceByID(ctx context.Context, workspace_id int64) (domain.Workspace, error) {
	var workspace = domain.Workspace{ID: workspace_id}
	result := r.db.WithContext(ctx).First(&workspace)

	if result.Error != nil {
		return workspace, fmt.Errorf("Workspace getting by ID failed: %s", result.Error)
	}

	return workspace, nil
}

func (r *WorkspaceRepository) GetAllWorkspaces(ctx context.Context) ([]domain.Workspace, error) {
	var workspaces []domain.Workspace
	result := r.db.WithContext(ctx).Find(&workspaces)

	if result.Error != nil {
		return workspaces, fmt.Errorf("Getting all workspaces failed: %s", result.Error)
	}

	return workspaces, nil
}
