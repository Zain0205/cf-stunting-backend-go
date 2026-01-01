package models

type Question struct {
	ID        uint `gorm:"primaryKey"`
	DomainID  uint
	Code      string
	Text      string
	CFPakar   float64 `gorm:"type:decimal(3,2)"`
	IsReverse bool
}
