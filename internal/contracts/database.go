package contracts

import "github.com/majedutd990/bookstore-api/internal/models"

type (
	Repository interface {
		UserRepository
	}
	UserRepository interface {
		CreateUser(user *models.User) error
	}
)
