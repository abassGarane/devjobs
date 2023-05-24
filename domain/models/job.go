package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	JobType      JobType            `json:"job_type,omitempty" bson:"job_type,omitempty,string"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Requirements string             `json:"requirements,omitempty" bson:"requirements,omitempty"`
	CompanyID    primitive.ObjectID `json:"company_id,omitempty" bson:"company_id,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
