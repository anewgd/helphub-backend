package storage

import (
	"context"
	"helphub-backend/internal/constants/model/dto"

	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, param dto.RegisterUser) (*dto.User, error)
	Update(ctx context.Context, id uuid.UUID, param dto.UpdateUser) (*dto.User, error)
	CheckUserExists(ctx context.Context, param dto.RegisterUser) (bool, error)
	GetAll(ctx context.Context) ([]dto.User, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}
