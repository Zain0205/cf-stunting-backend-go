package models

type Question struct {
	ID       uint `gorm:"primaryKey"`
	DomainID uint `gorm:"not null"`

	Code string `gorm:"type:varchar(50);not null"`
	Text string `gorm:"type:text;not null"`

	// Bobot keyakinan pakar (0.0 – 1.0)
	CFPakar float64 `gorm:"type:decimal(3,2);not null"`

	// Jika true → CF dibalik (untuk pertanyaan negatif)
	IsReverse bool `gorm:"default:false"`

	Domain Domain `gorm:"foreignKey:DomainID"`
}
