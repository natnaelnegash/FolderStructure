package storage

import (
	"context"
	"folderstructure/internal/constant/entity"
)

type User interface {
	CreateUser(ctx context.Context, data entity.User) error
}
