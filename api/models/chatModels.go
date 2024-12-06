package models

import "time"

type Chat struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Participants []string `json:"participants"`
	CreatedAt    time.Time `json:"createdAt" bson:"created_at"`
	Type         string    `json:"type" bson:"type"`
}