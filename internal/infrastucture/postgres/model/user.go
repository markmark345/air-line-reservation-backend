package model

import (
	"air-line-reservation-backend/internal/domain/entities"
	utils "air-line-reservation-backend/internal/domain/utils"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID        `pg:"user_id, pk"`
	Email     string           `pg:"email"`
	Password  string           `pg:"password"`
	Phone     string           `pg:"phone"`
	Region    string           `pg:"region"`
	Gender    utils.NullGender `pg:"gender"`
	Title     string           `pg:"title"`
	FirstName string           `pg:"first_name"`
	LastName  string           `pg:"last_name"`
	CreateAt  time.Time        `pg:"create_at"`
	UpdateAt  time.Time        `pg:"update_at"`
	Age       int              `pg:"age"`
}

func (u *User) ToDomain() *entities.User {
	return &entities.User{
		UserID:    u.UserID,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone,
		Region:    u.Region,
		Gender:    string(u.Gender.Gender),
		Title:     u.Title,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       int8(u.Age),
		CreateAt:  u.CreateAt,
		UpdateAt:  u.UpdateAt,
	}
}
