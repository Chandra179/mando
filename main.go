package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mando/config"


func main() {
	// Load configuration from environment variables (.env file)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize MongoDB connection
	mongoDB, err := config.NewMongoDB(cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	if err = mongoDB.NewCollection(context.Background(), "skills"); err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}

	log.Printf("Connected to MongoDB database: %s", cfg.MongoDB.Database)
	collection := mongoDB.Database.Collection("skills")


	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-quit
	log.Println("Shutting down server...")

	// Close MongoDB connection
	if err := mongoDB.Close(); err != nil {
		log.Printf("Error closing MongoDB connection: %v", err)
	}

	log.Println("Server exited properly")
}
