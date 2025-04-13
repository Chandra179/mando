package skills

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Skills struct {
	Collection *mongo.Collection
}

// @Description Model for adding skills
type AddSkillRequest struct {
	// List of skills to add
	Skills []string `json:"skills" binding:"required" example:"Go"`
}

func InitSkills(collection *mongo.Collection) *Skills {
	s := &Skills{
		Collection: collection,
	}
	return s
}

// AddSkillHandler godoc
// @Summary Add new skills
// @Description Add one or more skills
// @Tags skills
// @Accept json
// @Produce json
// @Param request body AddSkillRequest true "Skills to add"
// @Success 200 {object} map[string]string "message: Skills added successfully"
// @Failure 400 {object} map[string]string "error: error message"
// @Router /skills/add [post]
func (s *Skills) AddSkillHandler(c *gin.Context) {
	var req AddSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	s.AddSkill(req.Skills)
	c.JSON(200, gin.H{"message": "Skills added successfully"})
}

func (s *Skills) AddSkill(skills []string) {
	doc := bson.D{
		{Key: "skills", Value: skills},
		{Key: "created_at", Value: time.Now()},
	}
	result, err := s.Collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Printf("Error inserting skills: %v", err)
		return
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
