package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB represents a MongoDB client connection
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   MongoDBConfig
}

// NewMongoDB creates and returns a new MongoDB connection
func NewMongoDB(config MongoDBConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout)
	defer cancel()

	// Build connection URI if not explicitly provided
	uri := config.URI
	if uri == "mongodb://localhost:27017" && (config.Username != "" || config.Password != "") {
		uri = fmt.Sprintf("mongodb://%s:%s@localhost:27017", config.Username, config.Password)
	}

	// Configure client options
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(config.ConnectTimeout).
		SetMaxConnIdleTime(config.MaxConnIdleTime).
		SetMinPoolSize(config.MinPoolSize).
		SetMaxPoolSize(config.MaxPoolSize)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Successfully connected to MongoDB")
	return &MongoDB{
		Client:   client,
		Database: client.Database(config.Database),
		Config:   config,
	}, nil
}

// Close gracefully closes the MongoDB connection
func (m *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	log.Println("MongoDB connection closed")
	return nil
}
