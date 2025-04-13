package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// MongoDBConfig holds all MongoDB connection parameters
type MongoDBConfig struct {
	URI             string
	Database        string
	ConnectTimeout  time.Duration
	MaxConnIdleTime time.Duration
	MinPoolSize     uint64
	MaxPoolSize     uint64
	Username        string
	Password        string
}

// Config holds all application configurations
type Config struct {
	MongoDB MongoDBConfig
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Default values
	mongoConfig := MongoDBConfig{
		URI:             getEnv("MONGO_URI", "mongodb://mongodb:27017"),
		Database:        getEnv("MONGO_DATABASE", "mando"),
		ConnectTimeout:  time.Duration(getEnvAsInt("MONGO_CONNECT_TIMEOUT", 10)) * time.Second,
		MaxConnIdleTime: time.Duration(getEnvAsInt("MONGO_MAX_CONN_IDLE_TIME", 60)) * time.Second,
		MinPoolSize:     uint64(getEnvAsInt("MONGO_MIN_POOL_SIZE", 5)),
		MaxPoolSize:     uint64(getEnvAsInt("MONGO_MAX_POOL_SIZE", 100)),
		Username:        getEnv("MONGO_USERNAME", "root"),
		Password:        getEnv("MONGO_PASSWORD", "root"),
	}

	return &Config{
		MongoDB: mongoConfig,
	}, nil
}

// Helper function to read environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to read and convert environment variables to integers
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Could not parse %s environment variable as integer: %v", key, err)
		return defaultValue
	}

	return value
}
