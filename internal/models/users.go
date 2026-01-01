package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Category  string `gorm:"type:enum('PRAKONSEPSI','PERNAH_MELAHIRKAN','REMAJA_19')"`
	CreatedAt time.Time
}
