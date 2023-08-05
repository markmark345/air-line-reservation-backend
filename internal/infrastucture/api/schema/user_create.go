package schema

type CreateUser struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Phone     string `json:"phone" validate:"required,len=10"`
	Region    string `json:"region"`
	Gender    string `json:"gender" validate:"oneof=M N W"`
	Title     string `json:"title" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Age       int    `json:"age" validate:"required"`
}
