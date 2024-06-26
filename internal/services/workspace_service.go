package services

import (
	"context"
	"fmt"

	"github.com/gimmefear/dswv3/internal/domain"
)

type WorkspaceService struct {
	workspaceStorer domain.WorkspaceStorer
	k8sService      K8sService
	logger          domain.Logger
}

func NewWorkspaceService(workspaceStorer domain.WorkspaceStorer, k8sService K8sService, logger domain.Logger) *WorkspaceService {
	return &WorkspaceService{workspaceStorer: workspaceStorer, k8sService: k8sService, logger: logger}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, workspace domain.Workspace) error {
	err := s.workspaceStorer.CreateWorkspace(ctx, workspace)

	if err != nil {
		return fmt.Errorf("Error while creating a workspace write to DB: %s", err)
	}

	s.logger.Info("Creating a deployment for workspace", "workspace", workspace.Name)
	err = s.k8sService.createDeployment(ctx, workspace)

	if err != nil {
		return fmt.Errorf("Error while creating a deployment: %s", err)
	}

	s.logger.Info("Deployment created")

	return nil
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, workspace domain.Workspace) error {
	s.logger.Info("Deleting workspace", "workspace", workspace.Name)
	err := s.workspaceStorer.DeleteWorkspace(ctx, workspace)

	if err != nil {
		return fmt.Errorf("Error deleting workspace from DB: %s", err)
	}

	err = s.k8sService.deleteNamespace(ctx, fmt.Sprintf("project-%d", workspace.ProjectID))

	if err != nil {
		return fmt.Errorf("Error deleting namespace: %s", err)
	}

	return nil
}

func (s *WorkspaceService) GetWorkspaceByID(ctx context.Context, workspace_id int64) (domain.Workspace, error) {
	return s.workspaceStorer.GetWorkspaceByID(ctx, workspace_id)
}

func (s *WorkspaceService) GetAllWorkspaces(ctx context.Context) ([]domain.Workspace, error) {
	return s.workspaceStorer.GetAllWorkspaces(ctx)
}
