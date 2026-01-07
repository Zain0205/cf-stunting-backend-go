package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"

	"gorm.io/gorm"
)

func SeedAnswerMappings(db *gorm.DB) error {
	// =============================
	// LABEL STANDAR FREKUENSI
	// =============================

	frequencyLabels := map[string]string{
		"0": "Setiap hari",

		"1": "4–6 hari/minggu",

		"2": "1–3 hari/minggu",

		"3": "Tidak pernah",
	}

	normalCF := map[string]float64{
		"0": 0.0,

		"1": 0.4,

		"2": 0.7,

		"3": 1.0,
	}

	reversedCF := map[string]float64{
		"0": 1.0,

		"1": 0.7,

		"2": 0.4,

		"3": 0.0,
	}

	// =============================

	// SPECIAL ANSWER SET

	// =============================

	special := map[string]map[string]struct {
		Label string

		CF float64
	}{
		// PRAKONSEPSI C4

		"PRAKONSEPSI_C4": {
			"0": {"Air minum terlindung (PDAM/kemasan)", 0.0},

			"2": {"Air tidak terlindung / air permukaan", 1.0},
		},

		// PRAKONSEPSI C5

		"PRAKONSEPSI_C5": {
			"0": {"Jamban pribadi layak", 0.0},

			"1": {"Jamban komunal layak", 0.4},

			"2": {"Tidak punya jamban / tidak layak", 1.0},
		},

		// YA / TIDAK / TIDAK TAHU

		"YA_TIDAK_TT": {
			"0": {"Ya", 1.0},

			"1": {"Tidak", 0.0},

			"2": {"Tidak tahu", 0.4},
		},

		// YA / TIDAK (RISIKO)

		"YA_TIDAK_RISK": {
			"0": {"Ya", 0.8},

			"1": {"Tidak", 0.0},
		},
	}

	var questions []models.Question

	db.Preload("Domain").Preload("Domain.Category").Find(&questions)

	for _, q := range questions {

		category := q.Domain.Category.Code

		domain := q.Domain.Code

		code := q.Code

		// ====== PRAKONSEPSI SPECIAL ======

		if category == "PRAKONSEPSI" && code == "C4" {

			for k, v := range special["PRAKONSEPSI_C4"] {
				createAnswer(db, q.ID, k, v.Label, v.CF)
			}

			continue

		}

		if category == "PRAKONSEPSI" && code == "C5" {

			for k, v := range special["PRAKONSEPSI_C5"] {
				createAnswer(db, q.ID, k, v.Label, v.CF)
			}

			continue

		}

		// ====== REMAJA ======

		if category == "REMAJA_19" && domain == "A" {

			for k, v := range special["YA_TIDAK_TT"] {
				createAnswer(db, q.ID, k, v.Label, v.CF)
			}

			continue

		}

		if category == "REMAJA_19" {

			for k, v := range special["YA_TIDAK_RISK"] {
				createAnswer(db, q.ID, k, v.Label, v.CF)
			}

			continue

		}

		// ====== DEFAULT FREKUENSI ======

		for k, label := range frequencyLabels {

			cf := normalCF[k]

			if q.IsReverse {
				cf = reversedCF[k]
			}

			createAnswer(db, q.ID, k, label, cf)

		}

	}

	return nil
}

func createAnswer(db *gorm.DB, qID uint, key, label string, cf float64) {
	a := models.AnswerMapping{
		QuestionID: qID,

		AnswerKey: key,

		Label: label,

		CFEvidence: cf,
	}

	db.Where("question_id = ? AND answer_key = ?", qID, key).
		FirstOrCreate(&a)
}
