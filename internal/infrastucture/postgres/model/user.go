package model

import (
	"air-line-reservation-backend/internal/domain/entities"
	utils "air-line-reservation-backend/internal/domain/utils"

	"github.com/jackc/pgtype"
)

type GetUser struct {
	UserID    pgtype.UUID        `json:"user_id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Phone     pgtype.Text        `json:"phone"`
	Region    pgtype.Text        `json:"region"`
	Gender    utils.NullGender   `json:"gender"`
	Title     string             `json:"title"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	CreateAt  pgtype.Timestamptz `json:"create_at"`
	UpdateAt  pgtype.Timestamptz `json:"update_at"`
	Age       pgtype.Int2        `json:"age"`
}

type User struct {
	UserID    pgtype.UUID        `pg:"user_id, pk"`
	Email     string             `pg:"email"`
	Password  string             `pg:"password"`
	Phone     pgtype.Text        `pg:"phone"`
	Region    pgtype.Text        `pg:"region"`
	Gender    utils.NullGender   `pg:"gender"`
	Title     string             `pg:"title"`
	FirstName string             `pg:"first_name"`
	LastName  string             `pg:"last_name"`
	CreateAt  pgtype.Timestamptz `pg:"create_at"`
	UpdateAt  pgtype.Timestamptz `pg:"update_at"`
	Age       pgtype.Int2        `pg:"age"`
}

func (u *User) ToDomain() *entities.User {
	return &entities.User{
		UserID:    u.UserID,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone.String,
		Region:    u.Region.String,
		Gender:    u.Gender,
		Title:     u.Title,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       int8(u.Age.Int),
		CreateAt:  u.CreateAt,
		UpdateAt:  u.UpdateAt,
	}
}
