package services

import (
	"context"
	"fmt"

	"github.com/gimmefear/dswv3/internal/domain"
)

type WorkspaceService struct {
	workspaceStorer domain.WorkspaceStorer
    k8sService K8sService
}

func NewWorkspaceService(workspaceStorer domain.WorkspaceStorer, k8sService K8sService) *WorkspaceService {
    return &WorkspaceService{workspaceStorer: workspaceStorer, k8sService: k8sService}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, workspace domain.Workspace) error {
    err := s.workspaceStorer.CreateWorkspace(ctx, workspace)

    if err != nil {
        return fmt.Errorf("Error while creating a workspace write to DB: %s", err)
    }
    err = s.k8sService.createNamespace(workspace.Name)

    if err != nil {
        return fmt.Errorf("Error while creating a namespace: %s", err)
    }

    return nil
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
