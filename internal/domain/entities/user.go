package entities

import (
	"air-line-reservation-backend/internal/domain/utils"

	"github.com/jackc/pgtype"
)

type User struct {
	UserID    pgtype.UUID
	Email     string
	Password  string
	Phone     string
	Region    string
	Gender    utils.NullGender
	Title     string
	FirstName string
	LastName  string
	CreateAt  pgtype.Timestamptz
	UpdateAt  pgtype.Timestamptz
	Age       int8
}

// type User struct {
// 	UserID    pgtype.UUID        `json:"user_id"`
// 	Email     string             `json:"email"`
// 	Password  string             `json:"password"`
// 	Phone     pgtype.Text        `json:"phone"`
// 	Region    pgtype.Text        `json:"region"`
// 	Gender    utils.NullGender   `json:"gender"`
// 	Title     string             `json:"title"`
// 	FirstName string             `json:"first_name"`
// 	LastName  string             `json:"last_name"`
// 	CreateAt  pgtype.Timestamptz `json:"create_at"`
// 	UpdateAt  pgtype.Timestamptz `json:"update_at"`
// 	Age       pgtype.Int2        `json:"age"`
// }
