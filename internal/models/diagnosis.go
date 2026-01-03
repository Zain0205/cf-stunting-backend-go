package models

import "time"

type Diagnosis struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Category  string
	Result    string
	CreatedAt time.Time

	Answers []DiagnosisAnswer `gorm:"foreignKey:DiagnosisID"`
	Domains []DiagnosisDomain `gorm:"foreignKey:DiagnosisID"`
}
