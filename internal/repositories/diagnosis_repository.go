package repositories

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"gorm.io/gorm"
)

type DiagnosisRepository struct {
	DB *gorm.DB
}

func (r *DiagnosisRepository) CreateDiagnosis(
	d *models.Diagnosis,
	answers []models.DiagnosisAnswer,
	domains []models.DiagnosisDomain,
) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(d).Error; err != nil {
			return err
		}

		for i := range answers {
			answers[i].DiagnosisID = d.ID
		}

		for i := range domains {
			domains[i].DiagnosisID = d.ID
		}

		if err := tx.Create(&answers).Error; err != nil {
			return err
		}

		if err := tx.Create(&domains).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *DiagnosisRepository) GetByUserID(userID uint) ([]models.Diagnosis, error) {
	var diagnoses []models.Diagnosis

	err := r.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Preload("Answers").
		Preload("Domains").
		Find(&diagnoses).Error

	return diagnoses, err
}
