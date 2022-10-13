package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"not null;unique_index" json:"username"`
	Password string `gorm:"not null"`
	Task     []Task `json:"task"`
}
