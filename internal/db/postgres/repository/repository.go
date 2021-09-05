package repository

import (
	"github.com/majedutd990/bookstore-api/internal/config"
	"github.com/majedutd990/bookstore-api/pkg/log"
	"github.com/majedutd990/bookstore-api/pkg/translate"
	"gorm.io/gorm"
)

type Repository struct {
	Db         *gorm.DB
	Cfg        *config.Config
	Translator translate.Translator
	Logger     log.Logger
}
