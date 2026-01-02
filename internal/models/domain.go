package models

type Domain struct {
	ID         uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"not null"`

	Code string `gorm:"type:varchar(50);not null"`
	Name string `gorm:"type:varchar(100);not null"`

	Category Category `gorm:"foreignKey:CategoryID"`
}
