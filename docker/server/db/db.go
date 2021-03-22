package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

func Init() {

	var err error

	os.Setenv("TZ", "Asia/Tokyo")
	dbEntrypoint := fmt.Sprintf("tcp(%s:%s)", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	if os.Getenv("DB_SOCKET") != "" {
		dbEntrypoint = fmt.Sprintf("unix(%s)", os.Getenv("DB_SOCKET"))
	}
	DNS := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbEntrypoint,
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", os.Getenv("DB_DRIVER"))
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", os.Getenv("DB_DRIVER"))
	}

	Migrate()

	// Create initial user data if needed
	var count int64
	DB.Table("users").Count(&count)
	if count == 0 {
		CreateInitialUserData()
	}
}
