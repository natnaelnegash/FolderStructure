package initiator

import (
	"folderstructure/internal/service"
	"folderstructure/internal/service/user"
	"folderstructure/platform/logger"
)

type Service struct {
	User service.User
}

func initService(persistence *Persistence, log logger.Logger, pool WorkerPool) *Service {

	user := user.NewUser(log, persistence.User)
	return &Service{
		User: user,
	}
}
