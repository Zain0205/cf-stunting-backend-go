// Package handlers
package handlers

import (
	"log"

	"github.com/Zain0205/cf-stunting-backend-go/internal/services"
	"github.com/Zain0205/cf-stunting-backend-go/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetQuestions - GET /api/questions
// Ambil semua pertanyaan sesuai kategori user yang login
func GetQuestions(c *fiber.Ctx) error {
	// Ambil user dari JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	category := claims["category"].(string)

	// Ambil pertanyaan berdasarkan kategori
	questions, err := services.GetQuestionsByUserCategory(category)
	log.Println(questions)
	if err != nil {
		return utils.Error(c, 500, "Failed to fetch questions")
	}

	return utils.Success(c, fiber.Map{
		"category": category,
		"domains":  questions,
		"total":    countTotalQuestions(questions),
	})
}

// GetQuestionDetail - GET /api/questions/:code
// Ambil detail satu pertanyaan
func GetQuestionDetail(c *fiber.Ctx) error {
	// Ambil user dari JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	category := claims["category"].(string)

	questionCode := c.Params("code")

	question, err := services.GetQuestionDetail(category, questionCode)
	if err != nil {
		return utils.Error(c, 404, "Question not found")
	}

	return utils.Success(c, question)
}

// Helper function
func countTotalQuestions(domains []services.DomainWithQuestionsResponse) int {
	total := 0
	for _, d := range domains {
		total += len(d.Questions)
	}
	return total
}
