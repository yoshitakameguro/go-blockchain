package db

import (
    "github.com/bxcodec/faker/v3"
    "server/models"
)

func FakeUserWithWallet() *models.User {
    user := models.User{}
    faker.FakeData(&user)
    user.Wallet = models.NewWallet()
    user.Wallet.Amount = 10
    DB.Create(&user)
    return &user
}
