package models

import (
	"time"
)

type (
	Base struct {
		ID        uint `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
)