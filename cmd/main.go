package main

import (
	"fmt"

	"github.com/Zain0205/cf-stunting-backend-go/internal/config"
	"github.com/Zain0205/cf-stunting-backend-go/internal/database"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)

	database.Connect(dsn)

	database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Domain{},
		&models.Question{},
		&models.AnswerMapping{},
		&models.Diagnosis{},
		&models.DiagnosisAnswer{},
		&models.DiagnosisDomain{},
	)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "CF Stunting API running",
		})
	})

	app.Listen("0.0.0.0:8080")
}
