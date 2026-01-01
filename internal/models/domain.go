package models

type Domain struct {
	ID         uint `gorm:"primaryKey"`
	CategoryID uint
	Code       string
	Name       string
}
