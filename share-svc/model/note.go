package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Content   string `bson:"content" json:"content"`
	Deactivated bool `bson:"deactivated" json:"deactivated"`

	CreatedAt           int64              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt           int64              `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeactivatedAt       int64              `bson:"deactivated_at,omitempty" json:"deactivated_at,omitempty"`
}
