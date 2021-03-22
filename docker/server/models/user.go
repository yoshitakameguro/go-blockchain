package models

import (
    "time"
)

type User struct {
	ID        uint64      `gorm:"primary_key;auto_increment" json:"id"`
    Email     string      `gorm:"size:256;not null;unique" json:"email"`
    PrivateKey  string    `gorm:"size:256;not null;unique" json:"private_key"`
    PublichKey  string    `gorm:"size:256;not null;unique" json:"public_key"`
    BlockchainAddress  string    `gorm:"size:256;not null;unique" json:"blockchain_address"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
