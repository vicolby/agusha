package domain

import "context"

type Workspace struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ImageID   int64  `json:"imageID"`
	ProjectID int64  `json:"ProjectID"`
}

type WorkspaceStorer interface {
	CreateWorkspace(ctx context.Context, workspace Workspace) error
	DeleteWorkspace(ctx context.Context, workspace Workspace) error
	GetWorkspaceByID(ctx context.Context, workspace_id int64) (Workspace, error)
	GetAllWorkspaces(ctx context.Context) ([]Workspace, error)
}
