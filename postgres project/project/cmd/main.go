package main

import (
	"log"
	"os"
	"project/internal/delivery"
	"project/internal/entity"
	"project/internal/repository"
	"project/internal/usecase"
	"project/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the database connection string
	dsn := os.Getenv("DATABASE_DSN")

	// Initialize PostgreSQL connection
	db := database.NewPostgresConnection(dsn)

	// Automatically migrate the User entity
	db.AutoMigrate(&entity.User{})

	// Initialize repository, usecase, and HTTP router
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	router := gin.Default()

	// Register HTTP routes
	delivery.NewUserHandler(router, userUsecase)

	// Start the HTTP server on port 5432
	router.Run(":5432")
}
