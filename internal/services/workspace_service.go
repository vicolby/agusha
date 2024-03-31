package services

import (
	"context"

	"github.com/gimmefear/dswv3/internal/domain"
)

type WorkspaceService struct {
	workspaceStorer domain.WorkspaceStorer
}

func NewWorkspaceService(workspaceStorer domain.WorkspaceStorer) *WorkspaceService {
	return &WorkspaceService{workspaceStorer: workspaceStorer}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, workspace domain.Workspace) error {
	return s.workspaceStorer.CreateWorkspace(ctx, workspace)
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, workspace domain.Workspace) error {
	return s.workspaceStorer.DeleteWorkspace(ctx, workspace)
}

func (s *WorkspaceService) GetWorkspaceByID(ctx context.Context, workspace_id int64) (domain.Workspace, error) {
	return s.workspaceStorer.GetWorkspaceByID(ctx, workspace_id)
}

func (s *WorkspaceService) GetAllWorkspaces(ctx context.Context) ([]domain.Workspace, error) {
	return s.workspaceStorer.GetAllWorkspaces(ctx)
}
