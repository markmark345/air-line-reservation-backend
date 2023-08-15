package services

import (
	"air-line-reservation-backend/config"
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/domain/repositories"
	"air-line-reservation-backend/internal/domain/utils"
	"context"
	"time"

	"github.com/google/uuid"
	commomContext "github.com/markmark345/air-line-v1-common/api/contexts"
)

type UserService interface {
	GetUser(ctx context.Context) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) (*entities.Token, error)
}

type userService struct {
	userRepo repositories.UserRepository
	config   *config.Config
}

func NewUserService(
	userRepo repositories.UserRepository,
	config *config.Config,
) UserService {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

func (svc *userService) GetUser(ctx context.Context) (*entities.User, error) {
	userId, err := commomContext.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

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

func (svc *userService) CreateUser(ctx context.Context, user *entities.User) (*entities.Token, error) {
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.UserID = uuid.New()
	user.Password = hashPass
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	err = svc.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(time.Minute * 15).Unix()
	secrets := svc.config.Secrets.JwtKeyAccess

	jwtToken, err := utils.GennerateJWT(user.UserID.String(), user.Email, expirationTime, secrets)
	if err != nil {
		return nil, err
	}

	return &entities.Token{
		Authorization: jwtToken,
		ExpiresIn:     expirationTime,
		IsFirstLogin:  true,
	}, nil
}
