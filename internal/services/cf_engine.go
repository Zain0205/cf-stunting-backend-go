package services

import "errors"

// ================= CF CALCULATION =================

// CF item = CFpakar * CFevidence
func CalculateCFItem(cfPakar, cfEvidence float64) float64 {
	return cfPakar * cfEvidence
}

// Combine CF dalam satu domain
// CFcombine = CF1 + CF2(1-CF1)
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

// helper untuk menghitung jumlah domain di atas threshold
func CountAbove(domains map[string]float64, threshold float64) int {
	count := 0
	for _, v := range domains {
		if v >= threshold {
			count++
		}
	}
	return count
}

// ================= RISK EVALUATION =================

func EvaluateRisk(category string, cf map[string]float64) (string, error) {
	switch category {

	// ================= PRAKONSEPSI =================
	case "PRAKONSEPSI":

		A := cf["A"]
		B := cf["B"]
		C := cf["C"]
		D := cf["D"]

		// R1
		if A >= 0.70 && B >= 0.70 && C >= 0.70 && D >= 0.70 {
			return "Resiko SANGAT TINGGI", nil
		}

		// R2
		if A >= 0.70 && B >= 0.70 &&
			(C >= 0.50 || D >= 0.50) {
			return "Resiko TINGGI", nil
		}

		// R3
		if A >= 0.50 && B >= 0.50 &&
			C < 0.70 && D < 0.70 {
			return "Resiko SEDANG", nil
		}

		// R4
		if (A >= 0.50 || B >= 0.50 || C >= 0.50 || D >= 0.50) &&
			A < 0.70 && B < 0.70 && C < 0.70 && D < 0.70 {
			return "Resiko RINGAN", nil
		}

		// R5
		if A < 0.30 && B < 0.30 && C < 0.30 && D < 0.30 {
			return "Resiko RENDAH", nil
		}

	// ================= PERNAH MELAHIRKAN =================
	case "PERNAH_MELAHIRKAN":

		A := cf["A"]
		B := cf["B"]
		C := cf["C"]

		// R1
		if A >= 0.70 && B >= 0.70 && C >= 0.70 {
			return "Resiko SANGAT TINGGI", nil
		}

		// R2
		if CountAbove(cf, 0.70) >= 2 {
			return "Resiko TINGGI", nil
		}

		// R3
		if (A >= 0.50 || B >= 0.50 || C >= 0.50) &&
			A < 0.70 && B < 0.70 && C < 0.70 {
			return "Resiko SEDANG", nil
		}

		// R4
		if (A >= 0.30 || B >= 0.30 || C >= 0.30) &&
			A < 0.50 && B < 0.50 && C < 0.50 {
			return "Resiko RINGAN", nil
		}

		// R5
		if A < 0.30 && B < 0.30 && C < 0.30 {
			return "Resiko RENDAH", nil
		}

	// ================= REMAJA =================
	case "REMAJA_19":

		A := cf["A"]
		B := cf["B"]
		C := cf["C"]
		D := cf["D"]
		E := cf["E"]
		F := cf["F"]

		// R1
		if A >= 0.70 &&
			B >= 0.70 &&
			C >= 0.70 &&
			D >= 0.70 &&
			E >= 0.70 &&
			F >= 0.70 {
			return "Resiko SANGAT TINGGI", nil
		}

		// R2
		if A >= 0.70 &&
			B >= 0.70 &&
			(C >= 0.70 || D >= 0.70 || E >= 0.70 || F >= 0.70) {
			return "Resiko TINGGI", nil
		}

		// R3
		if (A >= 0.50 || B >= 0.50) &&
			CountAbove(cf, 0.50) >= 3 {
			return "Resiko SEDANG", nil
		}

		// R4
		if (A >= 0.50 ||
			B >= 0.50 ||
			C >= 0.50 ||
			D >= 0.50 ||
			E >= 0.50 ||
			F >= 0.50) &&
			A < 0.70 &&
			B < 0.70 &&
			C < 0.70 &&
			D < 0.70 &&
			E < 0.70 &&
			F < 0.70 {
			return "Resiko RINGAN", nil
		}

		// R5
		if A < 0.30 &&
			B < 0.30 &&
			C < 0.30 &&
			D < 0.30 &&
			E < 0.30 &&
			F < 0.30 {
			return "Resiko RENDAH", nil
		}
	}

	return "", errors.New("tidak dapat menentukan risiko")
}
