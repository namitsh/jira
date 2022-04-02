package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Description string             `json:"description" binding:"required"`
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" binding:"required"`
	Lead        string             `json:"lead" binding:"required"`
}
