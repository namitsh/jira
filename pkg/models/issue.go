package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// create a model here of all resources

type Issue struct {
	CreatedAt   time.Time          `json:"created_at"`
	Description string             `json:"description" binding:"required"`
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name"`
	Summary     string             `json:"summary"`
	Status      string             `json:"status"`
	IssueType   string             `json:"issue_type"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
