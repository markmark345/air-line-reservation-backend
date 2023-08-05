package services

import (
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/domain/repositories"
	"air-line-reservation-backend/internal/domain/utils"
	"context"
)

type UserService interface {
	GetUser(ctx context.Context, userId string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
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
	if err != nil {
		return nil, err
	}

	return &entities.User{
		UserID:    result.UserID,
		Email:     result.Email,
		Phone:     result.Phone,
		Region:    result.Region,
		Gender:    result.Gender,
		Title:     result.Title,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Age:       result.Age,
		CreateAt:  result.CreateAt,
		UpdateAt:  result.UpdateAt,
	}, err
}

func (svc *userService) CreateUser(ctx context.Context, user *entities.User) error {
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPass
	err = svc.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return err
}
