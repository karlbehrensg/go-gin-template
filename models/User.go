package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"not null;uniqueIndex" json:"username"`
	Name     string `json:"name"`
	Password string `gorm:"not null"`
}
