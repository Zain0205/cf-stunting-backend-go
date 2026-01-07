package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{Code: "PRAKONSEPSI", Name: "Perempuan Prakonsepsi"},

		{Code: "PERNAH_MELAHIRKAN", Name: "Perempuan Pernah Melahirkan"},

		{Code: "REMAJA_19", Name: "Remaja 19 Tahun"},
	}

	for _, c := range categories {
		if err := db.Where("code = ?", c.Code).FirstOrCreate(&c).Error; err != nil {
			return err
		}
	}

	return nil
}
