package models

type User struct {
    BaseField
    Wallet Wallet
    Email     string      `gorm:"size:256;not null;unique" json:"email" faker:"email"`
}
