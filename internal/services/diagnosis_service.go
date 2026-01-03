package services

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/database"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"github.com/Zain0205/cf-stunting-backend-go/internal/repositories"
)

type AnswerInput struct {
	QuestionCode string `json:"question_code"`
	AnswerKey    string `json:"answer_key"`
}

type DiagnosisService struct {
	Repo repositories.DiagnosisRepository
}

func NewDiagnosisService() *DiagnosisService {
	return &DiagnosisService{
		Repo: repositories.DiagnosisRepository{
			DB: database.DB,
		},
	}
}

func (s *DiagnosisService) CreateDiagnosis(
	userID uint,
	category string,
	answers []AnswerInput,
) (*models.Diagnosis, error) {
	var diag models.Diagnosis
	diag.UserID = userID
	diag.Category = category

	var diagAnswers []models.DiagnosisAnswer
	domainBuckets := map[string][]float64{}

	for _, ans := range answers {

		var q models.Question
		if err := database.DB.Preload("Domain").
			Where("code = ?", ans.QuestionCode).
			First(&q).Error; err != nil {
			return nil, err
		}

		var mapping models.AnswerMapping
		if err := database.DB.
			Where("question_id = ? AND answer_key = ?", q.ID, ans.AnswerKey).
			First(&mapping).Error; err != nil {
			return nil, err
		}

		evidence := mapping.CFEvidence
		if q.IsReverse {
			evidence = 1 - evidence
		}

		cfItem := CalculateCFItem(q.CFPakar, evidence)

		diagAnswers = append(diagAnswers, models.DiagnosisAnswer{
			QuestionCode: q.Code,
			AnswerKey:    ans.AnswerKey,
			CFItem:       cfItem,
		})

		domainBuckets[q.Domain.Code] = append(domainBuckets[q.Domain.Code], cfItem)
	}

	cfDomains := map[string]float64{}
	var domainRecords []models.DiagnosisDomain

	for domain, values := range domainBuckets {
		cf := CombineCF(values)
		cfDomains[domain] = cf

		domainRecords = append(domainRecords, models.DiagnosisDomain{
			DomainCode: domain,
			CFValue:    cf,
		})
	}

	result, err := EvaluateRisk(category, cfDomains)
	if err != nil {
		return nil, err
	}

	diag.Result = result

	if err := s.Repo.CreateDiagnosis(&diag, diagAnswers, domainRecords); err != nil {
		return nil, err
	}

	return &diag, nil
}

func (s *DiagnosisService) GetHistoryByUser(userID uint) ([]models.Diagnosis, error) {
	return s.Repo.GetByUserID(userID)
}
