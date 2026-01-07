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
					Label:      m.Label,
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
			Label:      m.Label,
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
