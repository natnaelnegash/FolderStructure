package users

import (
	"context"
	"folderstructure/internal/constant/entity"
	"folderstructure/internal/constant/errors"
	"folderstructure/internal/constant/model/db"
	"folderstructure/internal/constant/model/persistencedb"

	"folderstructure/platform/logger"

	"github.com/google/uuid"
)

type Users struct {
	log           logger.Logger
	persistencedb persistencedb.PersistenceDB
}

func NewUsers(log logger.Logger, persistencedb persistencedb.PersistenceDB) *Users {
	return &Users{
		log:           log,
		persistencedb: persistencedb,
	}

}

func (u *Users) CreateUser(ctx context.Context, data entity.User) error {
	err := u.persistencedb.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:   uuid.New(),
		Name: data.Name,
	})
	if err != nil {
		return errors.ErrInvalidUserInput.Wrap(err, "Unable to create.")
	}
	return nil
}
