package initiator

import (
	"helphub-backend/internal/handler/rest"
	"helphub-backend/internal/handler/rest/gin/user"
	"helphub-backend/platform/logger"
	"time"
)

type Handler struct {
	user rest.User
}

func InitHandler(module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		user: user.Init(log.Named("user-handler"), module.user, timeout),
	}
}
