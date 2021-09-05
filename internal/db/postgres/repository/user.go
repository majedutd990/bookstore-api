package repository

import (
	"github.com/majedutd990/bookstore-api/internal/models"
	"github.com/majedutd990/bookstore-api/pkg/errors"
	"github.com/majedutd990/bookstore-api/pkg/log"
	"github.com/majedutd990/bookstore-api/pkg/translate/messages"
)

func (r *Repository) CreateUser(user *models.User) error {
	if err := r.Db.Create(user).Error; err != nil {
		r.Logger.Error(log.LogFields{
			Section:  "Repository.user",
			Function: "Create User",
			Params:   user,
			Message:  "",
		})
		return errors.New(errors.KindUnexpected, messages.DBError)
	}
	return nil
}
