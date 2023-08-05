package postgres

import (
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/domain/repositories"
	utils "air-line-reservation-backend/internal/domain/utils"
	"air-line-reservation-backend/internal/infrastucture/postgres/model"
	"context"
	"errors"

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
	pgUser := &model.User{}

	err := repo.pg.ModelContext(ctx, pgUser).Where("user_id = ?", userId).First()
	if err == pg.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	result := pgUser.ToDomain()

	return result, nil
}

func (repo *userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	tx, err := repo.pg.Begin()
	if err != nil {
		return err
	}
	defer tx.Close()

	pgUser := &model.User{}
	pgUser.Email = user.Email
	pgUser.Password = user.Password
	pgUser.Phone = user.Phone
	pgUser.Region = user.Region
	pgUser.Gender = utils.NullGender{Gender: utils.Gender(user.Gender), Valid: true}
	pgUser.Title = user.Title
	pgUser.FirstName = user.FirstName
	pgUser.LastName = user.LastName
	pgUser.Age = int(pgUser.Age)

	_, err = repo.pg.ModelContext(ctx, pgUser).WherePK().Insert()

	if err == pg.ErrNoRows {
		tx.Rollback()
		return errors.New("create user error")
	}
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
