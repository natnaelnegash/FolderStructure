package service

import (
	"context"
	"folderstructure/internal/dto/user"
)

type User interface {
	Users(ctx context.Context, req user.UserRequest) (user.UserResponse, error)
}
