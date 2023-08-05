package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Region    string    `json:"region"`
	Gender    string    `json:"gender"`
	Title     string    `json:"title"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	Age       int8      `json:"age"`
}
