package db

import (
    "server/models"
)

func Migrate() {
	migrator := DB.Debug().Migrator()

	migrator.AutoMigrate(
        &models.User{},
        &models.Wallet{},
	)
}
