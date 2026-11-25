package user

import (
	"folderstructure/internal/storage"
	"folderstructure/platform/logger"
)

type User struct {
	log      logger.Logger
	userrepo storage.User
}

func NewUser(log logger.Logger,userrepo storage.User) *User {
	return &User{
		log: log,
		userrepo: userrepo,
	}
}
