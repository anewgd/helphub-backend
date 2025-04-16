package initiator

import (
	"helphub-backend/internal/module"
	"helphub-backend/internal/module/user"
	"helphub-backend/platform/logger"
)

type Module struct {
	user module.User
}

func InitModule(persistence Persistence, log logger.Logger) Module {
	return Module{
		user: user.Init(log.Named("user-module"), persistence.user),
	}
}
