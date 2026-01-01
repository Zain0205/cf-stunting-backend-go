package models

type DiagnosisAnswer struct {
	ID           uint `gorm:"primaryKey"`
	DiagnosisID  uint
	QuestionCode string
	AnswerKey    string
	CFItem       float64 `gorm:"type:decimal(4,3)"`
}
