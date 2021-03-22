package models

import (
    "time"
    "gorm.io/gorm"
)

type BaseTimeField struct {
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

