package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joshua468/church_event_management/config"
	"github.com/joshua468/church_event_management/db"
	"github.com/joshua468/church_event_management/internal/controller"
	"github.com/joshua468/church_event_management/internal/middleware"
	"github.com/joshua468/church_event_management/internal/repository"
	"github.com/joshua468/church_event_management/internal/service"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Connect to the database
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Defer the closing of the database connection
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("Could not get database object: %v", err)
	}
	defer sqlDB.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(dbConn)
	eventRepo := repository.NewEventRepository(dbConn)

	// Initialize services
	userService := service.NewUserService(userRepo)
	eventService := service.NewEventService(eventRepo)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	eventController := controller.NewEventController(eventService)

	// Setup Gin router
	router := gin.Default()

	// JWT Middleware
	router.Use(middleware.JWTAuthMiddleware())

	// Register routes
	userController.RegisterRoutes(router)
	eventController.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
