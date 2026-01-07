package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"

	"gorm.io/gorm"
)

func SeedDomains(db *gorm.DB) error {
	type seed struct {
		Category string

		Code string

		Name string
	}

	data := []seed{
		{"PRAKONSEPSI", "A", "Gizi dan Suplementasi"},

		{"PRAKONSEPSI", "B", "Ketahanan Pangan"},

		{"PRAKONSEPSI", "C", "Lingkungan dan Perilaku Risiko"},

		{"PRAKONSEPSI", "D", "Psikososial"},

		{"PERNAH_MELAHIRKAN", "A", "Pemberian Makan"},

		{"PERNAH_MELAHIRKAN", "B", "Lingkungan Fisik"},

		{"PERNAH_MELAHIRKAN", "C", "Psikososial"},

		{"REMAJA_19", "A", "Biologis Intergenerasional"},

		{"REMAJA_19", "B", "Pola Makan"},

		{"REMAJA_19", "C", "Infeksi"},

		{"REMAJA_19", "D", "Sanitasi dan Perilaku"},

		{"REMAJA_19", "E", "Ketahanan Pangan"},

		{"REMAJA_19", "F", "Lingkungan Sosial"},
	}

	for _, d := range data {

		var cat models.Category

		db.Where("code = ?", d.Category).First(&cat)

		domain := models.Domain{
			CategoryID: cat.ID,

			Code: d.Code,

			Name: d.Name,
		}

		db.Where("category_id = ? AND code = ?", cat.ID, d.Code).
			FirstOrCreate(&domain)

	}

	return nil
}
