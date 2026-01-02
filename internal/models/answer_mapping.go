// Package models
package models

type AnswerMapping struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint `gorm:"not null"`

	// contoh: SANGAT_YAKIN, YAKIN, RAGU, TIDAK
	AnswerKey string `gorm:"type:varchar(50);not null"`

	// Bobot evidence (0.0 – 1.0)
	CFEvidence float64 `gorm:"type:decimal(3,2);not null"`

	Question Question `gorm:"foreignKey:QuestionID"`
}
