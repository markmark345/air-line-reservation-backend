package sqlc_test

import (
	"context"
	"strconv"
	"testing"

	db "air-line-reservation-backend/db/sqlc"
	"air-line-reservation-backend/internal/domain/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

var arg = db.CreateUserParams {
	Email:     utils.RandomEmail(),
	Password:  utils.RandomString(8),
	Phone:     pgtype.Text{String: strconv.FormatInt(utils.RandomInt(1000000000, 9999999999), 10), Valid: true},
	Region:    pgtype.Text{String: utils.RandomString(5), Valid: true},
	Gender:    db.NullGender{Gender: db.Gender(utils.RandomGender()), Valid: true},
	Title:     utils.RandomString(3),
	FirstName: utils.RandomString(10),
	LastName:  utils.RandomString(10),
	Age:       pgtype.Int2{Int16: int16(utils.RandomInt(0, 100)), Valid: true},
}

var userParams = db.User {
	Email:     arg.Email,
	Password:  arg.Password,
	Phone:     arg.Phone,
	Region:    arg.Region,
	Gender:    arg.Gender,
	Title:     arg.Title,
	FirstName: arg.FirstName,
	LastName:  arg.LastName,
	Age:       arg.Age,
}

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
	params := db.GetUsersParams {
		Email: arg.Email,
		Password: arg.Password,
	}

	user, err := testQueries.GetUsers(context.Background(), params)

	userParams := db.User{
		UserID: user.UserID,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
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
	uid := pgtype.UUID{Bytes: userParams.UserID.Bytes, Valid: true}
	// fmt.Println("uid", uid.String())

	err := testQueries.DeleteUser(context.Background(), uid)
	// fmt.Println("aaaa", err)
	require.NoError(t, err)

	// getUser := db.GetUsersParams {
	// 	Email: arg.Email,
	// 	Password: arg.Password,
	// }

	// user2, err := testQueries.GetUsers(context.Background(), getUser)
	// fmt.Println("ssss",user)
	// require.Error(t, err)
	// require.Empty(t, user)
	// require.Nil(t, user)
}