package main

import (
	"flag"
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
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	shouldSeed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

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

	log.Println("🔄 Running migrations...")
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
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Migrations completed!")

	if *shouldSeed {
		log.Println("🌱 Seeding database...")
		if err := database.SeedAll(database.DB); err != nil {
			log.Fatal("❌ Failed to seed database:", err)
		}
		log.Println("✅ Database seeded successfully!")
	}

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

	app.Use(logger.New())
	app.Use(cors.New())

	// Public routes
	app.Post("/auth/register", handlers.Register)
	app.Post("/auth/login", handlers.Login)

	// Protected routes
	protected := app.Group("/api", middlewares.JWTProtected())

	// Question endpoints
	protected.Get("/questions", handlers.GetQuestions)
	protected.Get("/questions/:code", handlers.GetQuestionDetail)
	protected.Post("/diagnosis", handlers.CreateDiagnosis)
	protected.Get("/diagnosis", handlers.GetDiagnosisHistory)

	protected.Get("/profile", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		return c.JSON(fiber.Map{
			"user_id":  claims["user_id"],
			"category": claims["category"],
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "CF Stunting API running",
			"version": "1.0.0",
		})
	})

	log.Println("🚀 Server starting on port 8080...")
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
