package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Topic  string `gorm:"type:varchar(255);not null"`
	Detail string `gorm:"not null"`
}
