package models

type DiagnosisDomain struct {
	ID          uint `gorm:"primaryKey"`
	DiagnosisID uint
	DomainCode  string
	CFValue     float64 `gorm:"type:decimal(4,3)"`
}
