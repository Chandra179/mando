package skills

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Skills struct {
	Collection *mongo.Collection
}

func InitSKills(collection *mongo.Collection) *Skills {
	return &Skills{
		Collection: collection,
	}
}

func (s *Skills) AddSkill(skill string) {
	user := bson.D{
		{Key: "username", Value: "johndoe"},
		{Key: "email", Value: "johndoe@example.com"},
		{Key: "created_at", Value: time.Now()},
	}
	result, err := s.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}

func RemoveSkill(skill string) {
	// Remove skill from the list
}

func ListSkills() []string {
	// List all skills
	return []string{"Go", "Python", "Java"}
}
