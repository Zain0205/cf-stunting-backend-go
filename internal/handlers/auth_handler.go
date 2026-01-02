package handlers

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"github.com/Zain0205/cf-stunting-backend-go/internal/services"
	"github.com/Zain0205/cf-stunting-backend-go/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Name     string              `json:"name"`
	Phone    string              `json:"phone"`
	Password string              `json:"password"`
	Category models.UserCategory `json:"category"`
}

func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, 400, "Invalid request body")
	}

	if req.Name == "" || req.Phone == "" || req.Password == "" {
		return utils.Error(c, 400, "All fields are required")
	}

	if len(req.Password) < 8 {
		return utils.Error(c, 400, "Password minimum 8 characters")
	}

	if err := services.Register(req.Name, req.Phone, req.Password, req.Category); err != nil {
		return utils.Error(c, 400, err.Error())
	}

	return utils.Success(c, "User registered successfully")
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, 400, "Invalid request body")
	}

	token, err := services.Login(req.Phone, req.Password)
	if err != nil {
		return utils.Error(c, 401, err.Error())
	}

	return utils.Success(c, fiber.Map{
		"token": token,
	})
}
