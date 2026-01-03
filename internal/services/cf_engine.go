package services

import "errors"

func CalculateCFItem(cfPakar, cfEvidence float64) float64 {
	return cfPakar * cfEvidence
}

func CombineCF(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	result := values[0]
	for i := 1; i < len(values); i++ {
		result = result + values[i]*(1-result)
	}
	return result
}

func CountAbove(domains map[string]float64, threshold float64) int {
	count := 0
	for _, v := range domains {
		if v >= threshold {
			count++
		}
	}
	return count
}

/*
Risk Evaluation – SESUAI PDF
*/
func EvaluateRisk(category string, cf map[string]float64) (string, error) {
	switch category {

	// ================= PRAKONSEPSI =================
	case "PRAKONSEPSI":
		A, B, C, D := cf["A"], cf["B"], cf["C"], cf["D"]

		if A >= 0.70 && B >= 0.70 && C >= 0.70 && D >= 0.70 {
			return "SANGAT TINGGI", nil
		}

		if A >= 0.70 && B >= 0.70 && (C >= 0.50 || D >= 0.50) {
			return "TINGGI", nil
		}

		if A >= 0.50 && B >= 0.50 && C < 0.70 && D < 0.70 {
			return "SEDANG", nil
		}

		if (A >= 0.50 || B >= 0.50 || C >= 0.50 || D >= 0.50) &&
			A < 0.70 && B < 0.70 && C < 0.70 && D < 0.70 {
			return "RINGAN", nil
		}

		if A < 0.30 && B < 0.30 && C < 0.30 && D < 0.30 {
			return "RENDAH", nil
		}

	// ================= PERNAH MELAHIRKAN =================
	case "PERNAH_MELAHIRKAN":
		A, B, C := cf["A"], cf["B"], cf["C"]

		if A >= 0.70 && B >= 0.70 && C >= 0.70 {
			return "SANGAT TINGGI", nil
		}

		if CountAbove(cf, 0.70) >= 2 {
			return "TINGGI", nil
		}

		if (A >= 0.50 || B >= 0.50 || C >= 0.50) &&
			A < 0.70 && B < 0.70 && C < 0.70 {
			return "SEDANG", nil
		}

		if (A >= 0.30 || B >= 0.30 || C >= 0.30) &&
			A < 0.50 && B < 0.50 && C < 0.50 {
			return "RINGAN", nil
		}

		if A < 0.30 && B < 0.30 && C < 0.30 {
			return "RENDAH", nil
		}

	// ================= REMAJA 19 =================
	case "REMAJA_19":
		if CountAbove(cf, 0.70) == 6 {
			return "SANGAT TINGGI", nil
		}

		if cf["A"] >= 0.70 && cf["B"] >= 0.70 && CountAbove(cf, 0.70) >= 3 {
			return "TINGGI", nil
		}

		if (cf["A"] >= 0.50 || cf["B"] >= 0.50) && CountAbove(cf, 0.50) >= 3 {
			return "SEDANG", nil
		}

		if CountAbove(cf, 0.50) >= 1 && CountAbove(cf, 0.70) == 0 {
			return "RINGAN", nil
		}

		if CountAbove(cf, 0.30) == 0 {
			return "RENDAH", nil
		}
	}

	return "", errors.New("tidak dapat menentukan risiko")
}
