// Package models
package models

type AnswerMapping struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint `gorm:"not null;uniqueIndex:idx_question_answer"`

	// contoh: "0", "1", "SANGAT_YAKIN", "TIDAK"
	AnswerKey string `gorm:"type:varchar(50);not null;uniqueIndex:idx_question_answer"`

	// contoh: "Setiap hari", "4–6 hari/minggu"
	Label string `gorm:"type:varchar(100);not null"`

	// Bobot evidence (0.0 – 1.0)
	CFEvidence float64 `gorm:"type:decimal(3,2);not null"`

	Question Question `gorm:"foreignKey:QuestionID"`
}
