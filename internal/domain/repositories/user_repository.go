package repositories

import (
	"air-line-reservation-backend/internal/domain/entities"
	"context"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId string) (*entities.User, error)
}
