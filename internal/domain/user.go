package domain

import "context"

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStorer interface {
	CreateUser(ctx context.Context, user User) (*User, error)
}
