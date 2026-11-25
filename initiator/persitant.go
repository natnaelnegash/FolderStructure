package initiator

import (
	"folderstructure/internal/constant/model/persistencedb"
	"folderstructure/internal/storage"
	"folderstructure/internal/storage/users"

	"folderstructure/platform/logger"

	"github.com/go-redis/redis/v8"
)

type Persistence struct {
	User storage.User
}

func initPersistence(persistencdb *persistencedb.PersistenceDB, log logger.Logger, redisClient *redis.Client) *Persistence {
	user := users.NewUsers(log, *persistencdb)

	return &Persistence{
		User: user,
	}
}
