package test

import (
    "os"
    "fmt"
    "log"
    "testing"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    . "server/db"
    _ "github.com/joho/godotenv/autoload"
)


var r *gin.Engine = gin.Default()

func TestMain(m *testing.M) {
    gin.SetMode(gin.TestMode)
    initDB()
    os.Exit(m.Run())
}

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
