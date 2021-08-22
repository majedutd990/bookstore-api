package application

import (
	"github.com/majedutd990/bookstore-api/internal/config"
	"github.com/majedutd990/bookstore-api/pkg/logger/logrus"
	"github.com/majedutd990/bookstore-api/pkg/translator/i18n"
)

func Run(cfg *config.Config) error {

	logger, err := logrus.New(
		cfg.Log.InternalPath,
		cfg.Log.FileNamePattern,
		cfg.Log.MaxAge,
		cfg.Log.RotationTime,
		cfg.Log.MaxSize,
	)
	if err != nil {
		return err
	}
	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	_ = translator
	_ = logger
	return nil
}

// How to Use Logger
// log.Error(logger.LogFields{
// 	Section:  "application.go",
// 	Function: "Runs",
// 	Params:   "Config",
// 	Message:  err.Error(),
// })
