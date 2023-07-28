package postgres

import (
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/domain/repositories"
	"air-line-reservation-backend/internal/infrastucture/postgres/model"
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type userRepository struct {
	pg *pg.DB
	// redis     *redis.Client
	// cacheTTL  time.Duration
	// cfRedis   *config.Redis
	// cfSecrets *config.Secrets
	// logger    logger.Logger
}

func NewUserRepository(
	db *pg.DB,
) repositories.UserRepository {
	return &userRepository{
		db,
	}
}

func (repo *userRepository) GetUser(ctx context.Context, userId string) (*entities.User, error) {
	// user = &entities.User{}
	pgUser := &model.User{}

	err := repo.pg.ModelContext(ctx, pgUser).Where("user_id = ?", userId).First()
	if err == pg.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	result := pgUser.ToDomain()
	fmt.Println(result)
	// result.UserID = userId

	return result, nil
}
