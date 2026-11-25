package initiator

import (
	"folderstructure/internal/handler"
	"folderstructure/internal/handler/user"

	"folderstructure/platform/logger"
)

type Handler struct {
	User handler.User
}

func initHandler(service *Service, log logger.Logger) *Handler {
	user := user.NewUser(log, service.User)

	return &Handler{
		User: user,
	}
}
