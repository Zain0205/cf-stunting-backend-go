package models

type AnswerMapping struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint
	AnswerKey  string
	CFEvidence float64 `gorm:"type:decimal(3,2)"`
}
