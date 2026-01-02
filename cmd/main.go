package main

import (
	"fmt"
	"log"

	"github.com/Zain0205/cf-stunting-backend-go/internal/config"
	"github.com/Zain0205/cf-stunting-backend-go/internal/database"
	"github.com/Zain0205/cf-stunting-backend-go/internal/handlers"
	"github.com/Zain0205/cf-stunting-backend-go/internal/middlewares"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Build DSN and connect database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
	)
	database.Connect(dsn)

	// Auto migrate models
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Domain{},
		&models.Question{},
		&models.AnswerMapping{},
		&models.Diagnosis{},
		&models.DiagnosisAnswer{},
		&models.DiagnosisDomain{},
	); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}

	// Run database seeder
	log.Println("🌱 Seeding database...")
	if err := database.SeedAll(database.DB); err != nil {
		log.Fatal("Failed to seed database:", err)
	}
	log.Println("✅ Database seeded successfully!")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Public routes
	app.Post("/auth/register", handlers.Register)
	app.Post("/auth/login", handlers.Login)

	// Protected routes
	protected := app.Group("/api", middlewares.JWTProtected())
	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "ok"})
	})

	// Health check / root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "CF Stunting API running",
			"version": "1.0.0",
		})
	})

	// Start server
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
