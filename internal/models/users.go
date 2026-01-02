package models

import "time"

type UserCategory string

const (
	CategoryPrakonsepsi UserCategory = "PRAKONSEPSI"
	CategoryMelahirkan  UserCategory = "PERNAH_MELAHIRKAN"
	CategoryRemaja      UserCategory = "REMAJA_19"
)

type User struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"type:varchar(100);not null"`
	PhoneNumber string       `gorm:"type:varchar(20);uniqueIndex;not null"`
	Password    string       `gorm:"not null"`
	Category    UserCategory `gorm:"type:enum('PRAKONSEPSI','PERNAH_MELAHIRKAN','REMAJA_19');not null"`
	CreatedAt   time.Time
}
