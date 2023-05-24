package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name"`
	LogoURL string             `json:"logo_url,omitempty" bson:"logo_url"`
}
