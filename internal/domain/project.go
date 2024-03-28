package domain

import "context"

type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProjectStorer interface {
	Create(ctx context.Context, project Project) (*Project, error)
}
