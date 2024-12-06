package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`          // Auto-generated MongoDB ObjectID
	Username  string    `json:"username" bson:"username" validate:"required,min=3,max=50"` // Required, min length 3, max length 50
	Email     string    `json:"email" bson:"email" validate:"required,email"` // Required, must be a valid email
	Password string     `json:"-" bson:"password" validate:"required,min=4"`// Required, min length 6
	CreatedAt time.Time `json:"created_at" bson:"created_at"`     // Timestamp for when the user is created
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`     // Timestamp for when the user is updated
}
