package domain

import "context"

type Workspace struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	ImageID int64  `json:"imageID"`
}

type WorkspaceStorer interface {
	CreateWorkspace(ctx context.Context, workspace Workspace) error
}
