package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseField struct {
	ID uint `gorm:"primaryKey" faker:"-"`
	BaseTimeField
}

type BaseTimeField struct {
	CreatedAt time.Time      `faker:"-"`
	UpdatedAt time.Time      `faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
