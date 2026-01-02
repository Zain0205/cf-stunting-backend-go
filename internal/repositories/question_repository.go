// Package repositories
package repositories

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/database"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
)

// GetQuestionsByCategory - Ambil semua pertanyaan berdasarkan kategori user
func GetQuestionsByCategory(categoryCode string) ([]models.Question, error) {
	var questions []models.Question

	err := database.DB.
		Joins("JOIN domains ON domains.id = questions.domain_id").
		Joins("JOIN categories ON categories.id = domains.category_id").
		Where("categories.code = ?", categoryCode).
		Preload("Domain").
		Order("questions.code ASC").
		Find(&questions).Error

	return questions, err
}

// GetQuestionByCode - Ambil pertanyaan spesifik berdasarkan code
func GetQuestionByCode(categoryCode, questionCode string) (*models.Question, error) {
	var question models.Question

	err := database.DB.
		Joins("JOIN domains ON domains.id = questions.domain_id").
		Joins("JOIN categories ON categories.id = domains.category_id").
		Where("categories.code = ? AND questions.code = ?", categoryCode, questionCode).
		Preload("Domain").
		First(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// GetAnswerMappingsByQuestionID - Ambil mapping jawaban untuk pertanyaan tertentu
func GetAnswerMappingsByQuestionID(questionID uint) ([]models.AnswerMapping, error) {
	var mappings []models.AnswerMapping

	err := database.DB.
		Where("question_id = ?", questionID).
		Order("answer_key ASC").
		Find(&mappings).Error

	return mappings, err
}

// GetQuestionsByDomain - Ambil pertanyaan berdasarkan domain
func GetQuestionsByDomain(categoryCode, domainCode string) ([]models.Question, error) {
	var questions []models.Question

	err := database.DB.
		Joins("JOIN domains ON domains.id = questions.domain_id").
		Joins("JOIN categories ON categories.id = domains.category_id").
		Where("categories.code = ? AND domains.code = ?", categoryCode, domainCode).
		Preload("Domain").
		Order("questions.code ASC").
		Find(&questions).Error

	return questions, err
}

// GetDomainsByCategory - Ambil semua domain berdasarkan kategori
func GetDomainsByCategory(categoryCode string) ([]models.Domain, error) {
	var domains []models.Domain

	err := database.DB.
		Joins("JOIN categories ON categories.id = domains.category_id").
		Where("categories.code = ?", categoryCode).
		Order("domains.code ASC").
		Find(&domains).Error

	return domains, err
}
