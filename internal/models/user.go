package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"` // guardado como hash
	CreatedAt time.Time
	Urls      []URL `gorm:"foreignKey:UserID"`
}
