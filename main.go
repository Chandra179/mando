package main

import (
	"context"
	"log"
	"mando/config"
	"mando/skills"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "mando/docs" // Import the generated docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Mando API
// @version         1.0
// @description     This is the Mando server API documentation
// @host            localhost:8080
// @BasePath        /api
// @schemes         http
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	mongoDB, err := config.NewMongoDB(cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	if err = mongoDB.NewCollection(context.Background(), "skills"); err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}
	log.Printf("Connected to MongoDB database: %s", cfg.MongoDB.Database)
	collection := mongoDB.Database.Collection("skills")
	skillsService := skills.InitSkills(collection)

	// Set up HTTP server
	router := gin.Default()
	setupRoutes(router, skillsService)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    "0.0.0.0:" + cfg.HttpServer.Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	log.Printf("Server started on :%s", cfg.HttpServer.Port)

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
	if err := mongoDB.Close(); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}
	log.Println("Server exited properly")
}

func setupRoutes(router *gin.Engine, skillsService *skills.Skills) {
	// API routes
	api := router.Group("/api")
	{
		// Skills routes
		skillsGroup := api.Group("/skills")
		{
			skillsGroup.POST("/add", skillsService.AddSkillHandler)
		}
	}
}
