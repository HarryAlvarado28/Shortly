package models

import "time"

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	ShortID     string `gorm:"uniqueIndex"`
	OriginalURL string
	Clicks      int
	CreatedAt   time.Time
	ExpiresAt   *time.Time
}
