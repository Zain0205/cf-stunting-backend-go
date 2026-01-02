package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name string `gorm:"type:varchar(100);not null"`
}
