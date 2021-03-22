package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	. "server/db"
	"server/models"
)

var R *gin.Engine = gin.Default()

func initDB() {
	var err error
	DNS := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"password",
		"mysql-test",
		"3306",
		"go-blockchain-test",
	)

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", "mysql")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", "mysql")
	}

	Migrate()
}

func InitTest() {
	gin.SetMode(gin.TestMode)
	initDB()
}

func ClearDB() {
	DB.Where("1=1").Delete(&models.User{})
	DB.Where("1=1").Delete(&models.Wallet{})
}

func Request(requestMethod string, endpoint string, data io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(requestMethod, endpoint, data)
	res := httptest.NewRecorder()
	R.ServeHTTP(res, req)
	return res, err
}
