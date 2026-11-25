package user
import (
	"context"
	"folderstructure/internal/constant"
	"folderstructure/internal/dto/user"
	"folderstructure/internal/constant/errors"

	"go.uber.org/zap"
)

func (l *User) Users(ctx context.Context, req user.UserRequest) (user.UserResponse, error) {
	logCtx := constant.SERVICE_USER
	if req.Name == "" {
		l.log.Named(logCtx).Error(ctx, "Invalid payload", zap.Any("data", req))
		return user.UserResponse{}, errors.ErrInvalidUserInput.New("Invalid Payload")
	}

	return user.UserResponse{}, nil

}
