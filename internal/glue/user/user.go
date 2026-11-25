package user

import (
	"folderstructure/internal/glue/mainroute"
	"folderstructure/internal/handler"
	"folderstructure/platform/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoute(group chi.Router, h handler.User, log logger.Logger) {
	routes := []mainroute.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/v1/user",
			Handler: h.User,
		},
	}

	mainroute.RegisterRoute(group, routes, log)
}
