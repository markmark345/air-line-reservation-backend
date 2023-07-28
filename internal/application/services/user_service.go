package services

import (
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/domain/repositories"
	"context"
	"fmt"
)

type UserService interface {
	GetUser(ctx context.Context, userId string) (*entities.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(
	userRepo repositories.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (svc *userService) GetUser(ctx context.Context, userId string) (*entities.User, error) {
	result, err := svc.userRepo.GetUser(ctx, userId)
	fmt.Println("service: ", result)
	if err != nil {
		panic(err)
	}

	return result, err
}
