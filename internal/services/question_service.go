// Package services
package services

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/repositories"
)

// QuestionResponse - Response structure untuk pertanyaan
type QuestionResponse struct {
	ID         uint                   `json:"id"`
	Code       string                 `json:"code"`
	Text       string                 `json:"text"`
	DomainCode string                 `json:"domain_code"`
	DomainName string                 `json:"domain_name"`
	CFPakar    float64                `json:"cf_pakar"`
	IsReverse  bool                   `json:"is_reverse"`
	Options    []AnswerOptionResponse `json:"options"`
}

type AnswerOptionResponse struct {
	AnswerKey  string  `json:"answer_key"`
	CFEvidence float64 `json:"cf_evidence"`
	Label      string  `json:"label"` // Label untuk UI
}

type DomainWithQuestionsResponse struct {
	Code      string             `json:"code"`
	Name      string             `json:"name"`
	Questions []QuestionResponse `json:"questions"`
}

// GetQuestionsByUserCategory - Ambil semua pertanyaan untuk kategori user
func GetQuestionsByUserCategory(categoryCode string) ([]DomainWithQuestionsResponse, error) {
	// Ambil semua domain
	domains, err := repositories.GetDomainsByCategory(categoryCode)
	if err != nil {
		return nil, err
	}

	var response []DomainWithQuestionsResponse

	for _, domain := range domains {
		// Ambil pertanyaan per domain
		questions, err := repositories.GetQuestionsByDomain(categoryCode, domain.Code)
		if err != nil {
			return nil, err
		}

		var questionResponses []QuestionResponse

		for _, q := range questions {
			// Ambil answer mappings
			mappings, err := repositories.GetAnswerMappingsByQuestionID(q.ID)
			if err != nil {
				return nil, err
			}

			var options []AnswerOptionResponse
			for _, m := range mappings {
				options = append(options, AnswerOptionResponse{
					AnswerKey:  m.AnswerKey,
					CFEvidence: m.CFEvidence,
					Label:      getAnswerLabel(categoryCode, q.Code, m.AnswerKey),
				})
			}

			questionResponses = append(questionResponses, QuestionResponse{
				ID:         q.ID,
				Code:       q.Code,
				Text:       q.Text,
				DomainCode: domain.Code,
				DomainName: domain.Name,
				CFPakar:    q.CFPakar,
				IsReverse:  q.IsReverse,
				Options:    options,
			})
		}

		response = append(response, DomainWithQuestionsResponse{
			Code:      domain.Code,
			Name:      domain.Name,
			Questions: questionResponses,
		})
	}

	return response, nil
}

// getAnswerLabel - Generate label untuk option jawaban
func getAnswerLabel(categoryCode, questionCode, answerKey string) string {
	// 1. Cek jawaban standar (Boolean/Common)
	standardLabels := map[string]string{
		"YA":          "Ya",
		"TIDAK":       "Tidak",
		"TIDAK_TAHU":  "Tidak Tahu",
		"TIDAK_INGAT": "Tidak Ingat",
	}

	if label, ok := standardLabels[answerKey]; ok {
		return label
	}

	// 2. Mapping spesifik berdasarkan Kategori dan Kode Pertanyaan
	// Pastikan struktur map benar: map[Category][QuestionCode][AnswerKey]
	labelMap := map[string]map[string]map[string]string{
		"PRAKONSEPSI": {
			"A1": {"0": "Tidak pernah", "1": "1 kali", "2": "2 kali", "3": "≥ 3 kali"},
			"A2": {"0": "Setiap hari", "1": "4-6 hari", "2": "1-3 hari", "3": "Tidak pernah"},
			"A3": {"0": "≥ 4 tablet", "1": "2-3 tablet", "2": "1 tablet", "3": "Tidak pernah"},
			"A4": {"0": "Selalu", "1": "Sering", "2": "Jarang", "3": "Tidak pernah"},
			"A5": {"0": "Ada rencana jelas", "1": "Rencana belum tertulis", "2": "Pernah dengar", "3": "Tidak tahu"},
		},
		"PERNAH_MELAHIRKAN": {
			"A1": {"0": "≥ 6 bulan", "1": "4-5 bulan", "2": "1-3 bulan", "3": "Tidak ASI eksklusif"},
			"A2": {"0": "≥ 6 hari", "1": "4-5 hari", "2": "2-3 hari", "3": "0-1 hari"},
			// ... tambahkan lainnya sesuai laporan kemajuan hal 7-8 [cite: 82, 83]
		},
	}

	if cat, ok := labelMap[categoryCode]; ok {
		if q, ok := cat[questionCode]; ok {
			if label, ok := q[answerKey]; ok {
				return label
			}
		}
	}

	return "Opsi " + answerKey
}

// GetQuestionDetail - Ambil detail satu pertanyaan
func GetQuestionDetail(categoryCode, questionCode string) (*QuestionResponse, error) {
	question, err := repositories.GetQuestionByCode(categoryCode, questionCode)
	if err != nil {
		return nil, err
	}

	mappings, err := repositories.GetAnswerMappingsByQuestionID(question.ID)
	if err != nil {
		return nil, err
	}

	var options []AnswerOptionResponse
	for _, m := range mappings {
		options = append(options, AnswerOptionResponse{
			AnswerKey:  m.AnswerKey,
			CFEvidence: m.CFEvidence,
			Label:      getAnswerLabel(categoryCode, question.Code, m.AnswerKey),
		})
	}

	return &QuestionResponse{
		ID:         question.ID,
		Code:       question.Code,
		Text:       question.Text,
		DomainCode: question.Domain.Code,
		DomainName: question.Domain.Name,
		CFPakar:    question.CFPakar,
		IsReverse:  question.IsReverse,
		Options:    options,
	}, nil
}
