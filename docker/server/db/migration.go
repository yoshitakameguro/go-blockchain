package db

import (
	"gorm.io/gorm"
    "server/models"
)

func Migrate(DB *gorm.DB) {
	migrator := DB.Debug().Migrator()

	migrator.AutoMigrate(
        &models.User{},
        &models.Wallet{},
	)
}
