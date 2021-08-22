package application

import (
	"github.com/majedutd990/bookstore-api/internal/config"
	"github.com/majedutd990/bookstore-api/pkg/transloator/i18n"
)

func Run(cfg *config.Config) error {
	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	_ = translator
	return err
}
