package types

import "time"

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required,min=5,max=32"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}
