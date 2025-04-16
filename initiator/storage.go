package initiator

import (
	"helphub-backend/internal/constants/dbinstance"
	"helphub-backend/internal/storage"
	"helphub-backend/internal/storage/sqlc/user"
	"helphub-backend/platform/logger"
)

type Persistence struct {
	user storage.User
}

func InitPersistence(db dbinstance.DBInstance, log logger.Logger) Persistence {
	return Persistence{
		user: user.Init(db, log.Named("user-persistence")),
	}
}
