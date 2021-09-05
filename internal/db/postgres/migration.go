package postgres

import (
	"github.com/majedutd990/bookstore-api/internal/models"
	"github.com/majedutd990/bookstore-api/pkg/errors"
	"github.com/majedutd990/bookstore-api/pkg/log"
	"github.com/majedutd990/bookstore-api/pkg/translate/messages"
)

func (d *database) migration() error {
	if err := d.db.AutoMigrate(
		&models.User{},
		&models.Picture{},
		&models.Wallet{},
		&models.Category{},
		&models.Comment{},
		&models.Book{},
	); err != nil {
		d.logger.Error(log.LogFields{
			Section:  "postgres.migration",
			Function: "migration",
			Params:   nil,
			Message:  err.Error(),
		})
		return errors.New(errors.KindUnexpected, messages.DBError)
	}
	return nil
}
