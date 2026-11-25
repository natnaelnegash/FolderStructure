package user

import (
	"encoding/json"
	"folderstructure/internal/constant"
	"folderstructure/internal/constant/errors"
	"folderstructure/internal/constant/response"
	"folderstructure/internal/dto/user"
	"net/http"

	"go.uber.org/zap"
)

func (userHandler *User) User(w http.ResponseWriter, r *http.Request) {
	var ctxName = constant.HANDLER_USER

	var req user.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		userHandler.log.Named(ctxName).Error(r.Context(), "failed to decode request", zap.Error(err))
		response.SendErrorResponseFormated(w, errors.ErrInvalidUserInput.Wrap(err, "invalid request body"))
		return
	}
	res, err := userHandler.service.Users(r.Context(), req)
	if err != nil {
		response.SendErrorResponseFormated(w, err)
		return
	}

	userHandler.log.Named(ctxName).Info(r.Context(), "Successfully fetched cases.")

	response.SendSuccessResponse(w, http.StatusOK, response.USER_CREATED, res, "user created successfully!")

}
