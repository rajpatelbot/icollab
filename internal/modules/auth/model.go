package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID         uint   `gorm:"primaryKey; autoIncrement"`
	Email      string `gorm:"uniqueIndex; not null"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Location   string `gorm:"not null"`
	Avatar_URL string `gorm:"not null"`
}
