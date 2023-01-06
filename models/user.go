package models

import (
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `gorm:"type:bool;default:false"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.Password = hashPassword(u.Password)
	return nil
}

func hashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))

	if err != nil {
		panic(err)
	}

	return string(encoded)
}
