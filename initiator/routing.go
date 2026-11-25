package initiator

import (
	"folderstructure/internal/glue/user"

	"folderstructure/platform/logger"

	"github.com/go-chi/chi/v5"
)

func initRoute(grp chi.Router, handler *Handler, module *Service, logger logger.Logger) {
	user.UserRoute(grp, handler.User, logger)

}
