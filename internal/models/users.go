package models

import "github.com/majedutd990/bookstore-api/internal/models/types"

type (
	User struct {
		Base
		UserName              string `gorm:"unique; not null"`
		FirstName             string
		LastName              string
		Email                 string
		IsEmailVerified       bool
		PhoneNumber           string
		IsPhoneNumberVerified bool
		Gender                types.Gender
		Role                  types.Role
		Avatar                string
		Token                 string
	}
)
