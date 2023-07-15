package sqlc_test

import (
	"context"
	"testing"

	db "air-line-reservation-backend/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// userID := uuid.New()

	arg := db.CreateUserParams{
		Email:     "user01@gmail.com",
		Password:  "password",
		Phone:     pgtype.Text{},
		Region:    pgtype.Text{String: "thai"},
		Gender:    db.NullGender{Gender: "M"},
		Title:     "Mr.",
		FirstName: "testf",
		LastName:  "testl",
		Age:       pgtype.Int2{Int16: 24},
	}

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