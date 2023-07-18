package sqlc_test

import (
	"context"
	"strconv"
	"testing"

	"air-line-reservation-backend/internal/domain/utils"
	"air-line-reservation-backend/internal/infrastucture/postgres"
	"air-line-reservation-backend/internal/infrastucture/postgres/model"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

var arg = postgres.CreateUserParams {
	Email:     utils.RandomEmail(),
	Password:  utils.RandomString(8),
	Phone:     pgtype.Text{String: strconv.FormatInt(utils.RandomInt(1000000000, 9999999999), 10), Valid: true},
	Region:    pgtype.Text{String: utils.RandomString(5), Valid: true},
	Gender:    model.NullGender{Gender: model.Gender(utils.RandomGender()), Valid: true},
	Title:     utils.RandomString(3),
	FirstName: utils.RandomString(10),
	LastName:  utils.RandomString(10),
	Age:       pgtype.Int2{Int16: int16(utils.RandomInt(0, 100)), Valid: true},
}

var uid pgtype.UUID

func TestCreateUser(t *testing.T) {
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Region, user.Region)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Title, user.Title)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Age, user.Age)

	require.NotZero(t, user.CreateAt)
	require.NotZero(t, user.UpdateAt)
	require.NotZero(t, user.UserID)
}


func TestGetUser(t *testing.T) {
	params := postgres.GetUsersParams {
		Email: arg.Email,
		Password: arg.Password,
	}

	user, err := testQueries.GetUsers(context.Background(), params)

	userParams := model.User{
		UserID: user.UserID,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}

	uid = pgtype.UUID{
        Bytes: user.UserID.Bytes,
        Valid: true,
    }

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, params.Email, user.Email)
	require.Equal(t, params.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Region, user.Region)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Title, user.Title)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, userParams.CreateAt, user.CreateAt)
	require.Equal(t, userParams.UpdateAt, user.UpdateAt)
	require.Equal(t, userParams.UserID, user.UserID)
}

func TestDeleteUser(t *testing.T) {
	err2 := testQueries.DeleteUser(context.Background(), uid)
	require.NoError(t, err2)

	getUser := postgres.GetUsersParams {
		Email: arg.Email,
		Password: arg.Password,
	}

	user, err := testQueries.GetUsers(context.Background(), getUser)
	require.Error(t, err)
	require.Empty(t, user)
}