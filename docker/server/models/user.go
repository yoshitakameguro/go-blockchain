package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Wallet Wallet
    Email     string      `gorm:"size:256;not null;unique" json:"email"`
}
