package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
	Email          string
	Role           int `gorm:"default:0"` //默认0 普通用户
}

const passwordCost = 12

func (user *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(hashed)
	return nil
}

func (user *User) ComparePassword(password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
