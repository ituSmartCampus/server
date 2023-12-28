package models

import (
	"github.com/hlandau/passlib"
)

type User struct {
	Base
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
}

// func (user *User) SetPassword(db *gorm.DB) (err error) {

// 	hashedPassword, err := passlib.Hash(user.Password)
// 	db.Model(user).Update("Password", hashedPassword)
// 	return err
// }

func (user *User) CheckPassword(password string) error {
	err := passlib.VerifyNoUpgrade(password, user.Password)
	return err
}
