package repositories

import (
	"air-line-reservation-backend/internal/domain/entities"
	"context"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId string) (*entities.User, error)
	GetUserWithEmailAndPassword(ctx context.Context, email string, password string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
}
