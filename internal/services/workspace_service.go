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
