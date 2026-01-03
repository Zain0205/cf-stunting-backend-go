package handlers

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type DiagnosisRequest struct {
	Category string                 `json:"category"`
	Answers  []services.AnswerInput `json:"answers"`
}

func CreateDiagnosis(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var req DiagnosisRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	service := services.NewDiagnosisService()
	diag, err := service.CreateDiagnosis(userID, req.Category, req.Answers)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    diag,
	})
}

func GetDiagnosisHistory(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	service := services.NewDiagnosisService()
	history, err := service.GetHistoryByUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    history,
	})
}
