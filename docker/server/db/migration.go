package db

import (
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	migrator := DB.Debug().Migrator()

	migrator.AutoMigrate(
	)
}
