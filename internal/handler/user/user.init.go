package user

import (
	"folderstructure/internal/service"
	"folderstructure/platform/logger"
)

type User struct {
	log     logger.Logger
	service service.User
}
func NewUser(log     logger.Logger,service service.User)*User{
	return &User{
		log:log,
		service: service,
	}
}
