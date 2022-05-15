package dto

import "time"

type CreateUserDTO struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UserDto struct {
	ID               int64     `uri:"id" json:"id" binding:"required,min=1"`
	Email            string    `json:"email"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	PasswordChangeAt time.Time `json:"password_change_at"`
}
