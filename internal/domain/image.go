package domain

import "context"

type Image struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type ImageStorer interface {
	GetByID(ctx context.Context, id int64) (*Image, error)
}
